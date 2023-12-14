package need_approve

import (
	"fmt"
	"log"
	"path"
	"smile-service/entities"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func FolderDocumentCopy(cust_req_no *entities.DataCustomerNumbers) {
	bucket := "upload.file"

	// config, err := external.LoadDefaultAWSConfig()
	// if err != nil {
	// 	log.Fatalf("failed to load config, %v", err)
	// }

	// config.Credentials = NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentialsFromCreds(
				credentials.Value{
					AccessKeyID:     "AKIA3WB2X4HJGJTDWPGL",
					SecretAccessKey: "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",
				},
			),
		},
	))

	svc := s3.New(sess)

	var marker *string
	for {
		res, err := svc.ListObjects(
			&s3.ListObjectsInput{
				Bucket:    aws.String(bucket),
				Delimiter: aws.String("/"),
				Marker:    marker,
				Prefix:    aws.String("Benfarm/Smile/Customer/Documents/" + cust_req_no.CustomerRequestNo + "/"), // e.g. Must end with a "/" for a directory
			},
		)
		if err != nil {
			log.Fatal("Failed to list objects", err)
		}

		fmt.Println(len(res.Contents), *res.IsTruncated, res.NextMarker)

		for _, obj := range res.Contents {
			srcKey := "/" + bucket + "/" + *obj.Key
			destKey := "Benfarm/DocumentCustomer/" + cust_req_no.CustomerNo + "/" + path.Base(*obj.Key)
			_, err = svc.CopyObject(
				&s3.CopyObjectInput{
					Bucket:     aws.String(bucket),
					CopySource: aws.String(srcKey),
					Key:        aws.String(destKey),
				},
			)
			fmt.Println(srcKey, destKey)

			if err != nil {
				log.Printf("Failed to copy object: %v", err)
				continue
			}

			_, _ = svc.DeleteObject(
				&s3.DeleteObjectInput{
					Bucket: aws.String(bucket),
					Key:    obj.Key,
				},
			)
		}

		marker = res.NextMarker
		if marker == nil {
			break
		}
	}
}
