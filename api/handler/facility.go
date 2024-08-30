package handler

import (
	pb "api/genproto/gym"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Facility
// @Description Create Facility
// @Tags Facility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateFacilityRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /facility/create [post]
func (h *Handler) CreateFacility(ctx *gin.Context) {
	req := &pb.CreateFacilityRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Facility.CreateFacility(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Facility Created Successfully"})
}

// @Summary Update Facility
// @Description Update Facility
// @Tags Facility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateFacilityRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /facility/update [put]
func (h *Handler) UpdateFacility(ctx *gin.Context) {
	req := &pb.UpdateFacilityRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Facility.UpdateFacility(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Facility Updated Successfully"})
}

// @Summary Delete Facility
// @Description Delete Facility
// @Tags Facility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Facility ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /facility/delete/{id} [delete]
func (h *Handler) DeleteFacility(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteFacilityRequest{Id: id}

	_, err := h.Facility.DeleteFacility(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Facility Deleted Successfully"})
}

// @Summary Get Facility
// @Description Get an existing Facility record by ID
// @Tags Facility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Facility ID"
// @Success 200 {object} pb.GetFacilityResponse
// @Failure 400 {string} string "Error"
// @Router /facility/get/{id} [get]
func (h *Handler) GetFacility(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetFacilityRequest{Id: id}

	res, err := h.Facility.GetFacility(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary List Facilitys
// @Description List Facilitys
// @Tags Facility
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param name query string false "Name"
// @Success 200 {object} pb.ListFacilityResponse
// @Failure 400 {string} string "Bad Request"
// @Router /facility/list [get]
func (h *Handler) ListFacilitys(ctx *gin.Context) {
	defaultPage := 1

	pageStr := ctx.Query("page")
	name := ctx.Query("name")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	req := &pb.ListFacilityRequest{
		Page:       int32(page),
		Name:       name,
	}

	res, err := h.Facility.ListFacility(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
