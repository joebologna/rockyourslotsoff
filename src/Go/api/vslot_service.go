package api

import "context"

type VSlotService interface {
	Spin(context.Context, *SpinRequest) (*SpinResponse, error)
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
}

type SpinRequest struct {
}

type SpinResponse struct {
	Reels   [3]int
	Success bool
}

type ResetRequest struct{}
type ResetResponse struct{ Success bool }

type UpdateBalanceRequest struct{ Amount int }
type UpdateBalanceResponse struct{ Success bool }

type GetBalanceRequest struct{}
type GetBalanceResponse struct {
	Amount  int
	Success bool
}
