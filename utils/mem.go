package utils

import (
  "github.com/shirou/gopsutil/v3/mem"
  "httop/utils/math"
)

type Memory struct {
  Total float64
  Used float64
  Percent int
}

func GetMemory() Memory {
  m, _ := mem.VirtualMemory()
  return Memory{Total: math.ToGB(m.Total), Used: math.ToGB(m.Used), Percent: int(m.UsedPercent)}
}

