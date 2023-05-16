package db

const (
	usersTable    = "bshop.user"
	productsTable = "bshop.product"
	ordersTable   = "bshop.purchase"
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
	Amount       string  `json:"amount"`
}

type Order struct {
	Id           uint32    `json:"id"`
	DeliveryDate string    `json:"delivery_date"`
	TotalSum     int       `json:"total_sum"`
	City         string    `json:"city"`
	HouseNumber  uint16    `json:"house_number"`
	FlatNumber   uint16    `json:"flat_number"`
	Street       string    `json:"street"`
	Email        string    `json:"email"`
	Products     []Product `json:"chosen_products"`
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
