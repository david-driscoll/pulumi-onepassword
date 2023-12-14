// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DriverLicenseItem struct {
	pulumi.CustomResourceState

	Address                pulumi.StringPtrOutput    `pulumi:"address"`
	Attachments            OutputAttachmentMapOutput `pulumi:"attachments"`
	Category               pulumi.StringOutput       `pulumi:"category"`
	ConditionsRestrictions pulumi.StringPtrOutput    `pulumi:"conditionsRestrictions"`
	Country                pulumi.StringPtrOutput    `pulumi:"country"`
	DateOfBirth            pulumi.StringPtrOutput    `pulumi:"dateOfBirth"`
	ExpiryDate             pulumi.StringPtrOutput    `pulumi:"expiryDate"`
	Fields                 OutputFieldMapOutput      `pulumi:"fields"`
	FullName               pulumi.StringPtrOutput    `pulumi:"fullName"`
	Gender                 pulumi.StringPtrOutput    `pulumi:"gender"`
	Height                 pulumi.StringPtrOutput    `pulumi:"height"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id           pulumi.StringOutput        `pulumi:"id"`
	LicenseClass pulumi.StringPtrOutput     `pulumi:"licenseClass"`
	Notes        pulumi.StringPtrOutput     `pulumi:"notes"`
	Number       pulumi.StringPtrOutput     `pulumi:"number"`
	References   OutputReferenceArrayOutput `pulumi:"references"`
	Sections     OutputSectionMapOutput     `pulumi:"sections"`
	State        pulumi.StringPtrOutput     `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput  `pulumi:"title"`
	Urls  OutputUrlArrayOutput `pulumi:"urls"`
	Vault OutputVaultOutput    `pulumi:"vault"`
}

