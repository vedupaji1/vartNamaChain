package nama

import (
	"nama/x/nama/keeper"
	"nama/x/nama/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetNamaAdmin(ctx, "cosmos1qpz7fmms6z628l3dtcjhds54v06y88g2m07z3t")
	k.InitTotalNama(&ctx)
	k.SetNamaServiceCostData(&ctx, genState.NamaServiceCost)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
