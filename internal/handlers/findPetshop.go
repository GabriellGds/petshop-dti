package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GabriellGds/petshop-dti/internal/model"
	"github.com/GabriellGds/petshop-dti/internal/service"
	response "github.com/GabriellGds/petshop-dti/pkg"
)

func getQueryParameter(r *http.Request, key string) (int, error) {
	value := r.URL.Query().Get(key)
	return strconv.Atoi(value)
}

func ListPetshops(w http.ResponseWriter, r *http.Request) {
	var request model.Request
	request.Date = r.URL.Query().Get("date")
	smallDogs, err := getQueryParameter(r, "smallDogs")
	if err != nil {
		response.SendJSON(w, http.StatusBadRequest, model.ErrorResponse{Message: "invalid type"})
		return
	}
	request.SmallDogs = smallDogs
	bigDogs, err := getQueryParameter(r, "bigDogs")
	if err != nil {
		response.SendJSON(w, http.StatusBadRequest, model.ErrorResponse{Message: "invalid type"})
		return
	}
	request.BigDogs = bigDogs
	if request.BigDogs == 0 && request.SmallDogs == 0 {
		response.SendJSON(w, http.StatusBadRequest, model.ErrorResponse{Message: "No dogs provided, please provide at least one dog"})
	} 

	if err := request.Validate(); err != nil {
		response.SendJSON(w, http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	log.Print(request)

	petshop, price, err := service.ListPetshops(&request)
	if err != nil {
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	res := model.Response{Petshop: petshop.Name, TotalPrice: price}

	response.SendJSON(w, http.StatusOK, res)
}
