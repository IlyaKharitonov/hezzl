package models

import (
	"fmt"
	"time"

	"hezzlTestTask/constants"
)

type Item struct {
	ID          int       `json:"id,omitempty" db:"id"`
	CampaignID  int       `json:"campaignId,omitempty" db:"campaignId"`
	Name        string    `json:"name,omitempty" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	Priority    int       `json:"priority,omitempty" db:"priority"`
	Removed     bool      `json:"removed,omitempty" db:"removed"`
	CreatedAt   time.Time `json:"createdAt,omitempty" db:"createdAt"`
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
