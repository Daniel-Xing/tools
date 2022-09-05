package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var tr *http.Transport

var (
	remoteURL = "https://impact.ese.ic.ac.uk/ImpactEarth/cgi-bin/crater.cgi?dist=12371&distanceUnits=1&diam=111&diameterUnits=1&pdens=111&pdens_select=0&vel=1111&velocityUnits=1&theta=45&tdens=1000&wdepth=1111&wdepthUnits=1"
	localURL2 = "http://121.36.81.191/cgi-bin/crater.cgi?dist=111&distanceUnits=1&diam=111&diameterUnits=1&pdens=111000&pdens_select=0&vel=111&velocityUnits=1&theta=45&tdens=1000&wdepth=111&wdepthUnits=1"
	localURL  = "http://121.36.81.191:50052/simulator"
	localURL3 = "http://121.36.81.191:50052/simulatorWithRedis"
)

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
	}
}

func GetRe(URL string) {
	client := &http.Client{
		Transport: tr,
		Timeout:   100 * time.Minute,
	}
	defer client.CloseIdleConnections()

	content, err := client.Get(URL)
	if err != nil {
		log.Println(err)
		return
	}
	defer content.Body.Close()
}

func Post(postBody string, URL string) {
	client := &http.Client{
		Transport: tr,
		Timeout:   100 * time.Minute,
	}
	defer client.CloseIdleConnections()

	req, err := http.NewRequest("POST", URL, bytes.NewReader([]byte(postBody)))
	if err != nil {
		log.Fatal(err)
		return
	}

	content, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer content.Body.Close()
}

func main() {
	// read the args
	url := flag.String("url", "http://121.36.81.191:50052/simulator", "Url")
	postBody := flag.String("post_body", "", "json string of post body")
	requestTimes := flag.Int("requestTimes", 10, "requestTimes")
	requestMethod := flag.String("requestMethod", "post", "requestMethod")

	flag.Parse()

	start := time.Now()
	wg := sync.WaitGroup{}

	fmt.Printf("Run request for %d times to URL: %s \n", *requestTimes, *url)
	for i := 0; i < *requestTimes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if *requestMethod == "get" {
				GetRe(*url)
			} else if *requestMethod == "post" {
				Post(*postBody, *url)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Time Cost: ", time.Since(start).Milliseconds(), "ms")
	return
}
