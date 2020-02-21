package castle

import (
	"bytes"
	"errors"
	"flying-castle/cmd"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"net/http"
	"time"
)

type S3Backend struct {
	bucket     string
	svc        *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func (s S3Backend) Read(fileName string) (data []byte, err error) {
	v4.NewSigner(credentials.NewEnvCredentials(), func(signer *v4.Signer) {
		var request *http.Request
		request, err = http.NewRequest("GET", fileName, nil)
		_, err = signer.Sign(request, nil, "s3", "eu-west-3", time.Now())
		if err != nil {
			return
		}
		res, err := http.DefaultClient.Do(request)
		if err != nil {
			return
		}
		if res.StatusCode != 200 {
			err = errors.New(fmt.Sprintf("cannot read data: %s", res.Status))
		}
		data, err = ioutil.ReadAll(res.Body)
		return
	})
	return
}

func (s S3Backend) Write(fileName string, chunkData []byte) (string, error) {
	var objectInput = &s3manager.UploadInput{Bucket: &s.bucket, Key: aws.String(fileName), Body: bytes.NewReader(chunkData)}
	output, err := s.uploader.Upload(objectInput)
	if err != nil {
		return "", err
	}
	return output.Location, nil
}

var BucketNotFoundError = errors.New("bucket not found")

func init() {
	constructors["s3"] = NewS3Backend
}

func NewS3Backend(bucketName string, config *cmd.Config) (StorageBackend, error) {
	var accessId string
	var secret string
	var creds *credentials.Credentials
	if len(config.S3Credentials) == 0 {
		if len(accessId) == 0 || len(secret) == 0 {
			creds = credentials.NewEnvCredentials()
		} else {
			creds = credentials.NewStaticCredentials(config.S3AccessId, config.S3Secret, "")
		}
	} else {
		creds = credentials.NewSharedCredentials(config.S3Credentials, config.S3Profile)
	}
	if _, err := creds.Get(); err != nil || creds.IsExpired() {
		return S3Backend{}, errors.New("invalid s3 credentials")
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("eu-west-3"),
	}))
	svc := s3.New(sess)
	_, err := svc.HeadBucket(&s3.HeadBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		return nil, BucketNotFoundError
	}
	return S3Backend{bucket: bucketName, uploader: s3manager.NewUploader(sess), downloader: s3manager.NewDownloader(sess)}, nil

}
