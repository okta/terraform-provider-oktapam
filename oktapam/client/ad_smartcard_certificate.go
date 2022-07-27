package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

const (
	ADCertificateTypeSigningRequest = "certificate_signing_request"
	ADCertificateTypeSelfSigned     = "self_signed"
)

type ADSmartCardCertificate struct {
	ID               *string               `json:"id,omitempty"`
	DisplayName      *string               `json:"display_name,omitempty"`
	CommonName       *string               `json:"common_name,omitempty"`
	Type             *string               `json:"type,omitempty"`
	Details          *ADCertificateDetails `json:"details,omitempty"`
	Status           *string               `json:"status,omitempty"`
	EnterpriseSigned *bool                 `json:"enterprise_signed,omitempty"`
	CreatedAt        *time.Time            `json:"created_at,omitempty"`
	ExpiresAt        *time.Time            `json:"expires_at,omitempty"`
	Content          *string               `json:"content,omitempty"`
}

type ADCertificateDetails struct {
	Organization       *string `json:"organization,omitempty"`
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
	Locality           *string `json:"locality,omitempty"`
	Province           *string `json:"province,omitempty"`
	Country            *string `json:"country,omitempty"`
	TTLDays            *int64  `json:"ttl_days,omitempty"`
}

type UpdateADCertificateRequest struct {
	DisplayName *string `json:"display_name,omitempty"`
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
		flattenedCertDetails := make([]interface{}, 1)
		flattenedCertDetail := make(map[string]interface{})

		if t.Details.Organization != nil {
			flattenedCertDetail[attributes.Organization] = *t.Details.Organization
		}
		if t.Details.OrganizationalUnit != nil {
			flattenedCertDetail[attributes.OrganizationalUnit] = *t.Details.OrganizationalUnit
		}
		if t.Details.Locality != nil {
			flattenedCertDetail[attributes.Locality] = *t.Details.Locality
		}
		if t.Details.Province != nil {
			flattenedCertDetail[attributes.Province] = *t.Details.Province
		}
		if t.Details.Country != nil {
			flattenedCertDetail[attributes.Country] = *t.Details.Country
		}
		if t.Details.TTLDays != nil {
			flattenedCertDetail[attributes.TTLDays] = *t.Details.TTLDays
		}
		flattenedCertDetails[0] = flattenedCertDetail

		m[attributes.Details] = flattenedCertDetails
	}
	if t.Status != nil {
		m[attributes.Status] = *t.Status
	}
	if t.EnterpriseSigned != nil {
		m[attributes.EnterpriseSigned] = *t.EnterpriseSigned
	}
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
	if _, err := checkStatusCode(resp, http.StatusCreated); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}
	createdADCert := resp.Result().(*ADSmartCardCertificate)

	return createdADCert, nil
}

func (c OktaPAMClient) DeleteADSmartcardCertificate(ctx context.Context, certificateID string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/%s", url.PathEscape(c.Team), url.PathEscape(certificateID))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}

func (c OktaPAMClient) GetADSmartcardCertificate(ctx context.Context, certificateID string) (*ADSmartCardCertificate, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/%s", url.PathEscape(c.Team), url.PathEscape(certificateID))
	logging.Tracef("making GET request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetResult(&ADSmartCardCertificate{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		createdADCert := resp.Result().(*ADSmartCardCertificate)
		if createdADCert.Exists() {
			return createdADCert, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) UpdateADSmartcardCertificateName(ctx context.Context, certificateID string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/%s", url.PathEscape(c.Team), url.PathEscape(certificateID))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}


func (c OktaPAMClient) UploadADSmartcardCertificate(ctx context.Context, certificateId string, content string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/certificates/%s/upload", url.PathEscape(c.Team), url.PathEscape(certificateId))
	logging.Tracef("making GET request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetMultipartField("file", "cert-upload-tf", "multipart/form-data", strings.NewReader(content)).Post(requestURL)

	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusOK)
	return err
}
