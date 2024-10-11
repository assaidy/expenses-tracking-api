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

	// category pos
	CheckIfCategoryExists(category string) (bool, error)

	// expenses pos
	CreateExpnse(exp *models.Expense) (*models.Expense, error)
	GetAllExpensesByUserId(uid int) ([]*models.Expense, error)
	UpdateExpnse(exp *models.Expense) error
	DeleteExpenseById(id int) error
	CheckIfExpenseExists(id int) (bool, error)
}
