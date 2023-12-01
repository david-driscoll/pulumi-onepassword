// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getCreditCard(args: GetCreditCardArgs, opts?: pulumi.InvokeOptions): Promise<GetCreditCardResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("one-password-native-unoffical:index:GetCreditCard", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetCreditCardArgs {
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

export interface GetCreditCardResult {
    readonly additionalDetails?: outputs.creditCard.AdditionalDetailsSection;
    readonly attachments: {[key: string]: outputs.OutField};
    readonly cardholderName?: string;
    readonly category: enums.Category | string;
    readonly contactInformation?: outputs.creditCard.ContactInformationSection;
    readonly expiryDate?: string;
    readonly fields: {[key: string]: outputs.OutField};
    readonly notes?: string;
    readonly number?: string;
    readonly references: {[key: string]: outputs.OutField};
    readonly sections: {[key: string]: outputs.OutSection};
    /**
     * An array of strings of the tags assigned to the item.
     */
    readonly tags: string[];
    /**
     * The title of the item.
     */
    readonly title: string;
    readonly type?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid: string;
    readonly validFrom?: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault: string;
    readonly verificationNumber?: string;
}

export function getCreditCardOutput(args: GetCreditCardOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCreditCardResult> {
    return pulumi.output(args).apply(a => getCreditCard(a, opts))
}

export interface GetCreditCardOutputArgs {
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
