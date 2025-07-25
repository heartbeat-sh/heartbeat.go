# Heartbeat.sh Go Client

This is a Go client library for [heartbeat.sh](https://heartbeat.sh).

## Usage

Install with `go get github.com/heartbeat-sh/heartbeat.go`

```Go
package main

import (
    "fmt"
    "github.com/heartbeat-sh/heartbeat.go/heartbeatsh"
    "time"
)

func main() {
    client := heartbeatsh.NewClient("example")
    minute := time.Minute
    hour := time.Hour

    // Send a beat
    err := client.SendBeat("go", &minute, &hour)
    if err != nil {
        fmt.Printf("Unexpected error %v", err)
    }

    // Send beat with a time frame, beats will only be checked between 09:00 and 17:00
    err = client.SendBeatWithTimes("go", &minute, &hour, "09:00", "17:00")
    if err != nil {
        fmt.Printf("Unexpected error %v", err)
    }

    // Delete a beat
    err = client.DeleteBeat("go")
    if err != nil {
        fmt.Printf("Unexpected error %v", err)
    }
}
```
