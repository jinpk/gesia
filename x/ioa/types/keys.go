package types

const (
	// ModuleName defines the module name
	ModuleName = "ioa"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ioa"

	InternallyOwnedAccountPrefix = "ioa/"
)

var (
	ParamsKey = []byte("p_ioa")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
