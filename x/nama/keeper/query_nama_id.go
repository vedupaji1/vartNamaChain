package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NamaId(ctx context.Context, req *types.QueryNamaIdRequest) (*types.QueryNamaIdResponse, error) {
	fmt.Println("Received NamaId Query: ", req)
	return &types.QueryNamaIdResponse{
		NamaId: k.GetNamaId(sdkTypes.UnwrapSDKContext(ctx), req.Nama).Value,
	}, nil
}
