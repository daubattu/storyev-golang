package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:255;not null;" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *User) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Email = html.EscapeString(strings.TrimSpace(p.Email))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *User) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Email == "" {
		return errors.New("Required Email")
	}
	if p.Password == "" {
		return errors.New("Required Password")
	}
	return nil
}

func (p *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Create(&p).Error
	if err != nil {
		return &User{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &User{}, err
	// 	}
	// }
	return p, nil
}

func (p *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	Users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&Users).Error
	if err != nil {
		return &[]User{}, err
	}
	// if len(Users) > 0 {
	// 	for i, _ := range Users {
	// 		err := db.Debug().Model(&User{}).Where("id = ?", Users[i].AuthorID).Take(&Users[i].Author).Error
	// 		if err != nil {
	// 			return &[]User{}, err
	// 		}
	// 	}
	// }
	return &Users, nil
}

func (p *User) FindUserByID(db *gorm.DB, pid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &User{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &User{}, err
	// 	}
	// }
	return p, nil
}

func (p *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {

	var err error

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Updates(User{Name: p.Name, Email: p.Email, Password: p.Password, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &User{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &User{}, err
	// 	}
	// }
	return p, nil
}

func (p *User) DeleteUser(db *gorm.DB, pid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", pid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("User not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
