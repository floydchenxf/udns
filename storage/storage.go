package storage

import (
	"github.com/miekg/dns"
	"strings"
	"udns/logger"
)

var (
	ds dnsStorage
)

type dnsStorage interface {
	Query(dnsType uint16, name string) ([]dns.RR, error)
	Shutdown()
}

func Init(dataSource string) {
	if strings.HasPrefix(dataSource, "file://") {
		ds = newFileStorage(dataSource[7:]) // file:///xxx/yyy/c.txt => /xxx/yyy/c.txt
	} else {
		logger.Fatal("storage", "unsupported data source %s", dataSource)
	}
}

func Query(dnsType uint16, name string) ([]dns.RR, error) {
	return ds.Query(dnsType, name)
}

func Shutdown() {
	ds.Shutdown()
}
