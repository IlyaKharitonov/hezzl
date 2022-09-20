package item

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"context"

	"hezzlTestTask/config"
	"hezzlTestTask/internal/models"
	"hezzlTestTask/constants"
	"hezzlTestTask/internal/storage"
)

func Update(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPatch{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var(
		ctx = context.Background()
		item = &models.Item{}
	)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Update#1 \n Error: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(body, item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := item.BeforeQuery(constants.Update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err = storage.NewItemStorage(config.Config.Postgres.DB).Update(ctx, item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Update#2 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Update#2 \n Error: %s\n", err.Error())
		return
	}




}