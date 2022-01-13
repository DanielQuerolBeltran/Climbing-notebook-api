package climb

import (
	"net/http"

	entities "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
	service "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/services"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Id     string `json:"id"`
	Description     string `json:"description"`
	Area string `json:"area" binding:"required"`
	Grade string `json:"grade" binding:"required"`
	Date string `json:"date" binding:"required"`
}

func CreateHandler(climbRepository entities.ClimbRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		climb, err := service.NewSaveClimbService(climbRepository).Execute(
			ctx,
			req.Date,
			req.Grade,
			req.Description,
			req.Area,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, createRequest{
			Id: climb.Id().String(),
			Description: climb.Description().String(),
			Area: climb.Area().String(),
			Grade: climb.Grade().String(),
			Date: climb.Date().String(),
		})
	}
}