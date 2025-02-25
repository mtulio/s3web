package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

//S3web is the prototype for S3web application
type S3web struct {
	Port       string
	Host       string
	ServeVhost bool
}

//S3webOptions is the options used for flag
type S3webOptions struct {
	Port *string
	Host *string
}

//NewS3Web create a new web application
func NewS3Web(opts S3webOptions) (*S3web, error) {
	s := &S3web{
		Port: *opts.Port,
	}
	if opts.Host != nil {
		s.Host = *opts.Host
		s.ServeVhost = true
	}
	return s, nil
}

//Run the web app
func (s *S3web) Run() {
	http.HandleFunc("/", s.serveRoot)

	log.Printf("Listening on port %s...\n", s.Port)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		log.Println(err)
	}
}

func (s *S3web) serveRoot(w http.ResponseWriter, r *http.Request) {
	region := r.URL.Query().Get("region")
	if region != "" {
		defaultAWSRegion = region
	}

	// bucket := r.URL.Query().Get("bucket")
	// if bucket == "" {
	// 	http.Error(w, "Get 'bucket' not specified in url.", 400)
	// 	return
	// }
	bucket, object, ok := s.parseBucket(&w, r)
	if !ok {
		return
	}

	filename := filepath.Base(object)

	b, _, err := getFromS3(bucket, object)
	if err != nil {
		logHTTPError(r, &w, err, 404)
		return
	} else {
		logRequest(r)
	}

	// Detect file header (first 512B) and write header
	fileHeader := make([]byte, 512)
	byteReader := bytes.NewReader(b)
	byteReader.ReadAt(fileHeader, 0)
	mime := http.DetectContentType(fileHeader)
	found := true

	if mime == "application/octet-stream" {
		mime, found = fileGetMIME(filename)
	}

	if !found {
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	}
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", cap(b)))

	w.Write(b)

	return
}

func (s *S3web) parseBucket(w *http.ResponseWriter, r *http.Request) (
	bucket string, object string, ok bool,
) {
	ok = true

	if s.ServeVhost {
		bucket = strings.Split(r.Host, ":")[0]
		object = r.URL.Path

	} else {
		bucket = r.URL.Query().Get("bucket")
		if bucket == "" {
			http.Error(*w, "Get 'bucket' not specified in url.", 400)
			return "", "", !ok
		}

		object = r.URL.Query().Get("object")
		if object == "" {
			http.Error(*w, "Get 'object' not specified in url.", 400)
			return "", "", !ok
		}
	}

	return bucket, object, ok
}
