package repository

import (
	"backend-api-commerce/config"
	"backend-api-commerce/model"
)


func AddBlacklistToken(token model.TokenBlacklist) error {

	_, err := config.DBLocal.Exec(`
		INSERT INTO token_blacklist
		(
			token,
			expired_at
		)
		VALUES (?,?)
	`,
		token.Token,
		token.ExpiredAt,
	)

	return err
}


func IsTokenBlacklisted(token string) bool {

	var id int


	err := config.DBLocal.QueryRow(`
		SELECT id
		FROM token_blacklist
		WHERE token=?
	`,
	token).Scan(&id)


	return err == nil
}