package server

import (
	"channels/internal/repository"
	"channels/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine     *gin.Engine
	Services   *services.Services
	Repository *repository.Repository
}

func NewHandler(engine *gin.Engine, services *services.Services) *Handler {
	return &Handler{
		Engine:   engine,
		Services: services,
	}
}

func (h *Handler) Init() {
	h.Engine.GET("/check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Connected"})
	})
	h.Engine.GET("/get_users", h.GetUser)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	begin, err := strconv.Atoi(ctx.Query("begin"))
	end, err := strconv.Atoi(ctx.Query("end"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = h.Services.GetUser(begin, end)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Status": "Done!"})
}
