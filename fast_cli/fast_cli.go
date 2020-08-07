package fast_cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Get_urls() {
	fmt.Println("get_urls!")
	js_url := get_js_url()
	token := get_token(js_url)
	fmt.Println(token)
	client_ip, display_strings, target_urls := get_url_list(token)

	fmt.Println(client_ip)
	fmt.Println(display_strings)
	fmt.Println(target_urls)
}

func get_url_list(token string) (string, []string, []string) {
	s := []string{"https://api.fast.com/netflix/speedtest/v2?https=true&token=", token, "&urlCount=5"}
	url := strings.Join(s, "")
	response, err := http.Get(url)
	if err != nil {
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
	// client_location :=  obj["client"].(map[string]interface{})["ip"].(string)
	display_targets := []string{}
	target_urls := []string{}
	for _, target := range targets {
		target := target.(map[string]interface{})
		// element is the element from someSlice for where we are
		url := target["url"].(string)
		fmt.Println(url)
		location := target["location"].(map[string]interface{})
		city := location["city"].(string)
		country := location["country"].(string)
		s := []string{city, country, url}
		display_target := strings.Join(s, ", ")
		display_targets = append(display_targets, display_target)
		target_urls = append(target_urls, url)
	}

	return client_ip, display_targets, target_urls

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
