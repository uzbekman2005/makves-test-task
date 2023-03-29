package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/makves-test-task/api/models"
)

// @Summary		Ping pong
// @Description	Just to ensure that server is running
// @Tags		Ping
// @Produce		json
// @Success		200 				{object}  models.SuccessInfo
// @Router		/ [get]
func (h *handlerV1) AppIsRunning(c *gin.Context) {
	c.JSON(200, models.SuccessInfo{
		Message: "Server is running successfully",
	})
}
