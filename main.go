package main

import (
	"log"

	"github.com/RoniLandau/building-management-system/internal/routes"
	"github.com/RoniLandau/building-management-system/pkg/db"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // includes registering the PostgreSQL driver with the database/sql package - side effect
)

func main() {

	// Initialize the database connection
	db.InitDB()

	// Initialize a new Fiber app
	app := fiber.New()

	// Define routes
	routes.SetupRoutes(app)

	// Start the Fiber server
	log.Fatal(app.Listen(":3000"))

	// Defer closing the database connection
	defer db.CloseDB()
}
