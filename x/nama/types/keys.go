package types

const (
	// ModuleName defines the module name
	ModuleName = "nama"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nama"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NamaAdminKey       = "nama/namaAdmin"
	NamaServiceCostKey = "nama/namaServiceCost"
	NamaDataKey        = "nama/namaData"
	NamaUserBalanceKey = "nama/namaUserBalance"
	TotalNamaKey       = "nama/totalNama"
	NamaIdKey          = "nama/namaId"
)
