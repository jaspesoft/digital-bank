package pkg

import (
	fmt "fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	uuid "github.com/google/uuid"
	log "log"
	os "os"
	"path/filepath"
	time "time"

	"mime/multipart"
)

type (
	AWSS3Storage struct {
		s3Session *s3.S3
		bucket    string
	}
)

func NewAWSS3Storage(bucket string) *AWSS3Storage {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	if err != nil {
		log.Fatalf("Failed to initialize AWS session: %v", err)
	}

	return &AWSS3Storage{
		s3Session: s3.New(sess),
		bucket:    bucket,
	}
}

func (a *AWSS3Storage) UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {

	fileKey := generateFileName(header.Filename)

	// Configura los parámetros de subida
	uploadParams := &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(fileKey),
		Body:        file,
		ContentType: aws.String(header.Header.Get("Content-Type")),
	}

	// Sube el archivo a S3
	_, err := a.s3Session.PutObject(uploadParams)
	if err != nil {
		return "", err
	}

	return fileKey, nil
}

func (a *AWSS3Storage) DownloadFile(fileName string) (string, error) {
	req, _ := a.s3Session.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(fileName),
	})
	urlStr, err := req.Presign(1 * time.Hour)
	if err != nil {
		return "", fmt.Errorf("failed to sign request: %w", err)
	}

	return urlStr, nil
}

func generateFileName(fileName string) string {
	currentDate := time.Now()
	year := currentDate.Year()
	month := int(currentDate.Month())

	extension := filepath.Ext(fileName)
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), extension)

	return fmt.Sprintf("%d/%d/%s", year, month, newFileName)
}
