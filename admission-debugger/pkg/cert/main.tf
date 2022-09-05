package cert


func NewSelfSignedCertificate(hostname string) *SelfSignedCertificate {
	ca := newCACertificate()
	return newSelfSignedCertificate(ca)
}

