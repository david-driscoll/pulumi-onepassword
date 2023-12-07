// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDocument(ctx *pulumi.Context, args *GetDocumentArgs, opts ...pulumi.InvokeOption) (*GetDocumentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDocumentResult
	err := ctx.Invoke("one-password-native-unofficial:index:GetDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDocumentArgs struct {
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id *string `pulumi:"id"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetDocumentResult struct {
	Attachments map[string]OutputAttachment `pulumi:"attachments"`
	Category    string                      `pulumi:"category"`
	Fields      map[string]OutputField      `pulumi:"fields"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id         string                   `pulumi:"id"`
	Notes      *string                  `pulumi:"notes"`
	References []OutputReference        `pulumi:"references"`
	Sections   map[string]OutputSection `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string            `pulumi:"title"`
	Urls  []OutputUrl       `pulumi:"urls"`
	Vault map[string]string `pulumi:"vault"`
}

func GetDocumentOutput(ctx *pulumi.Context, args GetDocumentOutputArgs, opts ...pulumi.InvokeOption) GetDocumentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDocumentResult, error) {
			args := v.(GetDocumentArgs)
			r, err := GetDocument(ctx, &args, opts...)
			var s GetDocumentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDocumentResultOutput)
}

type GetDocumentOutputArgs struct {
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentArgs)(nil)).Elem()
}

type GetDocumentResultOutput struct{ *pulumi.OutputState }

func (GetDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentResult)(nil)).Elem()
}

func (o GetDocumentResultOutput) ToGetDocumentResultOutput() GetDocumentResultOutput {
	return o
}

func (o GetDocumentResultOutput) ToGetDocumentResultOutputWithContext(ctx context.Context) GetDocumentResultOutput {
	return o
}

func (o GetDocumentResultOutput) Attachments() OutputAttachmentMapOutput {
	return o.ApplyT(func(v GetDocumentResult) map[string]OutputAttachment { return v.Attachments }).(OutputAttachmentMapOutput)
}

func (o GetDocumentResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetDocumentResultOutput) Fields() OutputFieldMapOutput {
	return o.ApplyT(func(v GetDocumentResult) map[string]OutputField { return v.Fields }).(OutputFieldMapOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetDocumentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDocumentResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDocumentResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetDocumentResultOutput) References() OutputReferenceArrayOutput {
	return o.ApplyT(func(v GetDocumentResult) []OutputReference { return v.References }).(OutputReferenceArrayOutput)
}

func (o GetDocumentResultOutput) Sections() OutputSectionMapOutput {
	return o.ApplyT(func(v GetDocumentResult) map[string]OutputSection { return v.Sections }).(OutputSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetDocumentResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDocumentResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetDocumentResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetDocumentResultOutput) Urls() OutputUrlArrayOutput {
	return o.ApplyT(func(v GetDocumentResult) []OutputUrl { return v.Urls }).(OutputUrlArrayOutput)
}

func (o GetDocumentResultOutput) Vault() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDocumentResult) map[string]string { return v.Vault }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDocumentResultOutput{})
}
