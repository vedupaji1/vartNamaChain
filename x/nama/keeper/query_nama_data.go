package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	// sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NamaData(sdkCtx context.Context, req *types.QueryNamaDataRequest) (*types.QueryNamaDataResponse, error) {
	fmt.Println("Received NamaData Query: ", req)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Nama Id Is Not Passed")
	}
	ctx := sdkTypes.UnwrapSDKContext(sdkCtx)
	if req.NamaId > k.GetTotalNama(ctx).Value {
		return nil, status.Error(codes.InvalidArgument, "Invalid Nama Id")
	}
	return &types.QueryNamaDataResponse{
		Data: k.GetNamaData(ctx, req.NamaId),
	}, nil
}
