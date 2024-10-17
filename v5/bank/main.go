package main

import (
	"encoding/binary"
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"gofr.dev/pkg/gofr"
)

type Bank struct {
	Id      int     `json:"id"`
	Balance float32 `json:"balance"`
}

var db *badger.DB

func main() {
	// Open a BadgerDB database with logging disabled
	opts := badger.DefaultOptions("./badgerdb").WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// set initial balance
	err = updateBalance(db, 100.0)
	if err != nil {
		log.Fatal(err)
	}

	// get balance
	balance, err := getBalance(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Balance:", balance)

	a := gofr.New()

	err = a.AddRESTHandlers(&Bank{})
	if err != nil {
		log.Fatal(err)
	}
	a.Run()
}

func (b *Bank) Get(c *gofr.Context) (interface{}, error) {
	return getBalance(db)
}

func (b *Bank) Update(c *gofr.Context) (interface{}, error) {
	bal := c.Param("balance")
	if bal == "" {
		return nil, errors.New("balance is required")
	}
	balance, err := strconv.ParseFloat(bal, 32)
	if err != nil {
		return nil, err
	}
	if err = updateBalance(db, float32(balance)); err != nil {
		return nil, err
	}
	return balance, nil
}

func updateBalance(db *badger.DB, balance float32) (err error) {
	err = db.Update(func(txn *badger.Txn) error {
		balanceBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(balanceBytes, math.Float32bits(balance))
		return txn.Set([]byte("balance"), balanceBytes)
	})
	return err
}

func getBalance(db *badger.DB) (float32, error) {
	var balance float32
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("balance"))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			balance = math.Float32frombits(binary.LittleEndian.Uint32(val))
			return nil
		})
	})
	return balance, err
}
