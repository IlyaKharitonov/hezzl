package itemService

import (
	//"github.com/jackc/pgx/v4"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type postgresDB struct {
	*pgx.Conn
}

func NewPostgresDB(conn *pgx.Conn) postgresDB {
	return postgresDB{conn}
}

func (pdb postgresDB) Create(ctx context.Context, item *Item) (*Item, error) {

	row, err := pdb.Query(ctx, `
INSERT INTO items 
(name, campaign_id) 
VALUES ($1, $2)
RETURNING id`,
		item.Name, item.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Create #1 \n Error:%s \n", err.Error())
	}

	var id int
	for row.Next() {
		if err = row.Scan(&id); err != nil {
			return nil, fmt.Errorf("(pdb *postgresDB)Create #2 \n Error:%s \n", err.Error())
		}
	}
	//COALESCE(description,'') AS
	rows, err := pdb.Query(ctx, `
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
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Create #3 \n Error:%s \n", err.Error())
	}

	for rows.Next() {
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Name,
			&item.Description,
			&item.Priority,
			&item.Removed,
			&item.CreatedAt); err != nil {
			return nil, fmt.Errorf("(pdb *postgresDB)Create #4 \n Error:%s \n", err.Error())
		}
	}

	return item, nil
}

func (pdb postgresDB) GetList(ctx context.Context) ([]*Item, error) {
	var items = make([]*Item, 0)

	rows, err := pdb.Query(ctx, `
SELECT 
id,
campaign_id, 
name,
COALESCE(description,'') AS description,
priority,
removed,
created_at
FROM items`)
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)GetList #1 \n Error:%s \n", err.Error())
	}

	for rows.Next() {
		item := &Item{}
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Name,
			&item.Description,
			&item.Priority,
			&item.Removed,
			&item.CreatedAt); err != nil {
			return nil, fmt.Errorf("(pdb *postgresDB)GetList #2 \n Error:%s \n", err.Error())
		}
		items = append(items, item)
	}
	return items, nil

}

func (pdb postgresDB) Update(ctx context.Context, item *Item) (*Item, error) {
	args := make([]interface{}, 0)

	updateStr := "UPDATE items SET name = $1 "
	args = append(args, item.Name)

	if item.Description != "" {
		updateStr = updateStr + "description = $2 WHERE id = $3 AND campaign_id = $4"
		args = append(args, item.Description, item.ID, item.CampaignID)
	} else {
		updateStr = updateStr + "WHERE id = $2 AND campaign_id = $3"
		args = append(args, item.ID, item.CampaignID)
	}

	tx, err := pdb.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Update #1 \n Error:%s \n", err.Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	_, err = pdb.Exec(ctx, updateStr, args...)
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Update #2 \n Error:%s \n", err.Error())
	}

	rows, err := pdb.Query(ctx, `
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
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Update #3 \n Error:%s \n", err.Error())
	}

	for rows.Next() {
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Name,
			&item.Description,
			&item.Priority,
			&item.Removed,
			&item.CreatedAt); err != nil {
			return nil, fmt.Errorf("(pdb *postgresDB)Update #4 \n Error:%s \n", err.Error())
		}
	}

	tx.Commit(ctx)
	return item, nil
}

func (pdb postgresDB) Delete(ctx context.Context, item *Item) (*Item, error) {
	tx, err := pdb.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Delete #1 \n Error:%s \n", err.Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	_, err = pdb.Exec(ctx, `
UPDATE items SET removed = true 
WHERE id = $1 AND campaign_id = $2`,
		item.ID,
		item.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Delete#2 \n Error:%s \n", err.Error())
	}

	rows, err := pdb.Query(ctx, `
SELECT 
id,
campaign_id,
removed
FROM items
WHERE id = $1 AND campaign_id = $2`,
		item.ID, item.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("(pdb *postgresDB)Delete #3 \n Error:%s \n", err.Error())
	}

	item = &Item{}
	for rows.Next() {
		if err = rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Removed); err != nil {
			return nil, fmt.Errorf("(pdb *postgresDB)Delete #4 \n Error:%s \n", err.Error())
		}
	}

	return item, nil
}
