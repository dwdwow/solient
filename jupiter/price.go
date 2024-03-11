package jupiter

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func Price(token, vsToken string) (price TokenPrice, timeTaken float64, err error) {
	type Res struct {
		Data      map[string]TokenPrice `json:"data"`
		TimeTaken float64               `json:"timeTaken"`
	}
	v := url.Values{}
	v.Add("ids", token)
	if vsToken != "" {
		v.Add("vsToken", vsToken)
	}
	resp, err := resty.New().R().
		Get("https://price.jup.ag/v4/price?" + v.Encode())
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = fmt.Errorf("jupiter: http status code %v", resp.StatusCode())
		return
	}
	res := new(Res)
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		err = fmt.Errorf("jupiter: %w", err)
		return
	}
	timeTaken = res.TimeTaken
	price, ok := res.Data[token]
	if !ok {
		_vs := vsToken
		if vsToken == "" {
			_vs = "USDC"
		}
		err = fmt.Errorf("jupiter: no price of %v/%v", token, _vs)
	}
	return
}
