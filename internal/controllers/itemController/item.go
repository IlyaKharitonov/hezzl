package itemController

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"hezzlTestTask/constants"
	"hezzlTestTask/internal/services/itemService"
)

type (
	ItemService interface {
		Create(ctx context.Context, item *itemService.Item) (*itemService.Item, error)
		GetList(ctx context.Context) ([]*itemService.Item, error)
		Update(ctx context.Context, item *itemService.Item) (*itemService.Item, error)
		Delete(ctx context.Context, item *itemService.Item) (*itemService.Item, error)
	}

	controller struct {
		is ItemService
	}
)

func NewController(is ItemService) *controller {
	return &controller{is: is}
}

func (c *controller) Create(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var item = &itemService.Item{}

	err := unmarshalReq(item, req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Create#1 \n Error: %s\n", err.Error())
		return
	}

	if err := item.BeforeQuery(constants.Create); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err = c.is.Create(req.Context(), item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Create#2 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Create#3 \n Error: %s\n", err.Error())
		return
	}

}

func (c *controller) GetList(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	items, err := c.is.GetList(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.GetList#1 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.GetList#2 \n Error: %s\n", err.Error())
		return
	}

}

func (c *controller) Update(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPatch {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var item = &itemService.Item{}

	err := unmarshalReq(item, req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Update#1 \n Error: %s\n", err.Error())
		return
	}

	if err := item.BeforeQuery(constants.Update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err = c.is.Update(req.Context(), item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Update#2 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Update#3 \n Error: %s\n", err.Error())
		return
	}

}

func (c *controller) Delete(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var (
		params        = req.URL.Query()
		id, _         = strconv.Atoi(params.Get("id"))
		campaignId, _ = strconv.Atoi(params.Get("campaignId"))
		ctx           = context.Background()
		item          = &itemService.Item{ID: id, CampaignID: campaignId}
	)

	if err := item.BeforeQuery(constants.Delete); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := c.is.Delete(ctx, item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Delete#1 \n Error: %s\n", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("item.Delete#2 \n Error: %s\n", err.Error())
		return
	}

}

func unmarshalReq(model interface{}, body io.Reader) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &model)
}
