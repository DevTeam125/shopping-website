package product

type Photo struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ArticleID int    `json:"article_id"` // Link to Product ID
	Title     string `json:"title"`
	URL       string `json:"url"`
}
