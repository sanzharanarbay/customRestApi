package controllers

import (
	"GoApp/models"
	"GoApp/settings"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type App struct {
	Stg *settings.Settings
}

func (app *App) SaveProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	product:= &models.Product{}
	err := decoder.Decode(product)
	if err != nil {
		panic(err)
	}

	validation:= product.Validate()

	if(validation != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Wrong validation"})
		w.Write(resp)
	}
	saveStatus,errors := product.Save(app.Stg.Db)
	if(errors != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Not Saved in DB!"})
		w.Write(resp)
	}
	json.NewEncoder(w).Encode(saveStatus)
	return
}

func (app *App) GetProduct(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	id:=params["id"]
	product:= &models.Product{}
	int1, _ := strconv.Atoi(id)
	item, _ := product.GetProductById(int1, app.Stg.Db)
	json.NewEncoder(w).Encode(item)
	return
}

func (app *App) GetProducts(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	product:= &models.Product{}
	items, _ := product.GetProducts(app.Stg.Db)
	json.NewEncoder(w).Encode(items)
	return
}

func (app *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	id:=params["id"]
	int1, _ := strconv.Atoi(id)
	decoder := json.NewDecoder(r.Body)
	product:= &models.Product{}
	err := decoder.Decode(product)
	if err != nil {
		panic(err)
	}

	validation:= product.Validate()

	if(validation != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Wrong validation"})
		w.Write(resp)
	}
	_,errors := product.UpdateProduct(int1,app.Stg.Db)
	if(errors != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Not Saved in DB!"})
		w.Write(resp)
	}else{
		json.NewEncoder(w).Encode(map[string]interface{}{"status":"success", "message":"Product was successfully updated!"})
	}
	return
}

func (app *App) DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	id:=params["id"]
	int1, _ := strconv.Atoi(id)
	product:= &models.Product{}
	deleteStatus := product.DeleteProduct(int1, app.Stg.Db)
	if(deleteStatus != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Not Deleted in DB!"})
		w.Write(resp)
	}else{
		resp, _ := json.Marshal(map[string]interface{}{"status":"succcess", "message":"Product was successfully deleted!"})
		w.Write(resp)
	}
}

