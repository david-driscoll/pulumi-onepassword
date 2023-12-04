// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/emailaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EmailAccountItem struct {
	pulumi.CustomResourceState

	Attachments        OutputAttachmentMapOutput                       `pulumi:"attachments"`
	AuthMethod         pulumi.StringPtrOutput                          `pulumi:"authMethod"`
	Category           pulumi.StringOutput                             `pulumi:"category"`
	ContactInformation emailaccount.ContactInformationSectionPtrOutput `pulumi:"contactInformation"`
	Fields             OutputFieldMapOutput                            `pulumi:"fields"`
	Notes              pulumi.StringPtrOutput                          `pulumi:"notes"`
	Password           pulumi.StringPtrOutput                          `pulumi:"password"`
	PortNumber         pulumi.StringPtrOutput                          `pulumi:"portNumber"`
	References         OutputReferenceArrayOutput                      `pulumi:"references"`
	Sections           OutputSectionMapOutput                          `pulumi:"sections"`
	Security           pulumi.StringPtrOutput                          `pulumi:"security"`
	Server             pulumi.StringPtrOutput                          `pulumi:"server"`
	Smtp               emailaccount.SmtpSectionPtrOutput               `pulumi:"smtp"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title    pulumi.StringOutput    `pulumi:"title"`
	Type     pulumi.StringPtrOutput `pulumi:"type"`
	Urls     OutputUrlArrayOutput   `pulumi:"urls"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid  pulumi.StringOutput    `pulumi:"uuid"`
	Vault pulumi.StringMapOutput `pulumi:"vault"`
}

// NewEmailAccountItem registers a new resource with the given unique name, arguments, and options.
func NewEmailAccountItem(ctx *pulumi.Context,
	name string, args *EmailAccountItemArgs, opts ...pulumi.ResourceOption) (*EmailAccountItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Email Account")
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"password",
		"sections",
		"smtp",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource EmailAccountItem
	err := ctx.RegisterResource("one-password-native-unoffical:index:EmailAccountItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailAccountItem gets an existing EmailAccountItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailAccountItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailAccountItemState, opts ...pulumi.ResourceOption) (*EmailAccountItem, error) {
	var resource EmailAccountItem
	err := ctx.ReadResource("one-password-native-unoffical:index:EmailAccountItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailAccountItem resources.
type emailAccountItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type EmailAccountItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (EmailAccountItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailAccountItemState)(nil)).Elem()
}

type emailAccountItemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	AuthMethod  *string                          `pulumi:"authMethod"`
	// The category of the vault the item is in.
	Category           *string                                 `pulumi:"category"`
	ContactInformation *emailaccount.ContactInformationSection `pulumi:"contactInformation"`
	Fields             map[string]Field                        `pulumi:"fields"`
	Notes              *string                                 `pulumi:"notes"`
	Password           *string                                 `pulumi:"password"`
	PortNumber         *string                                 `pulumi:"portNumber"`
	Sections           map[string]Section                      `pulumi:"sections"`
	Security           *string                                 `pulumi:"security"`
	Server             *string                                 `pulumi:"server"`
	Smtp               *emailaccount.SmtpSection               `pulumi:"smtp"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    *string `pulumi:"title"`
	Type     *string `pulumi:"type"`
	Urls     []Url   `pulumi:"urls"`
	Username *string `pulumi:"username"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a EmailAccountItem resource.
type EmailAccountItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	AuthMethod  pulumi.StringPtrInput
	// The category of the vault the item is in.
	Category           pulumi.StringPtrInput
	ContactInformation emailaccount.ContactInformationSectionPtrInput
	Fields             FieldMapInput
	Notes              pulumi.StringPtrInput
	Password           pulumi.StringPtrInput
	PortNumber         pulumi.StringPtrInput
	Sections           SectionMapInput
	Security           pulumi.StringPtrInput
	Server             pulumi.StringPtrInput
	Smtp               emailaccount.SmtpSectionPtrInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    pulumi.StringPtrInput
	Type     pulumi.StringPtrInput
	Urls     UrlArrayInput
	Username pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (EmailAccountItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailAccountItemArgs)(nil)).Elem()
}

type EmailAccountItemInput interface {
	pulumi.Input

	ToEmailAccountItemOutput() EmailAccountItemOutput
	ToEmailAccountItemOutputWithContext(ctx context.Context) EmailAccountItemOutput
}

func (*EmailAccountItem) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailAccountItem)(nil)).Elem()
}

func (i *EmailAccountItem) ToEmailAccountItemOutput() EmailAccountItemOutput {
	return i.ToEmailAccountItemOutputWithContext(context.Background())
}

