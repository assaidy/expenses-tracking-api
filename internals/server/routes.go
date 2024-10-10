package server

import (
	"os"

	"github.com/assaidy/expenses-tracking-api/internals/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
)

func (s *FiberServer) RegisterRoutes() {
	var (
		userH = handlers.NewUserHandler(s.storage)
	)

	s.Post("/users/register", userH.HandleRegisterUser)
	s.Post("/users/login", userH.HandleLoginUser)

	// NOTE: this validates before our logging handler ie. it will not log errors
	// it sends 401: Invalid or expired JWT
	s.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))

	s.Get("/users", userH.HandleGetUserProfile)
	s.Put("/users", userH.HandleUpdateUser)
	s.Delete("/users", userH.HandleDeleteUser)
}
