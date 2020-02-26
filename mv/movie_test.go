package movie

import (
	"fmt"
	"log"
	"testing"
)

func TestAskMovie(t *testing.T) {
	testcase := []struct {
		TestDesc  string
		MovieName string
		ExpResult string
		ExpError  error
	}{
		{
			TestDesc:  "Ordianry test, use actual api request",
			MovieName: "super man",
			ExpResult: `{"Title":"Super Man","Year":"2010","Rated":"N/A","Released":"N/A","Runtime":"15 min","Genre":"Short, Drama","Director":"Peter O'Donoghue","Writer":"N/A","Actors":"Adam Gardiner, Tessa Mitchell, Niamh Peren, Matthew Sunderland","Plot":"A man in a hospital wheelchair is brutally attacked with a bottle of his own urine. The attacker has his reasons, but revenge is not always sweet. A story about mental fragility, the randomness of fate and the complex nature of blame.","Language":"English","Country":"Australia, New Zealand","Awards":"N/A","Poster":"N/A","Ratings":[{"Source":"Internet Movie Database","Value":"5.9/10"}],"Metascore":"N/A","imdbRating":"5.9","imdbVotes":"19","imdbID":"tt1725668","Type":"movie","DVD":"N/A","BoxOffice":"N/A","Production":"N/A","Website":"N/A","Response":"True"}`,
			ExpError:  nil,
		},
		{
			TestDesc:  "Ordianry test, use actual api request",
			MovieName: "asdfasdf",
			ExpResult: `{"Response":"False","Error":"Movie not found!"}`,
			ExpError:  fmt.Errorf("can't get movie info"),
		},
	}

	for _, c := range testcase {
		if c.ExpError != nil {
			continue
		}
		res, err := AskMovie(c.MovieName)
		if err != c.ExpError {
			log.Print("exp error is not equal in case: ", c.TestDesc)
		}

		if string(res) != c.ExpResult {
			log.Print("exp result is not equal to expect result in case: ", c.TestDesc)
		}
	}

}
