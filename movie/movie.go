package movie

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const authorization string = "cb132a69"

// AskMovie request user provide movie name to get movie information by applied omdb api
func AskMovie(moviename string) ([]byte, error) {
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/", nil)
	if err != nil {
		log.Printf("request " + moviename + " info failure")
		return nil, fmt.Errorf("can't get movie info")
	}

	q := req.URL.Query()
	// OMDB open api request parameter
	q.Add("apikey", authorization)
	q.Add("t", moviename)

	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print("get movie info failure", err)
		return nil, fmt.Errorf("API request failure")
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("read response failure")
		return nil, fmt.Errorf("Get API response failure")
	}

	return responseData, nil
}
