package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

func IfExistEnvVariable() int {
	envList := [7]string{"POSTGRES_HOST", "POSTGRES_USER", "PGPASSWORD", "MINIO_ENDPOINT", "MINIO_ACCESS_KEY", "MINIO_SECRET_KEY", "TIMER_MIN"}
	for i := 0; i < len(envList); i++ {
		if _, isVar := os.LookupEnv(envList[i]); isVar {
			continue
		} else {
			logrus.Errorf("[env] Missing environment variable: %s", envList[i])
			return -1
		}
	}
	return 0
}
