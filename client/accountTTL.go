// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetAccountTTL Returns the period of inactivity after which the account of the current user will automatically be deleted
func (client *Client) GetAccountTTL() (*tdlib.AccountTTL, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getAccountTtl",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var accountTTL tdlib.AccountTTL
	err = json.Unmarshal(result.Raw, &accountTTL)
	return &accountTTL, err

}
