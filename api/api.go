package api

import (
	"encoding/json"
	"fmt"
	"github.com/ZaneWithSpoon/fathomBack/config"
	"github.com/ZaneWithSpoon/fathomBack/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var apiServiceInstance *apiService

type apiService struct {
	running bool
	port string
	dbServiceInstance *db.DbService
}


func test(w http.ResponseWriter, r *http.Request) error {
	enableCors(&w)
	out, _ := json.MarshalIndent("tested", "", "  ")
	fmt.Fprintf(w, string(out))
	return nil
}

//this function only exists to support CORS. I hate CORS.
func corsOptions(w http.ResponseWriter, r *http.Request) error {
	enableCors(&w)
	out, _ := json.MarshalIndent("", "", "  ")
	fmt.Fprintf(w, string(out))
	return nil
}

func enableCors(w *http.ResponseWriter) {
	fmt.Println("enabling cors")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:443" will only work if you are accessing the server from your local machine.
	http.Redirect(w, r, "https://{insert_your_url}:443"+r.RequestURI, http.StatusMovedPermanently)
}

// Start
func StartAPI(db *db.DbService) {
	fmt.Println("server running localhost:3001")

	r := mux.NewRouter()

	//This returns some fake data
	r.Handle("/test", Handler{test}).Methods("GET")

	if config.IsDev() {
		fmt.Println("serving at localhost:3001");
		port := "3001" // this is the api port, but the app port is exposed at 3000
		err := http.ListenAndServe(":"+port, r)
		log.Fatal(err)
	} else {
		fmt.Println("serving at https://{insert_your_url}");
		go http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
		port := "443"
		err := http.ListenAndServeTLS(":"+port, "server.crt", "server.key", r)
		log.Fatal(err)
	}

}