package jupiter

type TokenPrice struct {
	Id            string  `json:"id"`
	MintSymbol    string  `json:"mintSymbol"`
	VsToken       string  `json:"vsToken"`
	VsTokenSymbol string  `json:"vsTokenSymbol"`
	Price         float64 `json:"price"`
}
