package db

const (
<<<<<<< HEAD
<<<<<<< HEAD
	usersTable    = "bshop.user"
	productsTable = "bshop.product"
=======
	usersTable = "bshop.user"
>>>>>>> 35fe851 (made some changes)
=======
	usersTable    = "bshop.user"
	productsTable = "bshop.product"
>>>>>>> 573a019 (finished with authorization, started with password recovery)
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
<<<<<<< HEAD
=======
>>>>>>> 573a019 (finished with authorization, started with password recovery)

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
<<<<<<< HEAD
=======
>>>>>>> 35fe851 (made some changes)
=======
>>>>>>> 573a019 (finished with authorization, started with password recovery)
