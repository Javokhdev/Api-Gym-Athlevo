package handler

import (
	budgeting "api/genproto/gym"

	"google.golang.org/grpc"
)

type Handler struct {
	Gym      budgeting.GymServiceClient
	Facility       budgeting.FacilityServiceClient
	GymFacility     budgeting.UniqueServiceClient
}

func NewBudgetingHandler(gymConn *grpc.ClientConn) *Handler {
	return &Handler{
		Gym:      budgeting.NewGymServiceClient(gymConn),
		Facility:       budgeting.NewFacilityServiceClient(gymConn),
		GymFacility:     budgeting.NewUniqueServiceClient(gymConn),
	}
}
