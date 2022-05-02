package types

import (
	"github.com/eligion/cosmos-sdk/codec"
	cryptocodec "github.com/eligion/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
