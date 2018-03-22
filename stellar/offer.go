package stellar

import (
	"github.com/stellar/go/clients/horizon"
	"TravelShipper/model"

	b "github.com/stellar/go/build"
	"log"
)

func CreateOffer(offer model.OfferModel) (string, error) {
	var rate b.Rate

	if offer.BuyingCode == "" {
		rate = b.Rate{
			Selling: b.CreditAsset(offer.SellingCode, offer.SellingIssuer),
			Buying:  b.NativeAsset(),
			Price:   b.Price(offer.Price),
		}
	} else {
		if offer.SellingCode == "" {
			rate = b.Rate{
				Selling: b.NativeAsset(),
				Buying:  b.CreditAsset(offer.BuyingCode, offer.BuyingIssuer),
				Price:   b.Price(offer.Price),
			}
		} else {
			rate = b.Rate{
				Selling: b.CreditAsset(offer.SellingCode, offer.SellingIssuer),
				Buying:  b.CreditAsset(offer.BuyingCode, offer.BuyingIssuer),
				Price:   b.Price(offer.Price),
			}
		}
	}
	//rate = b.Rate{
	//	Selling: b.CreditAsset(offer.SellingCode, offer.SellingIssuer),
	//	Buying:  b.NativeAsset(),
	//	Price:   b.Price(offer.Price)}

	//rate = b.Rate{
	//				Selling: b.NativeAsset(),
	//				Buying:  b.CreditAsset(offer.BuyingCode, offer.BuyingIssuer),
	//				Price:   b.Price(offer.Price),
	//			}

	log.Println(rate)

	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: offer.From},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.CreateOffer(rate, b.Amount(offer.Amount)))
	txe := tx.Sign(offer.From)

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

func CreateAssetOffer(offer model.OfferModel) (string, error) {
	var rate b.Rate

	if offer.BuyingCode == "" {
		rate = b.Rate{
			Selling: b.CreditAsset(offer.SellingCode, offer.SellingIssuer),
			Buying:  b.NativeAsset(),
			Price:   b.Price(offer.Price),
		}
	} else {
		if offer.SellingCode == "" {
			rate = b.Rate{
				Selling: b.NativeAsset(),
				Buying:  b.CreditAsset(offer.BuyingCode, offer.BuyingIssuer),
				Price:   b.Price(offer.Price),
			}
		} else {
			rate = b.Rate{
				Selling: b.CreditAsset(offer.SellingCode, offer.SellingIssuer),
				Buying:  b.CreditAsset(offer.BuyingCode, offer.BuyingIssuer),
				Price:   b.Price(offer.Price),
			}
		}
	}
	//rate = b.Rate{
	//	Selling: b.CreditAsset(offer.SellingCode, offer.SellingIssuer),
	//	Buying:  b.NativeAsset(),
	//	Price:   b.Price(offer.Price)}

	//rate = b.Rate{
	//				Selling: b.NativeAsset(),
	//				Buying:  b.CreditAsset(offer.BuyingCode, offer.BuyingIssuer),
	//				Price:   b.Price(offer.Price),
	//			}

	log.Println(rate)

	tx := b.Transaction(
		b.SourceAccount{AddressOrSeed: offer.From},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.CreatePassiveOffer(rate, b.Amount(offer.Amount)))
	txe := tx.Sign(offer.From)

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
