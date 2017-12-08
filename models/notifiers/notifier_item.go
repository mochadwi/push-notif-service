package notifiers

import "github.com/jinzhu/gorm"

//go:generate goqueryset -in notifier_item.go

// NotifierItem struct represent notifier model. Next line (gen:qs) is needed to autogenerate UserQuerySet.
// gen:qs
type NotifierItem struct {
	gorm.Model
	Name     string
	Provider string
	APIKey   string
}
