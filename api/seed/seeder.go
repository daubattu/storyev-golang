package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"storyev/api/models"
)

var users = []models.User{
	models.User{
		Name:     "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Name:     "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.User{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
