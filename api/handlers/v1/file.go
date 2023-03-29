package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makves-test-task/api/models"
)

// @Summary		Get All
// @Description	Used to get all informations.
// @Tags		File
// @Accept      json
// @Produce		json
// @Param       limit      query     int true "limit"
// @Param       page       query     int true "page"
// @Success		200 				{object}  models.GetAllFileInfoRes
// @Router		/get-all [get]
func (h *handlerV1) GetAllInfo(c *gin.Context) {
	page, err := h.ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.FailureInfo{
			Message: "Page should be int.",
		})
		h.log.Warn("Couldn't parse page")
		return
	}
	limit, err := h.ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.FailureInfo{
			Message: "Limit should be int.",
		})
		h.log.Warn("Couldn't parse limit")
		return
	}

	res, err := h.fileManager.GetAllInfo(&models.GetAllFileInfoReq{
		Page:  int64(page),
		Limit: int64(limit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.FailureInfo{
			Message: "Something went wrong in server side!",
		})
		h.log.Error("h.fileManager.GetAllInfo() err:" + err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary		Get by id
// @Description	Used to get information by id.
// @Tags		File
// @Accept      json
// @Produce		json
// @Param       id      path     int true "id"
// @Success		200 				{object}  models.GetAllFileInfoRes
// @Router		/get-by-id/{id} [get]
func (h *handlerV1) GetById(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.FailureInfo{
			Message: "id should be int.",
		})
		h.log.Warn("strconv.ParseInt(ids, 10, 64)")
		return
	}
	res, err := h.fileManager.GetById(&models.GetByidReq{
		Id: int(id),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.FailureInfo{
			Message: "Something went wrong in server side!",
		})
		h.log.Error("h.fileManager.GetById() err:" + err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
