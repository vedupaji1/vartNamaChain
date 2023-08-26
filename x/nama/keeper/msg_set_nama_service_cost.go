package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (msg msgServer) SetNamaServiceCost(ctx context.Context, req *types.MsgSetNamaServiceCost) (*types.MsgSetNamaServiceCostResponse, error) {
	fmt.Println("Received SetNamaServiceCost Message: ", req)
	sdkCtx := sdkTypes.UnwrapSDKContext(ctx)
	msg.Keeper.SetNamaServiceCostData(&sdkCtx, req.NewCost)
	return &types.MsgSetNamaServiceCostResponse{}, nil
}
