package nama_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nama/testutil/keeper"
	"nama/testutil/nullify"
	"nama/x/nama"
	"nama/x/nama/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NamaKeeper(t)
	nama.InitGenesis(ctx, *k, genesisState)
	got := nama.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
