package handler

import (
	pb "api/genproto/gym"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create GymFacility
// @Description Create GymFacility
// @Tags GymFacility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateUniqueRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /gymfacility/create [post]
func (h *Handler) CreateGymFacility(ctx *gin.Context) {
	req := &pb.CreateUniqueRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.GymFacility.CreateUnique(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "GymFacility Created Successfully"})
}

// @Summary Update GymFacility
// @Description Update GymFacility
// @Tags GymFacility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateUniqueRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /gymfacility/update [put]
func (h *Handler) UpdateGymFacility(ctx *gin.Context) {
	req := &pb.UpdateUniqueRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.GymFacility.UpdateUnique(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "GymFacility Updated Successfully"})
}

// @Summary Delete GymFacility
// @Description Delete GymFacility
// @Tags GymFacility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "GymFacility ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /gymfacility/delete/{id} [delete]
func (h *Handler) DeleteGymFacility(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteUniqueRequest{Id: id}

	_, err := h.GymFacility.DeleteUnique(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "GymFacility Deleted Successfully"})
}

// @Summary Get GymFacility
// @Description Get an existing GymFacility record by ID
// @Tags GymFacility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "GymFacility ID"
// @Success 200 {object} pb.GetUniqueResponse
// @Failure 400 {string} string "Error"
// @Router /gymfacility/get/{id} [get]
func (h *Handler) GetGymFacility(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetUniqueRequest{Id: id}

	res, err := h.GymFacility.GetUnique(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary List GymFacilitys
// @Description List GymFacilitys
// @Tags GymFacility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Success 200 {object} pb.ListUniquesResponse
// @Failure 400 {string} string "Bad Request"
// @Router /gymfacility/list [get]
func (h *Handler) ListGymFacilitys(ctx *gin.Context) {
	defaultPage := 1

	pageStr := ctx.Query("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	req := &pb.ListUniqueRequest{
		Page: int32(page),
	}

	res, err := h.GymFacility.ListUnique(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
