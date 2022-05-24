package blockchain

import (
	"errors"
	"strings"
	"teelgo/blockchain-go/utils"
	"time"
)

type Block struct {
	Hash         string
	PrevHash     string
	Height       int
	Difficulty   int
	Nonce        int
	Timestamp    int
	Transactions []*Tx
}

func persistBlock(b *Block) {
	dbStorage.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := dbStorage.FindBlock(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height, diff int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: diff,
		Nonce:      0,
	}
	block.Transactions = Mempool().TxToConfirm()
	block.mine()
	persistBlock(block)
	return block
}
