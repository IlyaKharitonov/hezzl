package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"

	"hezzlTestTask/internal/cache"
	"hezzlTestTask/internal/models"
)

type itemStorage struct {
	*pgx.Conn
}

func NewItemStorage(db *pgx.Conn)*itemStorage{
	return &itemStorage{db}
}

func(i *itemStorage)GetList(ctx context.Context)([]*models.Item,  error){
	var(
		items = make([]*models.Item, 0)
		redisStruct = cache.RedisStruct{Key: "items", Expires: time.Minute}
	)
	//
	err := redisStruct.Get(&items)
	switch err {
	case nil:
		return items, nil
	case redis.Nil:
	default:
		return nil,fmt.Errorf("storage.(i *itemStorage)GetList #1 \n Error:%s \n", err.Error() )
	}

	rows, err := i.Query(ctx, `
SELECT 
id,
campaign_id, 
name,
COALESCE(description,'') AS description,
priority,
removed,
created_at
FROM items`)
	if err != nil{
		return nil,fmt.Errorf("storage.(i *itemStorage)GetList #2 \n Error:%s \n", err.Error() )
	}

	for rows.Next(){
		item := &models.Item{}
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Name,
			&item.Description,
			&item.Priority,
			&item.Removed,
			&item.CreatedAt);
		err != nil{
			return nil,fmt.Errorf("storage.(i *itemStorage)GetList #3 \n Error:%s \n", err.Error() )
		}
		items = append(items, item)
	}

	err = redisStruct.Load(items)
	if err != nil{
		return nil,fmt.Errorf("storage.(i *itemStorage)GetList #4 \n Error:%s \n", err.Error() )
	}

	return items,nil
}

func(i *itemStorage)Update(ctx context.Context, item *models.Item)(*models.Item,error){
	args := make([]interface{},0)

	updateStr:= "UPDATE items SET name = $1 "
	args = append(args, item.Name)

	if item.Description != ""{
		updateStr = updateStr + "description = $2 WHERE id = $3 AND campaign_id = $4"
		args = append(args, item.Description, item.ID, item.CampaignID)
	} else {
		updateStr = updateStr + "WHERE id = $2 AND campaign_id = $3 COMMIT"
		args = append(args, item.ID, item.CampaignID)
	}

	tx, err := i.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
			return nil, fmt.Errorf("storage.(i *itemStorage)Update #1 \n Error:%s \n", err.Error() )
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	_,err = i.Exec(ctx, updateStr, args...)
	if err != nil {
		return nil, fmt.Errorf("storage.(i *itemStorage)Update #2 \n Error:%s \n", err.Error() )
	}

	rows, err := i.Query(ctx, `
SELECT 
id,
campaign_id, 
name,
COALESCE(description,'') AS description,
priority,
removed,
created_at
FROM items
WHERE id = $1`,
		item.ID)
	if err != nil{
		return nil,fmt.Errorf("storage.(i *itemStorage)Update #4 \n Error:%s \n", err.Error() )
	}

	for rows.Next(){
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Name,
			&item.Description,
			&item.Priority,
			&item.Removed,
			&item.CreatedAt);
			err != nil{
			return nil,fmt.Errorf("storage.(i *itemStorage)Update #5 \n Error:%s \n", err.Error())
		}
	}

	redisStruct := cache.RedisStruct{Key: "items"}
	if err = redisStruct.Delete(); err != nil {
		return nil,fmt.Errorf("storage.(i *itemStorage)Update #6 \n Error:%s \n", err.Error())
	}

	return item, nil

}

func(i *itemStorage)Create(ctx context.Context, item *models.Item)(*models.Item, error){

	row,err := i.Query(ctx, `
INSERT INTO items 
(name, campaign_id) 
VALUES ($1, $2)
RETURNING id`,
item.Name, item.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("storage.(i *itemStorage)Create #1 \n Error:%s \n", err.Error() )
	}

	var id int
	for row.Next() {
		if err = row.Scan(&id)
			err != nil {
			return nil, fmt.Errorf("storage.(i *itemStorage)Create #2 \n Error:%s \n", err.Error())
		}
	}
	//COALESCE(description,'') AS
	rows, err := i.Query(ctx, `
SELECT 
id,
campaign_id, 
name,
COALESCE(description,'') AS description,
priority,
removed,
created_at
FROM items
WHERE id = $1`,
id)
	if err != nil{
		return nil,fmt.Errorf("storage.(i *itemStorage)Create #3 \n Error:%s \n", err.Error() )
	}

	for rows.Next(){
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Name,
			&item.Description,
			&item.Priority,
			&item.Removed,
			&item.CreatedAt);
			err != nil{
			return nil,fmt.Errorf("storage.(i *itemStorage)Create #4 \n Error:%s \n", err.Error())
		}
	}

	return item, nil
}

func(i *itemStorage)Delete(ctx context.Context, item *models.Item)(*models.Item, error){

	tx, err := i.BeginTx(ctx,pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return nil, fmt.Errorf("storage.(i *itemStorage)Update #1 \n Error:%s \n", err.Error() )
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	_,err = i.Exec(ctx, `
UPDATE items SET removed = true 
WHERE id = $1 AND campaign_id = $2`,
		item.ID,
		item.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("storage.(i *itemStorage)Delete #1 \n Error:%s \n", err.Error() )
	}

	rows, err := i.Query(ctx, `
SELECT 
id,
campaign_id,
removed
FROM items
WHERE id = $1 AND campaign_id = $2`,
		item.ID, item.CampaignID)
	if err != nil{
		return nil,fmt.Errorf("storage.(i *itemStorage)Delete #2 \n Error:%s \n", err.Error() )
	}

	item = &models.Item{}
	for rows.Next(){
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Removed);
			err != nil{
			return nil,fmt.Errorf("storage.(i *itemStorage)Delete #3 \n Error:%s \n", err.Error())
		}
	}

	return item, nil
}