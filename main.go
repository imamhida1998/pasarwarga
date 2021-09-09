package main

import (
	"log"
	"net/http"
	"pasarwarga/api"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome server Pasarwarga"))
	}).Methods("GET")

	APIArtikel := api.ArtikelAPI{}
	router.HandleFunc("/cekartikel", APIArtikel.ListArtikel).Methods("GET")
	//	router.HandleFunc("/branch/id", branchApi.GetBranchByID).Methods("GET")
	router.HandleFunc("/postartikel", APIArtikel.CreateArtikel).Methods("POST")
	router.HandleFunc("/branch", APIArtikel.UpdateArtikel).Methods("PUT")
	//router.HandleFunc("/branch", APIArtikel.DeleteArtikelByArtikel).Methods("DELETE")

	KategoriApi := api.KategoriApi{}
	router.HandleFunc("/kategori", KategoriApi.CreateKategori).Methods("POST")

	//standard
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Server listening to 127.0.0.1:8000")
	err := http.ListenAndServe("127.0.0.1:8000", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
	}

}
