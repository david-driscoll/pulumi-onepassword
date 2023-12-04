// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class SecureNoteItem extends pulumi.CustomResource {
    /**
     * Get an existing SecureNoteItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecureNoteItemState, opts?: pulumi.CustomResourceOptions): SecureNoteItem {
        return new SecureNoteItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'one-password-native-unoffical:index:SecureNoteItem';

    /**
     * Returns true if the given object is an instance of SecureNoteItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecureNoteItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecureNoteItem.__pulumiType;
    }

    public readonly attachments!: pulumi.Output<{[key: string]: outputs.OutputAttachment}>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly fields!: pulumi.Output<{[key: string]: outputs.OutputField}>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly references!: pulumi.Output<outputs.OutputReference[]>;
    public readonly sections!: pulumi.Output<{[key: string]: outputs.OutputSection}>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The title of the item.
     */
    public readonly title!: pulumi.Output<string>;
    public readonly urls!: pulumi.Output<outputs.OutputUrl[] | undefined>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    public readonly vault!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a SecureNoteItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecureNoteItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecureNoteItemArgs | SecureNoteItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecureNoteItemState | undefined;
            resourceInputs["vault"] = state ? state.vault : undefined;
        } else {
            const args = argsOrState as SecureNoteItemArgs | undefined;
            if ((!args || args.vault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vault'");
            }
            resourceInputs["attachments"] = args ? args.attachments : undefined;
            resourceInputs["category"] = "Secure Note";
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["urls"] = args ? args.urls : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["attachments", "fields", "sections"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecureNoteItem.__pulumiType, name, resourceInputs, opts);
    }
}

export interface SecureNoteItemState {
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecureNoteItem resource.
 */
export interface SecureNoteItemArgs {
    attachments?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>}>;
    /**
     * The category of the vault the item is in.
     */
    category?: pulumi.Input<"Secure Note">;
    fields?: pulumi.Input<{[key: string]: pulumi.Input<inputs.FieldArgs>}>;
    notes?: pulumi.Input<string>;
    sections?: pulumi.Input<{[key: string]: pulumi.Input<inputs.SectionArgs>}>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    urls?: pulumi.Input<pulumi.Input<inputs.UrlArgs>[]>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}
