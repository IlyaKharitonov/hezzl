package storage

import (
	"database/sql"
)

type campaignStorage struct {
	db sql.DB
	tx sql.Tx
}

func NewCampaignStorage(db sql.DB, tx sql.Tx)*campaignStorage{
	return &campaignStorage{db, tx}
}

func(i *campaignStorage)Get(id int)error{
	return nil
}

func(i *campaignStorage)Update()error{
	return nil
}

func(i *campaignStorage)Create()error{
	return nil
}

func(i *campaignStorage)Delete()error{
	return nil
}
