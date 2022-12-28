package users

import (
	"bridge/users-service/pkg/logging"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var logger = logging.GetLogger()

type UserHandlers struct {
	UserService *UserService
}

func ProvideUserHandlers(u *UserService) *UserHandlers {
	return &UserHandlers{UserService: u}
}

func (h *UserHandlers) Find(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.UserService.Find(uint(id))

	if errors.Is(err, ErrUserNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err != nil {
		logger.Error(fmt.Sprintf("find user handler error: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandlers) Create(ctx *gin.Context) {
	var data CreateUserDTO
	err := ctx.Bind(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := h.UserService.Create(data)

	if errors.Is(err, ErrUserAlreadyExists) {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	if err != nil {
		logger.Error(fmt.Sprintf("create user handler error: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandlers) Update(ctx *gin.Context) {
	var data UpdateUserDTO
	err := ctx.Bind(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.UserService.Update(uint(id), data)

	if err != nil {
		logger.Error(fmt.Sprintf("update user handler error: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandlers) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	affected, err := h.UserService.Delete(uint(id))

	if errors.Is(err, ErrUserNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err != nil {
		logger.Error(fmt.Sprintf("delete user handler error: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"affected": affected})
}
