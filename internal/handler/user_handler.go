package handler

import (
	userusecase "github.com/alxhtp/monogo/internal/usecase/user"
	"github.com/alxhtp/monogo/pkg/dto"
	dtobase "github.com/alxhtp/monogo/pkg/dto/base"
	errorhelper "github.com/alxhtp/monogo/pkg/helper/error"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// init dtobase
var _ = dtobase.BaseRes{}

type userHandler struct {
	userUsecase userusecase.UserUsecase
}

func NewUserHandler(userUsecase userusecase.UserUsecase) *userHandler {
	return &userHandler{userUsecase: userUsecase}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.ReqCreateUser true "User"
// @Success 201 {object} dto.ResUserSingle
// @Router /users [post]
func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.ReqCreateUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":    false,
			"code":       fiber.StatusBadRequest,
			"message":    err.Error(),
			"stacktrace": errorhelper.ComposeStacktrace(err),
		})
	}

	res := h.userUsecase.CreateUser(c.Context(), &req)
	return c.Status(res.Code).JSON(res)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.ResUserSingle
// @Router /users/{id} [get]
func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res := h.userUsecase.GetUserByID(c.Context(), uuid.MustParse(id))
	return c.Status(res.Code).JSON(res)
}

// GetUsersByFilter godoc
// @Summary Get users by filter
// @Description Get users by filter
// @Tags User
// @Accept json
// @Produce json
// @Param ids query string false "User IDs, comma separated uuids"
// @Param name query string false "Name"
// @Param email query string false "Email"
// @Param status query int false "Status"
// @Param sex query string false "Sex"
// @Param address query string false "Address"
// @Param phone query string false "Phone"
// @Param include-deleted query bool false "Include Deleted"
// @Param show-count query bool false "Show Count"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param order-by query string false "Order By, default: +created_at"
// @Param created-at-gte query time.Time false "Created At Greater Than or Equal To"
// @Param created-at-lte query time.Time false "Created At Less Than or Equal To"
// @Param updated-at-gte query time.Time false "Updated At Greater Than or Equal To"
// @Param updated-at-lte query time.Time false "Updated At Less Than or Equal To"
// @Success 200 {object} dto.ResUserList
// @Router /users [get]
func (h *userHandler) GetUsersByFilter(c *fiber.Ctx) error {
	var req dto.ReqGetUser
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":    false,
			"code":       fiber.StatusBadRequest,
			"message":    err.Error(),
			"stacktrace": errorhelper.ComposeStacktrace(err),
		})
	}

	res := h.userUsecase.GetUsersByFilter(c.Context(), &req)
	return c.Status(res.Code).JSON(res)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.ReqUpdateUser true "User"
// @Success 200 {object} dto.ResUserSingle
// @Router /users/{id} [put]
func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.ReqUpdateUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":    false,
			"code":       fiber.StatusBadRequest,
			"message":    err.Error(),
			"stacktrace": errorhelper.ComposeStacktrace(err),
		})
	}

	res := h.userUsecase.UpdateUser(c.Context(), uuid.MustParse(id), &req)
	return c.Status(res.Code).JSON(res)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dtobase.BaseRes
// @Router /users/{id} [delete]
func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	res := h.userUsecase.DeleteUser(c.Context(), uuid.MustParse(id))
	return c.Status(res.Code).JSON(res)
}
