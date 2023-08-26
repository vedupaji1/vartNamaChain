package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetNewAdmin{}

const TypeSetNewAdmin = "set_new_admin"

func NewMsgSetNewAdmin(creator string, newAdmin string) *MsgSetNewAdmin {
	return &MsgSetNewAdmin{
		Creator:  creator,
		NewAdmin: newAdmin,
	}
}

func (msg *MsgSetNewAdmin) Route() string {
	return RouterKey
}

func (msg *MsgSetNewAdmin) Type() string {
	return TypeSetNamaServiceCost
}

func (msg *MsgSetNewAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetNewAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetNewAdmin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.NewAdmin); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new admin address (%s)", err)
	}
	return nil
}
