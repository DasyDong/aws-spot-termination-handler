package pkg

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	SpotInstanceTerminationUrl = "http://169.254.169.254/latest/meta-data/spot/termination-time"
)


// MonitorForSpotITNEvents continuously monitors metadata for spot ITNs and sends drain events to the passed in channel
func CheckForSpotInterruptionNotice() bool{
	log.Println("Started monitoring for spot ITN events")
	resp, err := getSpotTerminationStat()
	if err != nil {
		log.Fatalf("Unable to parse metadata response: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		log.Println("Sending drain event to the drain channel")
		return true
	}
	return false
}

func getSpotTerminationStat() (*http.Response, error) {
	httpReq := func() (*http.Response, error) {
		return http.Get(SpotInstanceTerminationUrl)
	}
	return retry(3, 2*time.Second, httpReq)
}

func retry(attempts int, sleep time.Duration, httpReq func() (*http.Response, error)) (*http.Response, error) {
	resp, err := httpReq()
	if err != nil {
		if attempts--; attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			log.Printf("Request failed. Attempts remaining: %d\n", attempts)
			log.Printf("Sleep for %s seconds\n", sleep)
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, httpReq)
		}

		log.Fatalln("Error getting response from instance metadata ", err.Error())
	}

	return resp, err
}
