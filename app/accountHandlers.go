package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andiahmads/bangking-service/dto"
	"github.com/andiahmads/bangking-service/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	//panggil service
	service service.AccountService
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(err)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusOK, account)
		}

	}

}
