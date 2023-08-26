package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (msg msgServer) SetNewAdmin(ctx context.Context, req *types.MsgSetNewAdmin) (*types.MsgSetNewAdminResponse, error) {
	fmt.Println("Received SetNewAdmin Message")
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if req.Creator != msg.Keeper.GetNamaAdmin(sdkCtx).Value {
		fmt.Println("Unauthorized Access")
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, "Only Nama Admin Can Set New Admin")
	}
	msg.Keeper.SetNamaAdmin(sdkCtx, req.NewAdmin)
	return &types.MsgSetNewAdminResponse{}, nil
}
