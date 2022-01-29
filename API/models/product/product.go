package product

import (
	"fmt"

	"github.com/DevTeam125/shopping-website/models"
	"gorm.io/gorm"
)

type Product struct {
	ID          *int      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" binding:"required"`
	Status      string    `json:"status"`
	Rating      string    `json:"rating"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Feature     []Feature `json:"feature" gorm:"-"`
}

func Init() {
	models.DB.AutoMigrate(&Product{}, &Photo{}, &Feature{})

}

func (p *Product) GetAllProductsBrief(pageNum int, pageSize int) ([]*Product, error) {
	//var products []*Product
	products := make([]*Product, 0)
	err := models.DB.Offset(pageNum).Limit(pageSize).Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	for i, v := range products {
		var res []Feature
		models.DB.Where("product_id = ?", v.ID).Find(&res)
		fmt.Println(res)
		products[i].Feature = res
	}

	return products, nil
}
func (p *Product) GetProductByID(id int) (*Product, error) {
	var product *Product
	err := models.DB.Where("ID = ?", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return product, nil
}

func (p *Product) SaveProduct() error {
	err := models.DB.Create(p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
