// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class WirelessRouterItem extends pulumi.CustomResource {
    /**
     * Get an existing WirelessRouterItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessRouterItemState, opts?: pulumi.CustomResourceOptions): WirelessRouterItem {
        return new WirelessRouterItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'one-password-native-unoffical:index:WirelessRouterItem';

    /**
     * Returns true if the given object is an instance of WirelessRouterItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessRouterItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessRouterItem.__pulumiType;
    }

    public readonly airPortId!: pulumi.Output<string | undefined>;
    public readonly attachedStoragePassword!: pulumi.Output<string | undefined>;
    public readonly attachments!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly baseStationName!: pulumi.Output<string | undefined>;
    public readonly baseStationPassword!: pulumi.Output<string | undefined>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly fields!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly networkName!: pulumi.Output<string | undefined>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly references!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly sections!: pulumi.Output<{[key: string]: outputs.OutSection}>;
    public readonly serverIpAddress!: pulumi.Output<string | undefined>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The title of the item.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    /**
     * The UUID of the vault the item is in.
     */
    public readonly vault!: pulumi.Output<string>;
    public readonly wirelessNetworkPassword!: pulumi.Output<string | undefined>;
    public readonly wirelessSecurity!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessRouterItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WirelessRouterItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessRouterItemArgs | WirelessRouterItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessRouterItemState | undefined;
            resourceInputs["vault"] = state ? state.vault : undefined;
        } else {
            const args = argsOrState as WirelessRouterItemArgs | undefined;
            if ((!args || args.vault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vault'");
            }
            resourceInputs["airPortId"] = args ? args.airPortId : undefined;
            resourceInputs["attachedStoragePassword"] = args?.attachedStoragePassword ? pulumi.secret(args.attachedStoragePassword) : undefined;
            resourceInputs["attachments"] = args ? args.attachments : undefined;
            resourceInputs["baseStationName"] = args ? args.baseStationName : undefined;
            resourceInputs["baseStationPassword"] = args?.baseStationPassword ? pulumi.secret(args.baseStationPassword) : undefined;
            resourceInputs["category"] = "Wireless Router";
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["networkName"] = args ? args.networkName : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["serverIpAddress"] = args ? args.serverIpAddress : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["wirelessNetworkPassword"] = args?.wirelessNetworkPassword ? pulumi.secret(args.wirelessNetworkPassword) : undefined;
            resourceInputs["wirelessSecurity"] = args ? args.wirelessSecurity : undefined;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["attachedStoragePassword", "attachments", "baseStationPassword", "fields", "references", "sections", "wirelessNetworkPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(WirelessRouterItem.__pulumiType, name, resourceInputs, opts);
    }

    getAttachment(args: WirelessRouterItem.GetAttachmentArgs): pulumi.Output<WirelessRouterItem.GetAttachmentResult> {
        return pulumi.runtime.call("one-password-native-unoffical:index:WirelessRouterItem/attachment", {
            "__self__": this,
            "name": args.name,
        }, this);
    }
}

export interface WirelessRouterItemState {
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessRouterItem resource.
 */
export interface WirelessRouterItemArgs {
    airPortId?: pulumi.Input<string>;
    attachedStoragePassword?: pulumi.Input<string>;
    attachments?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>}>;
    baseStationName?: pulumi.Input<string>;
    baseStationPassword?: pulumi.Input<string>;
    /**
     * The category of the vault the item is in.
     */
    category?: pulumi.Input<"Wireless Router">;
    fields?: pulumi.Input<{[key: string]: pulumi.Input<inputs.FieldArgs>}>;
    networkName?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    sections?: pulumi.Input<{[key: string]: pulumi.Input<inputs.SectionArgs>}>;
    serverIpAddress?: pulumi.Input<string>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
    wirelessNetworkPassword?: pulumi.Input<string>;
    wirelessSecurity?: pulumi.Input<string>;
}

export namespace WirelessRouterItem {
    /**
     * The set of arguments for the WirelessRouterItem.getAttachment method.
     */
    export interface GetAttachmentArgs {
        /**
         * The name or uuid of the attachment to get
         */
        name: pulumi.Input<string>;
    }

    /**
     * The results of the WirelessRouterItem.getAttachment method.
     */
    export interface GetAttachmentResult {
        /**
         * the value of the attachment
         */
        readonly value: string;
    }

}
