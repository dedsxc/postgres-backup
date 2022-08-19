package checksum

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetChecksum(filename string) ([]byte, string) {
	var dbName = strings.Split(filename, "_")[1]
	f, err := os.Open(filename)
	if err != nil {
		logrus.Error(err)
	}
	defer f.Close()

	var sum = sha256.New()
	if _, err := io.Copy(sum, f); err != nil {
		logrus.Fatal(err)
	}
	var checksum = sum.Sum(nil)
	logrus.Infof("[checksum] %x %s", checksum, dbName)
	return checksum, dbName
}

func CompareChecksum(sum1, sum2 []byte) bool {
	if bytes.Equal(sum1, sum2) {
		logrus.Infof("[checksum] checksum are identical. No upload needed")
		logrus.Infof("[checksum] checksum1=%x checksum2=%x", sum1, sum2)
		return true
	} else {
		logrus.Infof("[checksum] checksum are not identical. Upload started")
		logrus.Infof("[checksum] checksum1=%x checksum2=%x", sum1, sum2)
		return false
	}
}
