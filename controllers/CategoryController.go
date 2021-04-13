package controllers

import (
	"GoApp/models"
	"GoApp/settings"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CategoryController struct{
	Set *settings.Settings
}

func (c *CategoryController) SaveCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	cat:= &models.Category{}
	err := decoder.Decode(cat)
	if err != nil {
		panic(err)
	}

	validation:= cat.Validate()

	if(validation != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Wrong validation"})
		w.Write(resp)
	}
	saveStatus,errors := cat.Save(c.Set.Db)
	if(errors != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Not Saved in DB!"})
		w.Write(resp)
	}
	json.NewEncoder(w).Encode(saveStatus)
	return
}

func (c *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	id:=params["id"]
	cat:= &models.Category{}
	int1, _ := strconv.Atoi(id)
	item, _ := cat.GetCategoryById(int1, c.Set.Db)
	json.NewEncoder(w).Encode(item)
	return
}

func (c *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	cat:= &models.Category{}
	items, _ := cat.GetCategories(c.Set.Db)
	json.NewEncoder(w).Encode(items)
	return
}

func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	id:=params["id"]
	int1, _ := strconv.Atoi(id)
	decoder := json.NewDecoder(r.Body)
	cat:= &models.Category{}
	err := decoder.Decode(cat)
	if err != nil {
		panic(err)
	}

	validation:= cat.Validate()

	if(validation != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Wrong validation"})
		w.Write(resp)
	}
	_,errors := cat.UpdateCategory(int1,c.Set.Db)
	if(errors != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Not Saved in DB!"})
		w.Write(resp)
	}else{
		json.NewEncoder(w).Encode(map[string]interface{}{"status":"success", "message":"Product was successfully updated!"})
	}
	return
}

func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	id:=params["id"]
	int1, _ := strconv.Atoi(id)
	cat:= &models.Category{}
	deleteStatus := cat.DeleteCategory(int1, c.Set.Db)
	if(deleteStatus != nil){
		resp, _ := json.Marshal(map[string]interface{}{"status":"error", "message":"Not Deleted in DB!"})
		w.Write(resp)
	}else{
		resp, _ := json.Marshal(map[string]interface{}{"status":"succcess", "message":"Product was successfully deleted!"})
		w.Write(resp)
	}
}
