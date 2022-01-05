package climb

import (
	"net/http"

	entities "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
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

		climb, err := entities.NewClimb(
			"",
			req.Date,
			req.Grade,
			req.Description,
			req.Area,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := climbRepository.Save(ctx, climb); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}