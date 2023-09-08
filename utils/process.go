package utils

import (
  "github.com/shirou/gopsutil/v3/process"
  "sort"
  "restop/utils/math"
)

type Process struct {
  Name string
  Cpu float64
  Mem float64
}

func GetProc(sortBy string) []Process {
  allP, _ := process.Processes()
  var pSlice []Process
  for _, p := range(allP) {
    pName, _ := p.Name()
    pCpu, _ := p.CPUPercent()
    pMem, _ := p.MemoryPercent()
    pSlice = append(pSlice, Process{Name: pName, Cpu: math.Round(pCpu), Mem: math.Round(float64(pMem))})
  }
  if sortBy == "CPU" {
    sort.Slice(pSlice, func(i, j int) bool { return pSlice[i].Cpu > pSlice[j].Cpu })
  } else {
    sort.Slice(pSlice, func(i, j int) bool { return pSlice[i].Mem > pSlice[j].Mem })
  }
  return pSlice
}
