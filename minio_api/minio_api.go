package minio_api

import (
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

type Minio struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Filename  string
}

func New(endpoint string, accessKey string, secretKey string, filename string) Minio {
	var obj Minio
	obj.Endpoint = endpoint
	obj.AccessKey = accessKey
	obj.SecretKey = secretKey
	obj.Filename = filename
	return obj
}

func (m *Minio) UploadS3Server(pgHost string) {
	var bucketName = fmt.Sprintf("postgres-%s", strings.Split(pgHost, ".")[0])
	var ctx = context.Background()
	minioClient, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKey, m.SecretKey, ""),
		Secure: m.UseSSL,
	})
	if err != nil {
		logrus.Errorf("[minio_api] %s", err)
	}
	CreateBucketIfNotExist(ctx, minioClient, bucketName)
	var filePath = fmt.Sprintf("./%s", m.Filename)
	Upload(ctx, minioClient, bucketName, m.Filename, filePath)
}

func CreateBucketIfNotExist(ctx context.Context, minioClient *minio.Client, bucketName string) {
	var location = "us-east-1"
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			logrus.Infof("[minio_api] bucket %s already exist", bucketName)
		} else {
			logrus.Errorf("[minio_api] %s", err)
		}
	} else {
		logrus.Infof("[minio_api] bucket created %s", bucketName)
	}
}

func Upload(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, filePath string) {
	var contentType = "application/octet-stream"
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logrus.Errorf("[minio_api] %s", err)
	} else {
		logrus.Infof("[minio_api] Successfully uploaded %s of size %d in %s", objectName, info.Size, info.Bucket)
	}
}
