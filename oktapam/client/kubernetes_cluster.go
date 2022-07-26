package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/tomnomnom/linkheader"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type KubernetesCluster struct {
	ID            *string           `json:"id,omitempty"`
	Key           *string           `json:"key,omitempty"`
	Auth          *string           `json:"auth,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	OIDCIssuerURL *string
}

func (t KubernetesCluster) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	if t.ID != nil {
		m[attributes.ID] = *t.ID
	}

	if t.Key != nil {
		m[attributes.KubernetesClusterKey] = *t.Key
	}

	if t.OIDCIssuerURL != nil {
		m[attributes.OIDCIssuerURL] = *t.OIDCIssuerURL
	}

	if t.Auth != nil {
		m[attributes.KubernetesAuthMechanism] = *t.Auth
	}

	if t.Labels != nil {
		m[attributes.Labels] = t.Labels
	}

	return m
}

func (c OktaPAMClient) CreateKubernetesCluster(ctx context.Context, cluster KubernetesCluster) (*KubernetesCluster, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(cluster).SetResult(&KubernetesCluster{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	createdCluster := resp.Result().(*KubernetesCluster)

	return createdCluster, nil
}

func (c OktaPAMClient) GetKubernetesCluster(ctx context.Context, id string) (*KubernetesCluster, error) {
	if id == "" {
		return nil, fmt.Errorf("supplied blank kubernetes cluster id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters/%s", url.PathEscape(c.Team), url.PathEscape(id))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&KubernetesCluster{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	if statusCode, err := checkStatusCode(resp, http.StatusOK, http.StatusNotFound); err != nil {
		return nil, err
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	cluster := resp.Result().(*KubernetesCluster)

	cleanLabels := make(map[string]string)

	for l, v := range cluster.Labels {
		if !strings.HasPrefix(l, "api.") {
			continue
		} else {
			cleanLabels[l[4:]] = v
		}
	}

	cluster.Labels = cleanLabels

	oidcIssuerURL := fmt.Sprintf("%s/v1/teams/%s/kubernetes/clusters/%s",
		c.client.BaseURL,
		url.PathEscape(c.Team),
		url.PathEscape(id))

	cluster.OIDCIssuerURL = &oidcIssuerURL

	return cluster, nil
}

func (c OktaPAMClient) UpdateKubernetesCluster(ctx context.Context, clusterID string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters/%s", url.PathEscape(c.Team), url.PathEscape(clusterID))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (c OktaPAMClient) DeleteKubernetesCluster(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("supplied blank kubernetes cluster id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters/%s", url.PathEscape(c.Team), url.PathEscape(id))
	logging.Tracef("making DELETE request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err := checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return err
	}

	return nil
}

type KubernetesClusterListResponse struct {
	Clusters []KubernetesCluster `json:"list"`
}

func (c OktaPAMClient) ListKubernetesClusters(ctx context.Context) ([]KubernetesCluster, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/clusters", url.PathEscape(c.Team))
	clusters := make([]KubernetesCluster, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetResult(&KubernetesClusterListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		clusterGroupResponse := resp.Result().(*KubernetesClusterListResponse)
		clusters = append(clusters, clusterGroupResponse.Clusters...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""

		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
				break
			}
		}
	}

	return clusters, nil
}
