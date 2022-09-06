package controllers

import (
	"api-rest-go/models"
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
