package client

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

const (
	AD_CERTIFICATE_TYPE_SIGNING_REQUEST = "certificate_signing_request"
	AD_CERTIFICATE_TYPE_SELF_SIGNED     = "self_signed"
)

type ADSmartCardCertificate struct {
	ID               *string               `json:"id"`
	DisplayName      *string               `json:"display_name"`
	CommonName       *string               `json:"common_name"`
	Type             *string               `json:"type"`
	Details          *ADCertificateDetails `json:"details"`
	Status           *string               `json:"status"`
	EnterpriseSigned *bool                 `json:"enterprise_signed"`
	CreatedAt        *time.Time            `json:"created_at"`
	ExpiresAt        *time.Time            `json:"expires_at"`
	Content          *string               `json:"content"`
}

type ADCertificateDetails interface {
	CertType() string
}

type SelfSignedCertDetails struct {
	TTLDays int64
}

func (*SelfSignedCertDetails) CertType() string {
	return AD_CERTIFICATE_TYPE_SELF_SIGNED
}

type EnterpriseSignedCertDetails struct {
	Organization       string `json:"organization"`
	OrganizationalUnit string `json:"organizational_unit"`
	Locality           string `json:"locality"`
	Province           string `json:"province"`
	Country            string `json:"country"`
}

func (*EnterpriseSignedCertDetails) CertType() string {
	return AD_CERTIFICATE_TYPE_SIGNING_REQUEST
}

func (t ADSmartCardCertificate) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	if t.ID != nil {
		m[attributes.ID] = *t.ID
	}
	if t.DisplayName != nil {
		m[attributes.DisplayName] = *t.DisplayName
	}
	if t.CommonName != nil {
		m[attributes.CommonName] = *t.CommonName
	}
	if t.Details != nil {
		m[attributes.Details] = *t.Details
	}
	if t.Status != nil {
		m[attributes.Status] = *t.Status
	}
	if t.EnterpriseSigned != nil {
		m[attributes.EnterpriseSigned] = *t.EnterpriseSigned
	}
	//if t.CreatedAt != nil {
	//	m[attributes.CreatedAt] = *t.CreatedAt
	//}
	//if t.ExpiresAt != nil {
	//	m[attributes.ExpiresAt] = *t.ExpiresAt
	//}
	if t.Content != nil {
		m[attributes.Content] = *t.Content
	}
	return m
}

func (t ADSmartCardCertificate) Exists() bool {
	return utils.IsNonEmpty(t.ID)
}

func (c OktaPAMClient) CreateADSmartcardCertificate(ctx context.Context, adCert ADSmartCardCertificate) (*ADSmartCardCertificate, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/json", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adCert).SetResult(&ADSmartCardCertificate{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}
	createdADCert := resp.Result().(*ADSmartCardCertificate)

	return createdADCert, nil
}

func (c OktaPAMClient) DeleteADSmartcardCertificate(ctx context.Context, certificateId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/%s", url.PathEscape(c.Team), url.PathEscape(certificateId))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}

func (c OktaPAMClient) GetADSmartcardCertificate(ctx context.Context, certificateId string) (*ADSmartCardCertificate, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/%s", url.PathEscape(c.Team), url.PathEscape(certificateId))
	logging.Tracef("making GET request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetResult(&ADSmartCardCertificate{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		createdADCert := resp.Result().(*ADSmartCardCertificate)
		if createdADCert.Exists() {
			return createdADCert, nil
		}
		return nil, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}
