// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    public static class GetBankAccount
    {
        public static Task<GetBankAccountResult> InvokeAsync(GetBankAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBankAccountResult>("onepassword:index:GetBankAccount", args ?? new GetBankAccountArgs(), options.WithDefaults());

        public static Output<GetBankAccountResult> Invoke(GetBankAccountInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBankAccountResult>("onepassword:index:GetBankAccount", args ?? new GetBankAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBankAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public string? Title { get; set; }

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Input("uuid")]
        public string? Uuid { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public string Vault { get; set; } = null!;

        public GetBankAccountArgs()
        {
        }
    }

    public sealed class GetBankAccountInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public GetBankAccountInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBankAccountResult
    {
        public readonly string? AccountNumber;
        public readonly string? BankName;
        public readonly Pulumi.Onepassword.BankAccount.Outputs.BranchInformationSection? BranchInformation;
        public readonly string Category;
        public readonly ImmutableDictionary<string, Outputs.GetField> Fields;
        public readonly string? Iban;
        public readonly string? NameOnAccount;
        public readonly string? Notes;
        public readonly string? Pin;
        public readonly string? RoutingNumber;
        public readonly ImmutableDictionary<string, Outputs.GetSection> Sections;
        public readonly string? Swift;
        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The title of the item.
        /// </summary>
        public readonly string Title;
        public readonly string? Type;
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        public readonly string Vault;

        [OutputConstructor]
        private GetBankAccountResult(
            string? accountNumber,

            string? bankName,

            Pulumi.Onepassword.BankAccount.Outputs.BranchInformationSection? branchInformation,

            string category,

            ImmutableDictionary<string, Outputs.GetField> fields,

            string? iban,

            string? nameOnAccount,

            string? notes,

            string? pin,

            string? routingNumber,

            ImmutableDictionary<string, Outputs.GetSection> sections,

            string? swift,

            ImmutableArray<string> tags,

            string title,

            string? type,

            string uuid,

            string vault)
        {
            AccountNumber = accountNumber;
            BankName = bankName;
            BranchInformation = branchInformation;
            Category = category;
            Fields = fields;
            Iban = iban;
            NameOnAccount = nameOnAccount;
            Notes = notes;
            Pin = pin;
            RoutingNumber = routingNumber;
            Sections = sections;
            Swift = swift;
            Tags = tags;
            Title = title;
            Type = type;
            Uuid = uuid;
            Vault = vault;
        }
    }
}
