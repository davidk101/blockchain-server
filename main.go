package main

// blockchain - public DB that doesn't rely on trust between nodes of data
// here, blockchain nodes could produce incorrect data and the DB can fix itself

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func main() {

}
