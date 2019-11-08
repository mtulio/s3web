package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtulio/s3web/app"
)

var (
	opts = app.S3webOptions{}
)

func init() {
	opts.Port = flag.String("addr", ":8080", "web serve port.")
	opts.Host = flag.String("host", "localhost", "web serve host header.")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [--addr ':8080'] (--vhost <DNS>) \n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	s3web, _ := app.NewS3Web(opts)
	s3web.Run()
}
