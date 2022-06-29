package oktapam

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/client"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/descriptions"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
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
			},
			attributes.ID: {
				Type:     schema.TypeString,
				Computed: true,
				// Description is autogenerated
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
			attributes.Status: {
				Type:        schema.TypeString,
				Optional:    true, // TODO: Check this
				Description: descriptions.Status,
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
	userName := getStringPtr(attributes.Name, d, true)
	userType := getStringPtr(attributes.UserType, d, false) // TODO: user type must be set

	switch *userType {
	case string(client.UserTypeHuman):
		err := c.CreateHumanUser(ctx, *userName)
		if err != nil {
			return diag.FromErr(err)
		}
	case string(client.UserTypeService):
		err := c.CreateServiceUser(ctx, *userName)
		if err != nil {
			return diag.FromErr(err)
		}
	default:
		return diag.Errorf(errors.InvalidUserTypeError, *userType)
	}
	d.SetId(*userName)
	return resourceUserRead(ctx, d, m)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)

	userName := d.Id()
	userType := getStringPtr(attributes.UserType, d, true)

	var user *client.User
	var err error
	switch *userType {
	case string(client.UserTypeHuman):
		user, err = c.GetHumanUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	case string(client.UserTypeService):
		user, err = c.GetServiceUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	case "":
		// The human endpoint can return either service or type, this will be used in import cases
		user, err = c.GetHumanUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	default:
		return diag.Errorf(errors.InvalidUserTypeError, *userType)
	}

	// Allow for deleted users, as they are soft-deleted // TODO: Check this comment
	if user != nil {
		d.SetId(*user.Name)
		for key, value := range user.ToResourceMap() {
			d.Set(key, value)
		}
	} else {
		logging.Infof("service user %s does not exist", userName)
	}

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(client.OktaPAMClient)
	userName := d.Id()
	userType := getStringPtr(attributes.UserType, d, true) // TODO: user type must be set

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
			return diag.Errorf("could not create service user from supplied values")
		}

		switch *userType {
		case string(client.UserTypeHuman):
			err := c.UpdateHumanUser(ctx, userName, su)
			if err != nil {
				return diag.FromErr(err)
			}
		case string(client.UserTypeService):
			err := c.UpdateServiceUser(ctx, userName, su)
			if err != nil {
				return diag.FromErr(err)
			}
		case "":
			return diag.Errorf(errors.MissingUserTypeError)
		default:
			return diag.Errorf(errors.InvalidUserTypeError, *userType)
		}
	}
	d.SetId(userName)
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(client.OktaPAMClient)
	userName := d.Id()
	userType := getStringPtr(attributes.UserType, d, true)

	switch *userType {
	case string(client.UserTypeHuman):
		err := c.DeleteHumanUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	case string(client.UserTypeService):
		err := c.DeleteServiceUser(ctx, userName)
		if err != nil {
			return diag.FromErr(err)
		}
	case "":
		return diag.Errorf(errors.MissingUserTypeError)
	default:
		return diag.Errorf(errors.InvalidUserTypeError, *userType)
	}

	d.SetId("")
	return diags
}
