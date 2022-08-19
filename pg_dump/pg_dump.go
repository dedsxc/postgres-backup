package pg_dump

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type Psql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database []string
	Filename string
}

func New(host string, port string, username string, password string, db []string) Psql {
	var obj Psql
	obj.Host = host
	obj.Port = port
	obj.Username = username
	obj.Password = password
	obj.Database = db
	return obj
}

func (pg *Psql) Backup(datetime string) []string {
	var filename []string
	if len(pg.Database[0]) > 0 {
		for _, dbName := range pg.Database {
			pg.Filename = fmt.Sprintf("%s_%s.dump", datetime, dbName)
			var cmd = exec.Command("pg_dump", "-h", pg.Host, "-p", pg.Port, "-U", pg.Username, "-d", dbName, "--no-password", "-f", pg.Filename)
			logrus.Infof("[pg_dump][%s] Start backup", dbName)
			var res = BackupCmd(cmd)
			if res < 0 {
				logrus.Errorf("[pg_dump] Error during backup db: %s", dbName)
			} else {
				filename = append(filename, pg.Filename)
				logrus.Infof("[pg_dump][%s] Backup success", dbName)
			}
		}
	} else {
		pg.Filename = fmt.Sprintf("%s_ALL.dump", datetime)
		var cmd = exec.Command("pg_dumpall", "-h", pg.Host, "-p", pg.Port, "-U", pg.Username, "--no-password", "-f", pg.Filename)
		logrus.Info("[pg_dump][all] Start backup")
		var res = BackupCmd(cmd)
		if res < 0 {
			logrus.Error("[pg_dump] Error during backup all db\n")
		} else {
			filename = append(filename, pg.Filename)
			logrus.Info("[pg_dump][all] Backup success")
		}
	}
	return filename
}

func BackupCmd(cmd *exec.Cmd) int {
	var stderr strings.Builder
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		logrus.Error("[pg_dump] error during BackupCmd()")
		logrus.Errorf("[pg_dump] %s", stderr.String())
		return -1
	}
	return 0
}
