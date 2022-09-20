package item

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"hezzlTestTask/config"
	"hezzlTestTask/internal/storage"
)

func GetList(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodGet{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx := context.Background()

	items, err := storage.NewItemStorage(config.Config.Postgres.DB).GetList(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Get#1 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Get#2 \n Error: %s\n", err.Error())
		return
	}
}