package routers

import (
	"github.com/gorilla/mux"
	"TravelShipper/controllers"
	"TravelShipper/common"
)

func SetStellarRouters(router *mux.Router) *mux.Router {

	stellarRouter := mux.NewRouter()
	stellarRouter.HandleFunc("/stellar/balance", controllers.GetBalance).Methods("GET")
	stellarRouter.HandleFunc("/stellar/balance/{id}", controllers.GetBalanceFromID).Methods("GET")
	stellarRouter.HandleFunc("/stellar/account/create", controllers.CreateAccount).Methods("POST")
	stellarRouter.HandleFunc("/stellar/ico", controllers.CreateICO).Methods("POST")
	stellarRouter.HandleFunc("/stellar/transfer", controllers.Transfer).Methods("POST")
	stellarRouter.HandleFunc("/stellar/transferto3rd", controllers.TransferToThirdPerson).Methods("POST")
	stellarRouter.HandleFunc("/stellar/trustasset", controllers.TrustAsset).Methods("POST")
	stellarRouter.HandleFunc("/stellar/trustedasset", controllers.CheckTrust).Methods("POST")
	stellarRouter.HandleFunc("/stellar/addinformation", controllers.AddInformation).Methods("POST")
	stellarRouter.HandleFunc("/stellar/lockaccount", controllers.LockAccount).Methods("POST")
	stellarRouter.HandleFunc("/stellar/offer", controllers.CreateOffer).Methods("POST")
	stellarRouter.HandleFunc("/stellar/passiveoffer", controllers.CreatePassiveOffer).Methods("POST")
	router.PathPrefix("/stellar").Handler(common.AuthorizeRequest(stellarRouter))

	return router
}
