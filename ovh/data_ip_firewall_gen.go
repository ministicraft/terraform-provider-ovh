// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	ovhtypes "github.com/ovh/terraform-provider-ovh/v2/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IpFirewallDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				CustomType: ovhtypes.TfBoolType{},
				Computed:   true,
			},
			"ip": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Required:            true,
				Description:         "IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)",
				MarkdownDescription: "IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)",
			},
			"ip_on_firewall": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Required:            true,
				Description:         "IPv4 address (e.g., 192.0.2.0)",
				MarkdownDescription: "IPv4 address (e.g., 192.0.2.0)",
			},
			"state": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "Current state of your ip on firewall",
				MarkdownDescription: "Current state of your ip on firewall",
			},
		},
	}
}
