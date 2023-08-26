package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NamaCost(ctx context.Context, req *types.QueryNamaCostRequest) (*types.QueryNamaCostResponse, error) {
	fmt.Println("Received NamaCost Query: ", req)
	sdkCtx := sdkTypes.UnwrapSDKContext(ctx)
	return &types.QueryNamaCostResponse{
		Value: k.GetNamaCost(&sdkCtx).Value,
	}, nil
}
