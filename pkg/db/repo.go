package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/tanimutomo/sqlfile"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) executeQueriesFromFile(path string) error {
	s := sqlfile.New()

	if err := s.File(path); err != nil {
		return err
	}

	_, err := s.Exec(r.db)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CheckDataBaseAvailable() error {
	if err := r.db.Ping(); err != nil {
		return errDBUnavailable
	}
	return nil
}

func (r *Repository) MakeQueryRow(query string, args ...any) *sql.Row {
	return r.db.QueryRow(query, args...)
}

func (r *Repository) GetProducts() ([]Product, error) {
	query := fmt.Sprintf(`SELECT product_id, name, brand, category, rating,
							price, available, image_path FROM %s;`, productsTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var (
		id          uint32
		name        string
		brand       string
		category    string
		price       uint32
		rating      float64
		available   bool
		pathToImage string
		products    []Product
	)

	for rows.Next() {
		rows.Scan(&id, &name, &brand, &category, &rating,
			&price, &available, &pathToImage)
		products = append(
			products,
			NewProduct(
				id, category, name, brand, price,
				available, rating, pathToImage,
			),
		)
	}

	return products, nil
}

func (r *Repository) CheckEmailExists(email string) error {
	query := fmt.Sprintf(`SELECT user_id FROM %s WHERE email = '%s';`, usersTable, email)

	rows, err := r.db.Query(query)
	if err != nil {
		return err
	}

	var (
		rowsNumber int
		name       string
	)
	for rows.Next() {
		rows.Scan(&name)
		rowsNumber++
	}

	if rowsNumber == 0 {
		return errNoSuchEmail
	}

	return nil
}

func (r *Repository) UpdatePassword(email string, newPassword string) error {
	query := fmt.Sprintf(
		`UPDATE %s SET hashed_password = '%s' WHERE email = '%s';`,
		usersTable, newPassword, email)

	_, err := r.db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	query := fmt.Sprintf(
		`SELECT firstname, surname, email, phone_number, birthdate FROM %s WHERE email = '%s';`,
		usersTable, email)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	user := &User{}

	for rows.Next() {
		rows.Scan(
			&user.Firstname,
			&user.Surname,
			&user.Email,
			&user.PhoneNumber,
			&user.Birthday,
		)
	}

	return user, nil
}

func (r *Repository) UpdateUser(user *User, oldEmail string) error {
	query := fmt.Sprintf(
		`UPDATE %s 
		SET firstname = '%s',
			surname = '%s',
			phone_number = '%s',
			email = '%s',
			birthdate = '%s'
		WHERE email = '%s';
		`, usersTable, user.Firstname, user.Surname,
		user.PhoneNumber, user.Email,
		user.Birthday, oldEmail,
	)

	_, err := r.db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddOrder(order *Order) (int, error) {
	user, err := r.GetUserByEmail(strings.ReplaceAll(order.Email, "%40", "@"))
	if err != nil {
		return -1, err
	}

	query := fmt.Sprintf(
		`INSERT INTO %s(client_id, purchase_price,
						date, delivery_date, city,
						street, house_number, flat_number) VALUES (
							'%s', %d, NOW(), '%s', '%s', '%s', %d, %d
						) RETURNING purchase_id;`, ordersTable, strconv.Itoa(int(user.Id)),
		order.TotalSum, order.DeliveryDate, order.City,
		order.Street, order.HouseNumber, order.FlatNumber)

	rows, err := r.db.Query(query)
	if err != nil {
		return -1, err
	}

	var orderID int
	for rows.Next() {
		rows.Scan(&orderID)
	}

	return orderID, nil
}
