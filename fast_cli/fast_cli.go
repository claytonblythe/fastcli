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
	get_url_list(token)
}

func get_url_list(token string) {
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
	fmt.Println(obj)
	fmt.Println(obj["targets"])
	targets := obj["targets"].([]interface{})
	client := obj["client"]
	client_ip := obj["client"].(map[string]interface{})["ip"]
	fmt.Println(targets)
	fmt.Println(client)
	fmt.Println(client_ip)
	for _, target := range targets {
		target := target.(map[string]interface{})
		// element is the element from someSlice for where we are
		url := target["url"]
		fmt.Println(url)
		location := target["location"].(map[string]interface{})
		city := location["city"]
		country := location["country"]
		fmt.Println(city)
		fmt.Println(country)
	}

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
