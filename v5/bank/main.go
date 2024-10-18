package main

import (
	"encoding/binary"
	"log"
	"net/http"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/gin-gonic/gin"
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

	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin
	r := gin.Default()

	// Define routes
	r.GET("/balance", getBalanceHandler)
	r.POST("/balance", updateBalanceHandler)

	// Start the server
	if err := r.Run("127.0.0.1:9000"); err != nil {
		log.Fatal(err)
	}
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

func getBalanceHandler(c *gin.Context) {
	balance, err := getBalance(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func updateBalanceHandler(c *gin.Context) {
	balanceStr := c.PostForm("balance")
	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid balance value"})
		return
	}
	if err = updateBalance(db, balance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Balance updated successfully"})
}
