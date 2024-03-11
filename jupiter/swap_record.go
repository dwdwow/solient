package jupiter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func GetSwapRecords(tokenId string, year, month, day int64) (records []SwapRecord, err error) {
	resp, err := resty.New().R().Get(fmt.Sprintf("https://stats.jup.ag/token-ledger/%v/%v-%v-%v", tokenId, year, month, day))
	if err != nil {
		err = fmt.Errorf("jupiter: get swap records, %w", err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		err = fmt.Errorf("jupiter: get swap records, http code %v", resp.StatusCode())
		return
	}
	err = json.Unmarshal(resp.Body(), &records)
	return
}
