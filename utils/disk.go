package utils

import (
  "github.com/shirou/gopsutil/v3/disk"
  "restop/utils/math"
)

type Disk struct {
  Device string
  Path string
  Fstype string
  Total float64
  Free float64
  Used float64
  UsedPercent float64
}

func GetDisk() []Disk {
  var allDisks []Disk
  ad, _ := disk.Partitions(false)
  for _, d := range(ad) {
    u, _ := disk.Usage(d.Mountpoint)

    temp := Disk{Device: d.Device, Path: d.Mountpoint, Fstype: d.Fstype, 
      Total: math.ToGB(u.Total),
      Free: math.ToGB(u.Free), Used: math.ToGB(u.Used), UsedPercent: math.Round(u.UsedPercent)}
    allDisks = append(allDisks, temp)
  }
  return allDisks
}
