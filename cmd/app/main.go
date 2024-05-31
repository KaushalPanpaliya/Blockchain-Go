package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp                 int64  // timestamp of when block is created
	Data, PrevBlockHash, Hash []byte // data contains valuable information, prevBlockHash stored prev block hash, Hash is the hash of this block
}

type Blockchain struct {
	blocks []*Block // an array with each element of "pointer to block" data type
}

func (b *Block) SetHash() {
	var timestamp = []byte(strconv.FormatInt(b.Timestamp, 10))                       // FormatInt converts Timestamp to decimal number
	var headers = bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // all properties other than current block hash concatenated
	var hash = sha256.Sum256(headers)                                                // hash of current block generated

	b.Hash = hash[:]
}

