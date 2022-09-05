package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

type SelfSignedCertificate struct {
	Certificate    *x509.Certificate
	PublicKey      *rsa.PublicKey
	PrivateKey     *rsa.PrivateKey
	PrivateKeyPEM  *bytes.Buffer
	CertificatePEM *bytes.Buffer
	CACertificate  *CACertificate
}

func newSelfSignedCertificate(ca *CACertificate) *SelfSignedCertificate {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{""},
			Country:       []string{""},
			Province:      []string{""},
			Locality:      []string{""},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivKey, _ := rsa.GenerateKey(rand.Reader, 4096)

	certBytes, _ := x509.CreateCertificate(rand.Reader, cert, ca.Certificate, &certPrivKey.PublicKey, ca.PrivateKey)

	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})

	return &SelfSignedCertificate{
		Certificate:    cert,
		PublicKey:      &certPrivKey.PublicKey,
		PrivateKey:     certPrivKey,
		PrivateKeyPEM:  certPrivKeyPEM,
		CertificatePEM: certPEM,
		CACertificate:  ca,
	}
}
