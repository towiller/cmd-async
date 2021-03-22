package bdbroker

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
)

type Options struct {
	Addrs     []string
	Topics    []string
	GroupId   string
	EnableTls bool
	TlsConfig *tls.Config
	OffsetId  int64
}

type Option func(*Options)

func Addrs(addrs []string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func OffsetId(offsetId int64) Option {
	return func(o *Options) {
		o.OffsetId = offsetId
	}
}

func Topics(topics []string) Option {
	return func(o *Options) {
		o.Topics = topics
	}
}

func GroupId(groupId string) Option {
	return func(o *Options) {
		o.GroupId = groupId
	}
}

func WithTls(clientPemPath string, clientKeyPath string, caPemPath string) Option {
	return func(o *Options) {
		o.EnableTls = true
		o.TlsConfig = configTLS(clientPemPath, clientKeyPath, caPemPath)
	}
}

func configTLS(clientPemPath string, clientKeyPath string, caPemPath string) *tls.Config {
	checkFile(clientPemPath)
	checkFile(clientKeyPath)
	checkFile(caPemPath)

	clientPem, err := tls.LoadX509KeyPair(clientPemPath, clientKeyPath)
	if err != nil {
		log.Panic(err)
	}

	caPem, err := ioutil.ReadFile(caPemPath)
	if err != nil {
		log.Panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caPem)
	return &tls.Config{
		Certificates:       []tls.Certificate{clientPem},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}
}

func checkFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Size() == 0 {
		panic("Please replace " + path + " with your own. ")
	}
}
