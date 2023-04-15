package routes

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/nwochaadim/go-crm-fiber/lead"
)

func getLeads(c *fiber.Ctx) {
	if err := c.JSON(lead.GetAllLeads()); err != nil {
		handleError(err)
	}
}

func getLead(c *fiber.Ctx) {
	id, _ := strconv.ParseInt(c.Params("id"), 0, 0)

	if err := c.Status(fiber.StatusOK).JSON(lead.GetLead(id)); err != nil {
		handleError(err)
	}
}

func createLead(c *fiber.Ctx) {
	l := new(lead.Lead)

	if err := c.BodyParser(l); err != nil {
		handleError(err)
	}

	if err := c.Status(fiber.StatusOK).JSON(l.Create()); err != nil {
		handleError(err)
	}
}

func deleteLead(c *fiber.Ctx) {
	id, _ := strconv.ParseInt(c.Params("id"), 0, 0)

	lead.DeleteLead(id)

	if err := c.Status(fiber.StatusNoContent).JSON(nil); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	fmt.Printf("Encountered error, %v", err)
}

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", getLeads)
	app.Get("/api/v1/leads/:id", getLead)
	app.Post("/api/v1/leads", createLead)
	app.Delete("/api/v1/leads/:id", deleteLead)
}
