// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getAPICredential(args: GetAPICredentialArgs, opts?: pulumi.InvokeOptions): Promise<GetAPICredentialResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("onepassword:index:GetAPICredential", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetAPICredentialArgs {
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    uuid?: string;
    /**
     * The UUID of the vault the item is in.
     */
    vault: string;
}

export interface GetAPICredentialResult {
    readonly category: enums.Category | string;
    readonly credential?: string;
    readonly expires?: string;
    readonly fields: {[key: string]: outputs.GetField};
    readonly filename?: string;
    readonly hostname?: string;
    readonly notes?: string;
    readonly sections: {[key: string]: outputs.GetSection};
    /**
     * An array of strings of the tags assigned to the item.
     */
    readonly tags: string[];
    /**
     * The title of the item.
     */
    readonly title: string;
    readonly type?: string;
    readonly username?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid: string;
    readonly validFrom?: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault: string;
}

export function getAPICredentialOutput(args: GetAPICredentialOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAPICredentialResult> {
    return pulumi.output(args).apply(a => getAPICredential(a, opts))
}

export interface GetAPICredentialOutputArgs {
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    uuid?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}
