package math

import (
	"fmt"
	"math"
	"strconv"
)

func Round2(x float64) float64 {
  f, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", x), 64)
  return f
}

func Round(x float64) float64 {
  unit := 0.01
  if x > 0 {
      return Round2(float64(int64(x/unit+0.5)) * unit) 
  }
    return Round2(float64(int64(x/unit-0.5)) * unit) 
}


func ToGB(x uint64) float64 {
  y := float64(x) / 1000000000
  return math.Round(y)
}

func ToKB(x uint64) float64 {
  y := float64(x) / 1000
  return math.Round(y) 
}


func ToMB(x uint64) float64 {
  y := float64(x) / 1000000
  return Round(y) 
}
