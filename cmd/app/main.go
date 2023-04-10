package main

import (
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"reciept/internal/config"
	"reciept/internal/repository/psql"
	"reciept/internal/service"
	"reciept/internal/transport/rest"
	"reciept/pkg/database"
	"reciept/pkg/hash"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username_DB,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
		Password: cfg.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	hasher := hash.NewSHA1Hasher("salt")

	recipeRepo := psql.NewRecipe(db)
	recipeService := service.NewRecipe(recipeRepo)

	usersRepo := psql.NewUsers(db)
	tokensRepo := psql.NewTokens(db)
	usersService := service.NewUsers(usersRepo, tokensRepo, hasher, []byte("sample secret"))

	handler := rest.NewHandler(recipeService, usersService)

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: handler.InitRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
