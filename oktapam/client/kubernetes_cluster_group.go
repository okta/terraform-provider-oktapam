package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type KubernetesClusterGroup struct {
	ID              *string             `json:"id,omitempty"`
	GroupName       *string             `json:"group_name,omitempty"`
	ClusterSelector *string             `json:"cluster_selector,omitempty"`
	Claims          map[string][]string `json:"claims,omitempty"`
}

func (t KubernetesClusterGroup) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	if t.ID != nil {
		m[attributes.ID] = *t.ID
	}

	if t.GroupName != nil {
		m[attributes.GroupID] = *t.GroupName
	}

	if t.ClusterSelector != nil {
		m[attributes.ClusterSelector] = *t.ClusterSelector
	}

	if t.Claims != nil {
		claimsOut := make(map[string]string)
		for k, values := range t.Claims {
			claimsOut[k] = ""
			for _, claimValue := range values {
				if len(claimsOut[k]) > 0 {
					claimsOut[k] += ","
				}
				claimsOut[k] += claimValue
			}
		}
		m[attributes.Claims] = claimsOut
	}

	return m
}

type KubernetesClusterGroupListResponse struct {
	ClusterGroups []KubernetesClusterGroup `json:"list"`
}

func (c OktaPAMClient) ListKubernetesClusterGroups(ctx context.Context) ([]KubernetesClusterGroup, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/cluster_groups", url.PathEscape(c.Team))
	clusterGroups := make([]KubernetesClusterGroup, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetResult(&KubernetesClusterGroupListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return nil, err
		}

		clusterGroupResponse := resp.Result().(*KubernetesClusterGroupListResponse)
		clusterGroups = append(clusterGroups, clusterGroupResponse.ClusterGroups...)

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

	return clusterGroups, nil
}

func (c OktaPAMClient) GetKubernetesClusterGroup(ctx context.Context, id string) (*KubernetesClusterGroup, error) {
	if id == "" {
		return nil, errors.New("supplied blank kubernetes cluster group id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/cluster_groups/%s", url.PathEscape(c.Team), url.PathEscape(id))

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&KubernetesClusterGroup{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	if statusCode, err := checkStatusCode(resp, 200, 404); err != nil {
		return nil, err
	} else if statusCode == 404 {
		return nil, nil
	}

	clusterGroup := resp.Result().(*KubernetesClusterGroup)

	return clusterGroup, nil
}

func (c OktaPAMClient) CreateKubernetesClusterGroup(ctx context.Context, clusterGroup KubernetesClusterGroup) (*KubernetesClusterGroup, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/cluster_groups", url.PathEscape(c.Team))

	logging.Tracef("making POST request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(clusterGroup).SetResult(&KubernetesClusterGroup{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	createdClusterGroup := resp.Result().(*KubernetesClusterGroup)

	return createdClusterGroup, nil
}

func (c OktaPAMClient) UpdateKubernetesClusterGroup(ctx context.Context, clusterGroupID string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/cluster_groups/%s", url.PathEscape(c.Team), url.PathEscape(clusterGroupID))
	logging.Tracef("making PUT request to %s -> %+v", requestURL, updates)

	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaPAMClient) DeleteKubernetesClusterGroup(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("supplied blank kubernetes cluster group id")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/kubernetes/cluster_groups/%s",
		url.PathEscape(c.Team), url.PathEscape(id))

	logging.Tracef("making DELETE request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err := checkStatusCode(resp, 204, 404); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return err
	}

	return nil
}
