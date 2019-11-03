package app

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	defaultAWSRegion = "us-east-1"
)

func getFromS3(bucket, object string) ([]byte, int64, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(defaultAWSRegion),
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
