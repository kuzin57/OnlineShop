package db

const (
	usersTable                = "bshop.user"
	productsTable             = "bshop.product"
	productsDetailedInfoTable = "bshop.product_char"
	ordersTable               = "bshop.purchase"
	ordersProductsTable       = "bshop.purchase_product"
	onlyDateLength            = 10
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
	Id            uint32  `json:"id"`
	Category      string  `json:"category"`
	Name          string  `json:"name"`
	Brand         string  `json:"brand"`
	Price         uint32  `json:"price"`
	Available     bool    `json:"available"`
	Rating        float64 `json:"rating"`
	RatingAmount  uint64  `json:"rating_amount"`
	PathToImage   string  `json:"path_to_image"`
	Amount        uint32  `json:"amount"`
	Kcal          int     `json:"kcal"`
	Proteins      float32 `json:"proteins"`
	Fats          float32 `json:"fats"`
	Carbohydrates float32 `json:"carbohydrates"`
	Weight        float32 `json:"weight"`
	ExpireDate    string  `json:"expire_date"`
	Country       string  `json:"country"`
	Description   string  `json:"description"`
}

type Order struct {
	Id           uint32    `json:"id"`
	Date         string    `json:"date"`
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
	price uint32, available bool, rating float64, pathToImage string, amount uint32) Product {
	return Product{
		Id:          id,
		Category:    category,
		Name:        name,
		Brand:       brand,
		Price:       price,
		Available:   available,
		Rating:      rating,
		PathToImage: pathToImage,
		Amount:      amount,
	}
}

func NewOrder(
	orderID uint32,
	date, deliveryDate, city, street string,
	price, houseNumber, flatNumber uint32) *Order {
	return &Order{
		Id:           orderID,
		Date:         date,
		DeliveryDate: deliveryDate,
		TotalSum:     int(price),
		HouseNumber:  uint16(houseNumber),
		FlatNumber:   uint16(flatNumber),
		City:         city,
		Street:       street,
	}
}