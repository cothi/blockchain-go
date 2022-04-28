package p2p

import (
	"encoding/json"
	"fmt"
	"strings"
	blockchaint "tetgo/tetgocoin/blockchain"
	"tetgo/tetgocoin/utill"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
	MessageNewBlockNotify
	MessageNewTxNotify
	MessageNewPeerNotify
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utill.ToJSON(payload),
	}
	return utill.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	fmt.Printf("Sending newest block to %s\n", p.key)
	b, err := blockchaint.FindBlock(blockchaint.Blockchain().NewestHash)
	utill.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchaint.Blocks(blockchaint.Blockchain()))
	p.inbox <- m
}

func notifyNewBlock(b *blockchaint.Block, p *peer) {
	m := makeMessage(MessageNewBlockNotify, b)
	p.inbox <- m
}

func notifyNewTx(tx *blockchaint.Tx, p *peer) {
	m := makeMessage(MessageNewTxNotify, tx)
	p.inbox <- m
}

func notifyNewPeer(address string, p *peer) {
	m := makeMessage(MessageNewPeerNotify, address)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		fmt.Printf("Received the newest block from %s\n", p.key)
		var payload blockchaint.Block
		utill.HandleErr(json.Unmarshal(m.Payload, &payload))
		b, err := blockchaint.FindBlock(blockchaint.Blockchain().NewestHash)
		utill.HandleErr(err)
		if payload.Height >= b.Height {
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			requestAllBlocks(p)
		} else {
			sendNewestBlock(p)
		}
	case MessageAllBlocksRequest:
		fmt.Printf("%s wants all the blocks.\n", p.key)
		sendAllBlocks(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Received all the blocks from %s\n", p.key)
		var payload []*blockchaint.Block
		utill.HandleErr(json.Unmarshal(m.Payload, &payload))
	case MessageNewBlockNotify:
		var payload *blockchaint.Block
		utill.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchaint.Blockchain().AddPeerBlock(payload)
	case MessageNewTxNotify:
		var payload *blockchaint.Tx
		utill.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchaint.Mempool().AddPeerTx(payload)
	case MessageNewPeerNotify:
		var payload string
		utill.HandleErr(json.Unmarshal(m.Payload, &payload))
		parts := strings.Split(payload, ":")
		AddPeer(parts[0], parts[1], parts[2], false)
	}
}
