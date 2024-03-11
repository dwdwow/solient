package solient

type Block struct {
	BlockHeight       uint64       `json:"blockHeight"`
	BlockTime         int64        `json:"blockTime"`
	Blockhash         string       `json:"blockhash"`
	PreviousBlockhash string       `json:"previousBlockhash"`
	ParentSlot        uint64       `json:"parentSlot"`
	Rewards           BlockRewords `json:"rewards"`
	Transactions      []BlockTx    `json:"transactions"`
}

type BlockRewords struct {
	Commission  any    `json:"commission"`
	Lamports    int    `json:"lamports"`
	PostBalance int64  `json:"postBalance"`
	Pubkey      string `json:"pubkey"`
	RewardType  string `json:"rewardType"`
}

type BlockTx struct {
	Meta        BlockTxMeta
	Transaction BlockTxDetail
	Version     string
}

type BlockTxMeta struct {
	ComputeUnitsConsumed int    `json:"computeUnitsConsumed"`
	Err                  any    `json:"err"`
	Fee                  uint64 `json:"fee"`
	InnerInstructions    []any  `json:"innerInstructions"`
	LoadedAddresses      struct {
		Readonly []string `json:"readonly"`
		Writable []string `json:"writable"`
	} `json:"loadedAddresses"`
	LogMessages       []string      `json:"logMessages"`
	PostBalances      []int64       `json:"postBalances"`
	PostTokenBalances []interface{} `json:"postTokenBalances"`
	PreBalances       []int64       `json:"preBalances"`
	PreTokenBalances  []interface{} `json:"preTokenBalances"`
	Rewards           []any         `json:"rewards"`
	Status            struct {
		Ok interface{} `json:"Ok"`
	} `json:"status"`
}

type BlockTxMetaTokenBalance struct {
	AccountIndex  int64  `json:"accountIndex"`
	Mint          string `json:"mint"`
	Owner         string `json:"owner"`
	ProgramId     string `json:"programId"`
	UiTokenAmount struct {
		Amount         string  `json:"amount"`
		Decimals       int64   `json:"decimals"`
		UiAmount       float64 `json:"uiAmount"`
		UiAmountString string  `json:"uiAmountString"`
	} `json:"uiTokenAmount"`
}

type BlockTxDetail struct {
	Message struct {
		AccountKeys []string `json:"accountKeys"`
		Header      struct {
			NumReadonlySignedAccounts   int64 `json:"numReadonlySignedAccounts"`
			NumReadonlyUnsignedAccounts int64 `json:"numReadonlyUnsignedAccounts"`
			NumRequiredSignatures       int64 `json:"numRequiredSignatures"`
		} `json:"header"`
		Instructions []struct {
			Accounts       []int64 `json:"accounts"`
			Data           string  `json:"data"`
			ProgramIdIndex int64   `json:"programIdIndex"`
			StackHeight    any     `json:"stackHeight"`
		} `json:"instructions"`
		RecentBlockhash string `json:"recentBlockhash"`
	} `json:"message"`
	Signatures []string `json:"signatures"`
}
