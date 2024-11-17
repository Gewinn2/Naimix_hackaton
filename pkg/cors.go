package pkg

import "github.com/gofiber/fiber/v2"

// Middleware, который позволяет клиентам с разных источников безопасно взаимодействовать с вашим API.
// Он устанавливает необходимые заголовки CORS в HTTP-ответах.
func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		origin := c.Get("Origin")
		c.Set("Access-Control-Allow-Origin", origin)
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "*, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")

		// Проверка на предварительный запрос с методом OPTIONS.
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent) // 204 No Content
		}

		// Для всех остальных запросов вызывается следующий обработчик в цепочке.
		return c.Next()
	}
}
