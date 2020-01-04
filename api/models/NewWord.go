package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type NewWord struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Part      uint8     `gorm:"not null" json:"part"`
	StoryID   uint32    `gorm:"not null" json:"story_id"`
	Word      string    `gorm:"size:255;not null;" json:"word"`
	Type      string    `gorm:"size:255;not null;" json:"type"`
	Spelling  string    `gorm:"not null" json:"spelling"`
	AudioUS   string    `gorm:"not null" json:"audio_us"`
	AudioUK   string    `gorm:"not null" json:"audio_uk"`
	Example   string    `gorm:"not null" json:"example"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (nw *NewWord) FindNewWords(db *gorm.DB, partQuery string, storyIdQuery string) (*[]NewWord, error) {
	var err error
	newWords := []NewWord{}

	if partQuery != "" {
		part, err := strconv.ParseUint(partQuery, 10, 8)
		if err != nil {
			return &[]NewWord{}, err
		}
		db = db.Where("part = ?", uint8(part))
	}

	if storyIdQuery != "" {
		storyId, err := strconv.ParseUint(storyIdQuery, 10, 8)
		if err != nil {
			return &[]NewWord{}, err
		}
		db = db.Where("story_id = ?", uint8(storyId))
	}

	err = db.Limit(100).Find(&newWords).Error
	if err != nil {
		return &[]NewWord{}, err
	}
	return &newWords, nil
}

func (nw *NewWord) CreateNewWord(db *gorm.DB) (*NewWord, error) {
	var err error
	err = db.Create(
		&NewWord{Part: nw.Part, StoryID: nw.StoryID, Word: nw.Word, Type: nw.Type, Spelling: nw.Spelling, AudioUS: nw.AudioUS, AudioUK: nw.AudioUK, Example: nw.Example}).Error

	if err != nil {
		return nil, err
	}
	return nw, nil
}

func (nw *NewWord) UpdateNewWord(db *gorm.DB, uid uint32) (*NewWord, error) {

	var err error

	err = db.Debug().Model(&NewWord{}).Where("id = ?", uid).Updates(NewWord{Part: nw.Part, StoryID: nw.StoryID, Word: nw.Word, Type: nw.Type, Spelling: nw.Spelling, AudioUS: nw.AudioUS, AudioUK: nw.AudioUK, Example: nw.Example}).Error
	if err != nil {
		return &NewWord{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &User{}, err
	// 	}
	// }
	return nw, nil
}

func (nw *NewWord) DeleteNewWord(db *gorm.DB, id uint32) error {
	var err error
	err = db.Where("id = ?", id).Take(&NewWord{}).Delete(&NewWord{}).Error

	if err != nil {
		return err
	}
	return nil
}
