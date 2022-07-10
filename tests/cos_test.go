package tests

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"testing"
)

func TestUploadFile(t *testing.T) {
	u, _ := url.Parse("https://cloud-disk-1312836572.cos.ap-chengdu.myqcloud.com")
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
