package solstats

import (
	"errors"
	"fmt"
	"time"

	"github.com/dwdwow/cex/bnc"
	"github.com/dwdwow/solient"
)

const (
	JLPBaseTokenLocatedAddr = "AVzP2GeRmqGphJsMxWoqjpUifPpCret7LqWhD8NWQK49"
)

type JLPInnerValue struct {
	Time      int64
	QtySol    float64
	PriceSol  float64
	QtyEth    float64
	PriceEth  float64
	QtyBtc    float64
	PriceBtc  float64
	QtyUsdc   float64
	PriceUsdc float64
	QtyUsdt   float64
	PriceUsdt float64
	Supply    float64
}

func (j JLPInnerValue) InnerValue() float64 {
	if j.Supply == 0 {
		return 0
	}
	qps := [][]float64{
		{j.QtySol, j.PriceSol},
		{j.QtyEth, j.PriceEth},
		{j.QtyBtc, j.PriceBtc},
		{j.QtyUsdc, j.PriceUsdc},
		{j.QtyUsdt, j.PriceUsdt},
	}
	var total float64
	for _, qp := range qps {
		total += qp[0] * qp[1]
		fmt.Println(qp[0], qp[1], qp[0]*qp[1])
	}
	fmt.Println(total)
	return total / j.Supply
}

func GetJLPInnerValue(rpcUrl string, targetBlock uint64) (value JLPInnerValue, err error) {
	var ok bool
	for start := targetBlock; start > 0; start-- {
		var block solient.Block
		block, err = solient.GetBlock(rpcUrl, start)
		if err != nil {
			return
		}
		if value.Time == 0 {
			value.Time = block.BlockTime * 1000
		}
		for _, tx := range block.Transactions {
			for _, bal := range tx.Meta.PostTokenBalances {
				if bal.Owner != JLPBaseTokenLocatedAddr {
					continue
				}
				switch bal.Mint {
				case solient.SOL_ID:
					fmt.Println("sol id equal")
					if value.QtySol == 0 {
						value.QtySol = bal.UiTokenAmount.UiAmount
					}
				case solient.ETH_ID:
					fmt.Println("eth id equal")
					if value.QtyEth == 0 {
						value.QtyEth = bal.UiTokenAmount.UiAmount
					}
				case solient.BTC_ID:
					fmt.Println("btc id equal")
					if value.QtyBtc == 0 {
						value.QtyBtc = bal.UiTokenAmount.UiAmount
					}
				case solient.USDC_ID:
					fmt.Println("usdc id equal")
					if value.QtyUsdc == 0 {
						value.QtyUsdc = bal.UiTokenAmount.UiAmount
					}
				case solient.USDT_ID:
					fmt.Println("usdt id equal")
					if value.QtyUsdt == 0 {
						value.QtyUsdt = bal.UiTokenAmount.UiAmount
					}
				}
			}
			if value.QtySol != 0 &&
				value.QtyEth != 0 &&
				value.QtyBtc != 0 &&
				value.QtyUsdc != 0 &&
				value.QtyUsdt != 0 {
				ok = true
				break
			}
		}
		if value.QtySol != 0 &&
			value.QtyEth != 0 &&
			value.QtyBtc != 0 &&
			value.QtyUsdc != 0 &&
			value.QtyUsdt != 0 {
			ok = true
			break
		}
	}
	if !ok {
		err = errors.New("not ok")
		return
	}
	getPrice := func(token string) (price float64, err error) {
		klines, err := bnc.QuerySpotKline(token+"USDC", "1m", value.Time, value.Time+time.Minute.Milliseconds())
		if err != nil {
			return
		}
		if len(klines) == 0 {
			err = errors.New(token + " klines len is 0")
			return
		}
		return klines[0].ClosePrice, nil
	}
	psol, err := getPrice("SOL")
	if err != nil {
		return
	}
	peth, err := getPrice("ETH")
	if err != nil {
		return
	}
	pbtc, err := getPrice("BTC")
	if err != nil {
		return
	}
	value.PriceUsdc = 1
	value.PriceUsdt = 1
	value.PriceSol = psol
	value.PriceEth = peth
	value.PriceBtc = pbtc
	return
}

func GetCurrentJLPInnerValue() (value JLPInnerValue, err error) {
	value.Time = time.Now().UnixMilli()
	getBal := func(mint string) (bal float64, err error) {
		acct, err := solient.GetTokenAccountsByOwner(solient.AlchemyDocDemoRpc, mint, JLPBaseTokenLocatedAddr)
		return acct.Data.Parsed.Info.TokenAmount.UiAmount, err
	}
	value.QtySol, err = getBal(solient.SOL_ID)
	if err != nil {
		return
	}
	value.QtyEth, err = getBal(solient.ETH_ID)
	if err != nil {
		return
	}
	value.QtyBtc, err = getBal(solient.WBTC_ID)
	if err != nil {
		return
	}
	value.QtyUsdt, err = getBal(solient.USDT_ID)
	if err != nil {
		return
	}
	value.QtyUsdc, err = getBal(solient.USDC_ID)
	if err != nil {
		return
	}
	getPrice := func(token string) (price float64, err error) {
		klines, err := bnc.QuerySpotKline(token+"USDT", "1m", value.Time-time.Minute.Milliseconds()*2, value.Time)
		if err != nil {
			return
		}
		if len(klines) == 0 {
			err = errors.New(token + " klines len is 0")
			return
		}
		return klines[0].ClosePrice, nil
	}
	psol, err := getPrice("SOL")
	if err != nil {
		return
	}
	peth, err := getPrice("ETH")
	if err != nil {
		return
	}
	pbtc, err := getPrice("BTC")
	if err != nil {
		return
	}
	pusdc, err := getPrice("USDC")
	if err != nil {
		return
	}
	value.PriceUsdt = 1
	value.PriceUsdc = pusdc
	value.PriceSol = psol
	value.PriceEth = peth
	value.PriceBtc = pbtc
	supply, err := solient.GetTokenSupply(solient.AlchemyDocDemoRpc, solient.JLP_ID)
	if err != nil {
		return
	}
	value.Supply = supply.UiAmount
	return
}
