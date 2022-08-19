package main

import (
	"fmt"
	"os"
	"postgres-backup/checksum"
	"postgres-backup/env"
	"postgres-backup/minio_api"
	"postgres-backup/pg_dump"
	"strconv"
	"strings"
	"time"
)

var (
	psqlHost       = os.Getenv("POSTGRES_HOST")
	psqlPort       = os.Getenv("POSTGRES_PORT")
	psqlDb         = os.Getenv("POSTGRES_DB")
	psqlUser       = os.Getenv("POSTGRES_USER")
	psqlPassword   = os.Getenv("PGPASSWORD")
	minioEnpoint   = os.Getenv("MINIO_ENDPOINT")
	minioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	minioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	timerMin       = os.Getenv("TIMER_MIN")
)

func main() {
	if env.IfExistEnvVariable() > -1 {
		var sum = make(map[string][]byte)
		var toIntTimer, _ = strconv.Atoi(timerMin)
		var pg = pg_dump.New(psqlHost, psqlPort, psqlUser, psqlPassword, strings.Split(psqlDb, ","))
		for {
			var datetime = time.Now().Format("01-02-2006-15:04")
			var pgFiles = pg.Backup(datetime)
			for _, file := range pgFiles {
				var filePath = fmt.Sprintf("./%s", file)
				csum, dbname := checksum.GetChecksum(file)
				if !checksum.CompareChecksum(sum[dbname], csum) {
					sum[dbname] = csum
					var m = minio_api.New(minioEnpoint, minioAccessKey, minioSecretKey, file)
					m.UploadS3Server(psqlHost)
				}
				os.Remove(filePath)
			}
			time.Sleep(time.Duration(toIntTimer) * time.Minute)
		}
	}
	os.Exit(0)
}
