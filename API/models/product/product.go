package product

import (
	"github.com/DevTeam125/shopping-website/models"
	"gorm.io/gorm"
)

type Product struct {
	ID          *int      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" binding:"required"`
	Status      string    `json:"status"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Feature     []Feature `json:"feature" gorm:"-"`
	Photo       []Photo   `json:"photo" gorm:"-"`
}

func Init() {
	models.DB.AutoMigrate(&Product{}, &Photo{}, &Feature{})

}

func (p *Product) GetAllProductsBrief(pageNum int, pageSize int) (*[]Product, error) {
	//var products []*Product
	/*products := make([]*Product, 0)
	err := models.DB.Offset(pageNum).Limit(pageSize).Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	for i, v := range products {
		var res []Photo
		models.DB.Where("product_id = ?", v.ID).Find(&res)
		fmt.Println(res)
		products[i].Photo = res
	}

	return products, nil/*/
	type result struct {
		ID          int
		ID2         int
		Name        string
		Price       int
		Status      string
		Rating      int
		Description string
		ProductID   int
		Title       string
		URL         string
	}
	res := []result{}
	models.DB.Model(&Product{}).Select("products.id, products.name, products.price, products.status, products.rating, products.description, photos.product_id, photos.title, photos.url, photos.id as id2").Joins("inner join photos on photos.product_id = products.id").
		Scan(&res)
	//fmt.Printf("%+v\n", res)

	out := []Product{}
	var done bool
	for _, v := range res {
		done = false
		photo := Photo{
			Title: v.Title,
			URL:   v.URL,
			ID:    v.ID2,
		}

		for i, v2 := range out {
			if v.ID == *v2.ID {

				out[i].Photo = append(out[i].Photo, photo)
				done = true
			}

		}

		if !done {
			done = false
			vvv := v.ID
			out = append(out, Product{ID: &vvv, Name: v.Name, Price: v.Price, Status: v.Status, Rating: v.Rating, Description: v.Description, Photo: []Photo{photo}})
		}

	}

	return &out, nil

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
