package rest

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"reciept/internal/models"
	"strconv"
)

type Recipes interface {
	Create(recipe models.Recipe) error
	GetByID(id int64) (models.Recipe, error)
	GetAll() ([]models.Recipe, error)
	Delete(id int64) error
	Update(id int64, inp models.RecipeUpdate) error
	GetByIngredient(ingredient string) ([]models.Recipe, error)
	FilteredByTime(totalTime int) ([]models.Recipe, error)
}
type Rates interface {
	CreateRates(rate models.Rates) error
}
type User interface {
	SignUp(inp models.SignUpInput) error
	SignIn(inp models.SignInInput) (string, string, error)
	ParseToken(accessToken string) (int64, error)
	RefreshTokens(refreshToken string) (string, string, error)
}
type Handler struct {
	recipesService Recipes
	ratesService   Rates
	usersService   User
}

func NewHandler(recipe Recipes, users User, rates Rates) *Handler {
	return &Handler{
		recipesService: recipe,
		usersService:   users,
		ratesService:   rates,
	}
}
func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodGet)
		auth.HandleFunc("/refresh", h.refresh).Methods(http.MethodGet)
	}
	recipes := r.PathPrefix("/recipes").Subrouter()
	{
		recipes.HandleFunc("", h.getAllRecipes).Methods(http.MethodGet)
		recipes.HandleFunc("/total_time/{total_time}", h.getByTotalTime).Methods(http.MethodGet)

	}
	recipes = r.PathPrefix("/recipes").Subrouter()

	{

		recipes.Use(h.authMiddleware)
		recipes.HandleFunc("/rates", h.createRate).Methods(http.MethodPost)
		recipes.HandleFunc("/ingredient/{ingredient}", h.getByIngredient).Methods(http.MethodGet)
		recipes.HandleFunc("", h.createRecipe).Methods(http.MethodPost)
		recipes.HandleFunc("/{id:[0-9]+}", h.getRecipeByID).Methods(http.MethodGet)
		recipes.HandleFunc("/{id:[0-9]+}", h.deleteRecipe).Methods(http.MethodDelete)
		recipes.HandleFunc("/{id:[0-9]+}", h.updateRecipe).Methods(http.MethodPut)
	}

	return r
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
func getIngredientFromRequest(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	ingredient, err := vars["ingredient"]
	if err != true {
		return "", errors.New("Неправильно написанно ingredients")
	}
	if ingredient == "" {
		return "", errors.New("no ingredient is mention ")
	}
	return ingredient, nil
}
func getTotalTimeFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	totalTime, err := strconv.ParseInt(vars["total_time"], 10, 64)
	if err != nil {
		return 0, err
	}

	if totalTime == 0 {
		return 0, errors.New("totalTime can't be 0")
	}

	return int(totalTime), nil
}
