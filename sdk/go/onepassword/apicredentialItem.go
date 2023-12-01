// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type APICredentialItem struct {
	pulumi.CustomResourceState

	Attachments OutFieldMapOutput      `pulumi:"attachments"`
	Category    pulumi.StringOutput    `pulumi:"category"`
	Credential  pulumi.StringPtrOutput `pulumi:"credential"`
	Expires     pulumi.StringPtrOutput `pulumi:"expires"`
	Fields      OutFieldMapOutput      `pulumi:"fields"`
	Filename    pulumi.StringPtrOutput `pulumi:"filename"`
	Hostname    pulumi.StringPtrOutput `pulumi:"hostname"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	References  OutFieldMapOutput      `pulumi:"references"`
	Sections    OutSectionMapOutput    `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title    pulumi.StringOutput    `pulumi:"title"`
	Type     pulumi.StringPtrOutput `pulumi:"type"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid      pulumi.StringOutput    `pulumi:"uuid"`
	ValidFrom pulumi.StringPtrOutput `pulumi:"validFrom"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewAPICredentialItem registers a new resource with the given unique name, arguments, and options.
func NewAPICredentialItem(ctx *pulumi.Context,
	name string, args *APICredentialItemArgs, opts ...pulumi.ResourceOption) (*APICredentialItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("API Credential")
	if args.Credential != nil {
		args.Credential = pulumi.ToSecret(args.Credential).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"credential",
		"fields",
		"references",
		"sections",
	})
	opts = append(opts, secrets)
	var resource APICredentialItem
	err := ctx.RegisterResource("one-password-native-unoffical:index:APICredentialItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPICredentialItem gets an existing APICredentialItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPICredentialItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APICredentialItemState, opts ...pulumi.ResourceOption) (*APICredentialItem, error) {
	var resource APICredentialItem
	err := ctx.ReadResource("one-password-native-unoffical:index:APICredentialItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APICredentialItem resources.
type apicredentialItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type APICredentialItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (APICredentialItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*apicredentialItemState)(nil)).Elem()
}

type apicredentialItemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category   *string            `pulumi:"category"`
	Credential *string            `pulumi:"credential"`
	Expires    *string            `pulumi:"expires"`
	Fields     map[string]Field   `pulumi:"fields"`
	Filename   *string            `pulumi:"filename"`
	Hostname   *string            `pulumi:"hostname"`
	Notes      *string            `pulumi:"notes"`
	Sections   map[string]Section `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title     *string `pulumi:"title"`
	Type      *string `pulumi:"type"`
	Username  *string `pulumi:"username"`
	ValidFrom *string `pulumi:"validFrom"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a APICredentialItem resource.
type APICredentialItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category   pulumi.StringPtrInput
	Credential pulumi.StringPtrInput
	Expires    pulumi.StringPtrInput
	Fields     FieldMapInput
	Filename   pulumi.StringPtrInput
	Hostname   pulumi.StringPtrInput
	Notes      pulumi.StringPtrInput
	Sections   SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title     pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Username  pulumi.StringPtrInput
	ValidFrom pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (APICredentialItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apicredentialItemArgs)(nil)).Elem()
}

func (r *APICredentialItem) GetAttachment(ctx *pulumi.Context, args *APICredentialItemGetAttachmentArgs) (APICredentialItemGetAttachmentResultOutput, error) {
	out, err := ctx.Call("one-password-native-unoffical:index:APICredentialItem/attachment", args, APICredentialItemGetAttachmentResultOutput{}, r)
	if err != nil {
		return APICredentialItemGetAttachmentResultOutput{}, err
	}
	return out.(APICredentialItemGetAttachmentResultOutput), nil
}

type apicredentialItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the GetAttachment method of the APICredentialItem resource.
type APICredentialItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (APICredentialItemGetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apicredentialItemGetAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type APICredentialItemGetAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type APICredentialItemGetAttachmentResultOutput struct{ *pulumi.OutputState }

func (APICredentialItemGetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APICredentialItemGetAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o APICredentialItemGetAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v APICredentialItemGetAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type APICredentialItemInput interface {
	pulumi.Input

	ToAPICredentialItemOutput() APICredentialItemOutput
	ToAPICredentialItemOutputWithContext(ctx context.Context) APICredentialItemOutput
}

func (*APICredentialItem) ElementType() reflect.Type {
	return reflect.TypeOf((**APICredentialItem)(nil)).Elem()
}

func (i *APICredentialItem) ToAPICredentialItemOutput() APICredentialItemOutput {
	return i.ToAPICredentialItemOutputWithContext(context.Background())
}

func (i *APICredentialItem) ToAPICredentialItemOutputWithContext(ctx context.Context) APICredentialItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APICredentialItemOutput)
}

