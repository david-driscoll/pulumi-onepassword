// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnofficial
{
    public static class GetCryptoWallet
    {
        public static Task<GetCryptoWalletResult> InvokeAsync(GetCryptoWalletArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCryptoWalletResult>("one-password-native-unofficial:index:GetCryptoWallet", args ?? new GetCryptoWalletArgs(), options.WithDefaults());

        public static Output<GetCryptoWalletResult> Invoke(GetCryptoWalletInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCryptoWalletResult>("one-password-native-unofficial:index:GetCryptoWallet", args ?? new GetCryptoWalletInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCryptoWalletArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public string? Title { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public string Vault { get; set; } = null!;

        public GetCryptoWalletArgs()
        {
        }
    }

    public sealed class GetCryptoWalletInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public GetCryptoWalletInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCryptoWalletResult
    {
        public readonly ImmutableDictionary<string, Outputs.OutputAttachment> Attachments;
        public readonly string Category;
        public readonly ImmutableDictionary<string, Outputs.OutputField> Fields;
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string Id;
        public readonly string? Notes;
        public readonly string? Password;
        public readonly string? RecoveryPhrase;
        public readonly ImmutableArray<Outputs.OutputReference> References;
        public readonly ImmutableDictionary<string, Outputs.OutputSection> Sections;
        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The title of the item.
        /// </summary>
        public readonly string Title;
        public readonly ImmutableArray<Outputs.OutputUrl> Urls;
        public readonly ImmutableDictionary<string, string> Vault;
        public readonly Rocket.Surgery.OnePasswordNativeUnofficial.CryptoWallet.Outputs.WalletSection? Wallet;

        [OutputConstructor]
        private GetCryptoWalletResult(
            ImmutableDictionary<string, Outputs.OutputAttachment> attachments,

            string category,

            ImmutableDictionary<string, Outputs.OutputField> fields,

            string id,

            string? notes,

            string? password,

            string? recoveryPhrase,

            ImmutableArray<Outputs.OutputReference> references,

            ImmutableDictionary<string, Outputs.OutputSection> sections,

            ImmutableArray<string> tags,

            string title,

            ImmutableArray<Outputs.OutputUrl> urls,

            ImmutableDictionary<string, string> vault,

            Rocket.Surgery.OnePasswordNativeUnofficial.CryptoWallet.Outputs.WalletSection? wallet)
        {
            Attachments = attachments;
            Category = category;
            Fields = fields;
            Id = id;
            Notes = notes;
            Password = password;
            RecoveryPhrase = recoveryPhrase;
            References = references;
            Sections = sections;
            Tags = tags;
            Title = title;
            Urls = urls;
            Vault = vault;
            Wallet = wallet;
        }
    }
}
