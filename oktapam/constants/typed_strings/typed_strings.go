package typed_strings

type UserStatus string
type UserType string
type CertificateRequestType string

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusDisabled UserStatus = "DISABLED"
	UserStatusDeleted  UserStatus = "DELETED"

	UserTypeHuman   UserType = "human"
	UserTypeService UserType = "service"

	ADCertificateTypeSigningRequest CertificateRequestType = "certificate_signing_request"
	ADCertificateTypeSelfSigned     CertificateRequestType = "self_signed"
)

func (ut UserType) String() string {
	return string(ut)
}

func (ut UserStatus) String() string {
	return string(ut)
}

func (certType CertificateRequestType) String() string {
	return string(certType)
}
