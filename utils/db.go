package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/nobackend-repo/push-notif-service/models"
)

// Manager interface
type Manager interface {
	AddNotifier(notifier *models.NotifierItem) error
	ShowNotifier(notifier *[]models.NotifierItem) error
	// Add other methods
}

type manager struct {
	db *gorm.DB
}

// Mgr to manage database
var Mgr Manager

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=mochadwi dbname=nobackend_db sslmode=disable password=")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	db.LogMode(true)
	// defer db.Close()

	Mgr = &manager{db: db}

	db.Debug().AutoMigrate(&models.NotifierItem{})
}

func (mgr *manager) AddNotifier(notifier *models.NotifierItem) (err error) {
	mgr.db.Debug().AutoMigrate(&models.NotifierItem{})
	notifier.Create(mgr.db)
	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("[notifier] create: ")
		fmt.Println(notifier)
	}
	return
}

func (mgr *manager) ShowNotifier(notifier *[]models.NotifierItem) (err error) {
	mgr.db.Debug().AutoMigrate(&models.NotifierItem{}) // mgr.db.AutoMigrate(&models.NotifierItem{})
	tempNotifier := []models.NotifierItem{}
	if err := models.NewNotifierItemQuerySet(mgr.db).All(&tempNotifier); err != nil {
		fmt.Print("[notifier] query_all: ")
		fmt.Println(err)
		return err
	}
	notifier = &tempNotifier
	fmt.Print("[notifier] result: ")
	fmt.Println(notifier)
	return
}
