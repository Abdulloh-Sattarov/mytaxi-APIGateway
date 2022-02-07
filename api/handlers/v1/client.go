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

// CreateClient ...
// @Summary CreateClient
// @Description This API for creating a new client
// @Tags client
// @Accept  json
// @Produce  json
// @Param Client request body models.CUClient true "clientCreateRequest"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/ [post]
func (h *handlerV1) CreateClient(c *gin.Context) {
	var (
		body        pb.Client
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

	response, err := h.serviceManager.MyTaxiService().CreateClient(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create client", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetClient ...
// @Summary GetClient
// @Description This API for getting client detail
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/{id} [get]
func (h *handlerV1) GetClient(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MyTaxiService().GetClient(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get client", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateClient ...
// @Summary UpdateClient
// @Description This API for updating client
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Client request body models.CUClient true "clientUpdateRequest"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/{id} [put]
func (h *handlerV1) UpdateClient(c *gin.Context) {
	var (
		body        pb.Client
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

	response, err := h.serviceManager.MyTaxiService().UpdateClient(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update client", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteClient ...
// @Summary DeleteClient
// @Description This API for deleting client
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/{id} [delete]
func (h *handlerV1) DeleteClient(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MyTaxiService().DeleteClient(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete client", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
