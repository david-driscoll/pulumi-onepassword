// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    [OnepasswordResourceType("onepassword:index:Membership")]
    public partial class Membership : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Membership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Membership(string name, MembershipArgs? args = null, CustomResourceOptions? options = null)
            : base("onepassword:index:Membership", name, args ?? new MembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Membership(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("onepassword:index:Membership", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Membership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Membership Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Membership(name, id, options);
        }
    }

    public sealed class MembershipArgs : Pulumi.ResourceArgs
    {
        [Input("expiryDate")]
        public Input<string>? ExpiryDate { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("memberId")]
        public Input<string>? MemberId { get; set; }

        [Input("memberName")]
        public Input<string>? MemberName { get; set; }

        [Input("memberSince")]
        public Input<string>? MemberSince { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("pin")]
        public Input<string>? Pin { get; set; }

        [Input("telephone")]
        public Input<string>? Telephone { get; set; }

        [Input("website")]
        public Input<string>? Website { get; set; }

        public MembershipArgs()
        {
        }
    }
}
