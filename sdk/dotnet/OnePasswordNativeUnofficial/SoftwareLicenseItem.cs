// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnofficial
{
    [OnePasswordNativeUnofficialResourceType("one-password-native-unofficial:index:SoftwareLicenseItem")]
    public partial class SoftwareLicenseItem : Pulumi.CustomResource
    {
        [Output("attachments")]
        public Output<ImmutableDictionary<string, Outputs.OutputAttachment>> Attachments { get; private set; } = null!;

        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("customer")]
        public Output<Rocket.Surgery.OnePasswordNativeUnofficial.SoftwareLicense.Outputs.CustomerSection?> Customer { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableDictionary<string, Outputs.OutputField>> Fields { get; private set; } = null!;

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("licenseKey")]
        public Output<string?> LicenseKey { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("order")]
        public Output<Rocket.Surgery.OnePasswordNativeUnofficial.SoftwareLicense.Outputs.OrderSection?> Order { get; private set; } = null!;

        [Output("publisher")]
        public Output<Rocket.Surgery.OnePasswordNativeUnofficial.SoftwareLicense.Outputs.PublisherSection?> Publisher { get; private set; } = null!;

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

        [Output("urls")]
        public Output<ImmutableArray<Outputs.OutputUrl>> Urls { get; private set; } = null!;

        [Output("vault")]
        public Output<ImmutableDictionary<string, string>> Vault { get; private set; } = null!;

        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a SoftwareLicenseItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SoftwareLicenseItem(string name, SoftwareLicenseItemArgs? args = null, CustomResourceOptions? options = null)
            : base("one-password-native-unofficial:index:SoftwareLicenseItem", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private SoftwareLicenseItem(string name, Input<string> id, SoftwareLicenseItemState? state = null, CustomResourceOptions? options = null)
            : base("one-password-native-unofficial:index:SoftwareLicenseItem", name, state, MakeResourceOptions(options, id))
        {
        }

        private static SoftwareLicenseItemArgs? MakeArgs(SoftwareLicenseItemArgs? args)
        {
            args ??= new SoftwareLicenseItemArgs();
            args.Category = "Software License";
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
                    "attachments",
                    "fields",
                    "sections",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SoftwareLicenseItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SoftwareLicenseItem Get(string name, Input<string> id, SoftwareLicenseItemState? state = null, CustomResourceOptions? options = null)
        {
            return new SoftwareLicenseItem(name, id, state, options);
        }
    }

    public sealed class SoftwareLicenseItemArgs : Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputMap<AssetOrArchive>? _attachments;
        public InputMap<AssetOrArchive> Attachments
        {
            get => _attachments ?? (_attachments = new InputMap<AssetOrArchive>());
            set => _attachments = value;
        }

        /// <summary>
        /// The category of the vault the item is in.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("customer")]
        public Input<Rocket.Surgery.OnePasswordNativeUnofficial.SoftwareLicense.Inputs.CustomerSectionArgs>? Customer { get; set; }

        [Input("fields")]
        private InputMap<Inputs.FieldArgs>? _fields;
        public InputMap<Inputs.FieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputMap<Inputs.FieldArgs>());
            set => _fields = value;
        }

        [Input("licenseKey")]
        public Input<string>? LicenseKey { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("order")]
        public Input<Rocket.Surgery.OnePasswordNativeUnofficial.SoftwareLicense.Inputs.OrderSectionArgs>? Order { get; set; }

        [Input("publisher")]
        public Input<Rocket.Surgery.OnePasswordNativeUnofficial.SoftwareLicense.Inputs.PublisherSectionArgs>? Publisher { get; set; }

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

        [Input("urls")]
        private InputList<Union<Inputs.UrlArgs, string>>? _urls;
        public InputList<Union<Inputs.UrlArgs, string>> Urls
        {
            get => _urls ?? (_urls = new InputList<Union<Inputs.UrlArgs, string>>());
            set => _urls = value;
        }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault")]
        public Input<string>? Vault { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public SoftwareLicenseItemArgs()
        {
        }
    }

    public sealed class SoftwareLicenseItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public SoftwareLicenseItemState()
        {
        }
    }
}
