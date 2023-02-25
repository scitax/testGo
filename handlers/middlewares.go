package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PreprocessingMW(c *fiber.Ctx) error {
	if PreprocesStatus != true {
		return c.Status(fiber.StatusInternalServerError).SendString("prerporcessing")
	}
	return c.Next()
}

func RespTimeMW(c *fiber.Ctx) error {
	before := time.Now()
	if err := c.Next(); err != nil {
		return err
	}
	diff := time.Since(before)
	log.Printf("%d | %s | %s | %v", c.Response().StatusCode(),
		c.Method(), c.Path(), diff)
	log.Printf("")
	return nil
}

// func RespTimeMW(param string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		before := time.Now()
// 		if err := c.Next(); err != nil {
// 			return err
// 		}
// 		diff := time.Since(before)
// 		log.Printf("%d | %s | %s | %v", c.Response().StatusCode(),
// 			c.Method(), c.Path(), diff)
// 		log.Printf("")
// 		return nil
// 	}
// }
