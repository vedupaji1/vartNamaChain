package keeper

import (
	"nama/x/nama/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) SetNamaAdmin(ctx sdkTypes.Context, newAdmin string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	store.Set([]byte(types.NamaAdminKey), k.cdc.MustMarshal(&types.NamaAdmin{
		Value: newAdmin,
	}))
}

func (k *Keeper) GetNamaAdmin(ctx sdkTypes.Context) (res types.NamaAdmin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	k.cdc.MustUnmarshal(store.Get([]byte(types.NamaAdminKey)), &res)
	return res
}

func (k *Keeper) SetNamaServiceCostData(ctx *sdkTypes.Context, newCost uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	namaServiceCostData := types.NamaReserveCost{
		Value: newCost,
	}
	store.Set([]byte(types.NamaServiceCostKey), k.cdc.MustMarshal(&namaServiceCostData))
}

func (k *Keeper) GetNamaCost(ctx *sdkTypes.Context) (res types.NamaReserveCost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	k.cdc.MustUnmarshal(store.Get([]byte(types.NamaServiceCostKey)), &res)
	return res
}

func (k *Keeper) InitTotalNama(ctx *sdkTypes.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	store.Set([]byte(types.TotalNamaKey), k.cdc.MustMarshal(&types.TotalNama{
		Value: 0,
	}))
}

func (k *Keeper) GetTotalNama(ctx sdkTypes.Context) (res types.TotalNama) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	k.cdc.MustUnmarshal(store.Get([]byte(types.TotalNamaKey)), &res)
	return res
}

func (k *Keeper) SetNamaData(ctx sdkTypes.Context, namaData types.NamaData) uint64 {
	var totalNama types.TotalNama
	totalNamaStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	k.cdc.MustUnmarshal(totalNamaStore.Get([]byte(types.TotalNamaKey)), &totalNama)
	totalNama.Value += 1
	namaStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaDataKey))
	totalNamaBytesData := k.cdc.MustMarshal(&totalNama)
	namaStore.Set(totalNamaBytesData, k.cdc.MustMarshal(&namaData))
	totalNamaStore.Set([]byte(types.TotalNamaKey), totalNamaBytesData)
	prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaIdKey)).Set([]byte(namaData.Nama), totalNamaBytesData)
	userNamaBalanceStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaUserBalanceKey))
	userNamaBalance := types.TotalNama{}
	k.cdc.MustUnmarshal(userNamaBalanceStore.Get([]byte(namaData.Owner)), &userNamaBalance)
	userNamaBalance.Value++
	userNamaBalanceStore.Set([]byte(namaData.Owner), k.cdc.MustMarshal(&userNamaBalance))
	return totalNama.Value
}

func (k *Keeper) ChangeNamaOwnerAndPrice(ctx sdkTypes.Context, namaId types.NamaId, newOwner string, newPrice uint64) {
	namaIdBytes := k.cdc.MustMarshal(&namaId)
	namaData := types.NamaData{}
	namaStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaDataKey))
	k.cdc.MustUnmarshal(namaStore.Get(namaIdBytes), &namaData)
	userNamaBalanceStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaUserBalanceKey))
	currentNamaOwnerBalance := types.TotalNama{}
	k.cdc.MustUnmarshal(userNamaBalanceStore.Get([]byte(namaData.Owner)), &currentNamaOwnerBalance)
	currentNamaOwnerBalance.Value--
	userNamaBalanceStore.Set([]byte(namaData.Owner), k.cdc.MustMarshal(&currentNamaOwnerBalance))
	newNamaOwnerBalance := types.TotalNama{}
	k.cdc.MustUnmarshal(userNamaBalanceStore.Get([]byte(newOwner)), &newNamaOwnerBalance)
	newNamaOwnerBalance.Value++
	userNamaBalanceStore.Set([]byte(newOwner), k.cdc.MustMarshal(&newNamaOwnerBalance))
	namaData.Owner = newOwner
	namaData.PurchasedIn = newPrice
	namaStore.Set(namaIdBytes, k.cdc.MustMarshal(&namaData))
}

func (k *Keeper) GetNamaId(ctx sdkTypes.Context, nama string) (res types.NamaId) {
	k.cdc.MustUnmarshal(prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaIdKey)).Get([]byte(nama)), &res)
	return res
}

func (k *Keeper) GetNamaData(ctx sdkTypes.Context, namaId uint64) *types.NamaData {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaDataKey))
	namaData := types.NamaData{}
	k.cdc.MustUnmarshal(store.Get(k.cdc.MustMarshal(&types.TotalNama{Value: uint64(namaId)})), &namaData)
	return &namaData
}

func (k *Keeper) GetUserNamaBalance(ctx sdkTypes.Context, userAddress string) *types.TotalNama {
	totalNama := types.TotalNama{}
	k.cdc.MustUnmarshal(prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NamaUserBalanceKey)).Get([]byte(userAddress)), &totalNama)
	return &totalNama
}
