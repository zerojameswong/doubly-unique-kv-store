package impl

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"zerojameswong/kvstore/model"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
		return errors.New("DB open error")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return errors.New("DB ping error")
	}

	fmt.Println("Connected!")
	return nil
}

func GetStores() ([]model.Store, error) {
	var stores []model.Store
	rows, err := db.Query("SELECT * FROM stores")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var s model.Store
		if err := rows.Scan(&s.StoreId, &s.StoreName); err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStoreById(storeId int64) (*model.Store, error) {
	var store model.Store
	row := db.QueryRow("SELECT * FROM stores WHERE store_id = ?", storeId)

	if err := row.Scan(&store.StoreId, &store.StoreName); err != nil {
		return nil, err
	}

	return &store, nil
}

func GetEntriesForStore(storeId int64) ([]model.Entry, error) {
	var entries []model.Entry
	rows, err := db.Query("SELECT * FROM entries WHERE store_id = ?", storeId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var e model.Entry
		if err := rows.Scan(&e.EntryId, &e.Key, &e.Value, &e.StoreId); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil

}
