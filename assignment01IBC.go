package assignment01IBC

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	hasher := sha256.New()
	var t_bytes = []byte{}
	for i := 0; i < len(inputBlock.transactions); i++ {
		b := []byte(inputBlock.transactions[i])
		for j := 0; j < len(b); j++ {
			t_bytes = append(t_bytes, b[j])
		}
	}
	hasher.Write(t_bytes)
	currentHash := hex.EncodeToString(hasher.Sum(nil))
	return currentHash
}

func calculateHash_String(input []string) string {
	hasher := sha256.New()
	var t_bytes = []byte{}

	for i := 0; i < len(input); i++ {
		b := []byte(input[i])
		for j := 0; j < len(b); j++ {
			t_bytes = append(t_bytes, b[j])
		}
	}
	hasher.Write(t_bytes)
	currentHash := hex.EncodeToString(hasher.Sum(nil))
	return currentHash
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
	if chainHead == nil {
		return &Block{transactionsToInsert, chainHead, "", calculateHash_String(transactionsToInsert)}
	}
	b := &Block{transactionsToInsert, chainHead, "", calculateHash_String(transactionsToInsert)}
	chainHead.prevHash = CalculateHash(b)
	return b
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for c := chainHead; c != nil; c = c.prevPointer {
		for i := 0; i < len(c.transactions); i++ {
			if c.transactions[i] == oldTrans {
				c.transactions[i] = newTrans
			}
		}
	}
}

func ListBlocks(chainHead *Block) {
	for c := chainHead; c != nil; c = c.prevPointer {
		for i := 0; i < len(c.transactions); i++ {
			fmt.Print("Transaction: ", i, c.transactions[i])
		}
		fmt.Print("--->")
	}
	fmt.Println("Blockchain Start")
}

func VerifyChain(chainHead *Block) {

}
