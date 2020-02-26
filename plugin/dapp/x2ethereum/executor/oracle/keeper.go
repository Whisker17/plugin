package oracle

import (
	"encoding/json"
	dbm "github.com/33cn/chain33/common/db"
	types2 "github.com/33cn/chain33/types"
	"github.com/33cn/plugin/plugin/dapp/x2ethereum/executor/common"
	"github.com/33cn/plugin/plugin/dapp/x2ethereum/types"
	"strings"
)

// oracle KV:
// id -> prophecy
// address -> power
type Keeper struct {
	db dbm.KV
	// TODO: use this as param instead
	consensusNeeded float64 // The minimum % of stake needed to sign claims in order for consensus to occur
}

// NewKeeper creates new instances of the oracle Keeper
func NewKeeper(db dbm.KV, consensusNeeded float64) Keeper {
	if consensusNeeded <= 0 || consensusNeeded > 1 {
		panic(types.ErrMinimumConsensusNeededInvalid)
	}
	return Keeper{
		db:              db,
		consensusNeeded: consensusNeeded,
	}
}

// GetProphecy gets the entire prophecy data struct for a given id
func (k Keeper) GetProphecy(id string) (Prophecy, error) {
	if id == "" {
		return NewEmptyProphecy(), types.ErrInvalidIdentifier
	}
	//store to localDB
	bz, err := k.db.Get([]byte(id))
	if err != nil {
		return NewEmptyProphecy(), types.ErrProphecyGet
	} else if bz == nil {
		return NewEmptyProphecy(), types.ErrProphecyNotFound
	}
	var dbProphecy DBProphecy
	err = json.Unmarshal(bz, &dbProphecy)
	if err != nil {
		return NewEmptyProphecy(), types2.ErrUnmarshal
	}

	deSerializedProphecy, err := dbProphecy.DeserializeFromDB()
	if err != nil {
		return NewEmptyProphecy(), types.ErrinternalDB
	}
	return deSerializedProphecy, nil
}

// setProphecy saves a prophecy with an initial claim
func (k Keeper) setProphecy(prophecy Prophecy) error {
	if prophecy.ID == "" {
		return types.ErrInvalidIdentifier
	}
	if len(prophecy.ClaimValidators) == 0 {
		return types.ErrNoClaims
	}
	serializedProphecy, err := prophecy.SerializeForDB()
	if err != nil {
		return types.ErrinternalDB
	}
	serializedProphecyBytes, err := json.Marshal(serializedProphecy)
	if err != nil {
		return types2.ErrMarshal
	}
	_ = k.db.Set([]byte(prophecy.ID), serializedProphecyBytes)
	return nil
}

// ProcessClaim TODO: validator hasn't implement
func (k Keeper) ProcessClaim(claim types.OracleClaim) (Status, error) {
	if strings.TrimSpace(claim.Content) == "" {
		return Status{}, types.ErrInvalidClaim
	}
	prophecy, err := k.GetProphecy(claim.ID)
	if err != nil {
		if err != types.ErrProphecyNotFound {
			return Status{}, err
		}
		prophecy = NewProphecy(claim.ID)
	} else {
		if prophecy.Status.Text == StatusText(types.EthBridgeStatus_SuccessStatusText) || prophecy.Status.Text == StatusText(types.EthBridgeStatus_FailedStatusText) {
			return Status{}, types.ErrProphecyFinalized
		}
		if prophecy.ValidatorClaims[claim.ValidatorAddress] != "" {
			return Status{}, types.ErrDuplicateMessage
		}
	}
	prophecy.AddClaim(common.NewChain33Address(claim.ValidatorAddress), claim.Content)
	prophecy, err = k.processCompletion(prophecy)
	err = k.setProphecy(prophecy)
	if err != nil {
		return Status{}, err
	}
	return prophecy.Status, nil
}

//todo
//func (k Keeper) checkActiveValidator(validatorAddress Chain33Address) bool {
//	validator, found := a.GetValidator(validatorAddress)
//	if !found {
//		return false
//	}
//	bondStatus := validator.GetStatus()
//	return bondStatus == sdk.Bonded
//}
//
//func (k Keeper) GetValidator(addr Chain33Address) ()

// 计算该prophecy是否达标
func (k Keeper) processCompletion(prophecy Prophecy) (Prophecy, error) {
	address2power := make(map[string]float64)
	validatorArrays, err := k.GetValidatorArray()
	if err != nil {
		return prophecy, err
	}
	for _, validator := range validatorArrays {
		powerBytes, err := k.db.Get([]byte(validator))
		if err != nil {
			return prophecy, err
		}
		var power float64
		err = json.Unmarshal(powerBytes, &power)
		if err != nil {
			return prophecy, types2.ErrUnmarshal
		}
		address2power[validator] = power
	}
	highestClaim, highestClaimPower, totalClaimsPower := prophecy.FindHighestClaim(address2power)
	totalPower, err := k.GetLastTotalPower()
	if err != nil {
		return prophecy, err
	}
	highestConsensusRatio := highestClaimPower / totalPower
	remainingPossibleClaimPower := totalPower - totalClaimsPower
	highestPossibleClaimPower := highestClaimPower + remainingPossibleClaimPower
	highestPossibleConsensusRatio := highestPossibleClaimPower / totalPower
	if highestConsensusRatio >= k.consensusNeeded {
		prophecy.Status.Text = StatusText(types.EthBridgeStatus_SuccessStatusText)
		prophecy.Status.FinalClaim = highestClaim
	} else if highestPossibleConsensusRatio < k.consensusNeeded {
		prophecy.Status.Text = StatusText(types.EthBridgeStatus_FailedStatusText)
	}
	return prophecy, nil
}

// Load the last total validator power.
func (k Keeper) GetLastTotalPower() (power float64, err error) {
	b, err := k.db.Get(types.LastTotalPowerKey)
	if err != nil && err != types2.ErrNotFound {
		return 0, err
	} else if err == types2.ErrNotFound {
		return 0, nil
	}
	err = json.Unmarshal(b, &power)
	if err != nil {
		return 0, types2.ErrUnmarshal
	}
	return
}

// Set the last total validator power.
func (k Keeper) SetLastTotalPower() error {
	var totalPower float64
	validatorArrays, err := k.GetValidatorArray()
	if err != nil {
		return err
	}
	for _, validator := range validatorArrays {
		powerBytes, err := k.db.Get([]byte(validator))
		if err != nil {
			return err
		}
		var power float64
		err = json.Unmarshal(powerBytes, &power)
		if err != nil {
			return types2.ErrUnmarshal
		}
		totalPower += power
	}
	powerBytes, err := json.Marshal(totalPower)
	if err != nil {
		return types2.ErrMarshal
	}
	err = k.db.Set(types.LastTotalPowerKey, powerBytes)
	if err != nil {
		return types.ErrSetKV
	}
	return nil
}

func (k Keeper) GetValidatorArray() ([]string, error) {
	validatorsBytes, err := k.db.Get(types.ValidatorMaps)
	if err != nil {
		return nil, err
	}
	var validatorArrays []string
	err = json.Unmarshal(validatorsBytes, &validatorArrays)
	if err != nil {
		return nil, types2.ErrUnmarshal
	}
	return validatorArrays, nil
}
