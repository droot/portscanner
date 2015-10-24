package scanner

import (
  "flag"
  "fmt"
  "testing"
  "time"
)

var (
  ipAddr = flag.String("ip", "localhost",
    "ip address of target host")
  concurrencyLevel = flag.Int("concurrency", 1,
    "number of concurrent scan worker to launch")
)

func TestScanner(t *testing.T) {

  fmt.Printf("scanning host: %s with concurrency: %d \n",
    *ipAddr, *concurrencyLevel)

  t1 := time.Now()
  openPorts := scanPorts(*ipAddr, *concurrencyLevel)
  timeElapsed := time.Since(t1)

  if len(openPorts) > 0 {
    fmt.Printf("open ports are: %v \n", openPorts)
    fmt.Printf("scanning time: %v \n", timeElapsed)
  }
}
