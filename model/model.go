package model

type Store struct {
	StoreId   int64
	StoreName string
}

type Entry struct {
	EntryId  int64
	Key      string
	Value    string
	Category string
	StoreId  int64
}
