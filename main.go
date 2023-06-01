package main

import (
	"github.com/rexposadas/chains/models"
	"log"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var mutex = &sync.Mutex{}

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := models.Block{}
		genesisBlock = models.Block{0, t.String(), 0, models.CalculateHash(genesisBlock), ""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()

	log.Fatal(run())
}
