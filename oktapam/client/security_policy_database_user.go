package client

import "github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

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

func (d *DatabaseLabelsBasedSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)

	databaseLabelsM := make(map[string]any, len(d.DatabaseSelector.Labels))
	for k, v := range d.DatabaseSelector.Labels {
		databaseLabelsM[k] = v
	}
	m[attributes.DatabaseLabels] = databaseLabelsM

	return m
}

func (p *PasswordCheckoutDatabasePrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.Enabled] = *p.Enabled
	return m
}
