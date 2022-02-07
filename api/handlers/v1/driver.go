package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/abdullohsattorov/mytaxi-APIGateway/api/handlers/models"
	pb "github.com/abdullohsattorov/mytaxi-APIGateway/genproto"
	l "github.com/abdullohsattorov/mytaxi-APIGateway/pkg/logger"
)

// CreateDriver ...
// @Summary CreateDriver
// @Description This API for creating a new driver
// @Tags driver
// @Accept  json
// @Produce  json
// @Param Driver request body models.CUDriver true "driverCreateRequest"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/ [post]
func (h *handlerV1) CreateDriver(c *gin.Context) {
	var (
		body        pb.Driver
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MyTaxiService().CreateDriver(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create driver", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetDriver ...
// @Summary GetDriver
// @Description This API for getting driver detail
// @Tags driver
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/{id} [get]
func (h *handlerV1) GetDriver(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MyTaxiService().GetDriver(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateDriver ...
// @Summary UpdateDriver
// @Description This API for updating driver
// @Tags driver
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Driver request body models.CUDriver true "driverUpdateRequest"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/{id} [put]
func (h *handlerV1) UpdateDriver(c *gin.Context) {
	var (
		body        pb.Driver
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MyTaxiService().UpdateDriver(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update driver", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteDriver ...
// @Summary DeleteDriver
// @Description This API for deleting driver
// @Tags driver
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/{id} [delete]
func (h *handlerV1) DeleteDriver(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MyTaxiService().DeleteDriver(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete driver", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
