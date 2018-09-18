// blinkRpio.go
// Rich Robinson
// Sept 2018

package main

import (
    "os"
    "fmt"
    "time"
    "os/signal"
    "syscall"
    "math/rand"
    "github.com/richrarobi/blinkRpio"
    
)

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
    running := true
// initialise getout
    signalChannel := make(chan os.Signal, 2)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
    go func() {
        sig := <-signalChannel
        switch sig {
        case os.Interrupt:
            fmt.Println("Stopping on Interrupt")
            running = false
            return
        case syscall.SIGTERM:
            fmt.Println("Stopping on Terminate")
            running = false
            return
        }
    }()

    blinkRpio.Setup()
    blinkRpio.SetLuminance(1)
    blinkRpio.Clear()
    blinkRpio.Show()

    for running {
        pixel := rand.Intn(8)
// note the int parameter for brightness 0 to 31 (not a float)
// also reduced brightness looks better hence only up to 5 here
        blinkRpio.SetPixel( pixel, rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(3) )
        blinkRpio.Show()
        r,g,b,l := blinkRpio.GetPixel( pixel)
        fmt.Println("getPixel", pixel, r,g,b,l )
        delay(50)
    }
    
    fmt.Println("Stopping")
    blinkRpio.Exit()
}
