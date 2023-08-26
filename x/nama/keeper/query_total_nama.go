package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) TotalNama(ctx context.Context, req *types.QueryTotalNamaRequest) (*types.QueryTotalNamaResponse, error) {
	fmt.Println("Received TotalNama Query: ", req)
	return &types.QueryTotalNamaResponse{
		Value: k.GetTotalNama(sdkTypes.UnwrapSDKContext(ctx)).Value,
	}, nil
}
