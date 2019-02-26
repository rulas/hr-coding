package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	var result map[string]interface{}

	for i := 1; i < 600; i++ {
		url := fmt.Sprintf("http://chaostool.oraclecorp.com:8080/chaosTool/api/v1/jobId/%d", i)
		// fmt.Println("Requesthing this url: ", url)
		resp, err := http.Get(url)

		if err != nil {
			log.Fatalln("Failed miresably")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		json.Unmarshal([]byte(body), &result)

		fmt.Printf("%d, %s, %s, %s\n", result["id"], result["outageType"], result["outageStatus"], result)
	}
}
