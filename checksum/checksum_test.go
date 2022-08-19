package checksum

import (
	"testing"
)

func TestGetChecksum(t *testing.T) {
	GetChecksum("08-15-2022-10:25_mission_localhost.dump")
}

func TestCompareChecksum(t *testing.T) {
	var c = make(map[string][]byte)
	checksum, dbname := GetChecksum("08-15-2022-10:25_mission_localhost.dump")
	if !CompareChecksum(c[dbname], checksum) {
		c[dbname] = checksum
	}
}
