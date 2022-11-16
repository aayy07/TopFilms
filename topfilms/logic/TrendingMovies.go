package logic

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aayy07/topfilms/graph/model"
	// "github.com/opentracing/opentracing-go/log"
)

func TrendingMovies(ctx context.Context, filter string) (*model.TrendingMoviesResponse, error) {

	url:= os.Getenv("API") + filter + os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Println(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil{
		log.Println(err)
	}

	defer resp.Body.Close()

	responseBody, responseErr := ioutil.ReadAll(resp.Body)
	if responseErr != nil{
		log.Println(responseErr)
	}

	var response model.TrendingMoviesResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil{
		log.Println(err)
	}
	return &response, nil
}