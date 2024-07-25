package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//if len(os.Args) < 2 {
	//	os.Args = append(os.Args, "web")
	//}
	//mode := os.Args[1]
	//switch mode {
	//case "cli":
	//	cli.Main()
	//case "web":
	//	app.Main()
	//default:
	//	app.Main()
	//
	//}

	//fmt.Println("Running web server...")
	//
	//err := repository.ConnectDB()
	//if err != nil {
	//	fmt.Println("Could not connect to database")
	//	panic(err)
	//}
	//defer repository.DB.Close()

	app := fiber.New()

	app.Get("/", home)

	err := app.Listen(":8080")
	if err != nil {
		return
	}

}

func home(c *fiber.Ctx) error {
	c.Status(200)

	// set content type
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString("hello world ")
}
