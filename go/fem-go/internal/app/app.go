package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"project-1/internal/api"
	"project-1/internal/store"
	"project-1/migrations"
)

type Application struct {
	Logger 					*log.Logger
	WorkoutHandler 	*api.WorkoutHandler
	DB 							*sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if(err != nil) {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
		DB: pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available!\n")
}
