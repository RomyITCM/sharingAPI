package need_approve

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"smile-service/entities"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	execSPCustomerDocInsert = `exec [sp_smile_customer_documents_insert] $1,$2,$3,$4`
)

type RenamerInput struct {
	Size         int64
	SourceBucket string
	SourceKey    string
	DestBucket   string
	DestKey      string
}

// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.
type BucketBasics struct {
	S3Client *s3.Client
}

type S3 struct {
	client *s3.Client
	signer *s3.PresignClient
}

type PresignedURLArgs struct {
	Bucket string
	Key    string
	Expiry time.Duration
}

func NewCrossAccountConfigWithRole(ctx context.Context, roleARN string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return aws.Config{}, err
	}

	stsClient := sts.NewFromConfig(cfg)
	stsCreds := stscreds.NewAssumeRoleProvider(stsClient, "arn:aws:iam::803282084306:group/s3.admin")

	cfg.Credentials = aws.NewCredentialsCache(stsCreds)

	return cfg, nil
}

func (s S3) PresignedUploadURL(ctx context.Context, args PresignedURLArgs) (string, error) {
	input := s3.PutObjectInput{
		Bucket: aws.String(args.Bucket),
		Key:    aws.String(args.Key),
	}

	expiry := func(opts *s3.PresignOptions) {
		opts.Expires = args.Expiry
	}

	req, err := s.signer.PresignPutObject(ctx, &input, expiry)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	return req.URL, nil
}

func getPresignedURL() (string, error) {
	ctx := context.Background()

	awsConfig, err := NewCrossAccountConfigWithRole(ctx, "arn:aws:iam::803282084306:group/s3.admin")
	if err != nil {
		log.Fatal(err)
	}

	s3 := NewS3(awsConfig)

	// UPLOAD ------------------------------------------------------------------
	uploadArgs := PresignedURLArgs{
		Bucket: "upload.file",
		Key:    "picture.png",
		Expiry: time.Minute * 15,
	}

	uploadURL, err := s3.PresignedUploadURL(ctx, uploadArgs)
	if err != nil {
		log.Fatalln(err)
	}

	return uploadURL, err

}

func NewS3(cfg aws.Config) S3 {
	client := s3.NewFromConfig(cfg)

	return S3{
		client: client,
		signer: s3.NewPresignClient(client),
	}
}

func CopyToFolder(basics S3) error {
	_, err := basics.client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String("upload.file"),
		CopySource: aws.String(fmt.Sprintf("%v/%v", "upload.file", "Benfarm/Smile/Customer/Documents/1239r0392_2023-10-02_08_07_30.png")),
		Key:        aws.String(fmt.Sprintf("%v/%v", "Benfarm/DocumentCustomer/testCust/", "Benfarm/Smile/Customer/Documents/1239r0392_2023-10-02_08_07_30.png")),
	})
	if err != nil {
		log.Printf("Couldn't copy object from %v:%v to %v:%v/%v. Here's why: %v\n",
			"upload.file", "Benfarm/Smile/Customer/Documents/1239r0392_2023-10-02_08_07_30.png", "upload.file", "Benfarm/DocumentCustomer/testCust/", "Benfarm/Smile/Customer/Documents/1239r0392_2023-10-02_08_07_30.png", err)
	}
	return err
}

func DataCustomerDocInsert(ctx context.Context, db *sql.DB,
	data_doc *entities.InfoCustomerDocument, log *zap.Logger) error {

	ctxb := context.Background()

	dt := time.Now()
	filename := data_doc.DocImg + "_" + dt.Format("2006-01-02_15_04_05") + ".png"

	awsConfig, err := NewCrossAccountConfigWithRole(ctxb, "arn:aws:iam::803282084306:group/s3.admin")
	if err != nil {
		// log.Fatal(err)
	}

	rows, err3 := db.QueryContext(ctx, execSPCustomerDocInsert,
		data_doc.DocumentNo,
		data_doc.CustomerNo,
		filename,
		data_doc.CreatedBy,
	)

	if err3 != nil {
		log.Info(
			"Exec DB", zap.Any("Exec Db", err3),
		)

		return nil
	}
	defer rows.Close()

	for rows.Next() {
		if err3 := rows.Scan(); err3 != nil {
			log.Info("scan rows",
				zap.Any("rows", err3),
			)
			return nil
		}
	}
	if err3 = rows.Err(); err3 != nil {
		return nil
	}

	s3 := NewS3(awsConfig)

	err2 := CopyToFolder(s3)

	if err != nil {
		return err2
	}

	return nil
}
