package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"net/url"
)

type KubernetesClusterConnection struct {
	ID                *string `json:"-"`
	APIURL            *string `json:"api_url,omitempty"`
	PublicCertificate *string `json:"public_certificate,omitempty"`
}

func (t KubernetesClusterConnection) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	if t.APIURL != nil {
		m[attributes.KubernetesAPIURL] = *t.APIURL
	}

	if t.PublicCertificate != nil {
		m[attributes.PublicCertificate] = *t.PublicCertificate
	}

	return m
}

func (c OktaPAMClient) CreateKubernetesClusterConnection(ctx context.Context, clusterID string, clusterConnection KubernetesClusterConnection) (*KubernetesClusterConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters/%s", url.PathEscape(c.Team), url.PathEscape(clusterID))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(clusterConnection).SetResult(&KubernetesClusterConnection{}).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	if _, err := checkStatusCode(resp, 204); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	createdConnection := resp.Result().(*KubernetesClusterConnection)

	return createdConnection, nil
}

func (c OktaPAMClient) GetKubernetesClusterConnection(ctx context.Context, clusterID string) (*KubernetesClusterConnection, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("supplied blank kubernetes cluster connection id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters/%s", url.PathEscape(c.Team), url.PathEscape(clusterID))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&KubernetesClusterConnection{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	if statusCode, err := checkStatusCode(resp, 200, 404); err != nil {
		return nil, err
	} else if statusCode == 404 {
		return nil, nil
	}

	clusterConnection := resp.Result().(*KubernetesClusterConnection)

	return clusterConnection, nil
}

func (c OktaPAMClient) UpdateKubernetesClusterConnection(ctx context.Context, clusterID string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters/%s", url.PathEscape(c.Team), url.PathEscape(clusterID))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaPAMClient) DeleteKubernetesClusterConnection(ctx context.Context, clusterID string) error {
	if clusterID == "" {
		return errors.New("supplied blank kubernetes cluster connection id")
	}

	emptyConnection := map[string]interface{}{
		"api_url":            "",
		"public_certificate": "",
	}

	return c.UpdateKubernetesCluster(ctx, clusterID, emptyConnection)
}
