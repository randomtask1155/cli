package net

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

// GetNSSLogger returns an *io.Wri
func SetNSSLogger(tc *tls.Config) {
	filename := os.Getenv("CF_NSS_KEYLOGGER")
	if filename != "" {
		kl, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		tc.KeyLogWriter = kl
	}
}

func NewTLSConfig(trustedCerts []tls.Certificate, disableSSL bool) (TLSConfig *tls.Config) {
	TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS10,
	}
	SetNSSLogger(TLSConfig)

	if len(trustedCerts) > 0 {
		certPool := x509.NewCertPool()
		for _, tlsCert := range trustedCerts {
			cert, _ := x509.ParseCertificate(tlsCert.Certificate[0])
			certPool.AddCert(cert)
		}
		TLSConfig.RootCAs = certPool
	}

	TLSConfig.InsecureSkipVerify = disableSSL

	return
}
