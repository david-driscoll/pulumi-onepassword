// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getSocialSecurityNumber(args: GetSocialSecurityNumberArgs, opts?: pulumi.InvokeOptions): Promise<GetSocialSecurityNumberResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("one-password-native-unofficial:index:GetSocialSecurityNumber", {
        "id": args.id,
        "title": args.title,
        "vault": args.vault,
    }, opts);
}

export interface GetSocialSecurityNumberArgs {
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    id?: string;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: string;
    /**
     * The UUID of the vault the item is in.
     */
    vault: string;
}

export interface GetSocialSecurityNumberResult {
    readonly attachments: {[key: string]: outputs.OutputAttachment};
    readonly category: enums.Category | string;
    readonly fields: {[key: string]: outputs.OutputField};
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly id: string;
    readonly name?: string;
    readonly notes?: string;
    readonly number?: string;
    readonly references: outputs.OutputReference[];
    readonly sections: {[key: string]: outputs.OutputSection};
    /**
     * An array of strings of the tags assigned to the item.
     */
    readonly tags: string[];
    /**
     * The title of the item.
     */
    readonly title: string;
    readonly urls?: outputs.OutputUrl[];
    readonly vault: {[key: string]: string};
}

export function getSocialSecurityNumberOutput(args: GetSocialSecurityNumberOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSocialSecurityNumberResult> {
    return pulumi.output(args).apply(a => getSocialSecurityNumber(a, opts))
}

export interface GetSocialSecurityNumberOutputArgs {
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    id?: pulumi.Input<string>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}
