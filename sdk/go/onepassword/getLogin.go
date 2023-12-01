// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogin(ctx *pulumi.Context, args *GetLoginArgs, opts ...pulumi.InvokeOption) (*GetLoginResult, error) {
	var rv GetLoginResult
	err := ctx.Invoke("one-password-native-unoffical:index:GetLogin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLoginArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetLoginResult struct {
	Attachments map[string]OutField   `pulumi:"attachments"`
	Category    string                `pulumi:"category"`
	Fields      map[string]OutField   `pulumi:"fields"`
	Notes       *string               `pulumi:"notes"`
	Password    *string               `pulumi:"password"`
	References  map[string]OutField   `pulumi:"references"`
	Sections    map[string]OutSection `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title    string  `pulumi:"title"`
	Username *string `pulumi:"username"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func GetLoginOutput(ctx *pulumi.Context, args GetLoginOutputArgs, opts ...pulumi.InvokeOption) GetLoginResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLoginResult, error) {
			args := v.(GetLoginArgs)
			r, err := GetLogin(ctx, &args, opts...)
			var s GetLoginResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLoginResultOutput)
}

type GetLoginOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetLoginOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoginArgs)(nil)).Elem()
}

type GetLoginResultOutput struct{ *pulumi.OutputState }

func (GetLoginResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoginResult)(nil)).Elem()
}

func (o GetLoginResultOutput) ToGetLoginResultOutput() GetLoginResultOutput {
	return o
}

func (o GetLoginResultOutput) ToGetLoginResultOutputWithContext(ctx context.Context) GetLoginResultOutput {
	return o
}

func (o GetLoginResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v GetLoginResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o GetLoginResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoginResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetLoginResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetLoginResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetLoginResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoginResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLoginResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoginResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GetLoginResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetLoginResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetLoginResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetLoginResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetLoginResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoginResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetLoginResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoginResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetLoginResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoginResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetLoginResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoginResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o GetLoginResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoginResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoginResultOutput{})
}
