package typed_strings

type UserStatus string
type UserType string
type ADCertificateType string

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusDisabled UserStatus = "DISABLED"
	UserStatusDeleted  UserStatus = "DELETED"

	UserTypeHuman   UserType = "human"
	UserTypeService UserType = "service"

	ADCertificateTypeSigningRequest ADCertificateType = "certificate_signing_request"
	ADCertificateTypeSelfSigned     ADCertificateType = "self_signed"
)

func (ut UserType) String() string {
	return string(ut)
}

func (ut UserStatus) String() string {
	return string(ut)
}

func (certType ADCertificateType) String() string {
	return string(certType)
}
