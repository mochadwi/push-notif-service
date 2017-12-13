package models

import (
	"time"
)

//go:generate goqueryset -in notifier_item.go

// NotifierItem struct represent notifier model. Next line (gen:qs) is needed to autogenerate UserQuerySet.
// gen:qs
type NotifierItem struct {
	ID       uint   `json:"id"` // gorm primary
	Name     string `json:"name"`
	Provider string `json:"provider"`
	APIKey   string `json:"apiKey"`
	// gorm model
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
