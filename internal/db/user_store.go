package db

import (
	"context"
	"matryer/internal"
	"matryer/internal/types"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserStore interface {
	GetUserByID(int) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(context.Context, *types.User) (*User, error)
}

type User struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	Age       int32  `db:"age"`
}

type MySQLUserStore struct {
	logger *logrus.Logger
	client *sqlx.DB
}

func NewMySQLUserStore(client *sqlx.DB, logger *logrus.Logger) *MySQLUserStore {
	return &MySQLUserStore{
		client: client,
		logger: logger,
	}
}

func (s *MySQLUserStore) GetUserByID(id int) (*User, error) {
	return &User{}, nil
}

func (s *MySQLUserStore) GetUsers() ([]*User, error) {
	var users []*User
	err := s.client.Select(&users, "CALL spGetUsers()")
	if err != nil {
		s.logger.Printf("error get users: %v", err)
		return nil, internal.ErrSqlSelect.SetOrigin(err)
	}

	if len(users) == 0 {
		s.logger.Print("there are no user records")
		return nil, internal.ErrNoContent
	}

	return users, nil
}

func (s *MySQLUserStore) CreateUser(ctx context.Context, user *types.User) (*User, error) {
	var userID int64
	query := "CALL spCreateUser(?,?,?,?,?,?)"
	err := s.client.QueryRowxContext(ctx, query,
		user.Name,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Age,
	).Scan(&userID)
	if err != nil {
		s.logger.Errorf("error create user: %v", err)
		return nil, internal.ErrSqlQueryRow.SetOrigin(err)
	}

	return &User{
		ID:        userID,
		Name:      user.Name.Value(),
		FirstName: user.FirstName.Value(),
		LastName:  user.LastName.Value(),
		Email:     user.Email.Value(),
		Phone:     user.Phone.Value(),
		Age:       user.Age.Value(),
	}, nil
}
