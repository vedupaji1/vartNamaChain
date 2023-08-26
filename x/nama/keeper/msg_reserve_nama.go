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

func (msg msgServer) ReserveNama(ctx context.Context, req *types.MsgReserveNama) (*types.MsgReserveNamaResponse, error) {
	fmt.Println("Received ReserveNama Message: ", req)
	sdkCtx := sdkTypes.UnwrapSDKContext(ctx)
	creatorAddress, err := sdkTypes.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Invalid Creator Address")
	}
	if req.Price > msg.Keeper.bank.GetBalance(sdkCtx, creatorAddress, "token").Amount.Uint64() {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, "Creator Address Have Insufficient Balance")
	}
	if req.Price < msg.Keeper.GetNamaCost(&sdkCtx).Value {
		return nil, status.Error(codes.InvalidArgument, "Nama Price Must Be More Than Min Nama Cost")
	}
	if msg.Keeper.GetNamaId(sdkCtx, req.Nama).Value != 0 {
		return nil, status.Error(codes.InvalidArgument, "Nama Already Reserved")
	}

	// tempAddress, _ := sdkTypes.AccAddressFromBech32("cosmos128xyc0tel8a2gryn6sqvrxk2eh7958f8rckpx9")
	// err = msg.Keeper.bank.SendCoins(sdkCtx, creatorAddress, tempAddress, sdkTypes.NewCoins(sdkTypes.NewCoin("token", sdkmath.NewIntFromUint64(req.Price))))
	err = msg.Keeper.bank.SendCoinsFromAccountToModule(sdkCtx, creatorAddress, types.ModuleName, sdkTypes.NewCoins(sdkTypes.NewCoin("token", sdkmath.NewIntFromUint64(req.Price))))
	if err != nil {
		fmt.Println("Received Error From Bank Module: ", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	fmt.Println("TempOp10101")
	return &types.MsgReserveNamaResponse{
		NamaId: msg.Keeper.SetNamaData(sdkCtx, types.NamaData{
			Nama:        req.Nama,
			Owner:       req.Creator,
			PurchasedIn: req.Price,
		}),
	}, nil
}
