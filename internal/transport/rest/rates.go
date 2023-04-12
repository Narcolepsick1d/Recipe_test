package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"reciept/internal/models"
)

func (h *Handler) createRate(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logError("createRate", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var rate models.Rates
	if err = json.Unmarshal(reqBytes, &rate); err != nil {
		logError("createRate", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ratesService.UpdateRates(rate)
	if err != nil {
		logError("createRate", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) filteredByRating(w http.ResponseWriter, r *http.Request) {
	rate, err := getRateFromRequest(r)
	if err != nil {
		logError("getRateFromRequest", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recipes, err := h.ratesService.FilteredByRates(rate)
	if err != nil {
		logError("getAllRecipesByRate", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(recipes)
	if err != nil {
		logError("getAllRecipes", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)

}
