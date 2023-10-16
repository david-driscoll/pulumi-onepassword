// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/bankaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBankAccount(ctx *pulumi.Context, args *GetBankAccountArgs, opts ...pulumi.InvokeOption) (*GetBankAccountResult, error) {
	var rv GetBankAccountResult
	err := ctx.Invoke("onepassword:index:GetBankAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBankAccountArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetBankAccountResult struct {
	AccountNumber     *string                               `pulumi:"accountNumber"`
	BankName          *string                               `pulumi:"bankName"`
	BranchInformation *bankaccount.BranchInformationSection `pulumi:"branchInformation"`
	Category          string                                `pulumi:"category"`
	Fields            map[string]GetField                   `pulumi:"fields"`
	Iban              *string                               `pulumi:"iban"`
	NameOnAccount     *string                               `pulumi:"nameOnAccount"`
	Notes             *string                               `pulumi:"notes"`
	Pin               *string                               `pulumi:"pin"`
	RoutingNumber     *string                               `pulumi:"routingNumber"`
	Sections          map[string]GetSection                 `pulumi:"sections"`
	Swift             *string                               `pulumi:"swift"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string  `pulumi:"title"`
	Type  *string `pulumi:"type"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func GetBankAccountOutput(ctx *pulumi.Context, args GetBankAccountOutputArgs, opts ...pulumi.InvokeOption) GetBankAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBankAccountResult, error) {
			args := v.(GetBankAccountArgs)
			r, err := GetBankAccount(ctx, &args, opts...)
			var s GetBankAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBankAccountResultOutput)
}

type GetBankAccountOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetBankAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBankAccountArgs)(nil)).Elem()
}

type GetBankAccountResultOutput struct{ *pulumi.OutputState }

func (GetBankAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBankAccountResult)(nil)).Elem()
}

func (o GetBankAccountResultOutput) ToGetBankAccountResultOutput() GetBankAccountResultOutput {
	return o
}

func (o GetBankAccountResultOutput) ToGetBankAccountResultOutputWithContext(ctx context.Context) GetBankAccountResultOutput {
	return o
}

func (o GetBankAccountResultOutput) AccountNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.AccountNumber }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) BankName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.BankName }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) BranchInformation() bankaccount.BranchInformationSectionPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *bankaccount.BranchInformationSection { return v.BranchInformation }).(bankaccount.BranchInformationSectionPtrOutput)
}

func (o GetBankAccountResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetBankAccountResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetBankAccountResultOutput) Fields() GetFieldMapOutput {
	return o.ApplyT(func(v GetBankAccountResult) map[string]GetField { return v.Fields }).(GetFieldMapOutput)
}

func (o GetBankAccountResultOutput) Iban() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.Iban }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) NameOnAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.NameOnAccount }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) Pin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.Pin }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) RoutingNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.RoutingNumber }).(pulumi.StringPtrOutput)
}

func (o GetBankAccountResultOutput) Sections() GetSectionMapOutput {
	return o.ApplyT(func(v GetBankAccountResult) map[string]GetSection { return v.Sections }).(GetSectionMapOutput)
}

func (o GetBankAccountResultOutput) Swift() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.Swift }).(pulumi.StringPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetBankAccountResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBankAccountResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetBankAccountResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetBankAccountResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetBankAccountResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBankAccountResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetBankAccountResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetBankAccountResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o GetBankAccountResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetBankAccountResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBankAccountResultOutput{})
}
