package routes

import (
	"github.com/RoniLandau/building-management-system/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Buildings routes
	app.Get("/buildings", handlers.GetBuildings)
	app.Get("/buildings/:id", handlers.GetBuilding)
	app.Post("/buildings", handlers.CreateBuilding)
	app.Delete("/buildings/:id", handlers.DeleteBuilding)

	// Apartments routes
	app.Get("/apartments", handlers.GetApartments)
	app.Get("/apartments/:id", handlers.GetApartment)
	app.Get("/apartments/building/:buildingId", handlers.GetApartmentsByBuilding)
	app.Post("/apartments", handlers.CreateApartment)
	app.Delete("/apartments/:id", handlers.DeleteApartment)
}
