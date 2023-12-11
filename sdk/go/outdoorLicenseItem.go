// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OutdoorLicenseItem struct {
	pulumi.CustomResourceState

	ApprovedWildlife pulumi.StringPtrOutput    `pulumi:"approvedWildlife"`
	Attachments      OutputAttachmentMapOutput `pulumi:"attachments"`
	Category         pulumi.StringOutput       `pulumi:"category"`
	Country          pulumi.StringPtrOutput    `pulumi:"country"`
	Expires          pulumi.StringPtrOutput    `pulumi:"expires"`
	Fields           OutputFieldMapOutput      `pulumi:"fields"`
	FullName         pulumi.StringPtrOutput    `pulumi:"fullName"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id           pulumi.StringOutput        `pulumi:"id"`
	MaximumQuota pulumi.StringPtrOutput     `pulumi:"maximumQuota"`
	Notes        pulumi.StringPtrOutput     `pulumi:"notes"`
	References   OutputReferenceArrayOutput `pulumi:"references"`
	Sections     OutputSectionMapOutput     `pulumi:"sections"`
	State        pulumi.StringPtrOutput     `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title     pulumi.StringOutput    `pulumi:"title"`
	Urls      OutputUrlArrayOutput   `pulumi:"urls"`
	ValidFrom pulumi.StringPtrOutput `pulumi:"validFrom"`
	Vault     pulumi.StringMapOutput `pulumi:"vault"`
}

