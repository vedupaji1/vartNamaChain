package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReserveNama{}

const TypeReserveNama = "reserve_nama"

func NewReserveNama(creator string, nama string, price uint64) *MsgReserveNama {
	return &MsgReserveNama{
		Creator: creator,
		Nama:    nama,
		Price:   price,
	}
}

func (msg *MsgReserveNama) Route() string {
	return RouterKey
}

func (msg *MsgReserveNama) Type() string {
	return TypeReserveNama
}

func (msg *MsgReserveNama) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReserveNama) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReserveNama) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
