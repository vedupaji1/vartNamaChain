package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAcquireNama{}

const TypeAcquireNama = "acquire_Nama"

func NewAcquireNama(creator string, namaId uint64, price uint64) *MsgAcquireNama {
	return &MsgAcquireNama{
		Creator: creator,
		NamaId:  namaId,
		Price:   price,
	}
}

func (msg *MsgAcquireNama) Route() string {
	return RouterKey
}

func (msg *MsgAcquireNama) Type() string {
	return TypeAcquireNama
}

func (msg *MsgAcquireNama) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcquireNama) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcquireNama) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
