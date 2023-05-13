package db

const (
	usersTable    = "bshop.user"
	productsTable = "bshop.product"
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

type Product struct {
	Id           uint32  `json:"id"`
	Category     string  `json:"category"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	Price        uint32  `json:"price"`
	Available    bool    `json:"available"`
	Rating       float64 `json:"rating"`
	RatingAmount uint64  `json:"rating_amount"`
	PathToImage  string  `json:"path_to_image"`
}

func NewProduct(id uint32, category string, name string, brand string,
	price uint32, available bool, rating float64, pathToImage string) Product {
	return Product{
		Id:          id,
		Category:    category,
		Name:        name,
		Brand:       brand,
		Price:       price,
		Available:   available,
		Rating:      rating,
		PathToImage: pathToImage,
	}
}
