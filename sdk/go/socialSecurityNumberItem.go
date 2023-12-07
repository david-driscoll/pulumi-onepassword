// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SocialSecurityNumberItem struct {
	pulumi.CustomResourceState

	Attachments OutputAttachmentMapOutput `pulumi:"attachments"`
	Category    pulumi.StringOutput       `pulumi:"category"`
	Fields      OutputFieldMapOutput      `pulumi:"fields"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id         pulumi.StringOutput        `pulumi:"id"`
	Name       pulumi.StringPtrOutput     `pulumi:"name"`
	Notes      pulumi.StringPtrOutput     `pulumi:"notes"`
	Number     pulumi.StringPtrOutput     `pulumi:"number"`
	References OutputReferenceArrayOutput `pulumi:"references"`
	Sections   OutputSectionMapOutput     `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput    `pulumi:"title"`
	Urls  OutputUrlArrayOutput   `pulumi:"urls"`
	Vault pulumi.StringMapOutput `pulumi:"vault"`
}

// NewSocialSecurityNumberItem registers a new resource with the given unique name, arguments, and options.
func NewSocialSecurityNumberItem(ctx *pulumi.Context,
	name string, args *SocialSecurityNumberItemArgs, opts ...pulumi.ResourceOption) (*SocialSecurityNumberItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Social Security Number")
	if args.Number != nil {
		args.Number = pulumi.ToSecret(args.Number).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"number",
		"sections",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SocialSecurityNumberItem
	err := ctx.RegisterResource("one-password-native-unofficial:index:SocialSecurityNumberItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSocialSecurityNumberItem gets an existing SocialSecurityNumberItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSocialSecurityNumberItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SocialSecurityNumberItemState, opts ...pulumi.ResourceOption) (*SocialSecurityNumberItem, error) {
	var resource SocialSecurityNumberItem
	err := ctx.ReadResource("one-password-native-unofficial:index:SocialSecurityNumberItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SocialSecurityNumberItem resources.
type socialSecurityNumberItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type SocialSecurityNumberItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (SocialSecurityNumberItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*socialSecurityNumberItemState)(nil)).Elem()
}

type socialSecurityNumberItemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category *string            `pulumi:"category"`
	Fields   map[string]Field   `pulumi:"fields"`
	Name     *string            `pulumi:"name"`
	Notes    *string            `pulumi:"notes"`
	Number   *string            `pulumi:"number"`
	Sections map[string]Section `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string  `pulumi:"title"`
	Urls  []string `pulumi:"urls"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a SocialSecurityNumberItem resource.
type SocialSecurityNumberItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category pulumi.StringPtrInput
	Fields   FieldMapInput
	Name     pulumi.StringPtrInput
	Notes    pulumi.StringPtrInput
	Number   pulumi.StringPtrInput
	Sections SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	Urls  pulumi.StringArrayInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (SocialSecurityNumberItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*socialSecurityNumberItemArgs)(nil)).Elem()
}

type SocialSecurityNumberItemInput interface {
	pulumi.Input

	ToSocialSecurityNumberItemOutput() SocialSecurityNumberItemOutput
	ToSocialSecurityNumberItemOutputWithContext(ctx context.Context) SocialSecurityNumberItemOutput
}

func (*SocialSecurityNumberItem) ElementType() reflect.Type {
	return reflect.TypeOf((**SocialSecurityNumberItem)(nil)).Elem()
}

func (i *SocialSecurityNumberItem) ToSocialSecurityNumberItemOutput() SocialSecurityNumberItemOutput {
	return i.ToSocialSecurityNumberItemOutputWithContext(context.Background())
}

func (i *SocialSecurityNumberItem) ToSocialSecurityNumberItemOutputWithContext(ctx context.Context) SocialSecurityNumberItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SocialSecurityNumberItemOutput)
}

// SocialSecurityNumberItemArrayInput is an input type that accepts SocialSecurityNumberItemArray and SocialSecurityNumberItemArrayOutput values.
// You can construct a concrete instance of `SocialSecurityNumberItemArrayInput` via:
//
//	SocialSecurityNumberItemArray{ SocialSecurityNumberItemArgs{...} }
type SocialSecurityNumberItemArrayInput interface {
	pulumi.Input

	ToSocialSecurityNumberItemArrayOutput() SocialSecurityNumberItemArrayOutput
	ToSocialSecurityNumberItemArrayOutputWithContext(context.Context) SocialSecurityNumberItemArrayOutput
}

type SocialSecurityNumberItemArray []SocialSecurityNumberItemInput

