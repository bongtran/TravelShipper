package stellar

import (
	"fmt"
	b "github.com/stellar/go/build"
	horizon "github.com/stellar/go/clients/horizon"
	"log"
)

func SendPayment() {
	// address: GB6S3XHQVL6ZBAF6FIK62OCK3XTUI4L5Z5YUVYNBZUXZ4AZMVBQZNSAU
	from := "SAML6XU2QPSBDXI5OECBHXIN3XCPTEZKVRYSE7JGDQOQBVBIQFZYYLL6"

	// seed: SDLJZXOSOMKPWAK4OCWNNVOYUEYEESPGCWK53PT7QMG4J4KGDAUIL5LG
	to := "GDTU52EPFDXUYZEB27GWR2MMVL2MLYMKDUI4FMAFCGLJYUVWQ5S6IRUD"

	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: "0.1"},
		),
	)

	txe := tx.Sign(from)
	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}

func SendAPayment(from string, to string, amount string) (string, error) {
	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: amount},
		),
	)

	txe := tx.Sign(from)
	txeB64, err := txe.Base64()

	if err != nil {
		//panic(err)
		return "", err
	}

	//fmt.Printf("tx base64: %s", txeB64)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		//panic(err)
		return "", err
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)

	return resp.Links.Transaction.Href, err
}

func SendAnAssetPayment(from string, to string, issuer string, asset string, amount string) (string, error) {
	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		//b.Sequence{1},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			//b.Asset{asset, from, false},
			b.CreditAmount{Code: asset, Issuer: issuer, Amount: amount}))

	txe := tx.Sign(from)
	txeB64, err := txe.Base64()

	if err != nil {
		//panic(err)
		return "", err
	}
	//fmt.Printf("tx base64: %s", txeB64)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		//panic(err)
		return "", err
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)

	return resp.Links.Transaction.Href, err
}

func SendAnAssetPaymentToThirdPerson(from string, to string, issuer, asset string, amount string) (string, error) {
	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: to},
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.TestNetwork,
		b.Trust(asset, issuer, b.Limit("10000000")))

	txe := tx.Sign(to)

	txeB64, err := txe.Base64()

	if err != nil {
		return "", err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	log.Println(resp)

	if err != nil {
		return "", err
	}

	tx = b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		//b.Sequence{1},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			//b.Asset{asset, from, false},
			b.CreditAmount{Code: asset, Issuer: from, Amount: amount}))

	txe = tx.Sign(from)
	txeB64, err = txe.Base64()

	if err != nil {
		//panic(err)
		return "", err
	}
	//fmt.Printf("tx base64: %s", txeB64)

	resp, err = horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		//panic(err)
		return "", err
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)

	return resp.Links.Transaction.Href, err
}

func TrustAnAsset(to string, issuer, asset string, amount string) (string, error) {
	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: to},
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.TestNetwork,
		b.Trust(asset, issuer, b.Limit(amount)))

	txe := tx.Sign(to)

	txeB64, err := txe.Base64()

	if err != nil {
		return "", err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	log.Println(resp)
	return resp.Links.Transaction.Href, err
}

func initAsset(issuer string, receiving string, code string, limit string) (result string, err error) {

	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: receiving},
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.TestNetwork,
		b.Trust(code, issuer, b.Limit(limit)))

	txe := tx.Sign(receiving)

	txeB64, err := txe.Base64()

	if err != nil {
		return "", err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	log.Println(resp)

	if err != nil {
		return "", err
	}

	transaction := b.Transaction(
		b.SourceAccount{AddressOrSeed: issuer},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(b.Destination{AddressOrSeed: receiving},
			b.CreditAmount{Code: code, Issuer: issuer, Amount: limit}))

	txe = transaction.Sign(issuer)
	txeB64, err = txe.Base64()

	if err != nil {
		return "", err
	}

	resp, err = horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	result, err = AddHomeInformation(issuer)

	if err != nil {
		return "", err
	}

	result, err = LockAccount(issuer)
	if err != nil {
		return "", err
	}

	return resp.Links.Transaction.Href, err
}

func AddHomeInformation(issuingSeed string) (result string, err error) {
	transaction := b.Transaction(
		b.SourceAccount{AddressOrSeed: issuingSeed},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.SetOptions(b.HomeDomain("sunpannel.com")))

	txe := transaction.Sign(issuingSeed)

	txeB64, err := txe.Base64()
	if err != nil {
		return "", err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	if err != nil {
		return "", err
	}

	return resp.Links.Transaction.Href, err
}

func AddInformation(issuingSeed string, value string) (result string, err error) {
	transaction := b.Transaction(
		b.SourceAccount{AddressOrSeed: issuingSeed},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.SetOptions(b.HomeDomain(value)))

	txe := transaction.Sign(issuingSeed)

	txeB64, err := txe.Base64()
	if err != nil {
		return "", err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	if err != nil {
		return "", err
	}

	return resp.Links.Transaction.Href, err
}

func CheckTrust(checkAddress string, issuingAddress string, assetCode string) (bool, error) {
	result := false
	account, err := horizon.DefaultTestNetClient.LoadAccount(checkAddress)
	if err != nil {
		return false, err
	}

	for i := 0; i < len(account.Balances); i++ {
		balance := account.Balances[i]
		if balance.Asset.Type != "native" && balance.Asset.Code == assetCode && balance.Asset.Issuer == issuingAddress {
			result = true
		}
	}

	return result, err
}

func LockAccount(issuingSeed string) (result string, err error) {
	transaction := b.Transaction(
		b.SourceAccount{AddressOrSeed: issuingSeed},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.SetOptions(b.MasterWeight(0), b.SetThresholds(1, 1, 1)))

	txe := transaction.Sign(issuingSeed)

	txeB64, err := txe.Base64()
	if err != nil {
		return "", err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)

	if err != nil {
		return "", err
	}

	return resp.Links.Transaction.Href, err
}


