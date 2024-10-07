package client

import (
	"encoding/json"
	"fmt"
	"github.com/atko-pam/platform/backend/labels"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

type DatabaseBasedResourceSubSelectorType string

const (
	DatabaseBasedResourceSelectorType     = ResourceSelectorType("database_based_resource")
	DatabaseLabelsDatabaseSubSelectorType = DatabaseBasedResourceSubSelectorType("database_labels")
	PasswordCheckoutDatabasePrivilegeType = PrivilegeType("password_checkout_database")
)

// Privileges: password_checkout_database

var _ SecurityPolicyRulePrivilege = &PasswordCheckoutDatabasePrivilege{}

type PasswordCheckoutDatabasePrivilege struct {
	Enabled *bool `json:"password_checkout_database"`
}

func (p *PasswordCheckoutDatabasePrivilege) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == DatabaseBasedResourceSelectorType
}

func (p *PasswordCheckoutDatabasePrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.Enabled] = *p.Enabled
	return m
}

// Sub-Selectors: database_labels

type DatabaseLabelsDatabaseSelector struct {
	Labels labels.Labels `json:"labels"`
}

var _ DatabaseBasedResourceSubSelector = &DatabaseLabelsBasedSubSelector{}

type DatabaseLabelsBasedSubSelector struct {
	DatabaseSelector *DatabaseLabelsDatabaseSelector `json:"database_selector"`
}

func (*DatabaseLabelsBasedSubSelector) DatabaseBasedResourceSubSelectorType() DatabaseBasedResourceSubSelectorType {
	return DatabaseLabelsDatabaseSubSelectorType
}

func (d *DatabaseLabelsBasedSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)

	databaseLabelsM := make(map[string]any, len(d.DatabaseSelector.Labels))
	for k, v := range d.DatabaseSelector.Labels {
		databaseLabelsM[k] = v
	}
	m[attributes.DatabaseLabels] = databaseLabelsM

	return m
}

func (d *DatabaseLabelsBasedSubSelector) UnmarshalJSON(data []byte) error {
	tmp := struct {
		DatabaseSelector *DatabaseLabelsDatabaseSelector `json:"database_selector"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	d.DatabaseSelector = tmp.DatabaseSelector
	return nil
}

var _ json.Unmarshaler = &DatabaseBasedResourceSubSelectorContainer{}

type DatabaseBasedResourceSubSelectorContainer struct {
	SelectorType DatabaseBasedResourceSubSelectorType `json:"selector_type"`
	Selector     DatabaseBasedResourceSubSelector     `json:"selector"`
}

func (c *DatabaseBasedResourceSubSelectorContainer) UnmarshalJSON(data []byte) error {
	tmp := struct {
		SelectorType DatabaseBasedResourceSubSelectorType `json:"selector_type"`
		Selector     json.RawMessage                      `json:"selector"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	c.SelectorType = tmp.SelectorType
	switch tmp.SelectorType {
	case DatabaseLabelsDatabaseSubSelectorType:
		c.Selector = &DatabaseLabelsBasedSubSelector{}
	default:
		return fmt.Errorf("received unknown sub-selector type: %s", tmp.SelectorType)
	}
	if err := json.Unmarshal(tmp.Selector, c.Selector); err != nil {
		return err
	}

	return nil
}

var _ SecurityPolicyRuleResourceSelector = &DatabaseBasedResourceSelector{}

type DatabaseBasedResourceSelector struct {
	Selectors []DatabaseBasedResourceSubSelectorContainer `json:"selectors"`
}

func (s *DatabaseBasedResourceSelector) ResourceSelectorType() ResourceSelectorType {
	return DatabaseBasedResourceSelectorType
}

func (s *DatabaseBasedResourceSelector) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)

	databasesArr := make([]any, 1)
	subSelectorsMap := make(map[string]any, 3)
	databasesArr[0] = subSelectorsMap

	databaseLabelsArr := make([]any, 0, 1)

	for _, subSelector := range s.Selectors {
		switch subSelector.SelectorType {
		case DatabaseLabelsDatabaseSubSelectorType:
			databaseLabelsArr = append(databaseLabelsArr, subSelector.Selector.ToResourceMap())
		}
	}

	subSelectorsMap[attributes.LabelSelectors] = databaseLabelsArr
	m[attributes.Databases] = databasesArr
	return m
}

type DatabaseBasedResourceSubSelector interface {
	DatabaseBasedResourceSubSelectorType() DatabaseBasedResourceSubSelectorType
	ToResourceMap() map[string]any
}
