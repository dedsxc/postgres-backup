package pg_dump

import (
	"os"
	"testing"
	"time"
)

func TestBackup(t *testing.T) {
	var datetime = time.Now().Format("01-02-2006-15:04")
	p := Psql{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_HOST"),
		Database: []string{"postgres"}}
	p.Backup(datetime)
}
