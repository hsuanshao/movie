package mvinterface

// Movie is movie query module's interface
type Movie interface {
	// AskMovie request user provide movie name to get movie information by applied omdb api
	AskMovie(moviename string) ([]byte, error)
}
