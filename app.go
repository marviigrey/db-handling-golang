package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v", DbUser, DbPassword, DbName)
	 var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("couldnt open database.")
	}
	 app.Router = mux.NewRouter().StrictSlash(true)
	 app.handleRoutes()
	return nil
}

func (app *App) Run(address string) {
log.Fatal(http.ListenAndServe(address, app.Router))
}

func (app *App)handleRoutes() {

	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")

}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}
func sendError(w http.ResponseWriter, statusCode int, err string) {
	error_message := map[string]string{"error": err}
	sendResponse(w, statusCode, error_message)

}
func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}