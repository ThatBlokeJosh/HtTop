package server

import (
  "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
  "restop/utils"
)

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

func Listen() {



  app := fiber.New()
	
  app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

  app.Static("/", "./dist")

  app.Get("/api/mem", Mem)
  app.Get("/api/cpu", Cpu)
  app.Get("/api/proc", Proc)
  app.Get("/api/disk", Disk)
  app.Get("/api/net", Net)

  app.Listen(":3000")
}



