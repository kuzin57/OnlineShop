package db

import (
	"fmt"
)

type AuthPostgres struct {
	repo *Repository
}

func NewAuthPostgresService(repo *Repository) *AuthPostgres {
	return &AuthPostgres{
		repo: repo,
	}
}

func (a *AuthPostgres) CheckEmailUnique(email string) error {
	query := fmt.Sprintf(`SELECT * FROM bshop.user WHERE email='%s';`, email)
	row := a.repo.MakeQueryRow(query)
	if row.Err() != nil {
		return errEmailExists
	}

	return nil
}

func (a *AuthPostgres) CreateUser(user *User) (uint32, error) {
	if err := a.repo.CheckDataBaseAvailable(); err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`INSERT INTO %s (firstname, surname, phone_number,
		email, birthdate, hashed_password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`, usersTable)

	row := a.repo.MakeQueryRow(query, user.Firstname, user.Surname,
		user.PhoneNumber, user.Email, user.Birthday, user.Password)

	var userID uint32
	if err := row.Scan(&userID); err != nil {
		return 0, err
	}

	user.Id = userID
	return userID, nil
}

func (a *AuthPostgres) GetUser(email string, password string) (*User, error) {
	query := fmt.Sprintf(`SELECT * FROM bshop.user WHERE email='%s';`, email)
	row := a.repo.MakeQueryRow(query)
	if row == nil {
		return nil, errNoSuchUser
	}

	var user User
	if err := row.Scan(
		&user.Id,
		&user.Firstname,
		&user.Surname,
		&user.PhoneNumber,
		&user.Email,
		&user.Birthday,
		&user.Password,
	); err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errIncorrectPassword
	}

	return &user, nil
}
