package oktapam

import (
	"context"
	"fmt"
	"strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Description:   descriptions.ResourceUser,
		Schema: map[string]*schema.Schema{
			attributes.Name: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.Name,
			},
			attributes.UserType: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: descriptions.UserType,
				ValidateFunc: validation.StringInSlice(
					[]string{
						typed_strings.UserTypeHuman.String(),
						typed_strings.UserTypeService.String(),
					}, false),
			},
			attributes.Status: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions.Status,
				ValidateFunc: validation.StringInSlice(
					[]string{
						typed_strings.UserStatusActive.String(),
						typed_strings.UserStatusDisabled.String(),
						typed_strings.UserStatusDeleted.String(),
					}, false),
			},
			attributes.TeamName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.TeamName,
			},
			attributes.ServerUserName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.ServerUserName,
			},
			attributes.DeletedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions.DeletedAt,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	userName, userType, err := getRequiredAttributes(d)
	if err != nil {
		return diag.FromErr(err)
	}

	switch typed_strings.UserType(userType) {
	case typed_strings.UserTypeHuman:
		err = c.CreateHumanUser(ctx, userName)
	case typed_strings.UserTypeService:
		err = c.CreateServiceUser(ctx, userName)
	default:
		return diag.Errorf(errors.InvalidUserTypeError, userType)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createUserID(userName, userType))
	return resourceUserRead(ctx, d, m)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	userName, userType, err := getRequiredAttributes(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var user *client.User
	switch typed_strings.UserType(userType) {
	case typed_strings.UserTypeHuman:
		user, err = c.GetHumanUser(ctx, userName)
	case typed_strings.UserTypeService:
		user, err = c.GetServiceUser(ctx, userName)
	default:
		return diag.Errorf(errors.InvalidUserTypeError, userType)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	if user != nil {
		d.SetId(createUserID(userName, userType))
		for key, value := range user.ToResourceMap() {
			d.Set(key, value)
		}
	} else {
		logging.Infof("%s user %s does not exist", userType, userName)
	}

	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)

	userName, userType, err := getRequiredAttributes(d)
	if err != nil {
		return diag.FromErr(err)
	}

	changed := false
	updates := make(map[string]interface{})

	// NOTE: possible that other fields must be included in the update
	changeableAttributes := []string{
		attributes.Status,
	}

	requiredAttributes := []string{
		attributes.Name,
	}

	for _, attribute := range changeableAttributes {
		if d.HasChange(attribute) {
			updates[attribute] = d.Get(attribute)
			changed = true
		}
	}

	for _, attribute := range requiredAttributes {
		updates[attribute] = d.Get(attribute)
	}

	if changed {
		su, err := client.UserFromMap(updates)
		if err != nil {
			return diag.FromErr(err)
		} else if su == nil {
			return diag.Errorf("could not update %s user from supplied values", userType)
		}

		switch typed_strings.UserType(userType) {
		case typed_strings.UserTypeHuman:
			err := c.UpdateHumanUser(ctx, userName, su)
			if err != nil {
				return diag.FromErr(err)
			}
		case typed_strings.UserTypeService:
			err := c.UpdateServiceUser(ctx, userName, su)
			if err != nil {
				return diag.FromErr(err)
			}
		default:
			return diag.Errorf(errors.InvalidUserTypeError, userType)
		}
	}
	d.SetId(createUserID(userName, userType))
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	userName, userType, err := getRequiredAttributes(d)
	if err != nil {
		return diag.FromErr(err)
	}

	switch typed_strings.UserType(userType) {
	case typed_strings.UserTypeHuman:
		err := c.DeleteHumanUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	case typed_strings.UserTypeService:
		err := c.DeleteServiceUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	default:
		return diag.Errorf(errors.InvalidUserTypeError, userType)
	}

	d.SetId("")
	return nil
}

func getRequiredAttributes(d *schema.ResourceData) (string, string, error) {
	if d.Id() != "" {
		return parseUserID(d.Id())
	}

	userName := getStringPtr(attributes.Name, d, false)
	if userName == nil {
		return "", "", fmt.Errorf(errors.MissingAttributeError, attributes.Name)
	}

	userType := getStringPtr(attributes.UserType, d, false)
	if userType == nil {
		return "", "", fmt.Errorf(errors.MissingAttributeError, attributes.UserType)
	}

	return *userName, *userType, nil
}

func createUserID(userName, userType string) string {
	return fmt.Sprintf("%s|%s", userName, userType)
}

func parseUserID(resourceId string) (string, string, error) {
	split := strings.Split(resourceId, "|")
	if len(split) != 2 {
		return "", "", fmt.Errorf("oktapam_user id must be in the format of <user name>|<user type>, received: %s", resourceId)
	}
	return split[0], split[1], nil
}
