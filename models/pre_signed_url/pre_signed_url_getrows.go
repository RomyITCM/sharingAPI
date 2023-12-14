package pre_signed_url

import (
	"context"
	"fmt"
	"log"
	"os"
	"smile-service/entities"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetDataPreSignedURL(ctx context.Context, file_path, data string, log *zap.Logger) ([]*entities.DataPreSignedURL, error) {
	i, err := strconv.Atoi(data)
	if err != nil {
		log.Info("parse String to Int",
			zap.Any("rows", err),
		)
		return nil, err
	}

	preSignedUrls := make([]*entities.DataPreSignedURL, 0)
	x := 0
	for x < i {
		url := ""
		url, err = getPresignedURL(file_path)
		preSignedUrl := &entities.DataPreSignedURL{
			PresignedUrl: url,
		}

		if err != nil {
			log.Info("get preSignedURL",
				zap.Any("rows", err),
			)
			return nil, err
		}

		preSignedUrls = append(preSignedUrls, preSignedUrl)
		x++
	}

	return preSignedUrls, nil

}

func getPresignedURL(filename string) (string, error) {

	// Load the bucket name
	s3Bucket := os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars")
	}

	// snippet-start:[s3.go.generate_presigned_url.session]
	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})
	// snippet-end:[s3.go.generate_presigned_url.session]

	// snippet-start:[s3.go.generate_presigned_url.call]
	svc := s3.New(sess)

	r, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
	})

	// Create the pre-signed url with an expiry
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("Failed to generate a pre-signed url: ", err)
		return "", err
	}

	// Display the pre-signed url
	fmt.Println("Pre-signed URL", url)
	return url, nil
}
