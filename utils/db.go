package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	models "gitlab.com/nobackend-repo/push-notif-service/models"
)

// Manager interface
type Manager interface {
	AddNotifier(notifier *models.NotifierItem) error
	ShowNotifier(notifier []models.NotifierItem) error
	// Add other methods
}

type manager struct {
	db *gorm.DB
}

// Mgr to manage database
var Mgr Manager

func init() {
	db, err := gorm.Open("sqlite3", "./notifier.db")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	// defer db.Close()

	Mgr = &manager{db: db}
}

func (mgr *manager) AddNotifier(notifier *models.NotifierItem) (err error) {
	mgr.db.AutoMigrate(&notifier)
	mgr.db.Create(notifier)
	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("notifier: ")
		fmt.Println(notifier)
	}
	return
}

func (mgr *manager) ShowNotifier(notifier []models.NotifierItem) (err error) {
	mgr.db.AutoMigrate(&notifier)
	if errs := mgr.db.Find(&notifier).Error; errs != nil {
		return errs
	}
	return
}
