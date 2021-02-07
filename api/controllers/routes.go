package controllers

// InitRoutes ...
func (server *Server) InitRoutes() {

	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/savestat", server.SaveStatHandler).Methods("POST")
	server.Router.HandleFunc("/getstat", server.GetStatHandler(server.DB)).Methods("GET")
	server.Router.HandleFunc("/dropstat", server.DropStatHandler)
}
