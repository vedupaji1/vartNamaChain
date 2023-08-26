package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "nama/testutil/keeper"
	"nama/x/nama/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NamaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
