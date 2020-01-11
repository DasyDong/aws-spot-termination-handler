package main

import (
	"kube-spot-temination-handler/pkg"
	"os"
	"time"
)

const (
	secs = 5
)


func getStat(quit chan int) {

	if pkg.CheckForSpotInterruptionNotice() {
		//Get instanceID by aws or Get env from docker file
		//InstanceID, _ := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
		nodeName := os.Getenv("NODE_NAME")
		if nodeName != "" {
			pkg.Drain(nodeName)
			pkg.SlackNotify(nodeName)
			quit <- 1
		}
	}
}

func main() {
	quit := make(chan int, 1)
		timer := time.NewTicker(time.Duration(secs) * time.Second)
		for {
			select {
			case <-timer.C:
				go getStat(quit)
			case <-quit:
				return
			}
	}
}
