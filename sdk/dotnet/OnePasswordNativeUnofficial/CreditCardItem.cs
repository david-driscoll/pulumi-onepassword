// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnofficial
{
    [OnePasswordNativeUnofficialResourceType("one-password-native-unofficial:index:CreditCardItem")]
    public partial class CreditCardItem : Pulumi.CustomResource
    {
        [Output("additionalDetails")]
        public Output<Rocket.Surgery.OnePasswordNativeUnofficial.CreditCard.Outputs.AdditionalDetailsSection?> AdditionalDetails { get; private set; } = null!;

        [Output("attachments")]
        public Output<ImmutableDictionary<string, Outputs.OutputAttachment>> Attachments { get; private set; } = null!;

        [Output("cardholderName")]
        public Output<string?> CardholderName { get; private set; } = null!;

        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("contactInformation")]
        public Output<Rocket.Surgery.OnePasswordNativeUnofficial.CreditCard.Outputs.ContactInformationSection?> ContactInformation { get; private set; } = null!;

        [Output("expiryDate")]
        public Output<string?> ExpiryDate { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableDictionary<string, Outputs.OutputField>> Fields { get; private set; } = null!;

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("number")]
        public Output<string?> Number { get; private set; } = null!;

        [Output("references")]
        public Output<ImmutableArray<Outputs.OutputReference>> References { get; private set; } = null!;

        [Output("sections")]
        public Output<ImmutableDictionary<string, Outputs.OutputSection>> Sections { get; private set; } = null!;

        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The title of the item.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("urls")]
        public Output<ImmutableArray<Outputs.OutputUrl>> Urls { get; private set; } = null!;

        [Output("validFrom")]
        public Output<string?> ValidFrom { get; private set; } = null!;

        [Output("vault")]
        public Output<ImmutableDictionary<string, string>> Vault { get; private set; } = null!;

        [Output("verificationNumber")]
        public Output<string?> VerificationNumber { get; private set; } = null!;


        /// <summary>
        /// Create a CreditCardItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CreditCardItem(string name, CreditCardItemArgs? args = null, CustomResourceOptions? options = null)
            : base("one-password-native-unofficial:index:CreditCardItem", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private CreditCardItem(string name, Input<string> id, CreditCardItemState? state = null, CustomResourceOptions? options = null)
            : base("one-password-native-unofficial:index:CreditCardItem", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CreditCardItemArgs? MakeArgs(CreditCardItemArgs? args)
        {
            args ??= new CreditCardItemArgs();
            args.Category = "Credit Card";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/david-driscoll",
                AdditionalSecretOutputs =
                {
                    "additionalDetails",
                    "attachments",
                    "fields",
                    "sections",
                    "verificationNumber",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CreditCardItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CreditCardItem Get(string name, Input<string> id, CreditCardItemState? state = null, CustomResourceOptions? options = null)
        {
            return new CreditCardItem(name, id, state, options);
        }
    }

    public sealed class CreditCardItemArgs : Pulumi.ResourceArgs
    {
        [Input("additionalDetails")]
        public Input<Rocket.Surgery.OnePasswordNativeUnofficial.CreditCard.Inputs.AdditionalDetailsSectionArgs>? AdditionalDetails { get; set; }

        [Input("attachments")]
        private InputMap<AssetOrArchive>? _attachments;
        public InputMap<AssetOrArchive> Attachments
        {
            get => _attachments ?? (_attachments = new InputMap<AssetOrArchive>());
            set => _attachments = value;
        }

        [Input("cardholderName")]
        public Input<string>? CardholderName { get; set; }

        /// <summary>
        /// The category of the vault the item is in.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("contactInformation")]
        public Input<Rocket.Surgery.OnePasswordNativeUnofficial.CreditCard.Inputs.ContactInformationSectionArgs>? ContactInformation { get; set; }

        [Input("expiryDate")]
        public Input<string>? ExpiryDate { get; set; }

        [Input("fields")]
        private InputMap<Inputs.FieldArgs>? _fields;
        public InputMap<Inputs.FieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputMap<Inputs.FieldArgs>());
            set => _fields = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("number")]
        public Input<string>? Number { get; set; }

        [Input("references")]
        private InputList<Inputs.ReferenceArgs>? _references;
        public InputList<Inputs.ReferenceArgs> References
        {
            get => _references ?? (_references = new InputList<Inputs.ReferenceArgs>());
            set => _references = value;
        }

        [Input("sections")]
        private InputMap<Inputs.SectionArgs>? _sections;
        public InputMap<Inputs.SectionArgs> Sections
        {
            get => _sections ?? (_sections = new InputMap<Inputs.SectionArgs>());
            set => _sections = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("urls")]
        private InputList<Inputs.UrlArgs>? _urls;
        public InputList<Inputs.UrlArgs> Urls
        {
            get => _urls ?? (_urls = new InputList<Inputs.UrlArgs>());
            set => _urls = value;
        }

        [Input("validFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault")]
        public Input<string>? Vault { get; set; }

        [Input("verificationNumber")]
        private Input<string>? _verificationNumber;
        public Input<string>? VerificationNumber
        {
            get => _verificationNumber;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _verificationNumber = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CreditCardItemArgs()
        {
        }
    }

    public sealed class CreditCardItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public CreditCardItemState()
        {
        }
    }
}
