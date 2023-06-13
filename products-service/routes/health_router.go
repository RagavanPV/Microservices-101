package routes

import (
    "github.com/gofiber/fiber/v2"
	"encoding/json"
)


func HealthRouter(app *fiber.App) {
    app.Get("/actuator/health/liveness", func(c *fiber.Ctx) error {
		str := `{"status": "UP","components": {"livenessstate": {"status": "UP"}}}`
    	var res map[string]interface{}
    	json.Unmarshal([]byte(str), &res)
		return c.JSON(res)
	  })
	  app.Get("/actuator/health/readiness", func(c *fiber.Ctx) error {
		str := `{"status": "UP","components": {"readinessstate": {"status": "UP"}}}`
    	var res map[string]interface{}
    	json.Unmarshal([]byte(str), &res)
		return c.JSON(res)
	  })
}