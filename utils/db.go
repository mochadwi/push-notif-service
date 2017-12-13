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
	ShowAllNotifier(notifier *[]models.NotifierItem) error
	ShowNotifier(name string, notifier *models.NotifierItem) error
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

	var tempNotifier models.NotifierItem
	models.NewNotifierItemQuerySet(mgr.db).NameEq(notifier.Name).One(&tempNotifier)

	if tempNotifier.Name != notifier.Name {
		// Create
		notifier.Create(mgr.db)

		if errs := mgr.db.GetErrors(); len(errs) > 0 {
			err = errs[0]
			fmt.Print("[error] addnotifier - create query: ")
			fmt.Println(err)
			return err
		} // end Create

		return
	}

	fmt.Print("[error] addnotifier - duplicate found: ")
	fmt.Println("duplicate entry")
	return fmt.Errorf("%s", "duplicate entry")
}

func (mgr *manager) ShowAllNotifier(notifier *[]models.NotifierItem) (err error) {
	if err := models.NewNotifierItemQuerySet(mgr.db).All(notifier); err != nil {
		fmt.Print("[error] showallnotifier: ")
		fmt.Println(err)
		return err
	}
	return
}

func (mgr *manager) ShowNotifier(name string, notifier *models.NotifierItem) (err error) {
	if err := models.NewNotifierItemQuerySet(mgr.db).NameEq(name).One(notifier); err != nil {
		fmt.Print("[error] shownotifier: ")
		fmt.Println(err)
		return err
	}

	// fmt.Print("[success] shownotifier: ")
	// fmt.Println(err)
	return
}
