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
	err = h.ratesService.CreateRates(rate)
	if err != nil {
		logError("createRate", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
