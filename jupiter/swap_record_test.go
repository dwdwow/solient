package jupiter

import (
	"testing"

	"github.com/dwdwow/props"
	"github.com/dwdwow/solient"
)

func TestGetTokenLedger(t *testing.T) {
	res, err := GetTokenLedger(solient.SOL_ID, 2024, 3, 10)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(res)
}
