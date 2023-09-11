package utils

import (
  "github.com/shirou/gopsutil/v3/net"
  "httop/utils/math"
)

var lastSent float64
var lastRecv float64

type Net struct {
  Sent float64
  Recv float64
}

func GetNet() Net {
  n, _ := net.IOCounters(true)

  if lastSent == 0 {
    lastSent = math.ToKB(n[1].BytesSent)
    lastRecv = math.ToKB(n[1].BytesRecv)
  }

  sent := math.ToKB(n[1].BytesSent) - lastSent
  lastSent = math.ToKB(n[1].BytesSent)

  recv := math.ToKB(n[1].BytesRecv) - lastRecv
  lastRecv = math.ToKB(n[1].BytesRecv)
  
  return Net{Sent: sent, Recv: recv}
}
