package handlers

import (
	"strconv"

	"github.com/RoniLandau/building-management-system/internal/models"
	"github.com/RoniLandau/building-management-system/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// GetBuildings retrieves all buildings
func GetBuildings(c *fiber.Ctx) error {
	buildings, err := models.Buildings().All(c.Context(), db.Executor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(buildings)
}

// GetApartment retrieves a single apartment by ID
func GetBuilding(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	building, err := models.FindBuilding(c.Context(), db.Executor, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Building not found")
	}

	return c.JSON(building)
}

// CreateBuilding creates a new building
func CreateBuilding(c *fiber.Ctx) error {
	var building models.Building

	if err := c.BodyParser(&building); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := building.Insert(c.Context(), db.Executor, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(building)
}

// DeleteBuilding deletes a building by ID
func DeleteBuilding(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing ID parameter")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	// Create a new Apartment instance and set its ID
	building := &models.Building{ID: id}

	rowsAffected, err := building.Delete(c.Context(), db.Executor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// Check if the delete operation affected any rows
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Building not found")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
