package castle

import (
	"bytes"
	"crypto/rand"
	"flying-castle/app"
	db2 "flying-castle/db"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestS3Backend_Read(t *testing.T) {
	var bucket string
	var ok bool
	if bucket, ok = os.LookupEnv("FC_BUCKET"); !ok {
		t.Skip("no s3 bucket")
	}
	s3Reader, _ := NewS3Backend(bucket, &app.Config{
		DbUrl:    db2.SqliteMemory,
		DataPath: fmt.Sprintf("s3://%s", bucket),
	})
	var randomContent = make([]byte, 10)
	rand.Read(randomContent)
	v4.NewSigner(credentials.NewEnvCredentials(), func(signer *v4.Signer) {
		fileUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, "test")
		request, _ := http.NewRequest("PUT", fileUrl, bytes.NewReader(randomContent))
		_, err := signer.Sign(request, bytes.NewReader(randomContent), "s3", "eu-west-3", time.Now())
		if err != nil {
			panic(err)
		}
		res, _ := http.DefaultClient.Do(request)
		if res.StatusCode != 200 {
			panic(res.Status)
		}
		b, err := s3Reader.Read(fileUrl)
		if err != nil {
			t.Fail()
		}
		if !bytes.Equal(b, randomContent) {
			t.Fail()
		}
	})
}

func TestS3Backend_Write(t *testing.T) {
	var bucket string
	var ok bool
	if bucket, ok = os.LookupEnv("FC_BUCKET"); !ok {
		t.Skip("no s3 bucket")
	}
	s3Writer, err := NewS3Backend(bucket, &app.Config{
		DbUrl:    db2.SqliteMemory,
		DataPath: fmt.Sprintf("s3://%s", bucket),
	})
	if err != nil {
		t.Fatal(err)
	}
	object, err := s3Writer.Write("1", []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	v4.NewSigner(credentials.NewEnvCredentials(), func(signer *v4.Signer) {
		request, _ := http.NewRequest("GET", object, nil)
		_, err := signer.Sign(request, nil, "s3", "eu-west-3", time.Now())
		if err != nil {
			panic(err)
		}
		res, _ := http.DefaultClient.Do(request)
		if res.StatusCode != 200 {
			t.Fatal("Object is inaccessible", res.Status)
		}
		result, _ := ioutil.ReadAll(res.Body)
		if !bytes.Equal(result, []byte("hello")) {
			t.Fatal("different file")
		}
	})
}

func TestS3Backend_Delete(t *testing.T) {
	var bucket string
	var ok bool
	if bucket, ok = os.LookupEnv("FC_BUCKET"); !ok {
		t.Skip("no s3 bucket")
	}
	s3Reader, err := NewS3Backend(bucket, &app.Config{
		DbUrl:    db2.SqliteMemory,
		DataPath: fmt.Sprintf("s3://%s", bucket),
	})
	if err != nil {
		panic(err)
	}
	var randomContent = make([]byte, 10)
	rand.Read(randomContent)
	signer := v4.NewSigner(credentials.NewEnvCredentials())
	fileUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, "test")
	request, _ := http.NewRequest("PUT", fileUrl, bytes.NewReader(randomContent))
	_, err = signer.Sign(request, bytes.NewReader(randomContent), "s3", "eu-west-3", time.Now())
	if err != nil {
		panic(err)
	}
	res, _ := http.DefaultClient.Do(request)
	if res.StatusCode != 200 {
		panic(res.Status)
	}
	err = s3Reader.Delete([]string{fileUrl})
	if err != nil {
		t.Fail()
	}
	request, _ = http.NewRequest("GET", fileUrl, nil)
	_, err = signer.Sign(request, nil, "s3", "eu-west-3", time.Now())
	if err != nil {
		panic(err)
	}
	res, _ = http.DefaultClient.Do(request)
	if res.StatusCode == 200 {
		t.Fatal("Object still exists")
	}
}
