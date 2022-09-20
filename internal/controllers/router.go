package controllers

import (
	"net/http"

	"hezzlTestTask/internal/controllers/item"
)

func HandlersRegister(){
	http.HandleFunc("/item/create", item.Create)
	http.HandleFunc("/item/list", item.GetList)
	http.HandleFunc("/item/remove", item.Delete)
	http.HandleFunc("/item/update", item.Update)
}
