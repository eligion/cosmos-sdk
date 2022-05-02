package v039

import sdk "github.com/eligion/cosmos-sdk/types"

const (
	ModuleName = "crisis"
)

type (
	GenesisState struct {
		ConstantFee sdk.Coin `json:"constant_fee" yaml:"constant_fee"`
	}
)
