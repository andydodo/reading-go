package postgres

import (
	"errors"
	"fmt"
	"ginexamples"

	"github.com/jinzhu/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func newLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

// Store creates a link record in the table
func (l *LinkRepository) Store(link *ginexamples.Link) error {
	return l.db.Create(link).Error
}

func (l *LinkRepository) Find(id string) (*ginexamples.Link, error) {
	var link ginexamples.Link

	db := l.db.Where("id = ?", id)
	err := first(db, &link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (l *LinkRepository) FindByUserName(userName string) (*ginexamples.Link, error) {
	if userName == "" {
		return &ginexamples.Link{}, errors.New("not found")
	}
	return l.findBy("username", userName)
}

func (l *LinkRepository) findBy(key, value string) (*ginexamples.Link, error) {
	link := ginexamples.Link{}

	db := l.db.Where(fmt.Sprintf("%s = ?", key), value)
	err := first(db, &link)

	return &link, err
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("resource not found")
	}
	return err
}

func (l *LinkRepository) Update(link *ginexamples.Link) error {
	return l.db.Save(link).Error
}

func (l *LinkRepository) Delete(id string) (*ginexamples.Link, error) {
	var link ginexamples.Link

	db := l.db.Where("id = ?", id).Delete(&link)
	err := first(db, &link)
	if err != nil {
		return nil, err
	}
	return link, nil
}
