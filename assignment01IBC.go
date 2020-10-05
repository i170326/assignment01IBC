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
	b.prevHash = CalculateHash(chainHead)
	return b
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	pc := chainHead
	for c := chainHead; c != nil; c = c.prevPointer {
		for i := 0; i < len(c.transactions); i++ {
			if c.transactions[i] == oldTrans {
				c.transactions[i] = newTrans
				c.currentHash = CalculateHash(c)
				pc.prevHash = CalculateHash(c)
			}
		}
		pc = c
	}
}

func ListBlocks(chainHead *Block) {
	for c := chainHead; c != nil; c = c.prevPointer {
		//print("Hash Current: ", c.currentHash, " ")
		for i := 0; i < len(c.transactions); i++ {
			fmt.Print(c.transactions[i], " , ")
		}
		//print(" Hash Previous:", c.prevHash, " ", "Hash Current: ", c.currentHash, " ")
		//print("Hash Previous: %x , Hash Current: %x ", c.prevHash, c.currentHash)
		fmt.Print(" ---> ")
	}
	fmt.Println("Blockchain Start")
}

func VerifyChain(chainHead *Block) {
	/*cBlock := chainHead
	if CalculateHash(chainHead) != chainHead.currentHash {
		fmt.Println("Blockchain is compromised")
		return
	}
	for p := chainHead.prevPointer; p != nil; p = p.prevPointer {
		hash_p := CalculateHash(cBlock)
		hash_c := CalculateHash(p)
		if hash_p != cBlock.prevHash || hash_c != p.currentHash {
			fmt.Println("Blockchain is compromised")
			return
		}
		cBlock = p
	}
	fmt.Println("Blockchain Verified")*/
	for c := chainHead; c != nil; c = c.prevPointer {
		hash_c := CalculateHash(c)
		if c.prevPointer != nil {
			hash_p := CalculateHash(c.prevPointer)
			if hash_p != c.prevHash || hash_c != c.currentHash {
				fmt.Println("Blockchain is compromised")
				return
			}
		}
		if hash_c != c.currentHash {
			fmt.Println("Blockchain is compromised")
			return
		}
	}
	fmt.Println("Blockchain Verified")
	return
}
