package typed_strings

import (
	"fmt"
	"strings"
)

type SSHCertificateType string

const (
	CertTypeRsa      = "CERT_TYPE_RSA_01"
	CertTypeRsa256   = "CERT_TYPE_RSA_SHA2_256_01"
	CertTypeRsa512   = "CERT_TYPE_RSA_SHA2_512_01"
	CertTypeEd25519  = "CERT_TYPE_ED25519_01"
	CertTypeEcdsa521 = "CERT_TYPE_ECDSA_521_01"
	CertTypeEcdsa384 = "CERT_TYPE_ECDSA_384_01"
	CertTypeEcdsa256 = "CERT_TYPE_ECDSA_256_01"
)

func (certType SSHCertificateType) String() string {
	return string(certType)
}

// Valid Key Algorithms
var ValidCertTypes = []string{
	CertTypeEcdsa521,
	CertTypeEcdsa384,
	CertTypeEcdsa256,
	CertTypeEd25519,
	CertTypeRsa512,
	CertTypeRsa256,
	CertTypeRsa,
}

func SSHCertTypeListFormat() string {
	quoted := make([]string, len(ValidCertTypes))
	for i, certType := range ValidCertTypes {
		quoted[i] = fmt.Sprintf("`%s`", certType)
	}
	return strings.Join(quoted, ", ")
}
