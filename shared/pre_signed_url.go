package shared

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetPresignedURL(filename string) (string, error) {

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
