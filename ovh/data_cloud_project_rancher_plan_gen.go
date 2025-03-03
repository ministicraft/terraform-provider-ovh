// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudProjectRancherPlanDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"plans": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"cause": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Cause for an unavailability",
						MarkdownDescription: "Cause for an unavailability",
					},
					"message": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Human-readable description of the unavailability cause",
						MarkdownDescription: "Human-readable description of the unavailability cause",
					},
					"name": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Name of the plan",
						MarkdownDescription: "Name of the plan",
					},
					"status": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Status of the plan",
						MarkdownDescription: "Status of the plan",
					},
				},
				CustomType: CloudProjectRancherPlanType{
					ObjectType: types.ObjectType{
						AttrTypes: CloudProjectRancherPlanValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType: ovhtypes.NewTfListNestedType[CloudProjectRancherPlanValue](ctx),
			Computed:   true,
		},
		"project_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Project ID",
			MarkdownDescription: "Project ID",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type CloudProjectRancherPlanModel struct {
	Plans     ovhtypes.TfListNestedValue[CloudProjectRancherPlanValue] `tfsdk:"plans" json:"plans"`
	ProjectId ovhtypes.TfStringValue                                   `tfsdk:"project_id" json:"projectId"`
}

var _ basetypes.ObjectTypable = CloudProjectRancherPlanType{}

type CloudProjectRancherPlanType struct {
	basetypes.ObjectType
}

func (t CloudProjectRancherPlanType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectRancherPlanType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectRancherPlanType) String() string {
	return "CloudProjectRancherPlanType"
}

func (t CloudProjectRancherPlanType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	causeAttribute, ok := attributes["cause"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cause is missing from object`)

		return nil, diags
	}

	causeVal, ok := causeAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cause expected to be ovhtypes.TfStringValue, was: %T`, causeAttribute))
	}

	messageAttribute, ok := attributes["message"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`message is missing from object`)

		return nil, diags
	}

	messageVal, ok := messageAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`message expected to be ovhtypes.TfStringValue, was: %T`, messageAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return nil, diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectRancherPlanValue{
		Cause:   causeVal,
		Message: messageVal,
		Name:    nameVal,
		Status:  statusVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectRancherPlanValueNull() CloudProjectRancherPlanValue {
	return CloudProjectRancherPlanValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectRancherPlanValueUnknown() CloudProjectRancherPlanValue {
	return CloudProjectRancherPlanValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectRancherPlanValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectRancherPlanValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectRancherPlanValue Attribute Value",
				"While creating a CloudProjectRancherPlanValue value, a missing attribute value was detected. "+
					"A CloudProjectRancherPlanValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectRancherPlanValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectRancherPlanValue Attribute Type",
				"While creating a CloudProjectRancherPlanValue value, an invalid attribute value was detected. "+
					"A CloudProjectRancherPlanValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectRancherPlanValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectRancherPlanValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectRancherPlanValue Attribute Value",
				"While creating a CloudProjectRancherPlanValue value, an extra attribute value was detected. "+
					"A CloudProjectRancherPlanValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectRancherPlanValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectRancherPlanValueUnknown(), diags
	}

	causeAttribute, ok := attributes["cause"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cause is missing from object`)

		return NewCloudProjectRancherPlanValueUnknown(), diags
	}

	causeVal, ok := causeAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cause expected to be ovhtypes.TfStringValue, was: %T`, causeAttribute))
	}

	messageAttribute, ok := attributes["message"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`message is missing from object`)

		return NewCloudProjectRancherPlanValueUnknown(), diags
	}

	messageVal, ok := messageAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`message expected to be ovhtypes.TfStringValue, was: %T`, messageAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewCloudProjectRancherPlanValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewCloudProjectRancherPlanValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectRancherPlanValueUnknown(), diags
	}

	return CloudProjectRancherPlanValue{
		Cause:   causeVal,
		Message: messageVal,
		Name:    nameVal,
		Status:  statusVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectRancherPlanValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectRancherPlanValue {
	object, diags := NewCloudProjectRancherPlanValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewCloudProjectRancherPlanValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectRancherPlanType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectRancherPlanValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectRancherPlanValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectRancherPlanValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewCloudProjectRancherPlanValueMust(CloudProjectRancherPlanValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectRancherPlanType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectRancherPlanValue{}
}

var _ basetypes.ObjectValuable = CloudProjectRancherPlanValue{}

type CloudProjectRancherPlanValue struct {
	Cause   ovhtypes.TfStringValue `tfsdk:"cause" json:"cause"`
	Message ovhtypes.TfStringValue `tfsdk:"message" json:"message"`
	Name    ovhtypes.TfStringValue `tfsdk:"name" json:"name"`
	Status  ovhtypes.TfStringValue `tfsdk:"status" json:"status"`
	state   attr.ValueState
}

func (v *CloudProjectRancherPlanValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectRancherPlanValue CloudProjectRancherPlanValue

	var tmp JsonCloudProjectRancherPlanValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Cause = tmp.Cause
	v.Message = tmp.Message
	v.Name = tmp.Name
	v.Status = tmp.Status

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectRancherPlanValue) MergeWith(other *CloudProjectRancherPlanValue) {

	if (v.Cause.IsUnknown() || v.Cause.IsNull()) && !other.Cause.IsUnknown() {
		v.Cause = other.Cause
	}

	if (v.Message.IsUnknown() || v.Message.IsNull()) && !other.Message.IsUnknown() {
		v.Message = other.Message
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectRancherPlanValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"cause":   v.Cause,
		"message": v.Message,
		"name":    v.Name,
		"status":  v.Status,
	}
}
func (v CloudProjectRancherPlanValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 4)

	var val tftypes.Value
	var err error

	attrTypes["cause"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["message"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 4)

		val, err = v.Cause.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["cause"] = val

		val, err = v.Message.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["message"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Status.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["status"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v CloudProjectRancherPlanValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectRancherPlanValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectRancherPlanValue) String() string {
	return "CloudProjectRancherPlanValue"
}

func (v CloudProjectRancherPlanValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"cause":   ovhtypes.TfStringType{},
			"message": ovhtypes.TfStringType{},
			"name":    ovhtypes.TfStringType{},
			"status":  ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"cause":   v.Cause,
			"message": v.Message,
			"name":    v.Name,
			"status":  v.Status,
		})

	return objVal, diags
}

func (v CloudProjectRancherPlanValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectRancherPlanValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Cause.Equal(other.Cause) {
		return false
	}

	if !v.Message.Equal(other.Message) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	return true
}

func (v CloudProjectRancherPlanValue) Type(ctx context.Context) attr.Type {
	return CloudProjectRancherPlanType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectRancherPlanValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"cause":   ovhtypes.TfStringType{},
		"message": ovhtypes.TfStringType{},
		"name":    ovhtypes.TfStringType{},
		"status":  ovhtypes.TfStringType{},
	}
}
