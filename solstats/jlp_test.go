package solstats

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/dwdwow/props"
	"github.com/dwdwow/solient"
)

func TestGetJLPInnerValue(t *testing.T) {
	value, err := GetJLPInnerValue(solient.AlchemyDocDemoRpc, 233891743)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(value)
}

func TestGetCurrentJLPInnerValue(t *testing.T) {
	value, err := GetCurrentJLPInnerValue()
	props.PanicIfNotNil(err)
	props.PrintlnIndent(value)
	props.PrintlnIndent(value.InnerValue())
}

func TestX(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString("Q+br7gAAAABD5uvuAAAAAEPm6+4AAAAAQ+br7gAAAABD5uvuAAAAAEPm6+4AAAAA")
	props.PanicIfNotNil(err)
	fmt.Println(string(data))
}
