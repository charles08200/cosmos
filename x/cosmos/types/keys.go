package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmos"
)

var (
	ParamsKey = []byte("p_cosmos")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
