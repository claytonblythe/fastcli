package fastcli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"net/url"

	"github.com/fatih/color"

	"github.com/sparrc/go-ping"
)

func make_request(url string, results chan int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	total_read := 0
	for {
		buf := make([]byte, 16384)
		num_bytes, err := resp.Body.Read(buf)
		if err != nil {
			total_read = total_read + num_bytes
			break
		}
		total_read = total_read + num_bytes
	}
	results <- total_read
}

func Test_Speed() {
	color.HiGreen("\nConnecting to fast.com...\n")
	client_display, _, url_list, fast_endpoint := get_urls()
	color.HiGreen("\nTesting Download Speed...")
	color.HiGreen("Client: %s\n", client_display)
	color.HiGreen("Fast.com endpoint: %s\n\n", fast_endpoint)

	num_urls := len(url_list)
	results := make(chan int, num_urls*2)
	start := time.Now()
	for _, url := range url_list {
		go make_request(url, results)
	}

	total_data_downloaded := 0
	for i := 1; i <= num_urls; i++ {
		this_url_data_downloaded := <-results
		total_data_downloaded = total_data_downloaded + this_url_data_downloaded
	}
	duration := time.Since(start).Seconds()
	color.HiGreen("Duration: %.2f seconds\n", duration)
	mb := float64(total_data_downloaded) / float64(125000)
	megabytes := float64(total_data_downloaded) / float64(1048576)
	color.HiGreen("Data downloaded: %.2f MB\n", megabytes)
	mb_per_sec := mb / duration
	color.HiGreen("Speed: %.2f Mbps \n", mb_per_sec)

}

func get_urls() (string, []string, []string, string) {
	js_url := get_js_url()
	token := get_token(js_url)
	client_display, display_strings, url_list, fast_endpoint := get_url_list(token)
	return client_display, display_strings, url_list, fast_endpoint
}

func get_url_list(token string) (string, []string, []string, string) {
	s := []string{"https://api.fast.com/netflix/speedtest/v2?https=true&token=", token, "&urlCount=5"}
	fast_endpoint := strings.Join(s, "")
	response, err := http.Get(fast_endpoint)
	if err != nil {
		color.HiRed(fast_endpoint)
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	responseString := string(responseData)
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(responseString), &obj); err != nil {
		log.Fatal(err)
	}
	targets := obj["targets"].([]interface{})
	client_ip := obj["client"].(map[string]interface{})["ip"].(string)
	client_location := obj["client"].(map[string]interface{})["location"]
	client_city := client_location.(map[string]interface{})["city"].(string)
	client_country := client_location.(map[string]interface{})["country"].(string)
	client_display := strings.Join([]string{client_city, client_country, client_ip}, ", ")
	targets_display := []string{}
	target_urls := []string{}
	color.HiGreen("Connecting to test servers...\n")
	color.HiGreen("Server locations:")
	for _, target := range targets {
		target := target.(map[string]interface{})
		// element is the element from someSlice for where we are
		target_url := target["url"].(string)
		location := target["location"].(map[string]interface{})
		city := location["city"].(string)
		country := location["country"].(string)
		u, err := url.Parse(target_url)
		if err != nil {
			panic(err)
		}
		pinger, err := ping.NewPinger(u.Host)
		if err != nil {
			panic(err)
		}
		pinger.Count = 1
		pinger.Run()                 // blocks until finished
		stats := pinger.Statistics() // get send/receive/rtt stats
		s := []string{city, country, target_url, fmt.Sprintf("%vms Avg RTT", stats.AvgRtt.Milliseconds())}
		target_display := strings.Join(s, ", ")
		color.HiBlue(target_display)
		targets_display = append(targets_display, target_display)
		target_urls = append(target_urls, target_url)
	}

	return client_display, targets_display, target_urls, fast_endpoint

}

func get_token(url string) string {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	token := GetStringInBetween(responseString, "{https:!0,endpoint:apiEndpoint,token:", ",urlCount:5")
	token = token[1 : len(token)-1]
	return token
}

func get_js_url() string {
	url := "https://fast.com"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	out := GetStringInBetween(responseString, "<script src=", "></script>")
	out = out[1 : len(out)-1]
	s := []string{"https://fast.com", out}
	js_url := strings.Join(s, "")
	return js_url
}
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	if e == -1 {
		return
	}
	return str[s:e]
}
