package main

import (
	"GoApp/controllers"
	"GoApp/settings"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Id int64
	Name string
	Email string
}


func testHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var data Response
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	setting:= settings.Settings{}
	setting.Initialize()
	app:= controllers.App{&setting}
	router := mux.NewRouter()
	router.HandleFunc("/test", testHandler).Methods("POST")
	router.HandleFunc("/saveproduct", app.SaveProduct).Methods("POST")
	router.HandleFunc("/product/{id}", app.GetProduct).Methods("GET")
	router.HandleFunc("/products", app.GetProducts).Methods("GET")
	router.HandleFunc("/updateproduct/{id}", app.UpdateProduct).Methods("PUT")
	router.HandleFunc("/deleteproduct/{id}", app.DeleteProduct).Methods("DELETE")
	fmt.Println("Server listening on port 8000:")
	http.ListenAndServe(":8000", router)
}




