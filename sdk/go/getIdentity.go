// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/identity"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIdentity(ctx *pulumi.Context, args *GetIdentityArgs, opts ...pulumi.InvokeOption) (*GetIdentityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIdentityResult
	err := ctx.Invoke("one-password-native-unoffical:index:GetIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIdentityArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetIdentityResult struct {
	Address         *identity.AddressSection         `pulumi:"address"`
	Attachments     map[string]OutputAttachment      `pulumi:"attachments"`
	Category        string                           `pulumi:"category"`
	Fields          map[string]OutputField           `pulumi:"fields"`
	Identification  *identity.IdentificationSection  `pulumi:"identification"`
	InternetDetails *identity.InternetDetailsSection `pulumi:"internetDetails"`
	Notes           *string                          `pulumi:"notes"`
	References      []OutputReference                `pulumi:"references"`
	Sections        map[string]OutputSection         `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string      `pulumi:"title"`
	Urls  []OutputUrl `pulumi:"urls"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid  string            `pulumi:"uuid"`
	Vault map[string]string `pulumi:"vault"`
}

func GetIdentityOutput(ctx *pulumi.Context, args GetIdentityOutputArgs, opts ...pulumi.InvokeOption) GetIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIdentityResult, error) {
			args := v.(GetIdentityArgs)
			r, err := GetIdentity(ctx, &args, opts...)
			var s GetIdentityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIdentityResultOutput)
}

type GetIdentityOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityArgs)(nil)).Elem()
}

type GetIdentityResultOutput struct{ *pulumi.OutputState }

func (GetIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityResult)(nil)).Elem()
}

func (o GetIdentityResultOutput) ToGetIdentityResultOutput() GetIdentityResultOutput {
	return o
}

func (o GetIdentityResultOutput) ToGetIdentityResultOutputWithContext(ctx context.Context) GetIdentityResultOutput {
	return o
}

func (o GetIdentityResultOutput) Address() identity.AddressSectionPtrOutput {
	return o.ApplyT(func(v GetIdentityResult) *identity.AddressSection { return v.Address }).(identity.AddressSectionPtrOutput)
}

func (o GetIdentityResultOutput) Attachments() OutputAttachmentMapOutput {
	return o.ApplyT(func(v GetIdentityResult) map[string]OutputAttachment { return v.Attachments }).(OutputAttachmentMapOutput)
}

func (o GetIdentityResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetIdentityResultOutput) Fields() OutputFieldMapOutput {
	return o.ApplyT(func(v GetIdentityResult) map[string]OutputField { return v.Fields }).(OutputFieldMapOutput)
}

func (o GetIdentityResultOutput) Identification() identity.IdentificationSectionPtrOutput {
	return o.ApplyT(func(v GetIdentityResult) *identity.IdentificationSection { return v.Identification }).(identity.IdentificationSectionPtrOutput)
}

func (o GetIdentityResultOutput) InternetDetails() identity.InternetDetailsSectionPtrOutput {
	return o.ApplyT(func(v GetIdentityResult) *identity.InternetDetailsSection { return v.InternetDetails }).(identity.InternetDetailsSectionPtrOutput)
}

func (o GetIdentityResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIdentityResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetIdentityResultOutput) References() OutputReferenceArrayOutput {
	return o.ApplyT(func(v GetIdentityResult) []OutputReference { return v.References }).(OutputReferenceArrayOutput)
}

func (o GetIdentityResultOutput) Sections() OutputSectionMapOutput {
	return o.ApplyT(func(v GetIdentityResult) map[string]OutputSection { return v.Sections }).(OutputSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetIdentityResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIdentityResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetIdentityResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetIdentityResultOutput) Urls() OutputUrlArrayOutput {
	return o.ApplyT(func(v GetIdentityResult) []OutputUrl { return v.Urls }).(OutputUrlArrayOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetIdentityResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o GetIdentityResultOutput) Vault() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetIdentityResult) map[string]string { return v.Vault }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdentityResultOutput{})
}