// APICredentialItemArrayInput is an input type that accepts APICredentialItemArray and APICredentialItemArrayOutput values.
// You can construct a concrete instance of `APICredentialItemArrayInput` via:
//
//	APICredentialItemArray{ APICredentialItemArgs{...} }
type APICredentialItemArrayInput interface {
	pulumi.Input

	ToAPICredentialItemArrayOutput() APICredentialItemArrayOutput
	ToAPICredentialItemArrayOutputWithContext(context.Context) APICredentialItemArrayOutput
}

type APICredentialItemArray []APICredentialItemInput

func (APICredentialItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APICredentialItem)(nil)).Elem()
}

func (i APICredentialItemArray) ToAPICredentialItemArrayOutput() APICredentialItemArrayOutput {
	return i.ToAPICredentialItemArrayOutputWithContext(context.Background())
}

func (i APICredentialItemArray) ToAPICredentialItemArrayOutputWithContext(ctx context.Context) APICredentialItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APICredentialItemArrayOutput)
}

// APICredentialItemMapInput is an input type that accepts APICredentialItemMap and APICredentialItemMapOutput values.
// You can construct a concrete instance of `APICredentialItemMapInput` via:
//
//	APICredentialItemMap{ "key": APICredentialItemArgs{...} }
type APICredentialItemMapInput interface {
	pulumi.Input

	ToAPICredentialItemMapOutput() APICredentialItemMapOutput
	ToAPICredentialItemMapOutputWithContext(context.Context) APICredentialItemMapOutput
}

type APICredentialItemMap map[string]APICredentialItemInput

func (APICredentialItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APICredentialItem)(nil)).Elem()
}

func (i APICredentialItemMap) ToAPICredentialItemMapOutput() APICredentialItemMapOutput {
	return i.ToAPICredentialItemMapOutputWithContext(context.Background())
}

func (i APICredentialItemMap) ToAPICredentialItemMapOutputWithContext(ctx context.Context) APICredentialItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APICredentialItemMapOutput)
}

type APICredentialItemOutput struct{ *pulumi.OutputState }

func (APICredentialItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APICredentialItem)(nil)).Elem()
}

func (o APICredentialItemOutput) ToAPICredentialItemOutput() APICredentialItemOutput {
	return o
}

func (o APICredentialItemOutput) ToAPICredentialItemOutputWithContext(ctx context.Context) APICredentialItemOutput {
	return o
}

type APICredentialItemArrayOutput struct{ *pulumi.OutputState }

func (APICredentialItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APICredentialItem)(nil)).Elem()
}

func (o APICredentialItemArrayOutput) ToAPICredentialItemArrayOutput() APICredentialItemArrayOutput {
	return o
}

func (o APICredentialItemArrayOutput) ToAPICredentialItemArrayOutputWithContext(ctx context.Context) APICredentialItemArrayOutput {
	return o
}

func (o APICredentialItemArrayOutput) Index(i pulumi.IntInput) APICredentialItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *APICredentialItem {
		return vs[0].([]*APICredentialItem)[vs[1].(int)]
	}).(APICredentialItemOutput)
}

type APICredentialItemMapOutput struct{ *pulumi.OutputState }

func (APICredentialItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APICredentialItem)(nil)).Elem()
}

func (o APICredentialItemMapOutput) ToAPICredentialItemMapOutput() APICredentialItemMapOutput {
	return o
}

func (o APICredentialItemMapOutput) ToAPICredentialItemMapOutputWithContext(ctx context.Context) APICredentialItemMapOutput {
	return o
}

func (o APICredentialItemMapOutput) MapIndex(k pulumi.StringInput) APICredentialItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *APICredentialItem {
		return vs[0].(map[string]*APICredentialItem)[vs[1].(string)]
	}).(APICredentialItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APICredentialItemInput)(nil)).Elem(), &APICredentialItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*APICredentialItemArrayInput)(nil)).Elem(), APICredentialItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*APICredentialItemMapInput)(nil)).Elem(), APICredentialItemMap{})
	pulumi.RegisterOutputType(APICredentialItemOutput{})
	pulumi.RegisterOutputType(APICredentialItemGetAttachmentResultOutput{})
	pulumi.RegisterOutputType(APICredentialItemArrayOutput{})
	pulumi.RegisterOutputType(APICredentialItemMapOutput{})
}
