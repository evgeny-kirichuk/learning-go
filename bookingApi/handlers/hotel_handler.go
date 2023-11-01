package handlers

import (
	"learning_go/bookingApi/db"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

func NewHotelHandler(rs db.RoomStore, hs db.HotelStore) *HotelHandler {
	return &HotelHandler{roomStore: rs, hotelStore: hs}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.List(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}