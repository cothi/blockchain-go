package blockchain

import (
	"reflect"
	"testing"
)

func TestBlock(t *testing.T) {
	blockchain := GetBlockchain()
	blockchain.AddBlock("get")
	blockchain.GetBlock(200)
	blockchain.AllBlocks()
}

func TestBlock_calculateHash(t *testing.T) {
	type fields struct {
		Data     string
		Hash     string
		PrevHash string
		Height   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Data:     tt.fields.Data,
				Hash:     tt.fields.Hash,
				PrevHash: tt.fields.PrevHash,
				Height:   tt.fields.Height,
			}
			b.calculateHash()
		})
	}
}

func Test_getLastHash(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastHash(); got != tt.want {
				t.Errorf("getLastHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createBlock(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *Block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBlock(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blockchain_AddBlock(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &blockchain{
				blocks: tt.fields.blocks,
			}
			b.AddBlock(tt.args.data)
		})
	}
}

func TestGetBlockchain(t *testing.T) {
	tests := []struct {
		name string
		want *blockchain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBlockchain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockchain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blockchain_AllBlocks(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &blockchain{
				blocks: tt.fields.blocks,
			}
			if got := b.AllBlocks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blockchain.AllBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blockchain_GetBlock(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	type args struct {
		height int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Block
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &blockchain{
				blocks: tt.fields.blocks,
			}
			got, err := b.GetBlock(tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("blockchain.GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blockchain.GetBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
