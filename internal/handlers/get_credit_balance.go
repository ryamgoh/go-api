package handlers

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/ryamgoh/go-api/api"
	"github.com/ryamgoh/go-api/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetCreditBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CreditBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CreditDetails
	tokenDetails = (*database).GetUserCreditDetails(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CreditBalanceResponse{
		Balance: (*tokenDetails).Credits,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
