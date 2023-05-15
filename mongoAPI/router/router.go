package router

import "github.com/gorilla/mux"

func Router() *mux.Router { //returns the reference of the router this is the very important step

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("POST")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.MarkAsWathced).Methods("PUT")

	return router
}
