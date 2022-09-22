package itemService

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"hezzlTestTask/constants"
)

type (
	Item struct {
		ID          int       `json:"id,omitempty" db:"id"`
		CampaignID  int       `json:"campaignId,omitempty" db:"campaignId"`
		Name        string    `json:"name,omitempty" db:"name"`
		Description string    `json:"description,omitempty" db:"description"`
		Priority    int       `json:"priority,omitempty" db:"priority"`
		Removed     bool      `json:"removed,omitempty" db:"removed"`
		CreatedAt   time.Time `json:"createdAt,omitempty" db:"createdAt"`
	}

	db interface {
		Create(ctx context.Context, item *Item) (*Item, error)
		GetList(ctx context.Context) ([]*Item, error)
		Update(ctx context.Context, item *Item) (*Item, error)
		Delete(ctx context.Context, item *Item) (*Item, error)
	}

	statistic interface {
		Insert([]string) error
	}

	cache interface {
		GetList(key string) ([]*Item, error)
		Load(data interface{}, key string, expires time.Duration) error
		Delete(key string) error
	}

	broker interface {
		Send(log string)
		Read()
	}

	itemService struct {
		cache     cache
		db        db
		broker    broker
		statistic statistic
	}
)

func NewItem(db postgresDB, cache cache, broker broker, statistic statistic) *itemService {
	return &itemService{cache: cache, db: db, broker: broker, statistic: statistic}
}

func (is *itemService) Create(ctx context.Context, item *Item) (*Item, error) {
	item, err := is.db.Create(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)Create #1 \n Error:%s \n", err.Error())
	}

	return item, nil
}

func (is *itemService) GetList(ctx context.Context) ([]*Item, error) {
	keyRedis := "items"
	items, err := is.cache.GetList(keyRedis)
	switch err {
	case nil:
		return items, nil
	case redis.Nil:
	default:
		return nil, fmt.Errorf("(is *itemService)GetList #1 \n Error:%s \n", err.Error())

	}

	items, err = is.db.GetList(ctx)
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)GetList #2 \n Error:%s \n", err.Error())
	}

	err = is.cache.Load(items, keyRedis, time.Minute)
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)GetList #3 \n Error:%s \n", err.Error())
	}

	return items, nil
}

func (is *itemService) Update(ctx context.Context, item *Item) (*Item, error) {
	item, err := is.db.Update(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)Update #1 \n Error:%s \n", err.Error())
	}

	err = is.cache.Delete("items")
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)Update #2 \n Error:%s \n", err.Error())
	}

	return item, nil
}

func (is *itemService) Delete(ctx context.Context, item *Item) (*Item, error) {
	item, err := is.db.Delete(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)Delete #1 \n Error:%s \n", err.Error())
	}
	err = is.cache.Delete("items")
	if err != nil {
		return nil, fmt.Errorf("(is *itemService)Delete #2 \n Error:%s \n", err.Error())
	}

	return item, nil
}

func (i *Item) BeforeQuery(operationType int) error {
	switch operationType {
	case constants.Create:
		if i.CampaignID == 0 {
			return fmt.Errorf("models.(i *Item)BeforeQuery #1 \n Error: no CampaignID in item")
		}
		if i.Name == "" {
			return fmt.Errorf("models.(i *Item)BeforeQuery #2 \n Error: no Name in item")
		}
	case constants.Delete:
		if i.CampaignID == 0 {
			return fmt.Errorf("models.(i *Item)BeforeQuery #3 \n Error: no CampaignID in item")
		}
		if i.ID == 0 {
			return fmt.Errorf("models.(i *Item)BeforeQuery #4 \n Error: no ID in item")
		}

	case constants.Update:
		if i.CampaignID == 0 {
			return fmt.Errorf("models.(i *Item)BeforeQuery #5 \n Error: no CampaignID in item")
		}
		if i.ID == 0 {
			return fmt.Errorf("models.(i *Item)BeforeQuery #6 \n Error: no ID in item")
		}
		if i.Name == "" {
			return fmt.Errorf("models.(i *Item)BeforeQuery #7 \n Error: no ID in item")
		}

	case constants.Get:
	}

	return nil
}
