package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkmath "cosmossdk.io/math"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (msg msgServer) AcquireNama(ctx context.Context, req *types.MsgAcquireNama) (*types.MsgAcquireNamaResponse, error) {
	fmt.Println("Received AcquireNama Message: ", req)
	sdkCtx := sdkTypes.UnwrapSDKContext(ctx)
	creatorAddress, err := sdkTypes.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Invalid Creator Address")
	}
	if req.NamaId > msg.Keeper.GetTotalNama(sdkCtx).Value {
		return nil, status.Error(codes.InvalidArgument, "Invalid Nama Id")
	}
	namaData := msg.Keeper.GetNamaData(sdkCtx, req.NamaId)
	oldCreatorAddress, err := sdkTypes.AccAddressFromBech32(namaData.Owner)
	if err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Invalid Creator Address")
	}
	if req.Price <= msg.Keeper.GetNamaData(sdkCtx, req.NamaId).PurchasedIn {
		return nil, status.Error(codes.InvalidArgument, "Nama Price Must Be More Than Existing Nama Price")
	}
	if req.Price > msg.Keeper.bank.GetBalance(sdkCtx, creatorAddress, "token").Amount.Uint64() {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, "Creator Address Have Insufficient Balance")
	}
	if req.Creator == namaData.Owner {
		return nil, status.Error(codes.InvalidArgument, "Existing Owner Cannot Acquire Nama")
	}
	msg.Keeper.bank.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, oldCreatorAddress, sdkTypes.NewCoins(sdkTypes.NewCoin("token", sdkmath.NewIntFromUint64(namaData.PurchasedIn))))
	msg.Keeper.ChangeNamaOwnerAndPrice(sdkCtx, types.NamaId{
		Value: req.NamaId,
	}, req.Creator, req.Price)
	msg.Keeper.bank.SendCoinsFromAccountToModule(sdkCtx, creatorAddress, types.ModuleName, sdkTypes.NewCoins(sdkTypes.NewCoin("token", sdkmath.NewIntFromUint64(req.Price))))
	return &types.MsgAcquireNamaResponse{}, nil
}
