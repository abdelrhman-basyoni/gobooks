package sqlRepos

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/abdelrhman-basyoni/gobooks/config"
	"github.com/abdelrhman-basyoni/gobooks/models"
)

type UserRepo struct {
	Id       int    `bson:"_id" json:"id"`
	UserName string `bson:"username" json:"username"  validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`
	Email    string `bson:"email" json:"email"   validate:"required,email"`
}

var db = config.SqlDb

func innit() sql.Result {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE,
		password VARCHAR(255),
		email VARCHAR(255) UNIQUE
	)
`
	res, err := db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("Table 'users' Already exists")
		fmt.Println(err)
	}
	fmt.Println("Table 'users' created successfully")

	return res
}

var I = innit()

func (u *UserRepo) Create(username, password, email string) error {
	insertSQL := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3)
`
	_, err := db.Exec(insertSQL, username, password, email)
	return err
}

func (u *UserRepo) GetUserById(id string) (*models.User, error) {

	selectQuery := `SELECT * FROM users WHERE id = $1`
	fmt.Println("createdUser")

	row := db.QueryRow(selectQuery, id)

	var user models.User

	err := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {

	selectQuery := `SELECT * FROM users WHERE email = $1`
	fmt.Println("createdUser")

	row := db.QueryRow(selectQuery, email)

	var user models.User

	err := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	selectSQL := `
	SELECT id, username, password, email
	FROM users
`
	var users []models.User

	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (u *UserRepo) EditUser(id string, update map[string]interface{}) error {
	var placeholders []string
	var values []interface{}
	index := 1
	for key, value := range update {
		placeholders = append(placeholders, fmt.Sprintf("%s = $%d", key, index))
		values = append(values, value)
		index++
	}
	// Combine placeholders into a comma-separated string
	setClause := strings.Join(placeholders, ", ")

	// Define the SQL update statement
	updateSQL := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $%d
	`, setClause, index)
	fmt.Println(updateSQL)
	// Append the user ID as the last value
	values = append(values, id)

	// Execute the update query with placeholders
	_, err := db.Exec(updateSQL, values...)

	return err
}
