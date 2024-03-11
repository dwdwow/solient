package jupiter

import "time"

type TokenPrice struct {
	Id            string  `json:"id"`
	MintSymbol    string  `json:"mintSymbol"`
	VsToken       string  `json:"vsToken"`
	VsTokenSymbol string  `json:"vsTokenSymbol"`
	Price         float64 `json:"price"`
}

type SwapRecord struct {
	Id                  int              `json:"id"`
	Owner               string           `json:"owner"`
	ProgramId           string           `json:"programId"`
	Signature           string           `json:"signature"`
	Timestamp           time.Time        `json:"timestamp"`
	LegCount            int              `json:"legCount"`
	VolumeInUSD         string           `json:"volumeInUSD"`
	InSymbol            string           `json:"inSymbol"`
	InAmount            string           `json:"inAmount"`
	InAmountInDecimal   string           `json:"inAmountInDecimal"`
	InAmountInUSD       string           `json:"inAmountInUSD"`
	InMint              string           `json:"inMint"`
	OutSymbol           string           `json:"outSymbol"`
	OutAmount           string           `json:"outAmount"`
	OutAmountInDecimal  string           `json:"outAmountInDecimal"`
	OutAmountInUSD      string           `json:"outAmountInUSD"`
	OutMint             string           `json:"outMint"`
	ExactOutAmount      any              `json:"exactOutAmount"`
	ExactOutAmountInUSD any              `json:"exactOutAmountInUSD"`
	ExactInAmount       any              `json:"exactInAmount"`
	ExactInAmountInUSD  any              `json:"exactInAmountInUSD"`
	Instruction         any              `json:"instruction"`
	SwapData            []SwapRecordData `json:"swapData"`
	FeeTokenPubkey      any              `json:"feeTokenPubkey"`
	FeeOwner            any              `json:"feeOwner"`
	FeeSymbol           any              `json:"feeSymbol"`
	FeeAmount           any              `json:"feeAmount"`
	FeeAmountInDecimal  any              `json:"feeAmountInDecimal"`
	FeeAmountInUSD      any              `json:"feeAmountInUSD"`
	FeeMint             any              `json:"feeMint"`
	TokenLedger         string           `json:"tokenLedger"`
	TransferAuthority   any              `json:"transferAuthority"`
}

type SwapRecordData struct {
	Amm                string  `json:"amm"`
	InMint             string  `json:"inMint"`
	OutMint            string  `json:"outMint"`
	InAmount           string  `json:"inAmount"`
	InSymbol           string  `json:"inSymbol"`
	OutAmount          string  `json:"outAmount"`
	OutSymbol          string  `json:"outSymbol"`
	InAmountInUSD      float64 `json:"inAmountInUSD,string"`
	OutAmountInUSD     float64 `json:"outAmountInUSD,string"`
	InAmountInDecimal  float64 `json:"inAmountInDecimal,string"`
	OutAmountInDecimal float64 `json:"outAmountInDecimal,string"`
}
