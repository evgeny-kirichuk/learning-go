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

type HotelQuery struct {
	Rooms bool
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var query HotelQuery
	if err := c.QueryParser(&query); err != nil {
		return err
	}
	hotels, err := h.hotelStore.List(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}