package main

import (
	"encoding/binary"
	"log"
	"net/http"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *badger.DB

func main() {
	var err error

	// Open a BadgerDB database with logging disabled
	opts := badger.DefaultOptions("./badgerdb").WithLogger(nil)
	db, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set initial balance
	err = updateBalance(db, 100)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define routes
	e.GET("/balance", getBalanceHandler)
	e.POST("/balance", updateBalanceHandler)

	// Start the server
	e.Logger.Fatal(e.Start("127.0.0.1:9000"))
}

func updateBalance(db *badger.DB, balance int) error {
	return db.Update(func(txn *badger.Txn) error {
		balanceBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(balanceBytes, uint32(balance))
		return txn.Set([]byte("balance"), balanceBytes)
	})
}

func getBalance(db *badger.DB) (int, error) {
	var balance int
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("balance"))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			balance = int(binary.LittleEndian.Uint32(val))
			return nil
		})
	})
	return balance, err
}

func getBalanceHandler(c echo.Context) error {
	balance, err := getBalance(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int{"balance": balance})
}

func updateBalanceHandler(c echo.Context) error {
	balanceStr := c.FormValue("balance")
	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid balance value"})
	}
	if err = updateBalance(db, balance); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Balance updated successfully"})
}
