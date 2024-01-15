package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Stats string `json:"stats"`
	Body  int    `json: "body"`
}

func main() {
	fmt.Print("Hello World")

	app := fiber.New()

	listTodo := []Todo{}

	app.Use(func(c *fiber.Ctx) error {
		if c.Method() != "POST" && c.Method() != "GET" && c.Method() != "DELETE" && c.Method() != "PATCH" {
			return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
				"error": "Método não permitido",
			})
		}

		return c.Next()
	})

	app.Post("/api/tasks", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(listTodo) + 1
		listTodo = append(listTodo, *todo)

		return c.JSON(listTodo)
	})

	app.Patch("/api/tasks/:id/:status", func(c *fiber.Ctx) error {
		param := struct {
			ID     int    `params:"id"`
			STATUS string `params:"status"`
		}{}

		if err := c.QueryParser(&param); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		for index, item := range listTodo {
			if item.ID == param.ID {
				listTodo[index].Stats = SetStatus(param.STATUS)
				break
			}
		}

		return c.JSON(listTodo)
	})

	app.Get("/api/tasks/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.JSON(listTodo)
		}

		var foundItem Todo
		for _, todo := range listTodo {
			if todo.ID == id {
				foundItem = todo
				break
			}
		}

		return c.JSON(foundItem)

	})

	app.Delete("/api/tasks/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Id invalido",
			})
		}

		response, mensage := RemoveItem(&listTodo, id)

		return c.Status(response).JSON(fiber.Map{"mensage": mensage})
	})

	log.Fatal(app.Listen(":4000"))
}

func SetStatus(status string) string {
	switch status {
	case "Ongoing", "Completed":
		return status
	default:
		return "Starting"
	}
}
func RemoveItem(listTodo *[]Todo, id int) (int, string) {
	if id < 0 || id >= len(*listTodo) {
		return 400, "ID inválido!"
	}

	*listTodo = append((*listTodo)[:id], (*listTodo)[id+1:]...)

	return 200, "Item excluido com sucesso!"
}
