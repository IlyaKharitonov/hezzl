package item

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"context"
	"strconv"

	"hezzlTestTask/config"
	"hezzlTestTask/constants"
	"hezzlTestTask/internal/models"
	"hezzlTestTask/internal/storage"
)

func Delete(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodDelete{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var (
		params= req.URL.Query()
		id,_ = strconv.Atoi(params.Get("id"))
		campaignId,_ = strconv.Atoi(params.Get("campaignId"))
		ctx = context.Background()
		item = &models.Item{ID: id, CampaignID: campaignId}
	)

	if err := item.BeforeQuery(constants.Delete); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := storage.NewItemStorage(config.Config.Postgres.DB).Delete(ctx, item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Delete#2 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Delete#2 \n Error: %s\n", err.Error())
		return
	}


}