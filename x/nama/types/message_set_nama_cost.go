package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetNamaServiceCost{}

const TypeSetNamaServiceCost = "set_nama_service_cost"

func NewMsgSetNamaServiceCost(creator string, newCost uint64) *MsgSetNamaServiceCost {
	return &MsgSetNamaServiceCost{
		Creator: creator,
		NewCost: newCost,
	}
}

func (msg *MsgSetNamaServiceCost) Route() string {
	return RouterKey
}

func (msg *MsgSetNamaServiceCost) Type() string {
	return TypeSetNamaServiceCost
}

func (msg *MsgSetNamaServiceCost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetNamaServiceCost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetNamaServiceCost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
