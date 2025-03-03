// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"
)

func CloudProjectRancherDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"created_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Date of the managed Rancher service creation",
			MarkdownDescription: "Date of the managed Rancher service creation",
		},
		"current_state": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"bootstrap_password": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Sensitive:           true,
					Description:         "Bootstrap password of the managed Rancher service, returned only on creation",
					MarkdownDescription: "Bootstrap password of the managed Rancher service, returned only on creation",
				},
				"ip_restrictions": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"cidr_block": schema.StringAttribute{
								CustomType:          ovhtypes.TfStringType{},
								Computed:            true,
								Description:         "Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)",
								MarkdownDescription: "Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)",
							},
							"description": schema.StringAttribute{
								CustomType:          ovhtypes.TfStringType{},
								Computed:            true,
								Description:         "Description of the allowed CIDR block",
								MarkdownDescription: "Description of the allowed CIDR block",
							},
						},
						CustomType: CurrentStateIpRestrictionsType{
							ObjectType: types.ObjectType{
								AttrTypes: CurrentStateIpRestrictionsValue{}.AttributeTypes(ctx),
							},
						},
					},
					CustomType:          ovhtypes.NewTfListNestedType[CurrentStateIpRestrictionsValue](ctx),
					Computed:            true,
					Description:         "List of allowed CIDR blocks for a managed Rancher service's IP restrictions. When empty, any IP is allowed",
					MarkdownDescription: "List of allowed CIDR blocks for a managed Rancher service's IP restrictions. When empty, any IP is allowed",
				},
				"name": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Name of the managed Rancher service",
					MarkdownDescription: "Name of the managed Rancher service",
				},
				"networking": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"egress_cidr_blocks": schema.ListAttribute{
							CustomType:          ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
							Computed:            true,
							Description:         "Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.",
							MarkdownDescription: "Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.",
						},
					},
					CustomType: CurrentStateNetworkingType{
						ObjectType: types.ObjectType{
							AttrTypes: CurrentStateNetworkingValue{}.AttributeTypes(ctx),
						},
					},
					Computed:            true,
					Description:         "Networking properties of a managed Rancher service",
					MarkdownDescription: "Networking properties of a managed Rancher service",
				},
				"plan": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Plan of the managed Rancher service",
					MarkdownDescription: "Plan of the managed Rancher service",
				},
				"region": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Region of the managed Rancher service",
					MarkdownDescription: "Region of the managed Rancher service",
				},
				"url": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "URL of the managed Rancher service",
					MarkdownDescription: "URL of the managed Rancher service",
				},
				"usage": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"datetime": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Computed:            true,
							Description:         "Date of the sample",
							MarkdownDescription: "Date of the sample",
						},
						"orchestrated_vcpus": schema.Int64Attribute{
							CustomType:          ovhtypes.TfInt64Type{},
							Computed:            true,
							Description:         "Total number of vCPUs orchestrated by the managed Rancher service through the downstream clusters",
							MarkdownDescription: "Total number of vCPUs orchestrated by the managed Rancher service through the downstream clusters",
						},
					},
					CustomType: CurrentStateUsageType{
						ObjectType: types.ObjectType{
							AttrTypes: CurrentStateUsageValue{}.AttributeTypes(ctx),
						},
					},
					Computed:            true,
					Description:         "Latest metrics regarding the usage of the managed Rancher service",
					MarkdownDescription: "Latest metrics regarding the usage of the managed Rancher service",
				},
				"version": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Version of the managed Rancher service",
					MarkdownDescription: "Version of the managed Rancher service",
				},
			},
			CustomType: RancherCurrentStateType{
				ObjectType: types.ObjectType{
					AttrTypes: RancherCurrentStateValue{}.AttributeTypes(ctx),
				},
			},
			Computed:            true,
			Description:         "Current configuration applied to the managed Rancher service",
			MarkdownDescription: "Current configuration applied to the managed Rancher service",
		},
		"current_tasks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Identifier of the current task",
						MarkdownDescription: "Identifier of the current task",
					},
					"link": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Link to the task details",
						MarkdownDescription: "Link to the task details",
					},
					"status": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Current global status of the current task",
						MarkdownDescription: "Current global status of the current task",
					},
					"type": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Type of the current task",
						MarkdownDescription: "Type of the current task",
					},
				},
				CustomType: RancherCurrentTasksType{
					ObjectType: types.ObjectType{
						AttrTypes: RancherCurrentTasksValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType:          ovhtypes.NewTfListNestedType[RancherCurrentTasksValue](ctx),
			Computed:            true,
			Description:         "Asynchronous operations ongoing on the managed Rancher service",
			MarkdownDescription: "Asynchronous operations ongoing on the managed Rancher service",
		},
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Unique identifier",
			MarkdownDescription: "Unique identifier",
		},
		"project_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Project ID",
			MarkdownDescription: "Project ID",
		},
		"resource_status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status",
			MarkdownDescription: "Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status",
		},
		"target_spec": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"ip_restrictions": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"cidr_block": schema.StringAttribute{
								CustomType:          ovhtypes.TfStringType{},
								Computed:            true,
								Description:         "Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)",
								MarkdownDescription: "Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)",
							},
							"description": schema.StringAttribute{
								CustomType:          ovhtypes.TfStringType{},
								Computed:            true,
								Description:         "Description of the allowed CIDR block",
								MarkdownDescription: "Description of the allowed CIDR block",
							},
						},
						CustomType: TargetSpecIpRestrictionsType{
							ObjectType: types.ObjectType{
								AttrTypes: TargetSpecIpRestrictionsValue{}.AttributeTypes(ctx),
							},
						},
					},
					CustomType:          ovhtypes.NewTfListNestedType[TargetSpecIpRestrictionsValue](ctx),
					Computed:            true,
					Description:         "List of allowed CIDR blocks for a managed Rancher service's IP restrictions. When empty, any IP is allowed",
					MarkdownDescription: "List of allowed CIDR blocks for a managed Rancher service's IP restrictions. When empty, any IP is allowed",
				},
				"name": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Name of the managed Rancher service",
					MarkdownDescription: "Name of the managed Rancher service",
				},
				"plan": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan",
					MarkdownDescription: "Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan",
				},
				"version": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/version",
					MarkdownDescription: "Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/version",
				},
			},
			CustomType: RancherTargetSpecType{
				ObjectType: types.ObjectType{
					AttrTypes: RancherTargetSpecValue{}.AttributeTypes(ctx),
				},
			},
			Computed:            true,
			Description:         "Last target specification of the managed Rancher service",
			MarkdownDescription: "Last target specification of the managed Rancher service",
		},
		"updated_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Date of the last managed Rancher service update",
			MarkdownDescription: "Date of the last managed Rancher service update",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}
