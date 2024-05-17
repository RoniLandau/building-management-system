package handlers

import (
	"strconv"

	"github.com/RoniLandau/building-management-system/internal/models"
	"github.com/RoniLandau/building-management-system/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetApartments lists all apartments
func GetApartments(c *fiber.Ctx) error {
	apartments, err := models.Apartments().All(c.Context(), db.Executor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(apartments)
}

// GetApartment retrieves a single apartment by ID
func GetApartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	apartment, err := models.FindApartment(c.Context(), db.Executor, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Apartment not found")
	}

	return c.JSON(apartment)
}

// GetApartmentsByBuilding lists all apartments in a specific building
func GetApartmentsByBuilding(c *fiber.Ctx) error {
	buildingId, err := strconv.Atoi(c.Params("buildingId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid building ID")
	}

	apartments, err := models.Apartments(qm.Where("building_id =?", buildingId)).All(c.Context(), db.Executor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(apartments)
}

// CreateApartment adds a new apartment
func CreateApartment(c *fiber.Ctx) error {
	var apartment models.Apartment

	if err := c.BodyParser(&apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := apartment.Insert(c.Context(), db.Executor, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(apartment)
}

// DeleteApartment deletes an apartment by ID
func DeleteApartment(c *fiber.Ctx) error {
	// Extract the ID from the request parameters
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing ID parameter")
	}

	// Convert the ID string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	// Create a new Apartment instance and set its ID
	apartment := &models.Apartment{ID: id}

	// Call the Delete method on the Apartment instance
	rowsAffected, err := apartment.Delete(c.Context(), db.Executor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Check if the delete operation affected any rows
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Apartment not found")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
