package utils

import (
	//"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"

  "restop/utils/math"
)

type CPU struct {
  Percent float64
  Temp float64
}

func GetTemp() float64 {
  temps, _ := host.SensorsTemperatures()
  var avgTemp float64

  for i := range(temps) {
    avgTemp += temps[i].Temperature
  }

  return math.Round(avgTemp / float64(len(temps)))
}

func GetCPU() CPU {
  var c CPU
  p, _ := cpu.Percent(0, false)
  c.Percent = math.Round(p[0]) 
  c.Temp = GetTemp()

  return c
}



