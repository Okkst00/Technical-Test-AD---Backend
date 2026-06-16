package repository

import (
	"backend-api-commerce/config"
	"backend-api-commerce/model"
)


func CreateUser(user model.User) error {

	_, err := config.DBLocal.Exec(`
		INSERT INTO users
		(
			nama,
			email,
			password,
			role,
			status
		)
		VALUES(?,?,?,?,?)
	`,
		user.Nama,
		user.Email,
		user.Password,
		user.Role,
		user.Status,
	)

	return err
}


func FindUserByEmail(email string)(model.User,error){

	var user model.User

	err:=config.DBLocal.QueryRow(`
		SELECT
		id,
		nama,
		email,
		password,
		role,
		status
		FROM users
		WHERE email=?
	`,
	email).Scan(
		&user.ID,
		&user.Nama,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Status,
	)


	return user,err
}

type User struct {
	ID           int
	MembershipID int
}

func GetUserByID(userID int) (User, error) {
	var user User

	err := config.DBLocal.QueryRow(`
		SELECT id, membership_id
		FROM users
		WHERE id = ?
	`, userID).Scan(&user.ID, &user.MembershipID)

	return user, err
}