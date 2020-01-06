package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Story struct {
	ID         uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Name       string     `gorm:"size:255;not null;" json:"name"`
	Part       uint8      `gorm:"size:255;not null;" json:"part"`
	Audio      string     `gorm:"not null" json:"audio"`
	English    string     `gorm:"not null" json:"en"`
	Vietnamese string     `gorm:"not null" json:"vn"`
	NewWords   *[]NewWord `json:"new_words"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Story) FindAllStories(db *gorm.DB, partQuery string) (*[]Story, error) {
	var err error
	stories := []Story{}

	if partQuery != "" {
		part, err := strconv.ParseInt(partQuery, 10, 8)
		if err != nil {
			return &[]Story{}, err
		}
		db = db.Where("part = ?", int8(part))
	}

	err = db.Find(&stories).Error
	if err != nil {
		return &[]Story{}, err
	}
	return &stories, nil
}

func (s *Story) CreateStory(db *gorm.DB) (*Story, error) {
	var err error
	err = db.Create(&Story{Name: s.Name, Part: s.Part, Audio: s.Audio, English: s.English, Vietnamese: s.Vietnamese}).Error

	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Story) UpdateStory(db *gorm.DB, uid uint32) (*Story, error) {

	var err error

	err = db.Debug().Model(&Story{}).Where("id = ?", uid).Updates(Story{Name: s.Name, Part: s.Part, Audio: s.Audio, English: s.English, Vietnamese: s.Vietnamese}).Error
	if err != nil {
		return &Story{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &User{}, err
	// 	}
	// }
	return s, nil
}

func (s *Story) GetStoryById(db *gorm.DB, id uint32) (*Story, error) {
	var err error
	err = db.Debug().Model(&Story{}).Where("id = ?", id).Take(&s).Error
	if err != nil {
		return &Story{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &User{}, err
	// 	}
	// }
	return s, nil
}

func (s *Story) DeleteStory(db *gorm.DB, id uint32) error {
	var err error
	err = db.Where("id = ?", id).Take(&Story{}).Delete(&Story{}).Error

	if err != nil {
		return err
	}
	return nil
}
