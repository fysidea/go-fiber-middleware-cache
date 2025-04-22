package cache

import (
	"github.com/gofiber/fiber/v2"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ตัวอย่าง middleware ง่ายๆ
		c.Set("X-Cache", "MISS")
		return c.Next()
	}
}
