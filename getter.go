package solient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	AlchemyDocDemoRpc = "https://solana-mainnet.g.alchemy.com/v2/alch-demo"
)

type rpcReqBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

func newRpcReqBody(method string, params []any) ([]byte, error) {
	body := rpcReqBody{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  method,
		Params:  params,
	}
	return json.Marshal(body)
}

type rpcRespBody[Result any] struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Result  Result `json:"result"`
}

func post[Result any](url string, method string, params []any) (res Result, err error) {
	body, err := newRpcReqBody(method, params)
	if err != nil {
		return
	}

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = fmt.Errorf("http status code %v", resp.StatusCode())
		return
	}

	_res := &rpcRespBody[Result]{}

	err = json.Unmarshal(resp.Body(), _res)

	return _res.Result, err
}

func rpcPoster[Value any](url string, method string, params []any) (v Value, err error) {
	res, err := post[RpcRespResult[Value]](url, method, params)
	return res.Value, err
}

func GetBlock(rpcUrl string, block uint64) (Block, error) {
	return post[Block](
		rpcUrl,
		"getBlock",
		[]any{
			block,
			map[string]any{
				"encoding":                       "json",
				"transactionDetails":             "full",
				"rewards":                        false,
				"maxSupportedTransactionVersion": 0,
			},
		},
	)
}

func GetTokenAccountsByOwner(rpcUrl string, mint, owner string) (acct TokenAccount, err error) {
	res, err := rpcPoster[[]TokenAccountRespValue](
		rpcUrl,
		"getTokenAccountsByOwner",
		[]any{
			owner,
			map[string]string{"mint": mint},
			map[string]string{"encoding": "jsonParsed"},
		},
	)
	if err != nil {
		return
	}
	for _, v := range res {
		if v.Account.Data.Parsed.Info.Owner == owner && v.Account.Data.Parsed.Info.Mint == mint {
			return v.Account, nil
		}
	}
	return
}

func GetTokenSupply(rpcUrl string, mint string) (amt TokenAmount, err error) {
	return rpcPoster[TokenAmount](rpcUrl, "getTokenSupply", []any{mint})
}
