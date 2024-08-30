package handler

import (
	pb "api/genproto/gym"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Gym
// @Description Create Gym
// @Tags Gym
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateGymRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /gym/create [post]
func (h *Handler) CreateGym(ctx *gin.Context) {
	req := &pb.CreateGymRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Gym.CreateGym(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Gym Created Successfully"})
}

// @Summary Update Gym
// @Description Update Gym
// @Tags Gym
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateGymRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /gym/update [put]
func (h *Handler) UpdateGym(ctx *gin.Context) {
	req := &pb.UpdateGymRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Gym.UpdateGym(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Gym Updated Successfully"})
}

// @Summary Delete Gym
// @Description Delete Gym
// @Tags Gym
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Gym ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /gym/delete/{id} [delete]
func (h *Handler) DeleteGym(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteGymRequest{Id: id}

	_, err := h.Gym.DeleteGym(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Gym Deleted Successfully"})
}

// @Summary Get Gym
// @Description Get an existing Gym record by ID
// @Tags Gym
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Gym ID"
// @Success 200 {object} pb.GetGymResponse
// @Failure 400 {string} string "Error"
// @Router /gym/get/{id} [get]
func (h *Handler) GetGym(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetGymRequest{Id: id}

	res, err := h.Gym.GetGym(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary List Gyms
// @Description List Gyms
// @Tags Gym
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param name query string false "Name"
// @Param location query string false "Location"
// @Param type_sport query string false "Type Sport"
// @Param type_gender query string false "Type Gender"
// @Success 200 {object} pb.ListGymResponse
// @Failure 400 {string} string "Bad Request"
// @Router /gym/list [get]
func (h *Handler) ListGyms(ctx *gin.Context) {
	defaultPage := 1

	pageStr := ctx.Query("page")
	name := ctx.Query("name")
	location := ctx.Query("location")
	typeSport := ctx.Query("type_sport")
	typeGender := ctx.Query("type_gender")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	req := &pb.ListGymRequest{
		Page:       int32(page),
		Name:       name,
		Location:   location,
		TypeSport:  typeSport,
		TypeGender: typeGender,
	}

	res, err := h.Gym.ListGym(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
