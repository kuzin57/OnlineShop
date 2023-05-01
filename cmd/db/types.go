package db

const (
<<<<<<< HEAD
	usersTable    = "bshop.user"
	productsTable = "bshop.product"
=======
	usersTable = "bshop.user"
>>>>>>> 35fe851 (made some changes)
)

type User struct {
	Id          uint32 `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Birthday    string `json:"birthday"`
	Password    string `json:"password"`
}
<<<<<<< HEAD

type Product struct {
	Id           uint32  `json:"id"`
	Category     string  `json:"category"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	Price        uint32  `json:"price"`
	Available    bool    `json:"available"`
	Rating       float64 `json:"rating"`
	RatingAmount uint64  `json:"rating_amount"`
}

func NewProduct(category string, name string, brand string,
	price uint32, available bool, rating float64) Product {
	return Product{
		Category:  category,
		Name:      name,
		Brand:     brand,
		Price:     price,
		Available: available,
		Rating:    rating,
	}
}
=======
>>>>>>> 35fe851 (made some changes)
