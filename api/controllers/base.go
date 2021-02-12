package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //
)

// Server ...
type Server struct {
	DB     *sqlx.DB
	Router *mux.Router
}

// Init is a start function of a server
func (server *Server) Init(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = sqlx.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		fmt.Print()
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", Dbdriver)
	}

	server.Router = mux.NewRouter()

	server.InitRoutes()

}

// InitRoutes ...
func (server *Server) InitRoutes() {

	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/savestat", server.SaveStatHandler).Methods("POST")
	server.Router.HandleFunc("/getstat", server.GetStatHandler(server.DB)).Methods("GET")
	server.Router.HandleFunc("/dropstat", server.DropStatHandler)
}

// Run is a launch function
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
