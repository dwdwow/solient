package jupiter

import (
	"fmt"
	"testing"

	"github.com/dwdwow/props"
)

func TestPrice(t *testing.T) {
	price, ti, err := Price("SOL", "USDT")
	props.PanicIfNotNil(err)
	fmt.Println(ti)
	props.PrintlnIndent(price)
}
