package repositories

import (
	"database/sql"

	"github.com/Chethu16/Online-Payment-Application/internal/models"
	"github.com/Chethu16/Online-Payment-Application/pkg/queries"
)

type authenticationRepository struct {
	db *sql.DB
}

func NewAuthenticationRepository(conn *sql.DB) *authenticationRepository {
	return &authenticationRepository{
		conn,
	}
}

func (ar *authenticationRepository) Login(req models.LoginRequest) (models.LoginResponse, error) {
	query := queries.NewQuery(ar.db)
	token, err := query.LoginQuery(req)
	if err != nil {
		return models.LoginResponse{}, err
	}
	return models.LoginResponse{Token: token}, nil
}

func (ar *authenticationRepository) Register(req models.RegisterRequest) (models.RegisterResponse, error) {
	query := queries.NewQuery(ar.db)
	token, err := query.RegsiterQuery(req)
	if err != nil {
		return models.RegisterResponse{}, err
	}
	return models.RegisterResponse{Token: token}, nil
}
