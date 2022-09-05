package cert

import (
  "os"
  "encoding/base64"

	"github.com/wolffberg/admission-debugger/pkg/k8s"
)

func signCertificate(hostname string) *SelfSignedCertificate {
	ca := newCACertificate()
	return newSelfSignedCertificate(ca)
}

func CreateCertFiles(hostname string) {
  cert := signCertificate(hostname)

  err := os.WriteFile("./tls.crt", cert.CertificatePEM.Bytes(), 0644)
  err = os.WriteFile("./tls.key", cert.PrivateKeyPEM.Bytes(), 0644)
  if err != nil {
    panic("Could not create TLS certificate!")
  }
}

func CreateCertSecret(hostname string) error {
  cert := signCertificate(hostname)

	data := make(map[string]string)
	data["tls.crt"] = base64.StdEncoding.EncodeToString(cert.CertificatePEM.Bytes())
	data["tls.key"] = base64.StdEncoding.EncodeToString(cert.PrivateKeyPEM.Bytes())

	return k8s.NewSecret(data)
}

