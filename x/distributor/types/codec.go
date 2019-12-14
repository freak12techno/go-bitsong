package types

import "github.com/cosmos/cosmos-sdk/codec"

// ModuleCdc is the codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateDistributor{}, "go-bitsong/MsgCreateDistributor", nil)

	cdc.RegisterConcrete(DistributorVerifyProposal{}, "go-bitsong/DistributorVerifyProposal", nil)
}