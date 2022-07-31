package typed_strings

type ADCertificateType string

const (
	ADCertificateTypeSigningRequest ADCertificateType = "certificate_signing_request"
	ADCertificateTypeSelfSigned     ADCertificateType = "self_signed"
)

func (certType ADCertificateType) String() string {
	return string(certType)
}