func (i *EmailAccountItem) ToEmailAccountItemOutputWithContext(ctx context.Context) EmailAccountItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailAccountItemOutput)
}

// EmailAccountItemArrayInput is an input type that accepts EmailAccountItemArray and EmailAccountItemArrayOutput values.
// You can construct a concrete instance of `EmailAccountItemArrayInput` via:
//
//	EmailAccountItemArray{ EmailAccountItemArgs{...} }
type EmailAccountItemArrayInput interface {
	pulumi.Input

	ToEmailAccountItemArrayOutput() EmailAccountItemArrayOutput
	ToEmailAccountItemArrayOutputWithContext(context.Context) EmailAccountItemArrayOutput
}

type EmailAccountItemArray []EmailAccountItemInput

func (EmailAccountItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailAccountItem)(nil)).Elem()
}

func (i EmailAccountItemArray) ToEmailAccountItemArrayOutput() EmailAccountItemArrayOutput {
	return i.ToEmailAccountItemArrayOutputWithContext(context.Background())
}

func (i EmailAccountItemArray) ToEmailAccountItemArrayOutputWithContext(ctx context.Context) EmailAccountItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailAccountItemArrayOutput)
}

// EmailAccountItemMapInput is an input type that accepts EmailAccountItemMap and EmailAccountItemMapOutput values.
// You can construct a concrete instance of `EmailAccountItemMapInput` via:
//
//	EmailAccountItemMap{ "key": EmailAccountItemArgs{...} }
type EmailAccountItemMapInput interface {
	pulumi.Input

	ToEmailAccountItemMapOutput() EmailAccountItemMapOutput
	ToEmailAccountItemMapOutputWithContext(context.Context) EmailAccountItemMapOutput
}

type EmailAccountItemMap map[string]EmailAccountItemInput

func (EmailAccountItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailAccountItem)(nil)).Elem()
}

func (i EmailAccountItemMap) ToEmailAccountItemMapOutput() EmailAccountItemMapOutput {
	return i.ToEmailAccountItemMapOutputWithContext(context.Background())
}

func (i EmailAccountItemMap) ToEmailAccountItemMapOutputWithContext(ctx context.Context) EmailAccountItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailAccountItemMapOutput)
}

type EmailAccountItemOutput struct{ *pulumi.OutputState }

func (EmailAccountItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailAccountItem)(nil)).Elem()
}

func (o EmailAccountItemOutput) ToEmailAccountItemOutput() EmailAccountItemOutput {
	return o
}

func (o EmailAccountItemOutput) ToEmailAccountItemOutputWithContext(ctx context.Context) EmailAccountItemOutput {
	return o
}

type EmailAccountItemArrayOutput struct{ *pulumi.OutputState }

func (EmailAccountItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailAccountItem)(nil)).Elem()
}

func (o EmailAccountItemArrayOutput) ToEmailAccountItemArrayOutput() EmailAccountItemArrayOutput {
	return o
}

func (o EmailAccountItemArrayOutput) ToEmailAccountItemArrayOutputWithContext(ctx context.Context) EmailAccountItemArrayOutput {
	return o
}

func (o EmailAccountItemArrayOutput) Index(i pulumi.IntInput) EmailAccountItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailAccountItem {
		return vs[0].([]*EmailAccountItem)[vs[1].(int)]
	}).(EmailAccountItemOutput)
}

type EmailAccountItemMapOutput struct{ *pulumi.OutputState }

func (EmailAccountItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailAccountItem)(nil)).Elem()
}

func (o EmailAccountItemMapOutput) ToEmailAccountItemMapOutput() EmailAccountItemMapOutput {
	return o
}

func (o EmailAccountItemMapOutput) ToEmailAccountItemMapOutputWithContext(ctx context.Context) EmailAccountItemMapOutput {
	return o
}

func (o EmailAccountItemMapOutput) MapIndex(k pulumi.StringInput) EmailAccountItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailAccountItem {
		return vs[0].(map[string]*EmailAccountItem)[vs[1].(string)]
	}).(EmailAccountItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailAccountItemInput)(nil)).Elem(), &EmailAccountItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailAccountItemArrayInput)(nil)).Elem(), EmailAccountItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailAccountItemMapInput)(nil)).Elem(), EmailAccountItemMap{})
	pulumi.RegisterOutputType(EmailAccountItemOutput{})
	pulumi.RegisterOutputType(EmailAccountItemArrayOutput{})
	pulumi.RegisterOutputType(EmailAccountItemMapOutput{})
}
