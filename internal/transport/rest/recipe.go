package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reciept/internal/models"
)

func (h *Handler) getRecipeByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("getRecipeByID", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recipe, err := h.recipesService.GetByID(id)
	if err != nil {
		if errors.Is(err, models.ErrRecipeNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logError("getRecipeByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(recipe)
	if err != nil {
		logError("getRecipeByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) createRecipe(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logError("createRecipe", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var recipe models.Recipe
	if err = json.Unmarshal(reqBytes, &recipe); err != nil {
		logError("createRecipe", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.recipesService.Create(recipe)
	if err != nil {
		logError("createRecipe", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) deleteRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("deleteRecipe", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.recipesService.Delete(id)
	if err != nil {
		logError("deleteRecipe", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.recipesService.GetAll()
	if err != nil {
		logError("getAllRecipes", err)
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

func (h *Handler) updateRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("updateRecipe", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logError("updateRecipe", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp models.RecipeUpdate
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		logError("updateRecipe", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.recipesService.Update(id, inp)
	if err != nil {
		logError("updateRecipe", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *Handler) getByIngredient(w http.ResponseWriter, r *http.Request) {
	recipes, err := getIngredientFromRequest(r)
	if err != nil {
		logError("getByIngredients", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	list, err := h.recipesService.GetByIngredient(recipes)
	if err != nil {

		if errors.Is(err, models.ErrRecipeNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logError("getRecipeByIngredient", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(list)
	if err != nil {
		logError("getByIngredients", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
