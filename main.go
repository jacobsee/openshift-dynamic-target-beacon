package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	cron "github.com/robfig/cron"
	"k8s.io/client-go/rest"
)

type Data struct {
	Token string `json:"token"`
}

func main() {
	c := cron.New()
	c.AddFunc("*/5 * * * *", register)
	c.Start()
	for true {
		time.Sleep(time.Hour)
	}
}

func register() {
	dynamicTargetServer := os.Getenv("SERVER")
	key := os.Getenv("CLUSTER_URL")
	auth_token := os.Getenv("AUTH_TOKEN")
	debug := os.Getenv("DEBUG")
	if key == "" || auth_token == "" {
		log.Println("Missing required environment variable")
		os.Exit(1)
	}

	var bearer_token string
	if len(debug) == 0 {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		bearer_token = config.BearerToken
	} else {
		bearer_token = "DEBUGMODE"
	}

	data := Data{
		Token: bearer_token,
	}
	dataJson, _ := json.Marshal(data)
	client := http.Client{}
	formValues := url.Values{
		"kind": {"openshift"},
		"key":  {key},
		"data": {string(dataJson)},
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/register", dynamicTargetServer),
		strings.NewReader(formValues.Encode()),
	)
	req.Header.Add("Authorization", auth_token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
}
