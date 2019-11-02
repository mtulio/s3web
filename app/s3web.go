package app

import "log"

type S3web struct {
	Port string
}

type S3webOptions struct {
	Port *string
}

//NewS3Web create a new web application
func NewS3Web(opts S3webOptions) (*S3web, error) {
	s := &S3web{
		Port: *opts.Port,
	}
	return s, nil
}

//Run the web app
func (s *S3web) Run() {
	// expose the routes
	// serve the web server
	log.Println("Running...")
}
