package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateBlock(oldBlock Block, BPM int64) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Bpm = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

// SHA256 hasing
func CalculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%d%s", block.Index, block.Timestamp, block.Bpm, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
