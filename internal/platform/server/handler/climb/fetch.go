package climb

import (
	"net/http"

	entities "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
	"github.com/gin-gonic/gin"
)

type fetchRequest struct {
	Id    string    `json:"id"`
	Description string `json:"description"`
	Grade  string `json:"grade"`
	Date  string `json:"date"`
	Area  string `json:"area"`
}

func FetchHandler(climbRepository entities.ClimbRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var fetchRequests []fetchRequest
		climbs, err := climbRepository.Get(ctx, entities.ClimbId{})

		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}

		for _, climb := range climbs {
			fetchRequests = append(fetchRequests, 
				fetchRequest{
					Id: climb.Id().String(),
					Description: climb.Description().String(),
					Grade: climb.Grade().String(),
					Date: climb.Date().String(),
					Area: climb.Area().String(),
				},
			)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"length": len(climbs),
			"result": fetchRequests,
		})

		ctx.Status(http.StatusCreated)
	}
}