package server

import (
	"encoding/json"
  "io/ioutil"
	"httop/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
  Name string
  Password string
  Remember bool
}

func Mem(c *fiber.Ctx) error {
  m := utils.GetMemory()
  
  return c.JSON(m)
}

func Cpu(c *fiber.Ctx) error {
  cpu := utils.GetCPU()
  
  return c.JSON(cpu)
}

func Proc(c *fiber.Ctx) error {
  proc := utils.GetProc(c.Query("sortBy"))

  return c.JSON(proc)
}

func Disk(c *fiber.Ctx) error {
  disks := utils.GetDisk()

  return c.JSON(disks)
}

func Net(c *fiber.Ctx) error {
  net := utils.GetNet()

  return c.JSON(net)
}

func Kill(c *fiber.Ctx) error {
  utils.Kill(c.Query("name"))
  return c.SendStatus(200)
}

func Login(c *fiber.Ctx) error {
  req := c.Context().PostBody()
  body := User{} 
  err := json.Unmarshal(req, &body)
 
  content, _ := ioutil.ReadFile("conf.json")


	check := User{}
	err = json.Unmarshal(content, &check)

  if err != nil {
    log.Fatal(err)
  }

  if body.Name == check.Name && utils.CheckPasswordHash(body.Password, check.Password) {
    return c.JSON("/")
  }
  return c.JSON("/login")
}

func Listen() {
  
  app := fiber.New()
	
  app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

  app.Static("/", "./dist")

  app.Post("/api/login", Login)

  app.Get("/api/mem", Mem)
  app.Get("/api/cpu", Cpu)
  app.Get("/api/proc", Proc)
  app.Get("/api/disk", Disk)
  app.Get("/api/net", Net)
  app.Get("/api/proc/kill", Kill)

  app.Listen(":3000")
}



