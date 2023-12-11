// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnofficial
{
    public static class GetSSHKey
    {
        public static Task<GetSSHKeyResult> InvokeAsync(GetSSHKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSSHKeyResult>("one-password-native-unofficial:index:GetSSHKey", args ?? new GetSSHKeyArgs(), options.WithDefaults());

        public static Output<GetSSHKeyResult> Invoke(GetSSHKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSSHKeyResult>("one-password-native-unofficial:index:GetSSHKey", args ?? new GetSSHKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSSHKeyArgs : Pulumi.InvokeArgs
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

        public GetSSHKeyArgs()
        {
        }
    }

    public sealed class GetSSHKeyInvokeArgs : Pulumi.InvokeArgs
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

        public GetSSHKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSSHKeyResult
    {
        public readonly ImmutableDictionary<string, Outputs.OutputAttachment> Attachments;
        public readonly string Category;
        public readonly ImmutableDictionary<string, Outputs.OutputField> Fields;
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string Id;
        public readonly string? Notes;
        public readonly string? PrivateKey;
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

        [OutputConstructor]
        private GetSSHKeyResult(
            ImmutableDictionary<string, Outputs.OutputAttachment> attachments,

            string category,

            ImmutableDictionary<string, Outputs.OutputField> fields,

            string id,

            string? notes,

            string? privateKey,

            ImmutableArray<Outputs.OutputReference> references,

            ImmutableDictionary<string, Outputs.OutputSection> sections,

            ImmutableArray<string> tags,

            string title,

            ImmutableArray<Outputs.OutputUrl> urls,

            ImmutableDictionary<string, string> vault)
        {
            Attachments = attachments;
            Category = category;
            Fields = fields;
            Id = id;
            Notes = notes;
            PrivateKey = privateKey;
            References = references;
            Sections = sections;
            Tags = tags;
            Title = title;
            Urls = urls;
            Vault = vault;
        }
    }
}
