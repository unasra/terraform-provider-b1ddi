// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// InheritedDHCPConfigIgnoreItemList IgnoreItemList
//
// The inheritance configuration for a field that contains a list of _IgnoreItem_ objects
//
// swagger:model InheritedDHCPConfigIgnoreItemList
func dataSourceInheritedDHCPConfigIgnoreItemList() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance setting.
			//
			// Valid values are:
			// * _inherit_: Use the inherited value.
			// * _override_: Use the value set in the object.
			//
			// Defaults to _inherit_.
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The inheritance setting.\n\nValid values are:\n* _inherit_: Use the inherited value.\n* _override_: Use the value set in the object.\n\nDefaults to _inherit_.",
			},

			// The human-readable display name for the object referred to by _source_.
			// Read Only: true
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable display name for the object referred to by _source_.",
			},

			// The resource identifier.
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The inherited value.
			// Read Only: true
			"value": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcIgnoreItem(),
				Computed:    true,
				Description: "The inherited value.",
			},
		},
	}
}
