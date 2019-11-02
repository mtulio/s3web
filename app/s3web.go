package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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
	// expose the routes
	// serve the web server
	// log.Println("Running...")
	http.HandleFunc("/", s.serveRoot)

	log.Printf("Listening on port %s...\n", s.Port)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *S3web) serveRoot(w http.ResponseWriter, r *http.Request) {
	//First of check if Get is set in the URL
	Filename := r.URL.Query().Get("file")
	if Filename == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + Filename)

	err := getFromS3(&w, Filename)
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found on S3.", 404)
		return
	}
	err = getFromFile(&w, Filename)
	if err != nil {
		//File not found, send 404
		http.Error(w, "File read error.", 404)
		return
	}
	return
}

func getFromFile(w *http.ResponseWriter, filename string) error {
	//Check if file exists and open
	Openfile, err := os.Open(filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(*w, "File not found.", 404)
		return err
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	(*w).Header().Set("Content-Disposition", "attachment; filename="+filename)
	(*w).Header().Set("Content-Type", FileContentType)
	(*w).Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(*w, Openfile) //'Copy' the file to the client
	return nil
}

func getFromS3(w *http.ResponseWriter, filename string) error {
	// The session the S3 Downloader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", filename, err)
	}

	// buf := aws.NewWriteAtBuffer([]byte{})

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("linx-cloud-static"),
		Key:    aws.String("index.md"),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return nil
}
