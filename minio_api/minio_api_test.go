package minio_api

import "testing"

func TestUploadS3Server(t *testing.T) {
	m := Minio{
		Endpoint:  "localhost:9000",
		AccessKey: "minio",
		SecretKey: "miniosecret",
		Filename:  "test.txt",
		UseSSL:    false,
	}
	m.UploadS3Server("localhost")
}
