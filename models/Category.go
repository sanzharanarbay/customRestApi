package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Category struct{
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

func (c *Category) Validate() error  {
	if c.Name == ""{
		return errors.New("Name is required")
	}
	return nil
}

func (c *Category) Save(db *gorm.DB) (*Category, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) GetCategories(db *gorm.DB) (*[]Category, error) {
	cats := []Category{}
	if err := db.Debug().Table("categories").Preload("Products").Find(&cats).Error; err != nil {
		return &[]Category{}, err
	}
	return &cats, nil
}

func (c *Category) GetCategoryById(id int, db *gorm.DB) (*Category, error) {
	cat := &Category{}
	if err := db.Debug().Table("categories").Where("id = ?", id).Preload("Products").First(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (c *Category) UpdateCategory(id int, db *gorm.DB) (*Category, error) {
	if err := db.Debug().Table("categories").Where("id = ?", id).Updates(Category{
		Name:        c.Name,
	}).Error; err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) DeleteCategory(id int, db *gorm.DB) error {
	if err := db.Debug().Table("categories").Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		return err
	}
	return nil
}