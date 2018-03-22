package stellar

import (
	"github.com/stellar/go/clients/horizon"
	"encoding/json"
)

func GetBalance() (int, []byte) {
	status := 0
	var result []byte
	account, err := horizon.DefaultTestNetClient.LoadAccount("GC2L5F3SVSD656STKS6GJVWP727XSMGKL5Q43HPGBIIWWG7YPGOABOV6")

	if err != nil {
		status = 500
		result, _ = json.Marshal(err)
	}else {
		result, _ = json.Marshal(account.Balances)
		status = 200
	}

	return status, result
}

func GetBalanceFromID(seed string) (int, []byte) {
	status := 0
	var result []byte
	account, err := horizon.DefaultTestNetClient.LoadAccount(seed)

	if err != nil {
		status = 500
		result, _ = json.Marshal(err)
	}else {
		result, _ = json.Marshal(account.Balances)
		status = 200
	}

	return status, result
}
