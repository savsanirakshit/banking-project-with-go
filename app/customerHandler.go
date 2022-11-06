package app

import (
	"Banking/err"
	"Banking/logger"
	"Banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	restOps(w, r, err, customers)
}

func (ch *CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)
	customers, err := ch.service.GetCustomerById(customerId["customer_id"])
	restOps(w, r, err, customers)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post")
}

func restOps(w http.ResponseWriter, r *http.Request, err *err.AppError, response any) {
	if err != nil {
		responseFormatter(w, r, err.AsMessage(), err.Code)
	} else {
		responseFormatter(w, r, response, http.StatusOK)
	}
}

func responseFormatter(w http.ResponseWriter, r *http.Request, response any, statusCode int) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		w.WriteHeader(statusCode)
		if error := xml.NewEncoder(w).Encode(response); error != nil {
			logger.Error(error.Error())
			panic(error)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		if error := json.NewEncoder(w).Encode(response); error != nil {
			logger.Error(error.Error())
			panic(error)
		}
	}
}
