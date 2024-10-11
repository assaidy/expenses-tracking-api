package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/assaidy/expenses-tracking-api/internals/models"
	"github.com/assaidy/expenses-tracking-api/internals/storage"
	"github.com/assaidy/expenses-tracking-api/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type ExpenseHandler struct {
	st storage.Storage
}

func NewExpenseHandler(s storage.Storage) *ExpenseHandler {
	return &ExpenseHandler{
		st: s,
	}
}

func (h *ExpenseHandler) HandleCreateExpense(c *fiber.Ctx) error {
	uid, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	if ok, err := h.st.CheckIfUserExists(uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.UnauthorizedError()
	}

	req := models.ExpenseCreateOrUpdateRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.InvalidJsonRequestError()
	}

	if errs := utils.ValidateRequest(req); errs != nil {
		return utils.ValidationError(errs)
	}

	if ok, err := h.st.CheckIfCategoryExists(req.Category); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("category %s not found", req.Category))
	}

	exp := models.Expense{
		UserId:      uid,
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
		AddedAt:     time.Now().UTC(),
	}
	resExp, err := h.st.CreateExpnse(&exp)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "expense created successfully",
		Data:    fiber.Map{"expense": resExp},
	})
}

func (h *ExpenseHandler) HandleGetAllExpenses(c *fiber.Ctx) error {
	uid, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	if ok, err := h.st.CheckIfUserExists(uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.UnauthorizedError()
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	dateRange := c.Query("date_range", "")
	var startDate, endDate string
	switch dateRange {
	case "week":
		startDate = time.Now().AddDate(0, 0, -7).Format(time.RFC3339)
		endDate = time.Now().Format(time.RFC3339)
	case "month":
		startDate = time.Now().AddDate(0, -1, 0).Format(time.RFC3339)
		endDate = time.Now().Format(time.RFC3339)
	case "3months":
		startDate = time.Now().AddDate(0, -3, 0).Format(time.RFC3339)
		endDate = time.Now().Format(time.RFC3339)
	case "custom":
		startDate = c.Query("start_date")
		startDate = c.Query("end_date")
		if startDate == "" || endDate == "" {
			return utils.ApiError{
				Code:    fiber.StatusBadRequest,
				Message: "invalid date range for custom filter",
			}
		}
	}

	exps, err := h.st.GetAllExpensesByUserId(uid, page, limit, startDate, endDate)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "expense created successfully",
		Data:    fiber.Map{"expenses": exps, "page": page, "total": len(exps)},
	})
}

func (h *ExpenseHandler) HandleUpdateExpense(c *fiber.Ctx) error {
	uid, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	if ok, err := h.st.CheckIfUserExists(uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.UnauthorizedError()
	}

	req := models.ExpenseCreateOrUpdateRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.InvalidJsonRequestError()
	}

	if errs := utils.ValidateRequest(req); errs != nil {
		return utils.ValidationError(errs)
	}

	if ok, err := h.st.CheckIfCategoryExists(req.Category); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("category %s not found", req.Category))
	}

	// TODO: add last_updated attribute
	eid, _ := c.ParamsInt("id")
	if ok, err := h.st.CheckIfExpenseExists(eid, uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("expense with id %d not found", eid))
	}

	exp := models.Expense{
		Id:          eid,
		UserId:      uid,
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
	}
	resExp, err := h.st.UpdateExpnse(&exp)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "expense updated successfully",
		Data:    fiber.Map{"expense": resExp},
	})
}

func (h *ExpenseHandler) HandleDeleteExpense(c *fiber.Ctx) error {
	uid, ok := utils.GetUserIdFromContext(c)
	if !ok {
		return utils.UnauthorizedError()
	}

	if ok, err := h.st.CheckIfUserExists(uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.UnauthorizedError()
	}

	eid, _ := c.ParamsInt("id")
	if ok, err := h.st.CheckIfExpenseExists(eid, uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("expense with id %d not found", eid))
	}

	if err := h.st.DeleteExpenseById(eid); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "expense deleted successfully",
	})
}
