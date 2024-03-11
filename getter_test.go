package solient

import (
	"testing"

	"github.com/dwdwow/props"
)

func TestGetBlock(t *testing.T) {
	res, err := GetBlock(AlchemyDocDemoRpc, 253406744)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(res)
}

func TestGetTokenAccountsByOwner(t *testing.T) {
	res, err := GetTokenAccountsByOwner(AlchemyDocDemoRpc, "So11111111111111111111111111111111111111112", "AVzP2GeRmqGphJsMxWoqjpUifPpCret7LqWhD8NWQK49")
	props.PanicIfNotNil(err)
	props.PrintlnIndent(res)
}

func TestGetTokenSupply(t *testing.T) {
	res, err := GetTokenSupply(AlchemyDocDemoRpc, "27G8MtK7VtTcCHkpASjSDdkWWYfoqT6ggEuKidVJidD4")
	props.PanicIfNotNil(err)
	props.PrintlnIndent(res)
}
