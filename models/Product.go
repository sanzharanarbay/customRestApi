package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Category    int    `gorm:"not null"                 json:"category"`
	Price int `gorm:"not null"                 json:"price"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("Name is required")
	}
	if p.Price < 0 {
		return errors.New("Price  is invalid")
	}

	if p.Category < 0 {
		return errors.New("Category  is invalid")
	}

	return nil
}

func (p *Product) Save(db *gorm.DB) (*Product, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) GetProducts(db *gorm.DB) (*[]Product, error) {
	var products []Product
	if err := db.Debug().Table("products").Find(&products).Error; err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) GetProductById(id int, db *gorm.DB) (*Product, error) {
	product := &Product{}
	if err := db.Debug().Table("products").Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) UpdateProduct(id int, db *gorm.DB) (*Product, error) {
	if err := db.Debug().Table("products").Where("id = ?", id).Updates(Product{
		Name:        p.Name,
		Price: p.Price,
		Category:    p.Category}).Error; err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteProduct(id int, db *gorm.DB) error {
	if err := db.Debug().Table("products").Where("id = ?", id).Delete(&Product{}).Error; err != nil {
		return err
	}
	return nil
}
