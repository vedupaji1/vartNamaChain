package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NamaAdmin(ctx context.Context, req *types.QueryNamaAdminRequest) (*types.QueryNamaAdminResponse, error) {
	fmt.Println("Received NamaAdmin Query: ", req)
	return &types.QueryNamaAdminResponse{
		Value: k.GetNamaAdmin(sdkTypes.UnwrapSDKContext(ctx)).Value,
	}, nil
}
