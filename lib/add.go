package lib

import (
	"log"
)

type Config struct {
	ACMEDomains       []string
	AddLocal          string
	Embedded          bool
	KubeConfig        string
	HTTPListenPort    int
	HTTPSListenPort   int
	K8sMode           string
	Debug             bool
	NoCACerts         bool
	AuditLogPath      string
	AuditLogMaxage    int
	AuditLogMaxsize   int
	AuditLogMaxbackup int
	AuditLevel        int
	Features          string
}

func Run(f string) int {
	log.Println(f)
	return 0
}

func Add(f string) int {
        log.Println(f)
        return 0
}
