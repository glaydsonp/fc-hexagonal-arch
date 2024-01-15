package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/glaydsonp/go-hexagonal/application"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	// r.Handle("/product", n.With(
	// 	negroni.Wrap(createProduct(service)),
	// )).Methods("POST", "OPTIONS")

	// r.Handle("/product/{id}", n.With(
	// 	negroni.Wrap(updateProduct(service)),
	// )).Methods("PUT", "OPTIONS")

	// r.Handle("/product/{id}", n.With(
	// 	negroni.Wrap(deleteProduct(service)),
	// )).Methods("DELETE", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

// func createProduct(service application.ProductServiceInterface) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		product, err := service.Create(r.FormValue("name"), r.FormValue("price"))
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(product))
// 	})
// }

// func updateProduct(service application.ProductServiceInterface) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		product, err := service.Update(id, r.FormValue("name"), r.FormValue("price"))
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(product))
// 	})
// }

// func deleteProduct(service application.ProductServiceInterface) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		err := service.Delete(id)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 	})
// }
