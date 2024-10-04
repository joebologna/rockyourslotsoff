//go:generate frodo client vslot_service.go
//go:generate frodo client vslot_service.go --language=js
//go:generate frodo gateway vslot_service.go

package api

import (
	"context"
	"slots/vslot"

	"github.com/monadicstack/frodo/rpc/errors"
)

type VSlotServiceHandler struct {
	MyVSlot *vslot.MyVSlot
}

// Spin performs a spin on the virtual slot machine.
func (v *VSlotServiceHandler) Spin(context.Context, *SpinRequest) (resp *SpinResponse, err error) {
	reels, isWinner, err := v.MyVSlot.Spin()
	if err != nil {
		return &SpinResponse{}, errors.BadRequest(err.Error())
	}
	return &SpinResponse{Reels: reels, IsWinner: isWinner, Success: true}, nil
}

// GetBalance returns the current balance of the virtual slot machine.
func (v *VSlotServiceHandler) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return &GetBalanceResponse{Amount: v.MyVSlot.GetBalance(), Success: true}, nil
}

// GetReels returns the current reels of the virtual slot machine.
func (v *VSlotServiceHandler) GetReels(context.Context, *GetReelsRequest) (*GetReelsResponse, error) {
	return &GetReelsResponse{Reels: v.MyVSlot.GetReels(), Success: true}, nil
}

// Reset resets the virtual slot machine.
func (v *VSlotServiceHandler) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	v.MyVSlot.Reset()
	return &ResetResponse{Success: true}, nil
}

// UpdateBalance updates the balance of the virtual slot machine.
func (v *VSlotServiceHandler) UpdateBalance(_ context.Context, req *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	v.MyVSlot.UpdateBalance(req.Amount)
	return &UpdateBalanceResponse{Amount: v.MyVSlot.GetBalance(), Success: true}, nil
}
