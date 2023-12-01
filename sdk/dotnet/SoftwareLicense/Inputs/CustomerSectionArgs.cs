// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OnePasswordNativeUnoffical.SoftwareLicense.Inputs
{

    public sealed class CustomerSectionArgs : Pulumi.ResourceArgs
    {
        [Input("company")]
        public Input<string>? Company { get; set; }

        [Input("licensedTo")]
        public Input<string>? LicensedTo { get; set; }

        [Input("registeredEmail")]
        public Input<string>? RegisteredEmail { get; set; }

        public CustomerSectionArgs()
        {
        }
    }
}
