package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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
	http.HandleFunc("/", s.serveRoot)

	log.Printf("Listening on port %s...\n", s.Port)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *S3web) serveRoot(w http.ResponseWriter, r *http.Request) {
	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		http.Error(w, "Get 'bucket' not specified in url.", 400)
		return
	}

	object := r.URL.Query().Get("object")
	if object == "" {
		http.Error(w, "Get 'object' not specified in url.", 400)
		return
	}
	filename := filepath.Base(object)

	b, _, err := getFromS3(bucket, object)
	if err != nil {
		_logHTTPError(r, &w, err, 404)
		return
	} else {
		_logRequest(r)
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

func getFromS3(bucket, object string) ([]byte, int64, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	downloader := s3manager.NewDownloader(sess)
	buf := aws.NewWriteAtBuffer([]byte{})

	n, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return nil, n, fmt.Errorf("failed to download file, %v", err)
	}

	return buf.Bytes(), n, nil
}
