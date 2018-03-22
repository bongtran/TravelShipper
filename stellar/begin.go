package stellar

import (
	"log"
	"github.com/stellar/go/keypair"
)

func CreateKeyPair() (*keypair.Full, error){
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SEED: " + pair.Seed())
	// SAV76USXIJOBMEQXPANUOQM6F5LIOTLPDIDVRJBFFE2MDJXG24TAPUU7
	log.Println("ADDRESS: " + pair.Address())
	// GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB

	return pair, err
}