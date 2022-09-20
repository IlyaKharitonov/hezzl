package item

import (
	"context"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"hezzlTestTask/config"
	"hezzlTestTask/constants"
	"hezzlTestTask/internal/models"
	"hezzlTestTask/internal/storage"
)

func Create(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var (
		item = &models.Item{}
		ctx = context.Background()
	)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Create#1 \n Error: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(body, item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := item.BeforeQuery(constants.Create); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err = storage.NewItemStorage(config.Config.Postgres.DB).Create(ctx, item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Create#2 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Create#2 \n Error: %s\n", err.Error())
		return
	}
}