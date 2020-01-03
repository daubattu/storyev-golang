package controllers

import (
	"net/http"
	"storyev/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.StrictSlash(true)

	s.Router.Handle("/", http.FileServer(http.Dir("./public")))

	s.Router.
		PathPrefix("/css/").
		Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("."+"/css/"))))

	s.Router.
		PathPrefix("/files/").
		Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("."+"/files/"))))

	// Handle all preflight request
	s.Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	s.Router.HandleFunc("/stories", middlewares.SetMiddlewareJSON(s.GetStories)).Methods("GET")
	s.Router.HandleFunc("/stories", middlewares.SetMiddlewareJSON(s.CreateStory)).Methods("POST")
	s.Router.HandleFunc("/stories/{id}", middlewares.SetMiddlewareJSON(s.DeleteStory)).Methods("DELETE")

	s.Router.HandleFunc("/newwords", middlewares.SetMiddlewareJSON(s.GetNewWords)).Methods("GET")
	s.Router.HandleFunc("/newwords", middlewares.SetMiddlewareJSON(s.CreateNewWord)).Methods("POST")
	s.Router.HandleFunc("/newwords/{id}", middlewares.SetMiddlewareJSON(s.DeleteNewWord)).Methods("DELETE")

	s.Router.HandleFunc("/upload", middlewares.SetMiddlewareFormData(s.UploadFile)).Methods("POST")
}
