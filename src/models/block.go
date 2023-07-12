package models

type Block struct {
	Data          string
	PrevBlockHash []byte
	Hash          []byte
}
