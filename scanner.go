package scanner

import (
  "fmt"
  "net"
  "strings"
  "sync"
  "time"
)

// scanPort performs a TCP connect based scanning on given IP/port.
func scanPort(ip string, port uint16) bool {
  addr := fmt.Sprintf("%s:%d", ip, port)
  timeout := time.Duration(5) * time.Second

  conn, err := net.DialTimeout("tcp", addr, timeout)
  if err != nil {
    if strings.Contains(err.Error(), "connection refused") {
      // this is expected on a non listening port
    } else {
      fmt.Println(err)
    }
    return false
  }
  conn.Close()
  return true
}

// scanPorts performs scans on entire possible port range and returns
// slice of open ports.
func scanPorts(ip string, concurrencyLevel int) []uint16 {
  var ports []uint16
  maxPort := 1<<16 - 1
  for ii := 1; ii <= maxPort; ii++ {
    ports = append(ports, uint16(ii))
  }

  portRanges := partition(ports, concurrencyLevel)

  var wg sync.WaitGroup
  // result slice with slot for each scan worker
  openPortResult := make([][]uint16, concurrencyLevel)

  wg.Add(concurrencyLevel)
  for ii := 0; ii < concurrencyLevel; ii++ {
    // launch scan workers with their port ranges to scan
    go func(idx int, portRange []uint16) {
      var openPorts []uint16
      for _, p := range portRange {
        isPortOpen := scanPort(ip, p)
        if isPortOpen {
          openPorts = append(openPorts, p)
        }
      }
      openPortResult[idx] = openPorts
      wg.Done()
    }(ii, portRanges[ii])
  }
  wg.Wait()

  // time to collect results
  var result []uint16
  for _, pp := range openPortResult {
    if len(pp) > 0 {
      result = append(result, pp...)
    }
  }
  return result
}