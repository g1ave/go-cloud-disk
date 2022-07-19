package tests

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestUploadFile(t *testing.T) {
	u, _ := url.Parse(testConfig.COS.BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  testConfig.COS.SecretId,
			SecretKey: testConfig.COS.SecretKey,
		},
	})

	key := "exampleobject.jpeg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./upload/file_to_upload.jpeg", nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartUploadInit(t *testing.T) {
	u, _ := url.Parse(testConfig.COS.BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  testConfig.COS.SecretId,
			SecretKey: testConfig.COS.SecretKey,
		},
	})
	name := "exampleobject"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), name, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID
	fmt.Println(UploadID) // 16582109544148b3f513ac65b6760910c3826b83806e35ceefeb17b6c634e41e97cc03dc0b
}

func TestMultipartUpload(t *testing.T) {
	u, _ := url.Parse(testConfig.COS.BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  testConfig.COS.SecretId,
			SecretKey: testConfig.COS.SecretKey,
		},
	})
	uploadID := "16582109544148b3f513ac65b6760910c3826b83806e35ceefeb17b6c634e41e97cc03dc0b"
	key := "exampleobject"
	f := strings.NewReader("test hello")
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadID, 1, f, nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag) // 0c9a8eac3e65ba7a73a4c7bbb25ba030
}

func TestMultipartUploadCompleted(t *testing.T) {
	u, _ := url.Parse(testConfig.COS.BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  testConfig.COS.SecretId,
			SecretKey: testConfig.COS.SecretKey,
		},
	})
	UploadID := "16582109544148b3f513ac65b6760910c3826b83806e35ceefeb17b6c634e41e97cc03dc0b"
	key := "exampleobject"
	PartETag := "0c9a8eac3e65ba7a73a4c7bbb25ba030"
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: PartETag},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