func (SocialSecurityNumberItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SocialSecurityNumberItem)(nil)).Elem()
}

func (i SocialSecurityNumberItemArray) ToSocialSecurityNumberItemArrayOutput() SocialSecurityNumberItemArrayOutput {
	return i.ToSocialSecurityNumberItemArrayOutputWithContext(context.Background())
}

func (i SocialSecurityNumberItemArray) ToSocialSecurityNumberItemArrayOutputWithContext(ctx context.Context) SocialSecurityNumberItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SocialSecurityNumberItemArrayOutput)
}

// SocialSecurityNumberItemMapInput is an input type that accepts SocialSecurityNumberItemMap and SocialSecurityNumberItemMapOutput values.
// You can construct a concrete instance of `SocialSecurityNumberItemMapInput` via:
//
//	SocialSecurityNumberItemMap{ "key": SocialSecurityNumberItemArgs{...} }
type SocialSecurityNumberItemMapInput interface {
	pulumi.Input

	ToSocialSecurityNumberItemMapOutput() SocialSecurityNumberItemMapOutput
	ToSocialSecurityNumberItemMapOutputWithContext(context.Context) SocialSecurityNumberItemMapOutput
}

type SocialSecurityNumberItemMap map[string]SocialSecurityNumberItemInput

func (SocialSecurityNumberItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SocialSecurityNumberItem)(nil)).Elem()
}

func (i SocialSecurityNumberItemMap) ToSocialSecurityNumberItemMapOutput() SocialSecurityNumberItemMapOutput {
	return i.ToSocialSecurityNumberItemMapOutputWithContext(context.Background())
}

func (i SocialSecurityNumberItemMap) ToSocialSecurityNumberItemMapOutputWithContext(ctx context.Context) SocialSecurityNumberItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SocialSecurityNumberItemMapOutput)
}

type SocialSecurityNumberItemOutput struct{ *pulumi.OutputState }

func (SocialSecurityNumberItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SocialSecurityNumberItem)(nil)).Elem()
}

func (o SocialSecurityNumberItemOutput) ToSocialSecurityNumberItemOutput() SocialSecurityNumberItemOutput {
	return o
}

func (o SocialSecurityNumberItemOutput) ToSocialSecurityNumberItemOutputWithContext(ctx context.Context) SocialSecurityNumberItemOutput {
	return o
}

type SocialSecurityNumberItemArrayOutput struct{ *pulumi.OutputState }

func (SocialSecurityNumberItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SocialSecurityNumberItem)(nil)).Elem()
}

func (o SocialSecurityNumberItemArrayOutput) ToSocialSecurityNumberItemArrayOutput() SocialSecurityNumberItemArrayOutput {
	return o
}

func (o SocialSecurityNumberItemArrayOutput) ToSocialSecurityNumberItemArrayOutputWithContext(ctx context.Context) SocialSecurityNumberItemArrayOutput {
	return o
}

func (o SocialSecurityNumberItemArrayOutput) Index(i pulumi.IntInput) SocialSecurityNumberItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SocialSecurityNumberItem {
		return vs[0].([]*SocialSecurityNumberItem)[vs[1].(int)]
	}).(SocialSecurityNumberItemOutput)
}

type SocialSecurityNumberItemMapOutput struct{ *pulumi.OutputState }

func (SocialSecurityNumberItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SocialSecurityNumberItem)(nil)).Elem()
}

func (o SocialSecurityNumberItemMapOutput) ToSocialSecurityNumberItemMapOutput() SocialSecurityNumberItemMapOutput {
	return o
}

func (o SocialSecurityNumberItemMapOutput) ToSocialSecurityNumberItemMapOutputWithContext(ctx context.Context) SocialSecurityNumberItemMapOutput {
	return o
}

func (o SocialSecurityNumberItemMapOutput) MapIndex(k pulumi.StringInput) SocialSecurityNumberItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SocialSecurityNumberItem {
		return vs[0].(map[string]*SocialSecurityNumberItem)[vs[1].(string)]
	}).(SocialSecurityNumberItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SocialSecurityNumberItemInput)(nil)).Elem(), &SocialSecurityNumberItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*SocialSecurityNumberItemArrayInput)(nil)).Elem(), SocialSecurityNumberItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SocialSecurityNumberItemMapInput)(nil)).Elem(), SocialSecurityNumberItemMap{})
	pulumi.RegisterOutputType(SocialSecurityNumberItemOutput{})
	pulumi.RegisterOutputType(SocialSecurityNumberItemArrayOutput{})
	pulumi.RegisterOutputType(SocialSecurityNumberItemMapOutput{})
}
