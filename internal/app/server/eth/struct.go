package eth

type Block struct {
	Number     int64  `json:"block_num"`
	Hash       string `json:"block_hash"`
	Time       int64  `json:"block_time"`
	ParentHash string `json:"parent_hash"`
}
