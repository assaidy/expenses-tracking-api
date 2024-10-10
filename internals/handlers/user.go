package handlers

import (
	"fmt"
	"time"

	"github.com/assaidy/expenses-tracking-api/internals/models"
	"github.com/assaidy/expenses-tracking-api/internals/storage"
	"github.com/assaidy/expenses-tracking-api/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	st storage.Storage
}

func NewUserHandler(s storage.Storage) *UserHandler {
	return &UserHandler{
		st: s,
	}
}

func (h *UserHandler) HandleRegisterUser(c *fiber.Ctx) error {
	req := models.UserRegisterOrUpdateRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.InvalidJsonRequestError()
	}

	if errs := utils.ValidateRequest(req); errs != nil {
		return utils.ValidationError(errs)
	}

	if ok, err := h.st.CheckUsernameAndEmailConflict(req.Username, req.Email); err != nil {
		return utils.InternalServerError(err)
	} else if ok {
		return utils.ConflictError("username or email already exists")
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password, // TODO: hash the password
		Email:    req.Email,
		JoinedAt: time.Now().UTC(),
	}
	if err := h.st.CreateUser(&user); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "user created successfully",
	})
}

func (h *UserHandler) HandleLoginUser(c *fiber.Ctx) error {
	req := models.UserLoginRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.InvalidJsonRequestError()
	}

	if errs := utils.ValidateRequest(req); errs != nil {
		return utils.ValidationError(errs)
	}

	user, err := h.st.GetUserByUsername(req.Username)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if user == nil || req.Password != user.Password { // TODO: unhash the password
		return utils.UnauthorizedError()
	}

	tokenStr, err := utils.GenerateJwtToken(user.Id, user.Username)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "logged in successfully",
		Data:    fiber.Map{"token": tokenStr},
	})
}

func (h *UserHandler) HandleGetUserProfile(c *fiber.Ctx) error {
	id, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	user, err := h.st.GetUserById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if user == nil {
		return utils.NotFoundError(fmt.Sprintf("user with id %d not found", id))
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "user found",
		Data:    fiber.Map{"user": user},
	})
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	req := models.UserRegisterOrUpdateRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.InvalidJsonRequestError()
	}

	if errs := utils.ValidateRequest(req); errs != nil {
		return utils.ValidationError(errs)
	}

	id, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	user, err := h.st.GetUserById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if user == nil {
		return utils.UnauthorizedError()
	}

	if req.Username != user.Username {
		if ok, err := h.st.CheckUsernameConflict(req.Username); err != nil {
			return utils.InternalServerError(err)
		} else if ok {
			return utils.ConflictError("username already exists")
		}
	}
	if req.Email != user.Email {
		if ok, err := h.st.CheckEmailConflict(req.Email); err != nil {
			return utils.InternalServerError(err)
		} else if ok {
			return utils.ConflictError("email already exists")
		}
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Username = req.Username
	user.Password = req.Password

	err = h.st.UpdateUser(user)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiError{
		Message: "updated successfully",
	})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	if ok, err := h.st.CheckIfUserExists(id); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.UnauthorizedError()
	}

	if err := h.st.DeleteUserById(id); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}