// NewOutdoorLicenseItem registers a new resource with the given unique name, arguments, and options.
func NewOutdoorLicenseItem(ctx *pulumi.Context,
	name string, args *OutdoorLicenseItemArgs, opts ...pulumi.ResourceOption) (*OutdoorLicenseItem, error) {
	if args == nil {
		args = &OutdoorLicenseItemArgs{}
	}

	args.Category = pulumi.StringPtr("Outdoor License")
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"sections",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource OutdoorLicenseItem
	err := ctx.RegisterResource("one-password-native-unofficial:index:OutdoorLicenseItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutdoorLicenseItem gets an existing OutdoorLicenseItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutdoorLicenseItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutdoorLicenseItemState, opts ...pulumi.ResourceOption) (*OutdoorLicenseItem, error) {
	var resource OutdoorLicenseItem
	err := ctx.ReadResource("one-password-native-unofficial:index:OutdoorLicenseItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutdoorLicenseItem resources.
type outdoorLicenseItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type OutdoorLicenseItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (OutdoorLicenseItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*outdoorLicenseItemState)(nil)).Elem()
}

type outdoorLicenseItemArgs struct {
	ApprovedWildlife *string                          `pulumi:"approvedWildlife"`
	Attachments      map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category     *string            `pulumi:"category"`
	Country      *string            `pulumi:"country"`
	Expires      *string            `pulumi:"expires"`
	Fields       map[string]Field   `pulumi:"fields"`
	FullName     *string            `pulumi:"fullName"`
	MaximumQuota *string            `pulumi:"maximumQuota"`
	Notes        *string            `pulumi:"notes"`
	References   []Reference        `pulumi:"references"`
	Sections     map[string]Section `pulumi:"sections"`
	State        *string            `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title     *string `pulumi:"title"`
	Urls      []Url   `pulumi:"urls"`
	ValidFrom *string `pulumi:"validFrom"`
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

// The set of arguments for constructing a OutdoorLicenseItem resource.
type OutdoorLicenseItemArgs struct {
	ApprovedWildlife pulumi.StringPtrInput
	Attachments      pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category     pulumi.StringPtrInput
	Country      pulumi.StringPtrInput
	Expires      pulumi.StringPtrInput
	Fields       FieldMapInput
	FullName     pulumi.StringPtrInput
	MaximumQuota pulumi.StringPtrInput
	Notes        pulumi.StringPtrInput
	References   ReferenceArrayInput
	Sections     SectionMapInput
	State        pulumi.StringPtrInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title     pulumi.StringPtrInput
	Urls      UrlArrayInput
	ValidFrom pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringPtrInput
}

func (OutdoorLicenseItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outdoorLicenseItemArgs)(nil)).Elem()
}

type OutdoorLicenseItemInput interface {
	pulumi.Input

	ToOutdoorLicenseItemOutput() OutdoorLicenseItemOutput
	ToOutdoorLicenseItemOutputWithContext(ctx context.Context) OutdoorLicenseItemOutput
}

func (*OutdoorLicenseItem) ElementType() reflect.Type {
	return reflect.TypeOf((**OutdoorLicenseItem)(nil)).Elem()
}

func (i *OutdoorLicenseItem) ToOutdoorLicenseItemOutput() OutdoorLicenseItemOutput {
	return i.ToOutdoorLicenseItemOutputWithContext(context.Background())
}

func (i *OutdoorLicenseItem) ToOutdoorLicenseItemOutputWithContext(ctx context.Context) OutdoorLicenseItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutdoorLicenseItemOutput)
}

// OutdoorLicenseItemArrayInput is an input type that accepts OutdoorLicenseItemArray and OutdoorLicenseItemArrayOutput values.
// You can construct a concrete instance of `OutdoorLicenseItemArrayInput` via:
//
//	OutdoorLicenseItemArray{ OutdoorLicenseItemArgs{...} }
type OutdoorLicenseItemArrayInput interface {
	pulumi.Input

	ToOutdoorLicenseItemArrayOutput() OutdoorLicenseItemArrayOutput
	ToOutdoorLicenseItemArrayOutputWithContext(context.Context) OutdoorLicenseItemArrayOutput
}

type OutdoorLicenseItemArray []OutdoorLicenseItemInput

func (OutdoorLicenseItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutdoorLicenseItem)(nil)).Elem()
}

func (i OutdoorLicenseItemArray) ToOutdoorLicenseItemArrayOutput() OutdoorLicenseItemArrayOutput {
	return i.ToOutdoorLicenseItemArrayOutputWithContext(context.Background())
}

func (i OutdoorLicenseItemArray) ToOutdoorLicenseItemArrayOutputWithContext(ctx context.Context) OutdoorLicenseItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutdoorLicenseItemArrayOutput)
}

// OutdoorLicenseItemMapInput is an input type that accepts OutdoorLicenseItemMap and OutdoorLicenseItemMapOutput values.
// You can construct a concrete instance of `OutdoorLicenseItemMapInput` via:
//
//	OutdoorLicenseItemMap{ "key": OutdoorLicenseItemArgs{...} }
type OutdoorLicenseItemMapInput interface {
	pulumi.Input

	ToOutdoorLicenseItemMapOutput() OutdoorLicenseItemMapOutput
	ToOutdoorLicenseItemMapOutputWithContext(context.Context) OutdoorLicenseItemMapOutput
}

type OutdoorLicenseItemMap map[string]OutdoorLicenseItemInput

func (OutdoorLicenseItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutdoorLicenseItem)(nil)).Elem()
}

func (i OutdoorLicenseItemMap) ToOutdoorLicenseItemMapOutput() OutdoorLicenseItemMapOutput {
	return i.ToOutdoorLicenseItemMapOutputWithContext(context.Background())
}

func (i OutdoorLicenseItemMap) ToOutdoorLicenseItemMapOutputWithContext(ctx context.Context) OutdoorLicenseItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutdoorLicenseItemMapOutput)
}

type OutdoorLicenseItemOutput struct{ *pulumi.OutputState }

func (OutdoorLicenseItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutdoorLicenseItem)(nil)).Elem()
}

func (o OutdoorLicenseItemOutput) ToOutdoorLicenseItemOutput() OutdoorLicenseItemOutput {
	return o
}

func (o OutdoorLicenseItemOutput) ToOutdoorLicenseItemOutputWithContext(ctx context.Context) OutdoorLicenseItemOutput {
	return o
}

type OutdoorLicenseItemArrayOutput struct{ *pulumi.OutputState }

func (OutdoorLicenseItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutdoorLicenseItem)(nil)).Elem()
}

func (o OutdoorLicenseItemArrayOutput) ToOutdoorLicenseItemArrayOutput() OutdoorLicenseItemArrayOutput {
	return o
}

func (o OutdoorLicenseItemArrayOutput) ToOutdoorLicenseItemArrayOutputWithContext(ctx context.Context) OutdoorLicenseItemArrayOutput {
	return o
}

func (o OutdoorLicenseItemArrayOutput) Index(i pulumi.IntInput) OutdoorLicenseItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OutdoorLicenseItem {
		return vs[0].([]*OutdoorLicenseItem)[vs[1].(int)]
	}).(OutdoorLicenseItemOutput)
}

type OutdoorLicenseItemMapOutput struct{ *pulumi.OutputState }

func (OutdoorLicenseItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutdoorLicenseItem)(nil)).Elem()
}

func (o OutdoorLicenseItemMapOutput) ToOutdoorLicenseItemMapOutput() OutdoorLicenseItemMapOutput {
	return o
}

func (o OutdoorLicenseItemMapOutput) ToOutdoorLicenseItemMapOutputWithContext(ctx context.Context) OutdoorLicenseItemMapOutput {
	return o
}

func (o OutdoorLicenseItemMapOutput) MapIndex(k pulumi.StringInput) OutdoorLicenseItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OutdoorLicenseItem {
		return vs[0].(map[string]*OutdoorLicenseItem)[vs[1].(string)]
	}).(OutdoorLicenseItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OutdoorLicenseItemInput)(nil)).Elem(), &OutdoorLicenseItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutdoorLicenseItemArrayInput)(nil)).Elem(), OutdoorLicenseItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutdoorLicenseItemMapInput)(nil)).Elem(), OutdoorLicenseItemMap{})
	pulumi.RegisterOutputType(OutdoorLicenseItemOutput{})
	pulumi.RegisterOutputType(OutdoorLicenseItemArrayOutput{})
	pulumi.RegisterOutputType(OutdoorLicenseItemMapOutput{})
}
