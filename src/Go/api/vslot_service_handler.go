//go:generate frodo client vslot_service.go
//go:generate frodo client vslot_service.go --language=js
//go:generate frodo gateway vslot_service.go

package api

import (
	"context"
	"slots/vslot"
)

type VSlotServiceHandler struct {
	MyVSlot *vslot.MyVSlot
}

// GetBalance implements VSlotService.
func (v *VSlotServiceHandler) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return &GetBalanceResponse{Amount: v.MyVSlot.GetBalance(), Success: true}, nil
}

// Reset implements VSlotService.
func (v *VSlotServiceHandler) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	v.MyVSlot.Reset()
	return &ResetResponse{Success: true}, nil
}

// Spin implements VSlotService.
func (v *VSlotServiceHandler) Spin(context.Context, *SpinRequest) (resp *SpinResponse, err error) {
	reels := v.MyVSlot.Spin()
	return &SpinResponse{Reels: reels, Success: true}, nil
}

// UpdateBalance implements VSlotService.
func (v *VSlotServiceHandler) UpdateBalance(_ context.Context, req *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	v.MyVSlot.UpdateBalance(req.Amount)
	return &UpdateBalanceResponse{Success: true}, nil
}
