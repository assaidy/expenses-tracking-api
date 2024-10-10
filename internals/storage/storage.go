package storage

import "github.com/assaidy/expenses-tracking-api/internals/models"

type Storage interface {
	// user ops
	CheckUsernameAndEmailConflict(username string, email string) (bool, error)
	CheckUsernameConflict(username string) (bool, error)
	CheckEmailConflict(email string) (bool, error)
	CheckIfUserExists(id int) (bool, error)
	CreateUser(user *models.User) error
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(newUser *models.User) error
	DeleteUserById(id int) error

	// xxx pos
}
