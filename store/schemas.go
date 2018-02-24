package store

type RecieveOffer struct {
	SocketID string `json:"address"`
	Address  string `json:"address"`
	Amount   int64  `json:"amount"`
	ChainID  int    `json:"chainid"`
}

type SendAbility struct {
	SocketID string   `json:"socketid"`
	BtIDs    []string `json:"btids`
}