// NewDriverLicenseItem registers a new resource with the given unique name, arguments, and options.
func NewDriverLicenseItem(ctx *pulumi.Context,
	name string, args *DriverLicenseItemArgs, opts ...pulumi.ResourceOption) (*DriverLicenseItem, error) {
	if args == nil {
		args = &DriverLicenseItemArgs{}
	}

	args.Category = pulumi.StringPtr("Driver License")
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"sections",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource DriverLicenseItem
	err := ctx.RegisterResource("one-password-native-unofficial:index:DriverLicenseItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDriverLicenseItem gets an existing DriverLicenseItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDriverLicenseItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DriverLicenseItemState, opts ...pulumi.ResourceOption) (*DriverLicenseItem, error) {
	var resource DriverLicenseItem
	err := ctx.ReadResource("one-password-native-unofficial:index:DriverLicenseItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DriverLicenseItem resources.
type driverLicenseItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type DriverLicenseItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (DriverLicenseItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*driverLicenseItemState)(nil)).Elem()
}

type driverLicenseItemArgs struct {
	Address     *string                          `pulumi:"address"`
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category               *string            `pulumi:"category"`
	ConditionsRestrictions *string            `pulumi:"conditionsRestrictions"`
	Country                *string            `pulumi:"country"`
	DateOfBirth            *string            `pulumi:"dateOfBirth"`
	ExpiryDate             *string            `pulumi:"expiryDate"`
	Fields                 map[string]Field   `pulumi:"fields"`
	FullName               *string            `pulumi:"fullName"`
	Gender                 *string            `pulumi:"gender"`
	Height                 *string            `pulumi:"height"`
	LicenseClass           *string            `pulumi:"licenseClass"`
	Notes                  *string            `pulumi:"notes"`
	Number                 *string            `pulumi:"number"`
	References             []string           `pulumi:"references"`
	Sections               map[string]Section `pulumi:"sections"`
	State                  *string            `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string       `pulumi:"title"`
	Urls  []interface{} `pulumi:"urls"`
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

// The set of arguments for constructing a DriverLicenseItem resource.
type DriverLicenseItemArgs struct {
	Address     pulumi.StringPtrInput
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category               pulumi.StringPtrInput
	ConditionsRestrictions pulumi.StringPtrInput
	Country                pulumi.StringPtrInput
	DateOfBirth            pulumi.StringPtrInput
	ExpiryDate             pulumi.StringPtrInput
	Fields                 FieldMapInput
	FullName               pulumi.StringPtrInput
	Gender                 pulumi.StringPtrInput
	Height                 pulumi.StringPtrInput
	LicenseClass           pulumi.StringPtrInput
	Notes                  pulumi.StringPtrInput
	Number                 pulumi.StringPtrInput
	References             pulumi.StringArrayInput
	Sections               SectionMapInput
	State                  pulumi.StringPtrInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	Urls  pulumi.ArrayInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringPtrInput
}

func (DriverLicenseItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*driverLicenseItemArgs)(nil)).Elem()
}

type DriverLicenseItemInput interface {
	pulumi.Input

	ToDriverLicenseItemOutput() DriverLicenseItemOutput
	ToDriverLicenseItemOutputWithContext(ctx context.Context) DriverLicenseItemOutput
}

func (*DriverLicenseItem) ElementType() reflect.Type {
	return reflect.TypeOf((**DriverLicenseItem)(nil)).Elem()
}

func (i *DriverLicenseItem) ToDriverLicenseItemOutput() DriverLicenseItemOutput {
	return i.ToDriverLicenseItemOutputWithContext(context.Background())
}

func (i *DriverLicenseItem) ToDriverLicenseItemOutputWithContext(ctx context.Context) DriverLicenseItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriverLicenseItemOutput)
}

// DriverLicenseItemArrayInput is an input type that accepts DriverLicenseItemArray and DriverLicenseItemArrayOutput values.
// You can construct a concrete instance of `DriverLicenseItemArrayInput` via:
//
//	DriverLicenseItemArray{ DriverLicenseItemArgs{...} }
type DriverLicenseItemArrayInput interface {
	pulumi.Input

	ToDriverLicenseItemArrayOutput() DriverLicenseItemArrayOutput
	ToDriverLicenseItemArrayOutputWithContext(context.Context) DriverLicenseItemArrayOutput
}

type DriverLicenseItemArray []DriverLicenseItemInput

func (DriverLicenseItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DriverLicenseItem)(nil)).Elem()
}

func (i DriverLicenseItemArray) ToDriverLicenseItemArrayOutput() DriverLicenseItemArrayOutput {
	return i.ToDriverLicenseItemArrayOutputWithContext(context.Background())
}

func (i DriverLicenseItemArray) ToDriverLicenseItemArrayOutputWithContext(ctx context.Context) DriverLicenseItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriverLicenseItemArrayOutput)
}

// DriverLicenseItemMapInput is an input type that accepts DriverLicenseItemMap and DriverLicenseItemMapOutput values.
// You can construct a concrete instance of `DriverLicenseItemMapInput` via:
//
//	DriverLicenseItemMap{ "key": DriverLicenseItemArgs{...} }
type DriverLicenseItemMapInput interface {
	pulumi.Input

	ToDriverLicenseItemMapOutput() DriverLicenseItemMapOutput
	ToDriverLicenseItemMapOutputWithContext(context.Context) DriverLicenseItemMapOutput
}

type DriverLicenseItemMap map[string]DriverLicenseItemInput

func (DriverLicenseItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DriverLicenseItem)(nil)).Elem()
}

func (i DriverLicenseItemMap) ToDriverLicenseItemMapOutput() DriverLicenseItemMapOutput {
	return i.ToDriverLicenseItemMapOutputWithContext(context.Background())
}

func (i DriverLicenseItemMap) ToDriverLicenseItemMapOutputWithContext(ctx context.Context) DriverLicenseItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriverLicenseItemMapOutput)
}

type DriverLicenseItemOutput struct{ *pulumi.OutputState }

func (DriverLicenseItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DriverLicenseItem)(nil)).Elem()
}

func (o DriverLicenseItemOutput) ToDriverLicenseItemOutput() DriverLicenseItemOutput {
	return o
}

func (o DriverLicenseItemOutput) ToDriverLicenseItemOutputWithContext(ctx context.Context) DriverLicenseItemOutput {
	return o
}

type DriverLicenseItemArrayOutput struct{ *pulumi.OutputState }

func (DriverLicenseItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DriverLicenseItem)(nil)).Elem()
}

func (o DriverLicenseItemArrayOutput) ToDriverLicenseItemArrayOutput() DriverLicenseItemArrayOutput {
	return o
}

func (o DriverLicenseItemArrayOutput) ToDriverLicenseItemArrayOutputWithContext(ctx context.Context) DriverLicenseItemArrayOutput {
	return o
}

func (o DriverLicenseItemArrayOutput) Index(i pulumi.IntInput) DriverLicenseItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DriverLicenseItem {
		return vs[0].([]*DriverLicenseItem)[vs[1].(int)]
	}).(DriverLicenseItemOutput)
}

type DriverLicenseItemMapOutput struct{ *pulumi.OutputState }

func (DriverLicenseItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DriverLicenseItem)(nil)).Elem()
}

func (o DriverLicenseItemMapOutput) ToDriverLicenseItemMapOutput() DriverLicenseItemMapOutput {
	return o
}

func (o DriverLicenseItemMapOutput) ToDriverLicenseItemMapOutputWithContext(ctx context.Context) DriverLicenseItemMapOutput {
	return o
}

func (o DriverLicenseItemMapOutput) MapIndex(k pulumi.StringInput) DriverLicenseItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DriverLicenseItem {
		return vs[0].(map[string]*DriverLicenseItem)[vs[1].(string)]
	}).(DriverLicenseItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DriverLicenseItemInput)(nil)).Elem(), &DriverLicenseItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*DriverLicenseItemArrayInput)(nil)).Elem(), DriverLicenseItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DriverLicenseItemMapInput)(nil)).Elem(), DriverLicenseItemMap{})
	pulumi.RegisterOutputType(DriverLicenseItemOutput{})
	pulumi.RegisterOutputType(DriverLicenseItemArrayOutput{})
	pulumi.RegisterOutputType(DriverLicenseItemMapOutput{})
}
