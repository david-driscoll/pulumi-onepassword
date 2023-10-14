// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'onepassword:index:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }


    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["connectionOptions"] = args ? args.connectionOptions : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sid"] = args ? args.sid : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    alias?: pulumi.Input<string>;
    connectionOptions?: pulumi.Input<string>;
    database?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    sid?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}
