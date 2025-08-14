package queries

import (
	"github.com/Chethu16/Online-Payment-Application/internal/models"
	"github.com/Chethu16/Online-Payment-Application/pkg/utils"
	"github.com/google/uuid"
)

func (q *query) LoginQuery(req models.LoginRequest) (string, error) {
	var userName, userId, password string
	query := `SELECT user_name , user_id , user_password FROM users WHERE user_email=$1`
	row := q.db.QueryRow(query, req.Email)
	if err := row.Scan(&userName, &userId, &password); err != nil {
		return "", err
	}
	if err := utils.ComparePassword(password, req.Password); err != nil {
		return "", err
	}
	token, err := utils.CreateToken(userId, userName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (q *query) RegsiterQuery(req models.RegisterRequest) (string, error) {
	var userId, password string
	userId = uuid.NewString()
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}
	token, err := utils.CreateToken(userId, req.UserName)
	if err != nil {
		return "", err
	}
	query := `INSERT INTO users(user_id,user_name,user_email,user_password) VALUES($1,$2,$3,$4)`
	_, err = q.db.Exec(query, userId, req.UserName, req.Email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
