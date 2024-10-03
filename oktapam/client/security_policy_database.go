package client

import (
	"encoding/json"
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

type DatabaseBasedResourceSubSelectorType string

const (
	DatabaseBasedResourceSelectorType    = ResourceSelectorType("database_based_resource")
	DatabaseLabelDatabaseSubSelectorType = DatabaseBasedResourceSubSelectorType("database_label")
)

type DatabaseLabelDatabaseSelector struct {
	Labels map[string]string `json:"labels"`
}

var _ DatabaseBasedResourceSubSelector = &DatabaseLabelBasedSubSelector{}

type DatabaseLabelBasedSubSelector struct {
	DatabaseSelector *DatabaseLabelDatabaseSelector `json:"database_selector"`

	AccountSelectorType AccountSelectorType `json:"account_selector_type"`
	AccountSelector     AccountSelector     `json:"account_selector"`
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

	individualDatabasesArr := make([]any, 0, len(s.Selectors))
	databaseLabelsArr := make([]any, 0, 1)

	for _, subSelector := range s.Selectors {
		switch subSelector.SelectorType {
		case DatabaseLabelDatabaseSubSelectorType:
			databaseLabelsArr = append(databaseLabelsArr, subSelector.Selector.ToResourceMap())
		}
	}

	subSelectorsMap[attributes.Database] = individualDatabasesArr

	m[attributes.Databases] = databasesArr
	return m
}

func (*DatabaseLabelBasedSubSelector) DatabaseBasedResourceSubSelectorType() DatabaseBasedResourceSubSelectorType {
	return DatabaseLabelDatabaseSubSelectorType
}

func (dbLabelSubSelector *DatabaseLabelBasedSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)

	databaseLabelsM := make(map[string]any, len(dbLabelSubSelector.DatabaseSelector.Labels))
	for k, v := range dbLabelSubSelector.DatabaseSelector.Labels {
		databaseLabelsM[k] = v
	}
	m[attributes.DatabaseLabels] = databaseLabelsM

	usernamesArr := make([]any, 0)

	if dbLabelSubSelector.AccountSelectorType == UsernameAccountSelectorType {
		usernamesArr = stringSliceToInterfaceSlice(dbLabelSubSelector.AccountSelector.(*UsernameAccountSelector).Usernames)
	}

	m[attributes.Accounts] = usernamesArr

	return m
}

func (dbLabelSubSelector *DatabaseLabelBasedSubSelector) UnmarshalJSON(data []byte) error {
	tmp := struct {
		DatabaseSelector *DatabaseLabelDatabaseSelector `json:"database_selector"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	dbLabelSubSelector.DatabaseSelector = tmp.DatabaseSelector
	return nil
}

type DatabaseBasedResourceSubSelector interface {
	DatabaseBasedResourceSubSelectorType() DatabaseBasedResourceSubSelectorType
	ToResourceMap() map[string]any
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
	case DatabaseLabelDatabaseSubSelectorType:
		c.Selector = &DatabaseLabelBasedSubSelector{}
	default:
		return fmt.Errorf("received unknown sub-selector type: %s", tmp.SelectorType)
	}
	if err := json.Unmarshal(tmp.Selector, c.Selector); err != nil {
		return err
	}

	return nil
}
