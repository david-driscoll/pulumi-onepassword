// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    public static class GetWirelessRouter
    {
        public static Task<GetWirelessRouterResult> InvokeAsync(GetWirelessRouterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWirelessRouterResult>("onepassword:index:GetWirelessRouter", args ?? new GetWirelessRouterArgs(), options.WithDefaults());

        public static Output<GetWirelessRouterResult> Invoke(GetWirelessRouterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWirelessRouterResult>("onepassword:index:GetWirelessRouter", args ?? new GetWirelessRouterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWirelessRouterArgs : Pulumi.InvokeArgs
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

        public GetWirelessRouterArgs()
        {
        }
    }

    public sealed class GetWirelessRouterInvokeArgs : Pulumi.InvokeArgs
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

        public GetWirelessRouterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWirelessRouterResult
    {
        public readonly string? AirPortId;
        public readonly string? AttachedStoragePassword;
        public readonly string? BaseStationName;
        public readonly string? BaseStationPassword;
        public readonly string Category;
        public readonly ImmutableDictionary<string, Outputs.GetField> Fields;
        public readonly string? NetworkName;
        public readonly string? Notes;
        public readonly ImmutableDictionary<string, Outputs.GetSection> Sections;
        public readonly string? ServerIpAddress;
        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The title of the item.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        public readonly string Vault;
        public readonly string? WirelessNetworkPassword;
        public readonly string? WirelessSecurity;

        [OutputConstructor]
        private GetWirelessRouterResult(
            string? airPortId,

            string? attachedStoragePassword,

            string? baseStationName,

            string? baseStationPassword,

            string category,

            ImmutableDictionary<string, Outputs.GetField> fields,

            string? networkName,

            string? notes,

            ImmutableDictionary<string, Outputs.GetSection> sections,

            string? serverIpAddress,

            ImmutableArray<string> tags,

            string title,

            string uuid,

            string vault,

            string? wirelessNetworkPassword,

            string? wirelessSecurity)
        {
            AirPortId = airPortId;
            AttachedStoragePassword = attachedStoragePassword;
            BaseStationName = baseStationName;
            BaseStationPassword = baseStationPassword;
            Category = category;
            Fields = fields;
            NetworkName = networkName;
            Notes = notes;
            Sections = sections;
            ServerIpAddress = serverIpAddress;
            Tags = tags;
            Title = title;
            Uuid = uuid;
            Vault = vault;
            WirelessNetworkPassword = wirelessNetworkPassword;
            WirelessSecurity = wirelessSecurity;
        }
    }
}
