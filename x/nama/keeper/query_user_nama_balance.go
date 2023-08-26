package keeper

import (
	"context"
	"fmt"
	"nama/x/nama/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) UserNamaBalance(ctx context.Context, req *types.QueryUserNamaBalanceRequest) (*types.QueryUserNamaBalanceResponse, error) {
	fmt.Println("Received UserNamaBalance Query: ", req)
	if _, err := sdkTypes.AccAddressFromBech32(req.UserAddress); err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Invalid User Address")
	}
	return &types.QueryUserNamaBalanceResponse{
		Value: k.GetUserNamaBalance(sdkTypes.UnwrapSDKContext(ctx), req.UserAddress).Value,
	}, nil
}
