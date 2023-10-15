import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import { ItemType, ResourceTypes, ItemTypeNames } from './types'
import { FieldAssignment, FieldAssignmentType, FieldPurpose, FieldTypeSelector, item } from "@1password/op-js"
import { flatMap, isObject, last, max, pick, snakeCase, split } from "lodash";
import { createHash, randomBytes } from "crypto";
import { RNGFactory, Random } from "random/dist/cjs/random";

const propertiesToIgnore = ['category', 'fields', 'sections', 'tags', 'title', 'vault'];


// import { StaticPage, StaticPageArgs } from "./staticPage";

export class Provider implements provider.Provider {
    resolvedSchema: typeof import('./schema.json');
    constructor(readonly version: string, readonly schema: string) {
        this.resolvedSchema = JSON.parse(schema);
    }

    /**
     * Construct creates a new component pulumi.
     *
     * @param name The name of the resource to create.
     * @param type The type of the resource to create.
     * @param inputs The inputs to the pulumi.
     * @param options the options for the pulumi.
     */
    async construct(name: string, type: string, inputs: pulumi.Inputs,
        options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {

        // TODO: Add support for additional component resources here.
        switch (type) {
            default:
                throw new Error(`unknown resource type ${type} name ${name}`);
        }
    }

    /**
     * Create allocates a new instance of the provided resource and returns its unique ID afterwards.
     * If this call fails, the resource must not have been created (i.e., it is "transactional").
     *
     * @param inputs The properties to set during creation.
     */
    async create(urn: pulumi.URN, inputs: CommonProperties): Promise<provider.CreateResult> {
        const resourceType = getResourceTypeFromUrn(urn);
        if (!resourceType) throw new Error(`unknown resource type ${urn}`);
        const resourceScheme = this.resolvedSchema.resources[resourceType];

        const assignments: FieldAssignment[] = [];
        // TODO: well known fields, sections

        const fields: Record<string, Field> = {}
        const sections: Record<string, Section> = {}
        if (inputs.fields) {
            Object.assign(fields, inputs.fields);
            // assignments.push(...assignFields(inputs.fields));
        }
        if (inputs.sections) {
            Object.assign(sections, inputs.sections);
            // assignments.push(...assignSections(inputs.sections));
        }

        assignments.push(...assignFields(fields));
        assignments.push(...assignSections(sections));

        for (const [name, propScheme] of Object.entries(resourceScheme.inputProperties).filter(z => !propertiesToIgnore.includes(z[0]))) {
            if (inputs[name] != null) {
                if (propScheme.type === "object") {
                    // Can we have deeply nested sections? I dunno.
                    // sections[name] = { fields: {} };
                    // for (const [sName, sPropSchema] of Object.entries(propScheme.properties as Record<string, any>)) {
                    //     if (inputs?.[name]?.fields?.[sName]) {
                    //         sections[name].fields[sName] = {
                    //             value: inputs?.[name]?.fields?.[sName],
                    //             type: sPropSchema.kind,
                    //             purpose: sPropSchema.purpose
                    //         };
                    //     }
                    // }
                } else {
                    fields[name] = {
                        value: inputs[name],
                        type: propScheme.kind,
                        purpose: propScheme.purpose
                    };
                }
            }
        }

        const name = newUniqueName(getNameFromUrn(urn));

        const result = item.create(
            assignments,
            {
                category: resourceType === ItemType.Item ? 'Secure Note' : ItemTypeNames[resourceType] as any,
                vault: inputs.vault,
                tags: inputs.tags ?? [],
                title: inputs.title ?? newUniqueName(getNameFromUrn(urn)),
            }
        )

        return {
            id: result.id,
            outs: {
                resourceScheme,
                assignments,
                fields: JSON.stringify(fields),
                sections,
                inputs,
                id: result.id,
                result,
                name,
                urn,
                test: 'data',
                type: resourceType
            }
        }

    }

    /**
     * Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
     *
     * @param id The ID of the resource to delete.
     * @param props The current properties on the pulumi.
     */
    async delete(id: pulumi.ID, urn: pulumi.URN, props: CommonProperties): Promise<void> {
        const resourceType = getResourceTypeFromUrn(urn);
        if (!resourceType) throw new Error(`unknown resource type ${urn}`);
        item.delete(id);
    }
    /**
     * Update updates an existing resource with new values.
     *
     * @param id The ID of the resource to update.
     * @param olds The old values of properties to update.
     * @param news The new values of properties to update.
     */
    // async update(id: pulumi.ID, urn: pulumi.URN, olds: any, news: any): Promise<provider.UpdateResult> {

    // }

    /**
     * Check validates that the given property bag is valid for a resource of the given type.
     *
     * @param olds The old input properties to use for validation.
     * @param news The new input properties to use for validation.
     */
    async check(urn: pulumi.URN, olds: CommonProperties, news: CommonProperties): Promise<provider.CheckResult> {
        const failures: provider.CheckFailure[] = [];
        const resourceType = getResourceTypeFromUrn(urn);
        if (!resourceType) return {};

        const typeName = ItemTypeNames[resourceType]
        if (resourceType !== ItemType.Item && news.category !== typeName) {
            failures.push({
                property: "category",
                reason: `Category must be ${typeName}`
            })
        }
        if (news.fields) {
            Object.entries(news.fields)
                .filter(z => z[0].includes('.') || z[0].includes('\\') || z[0].includes('='))
                .forEach(field => failures.push({
                    property: `fields.${field[0]}`,
                    reason: 'Field labels cannot contain a period, equals sign or backslash'
                }));
        }
        if (news.sections) {
            for (const section of Object.entries(news.sections)) {
                if (section[0].includes('.') || section[0].includes('\\') || section[0].includes('=')) {
                    failures.push({
                        property: `sections.${section[0]}`,
                        reason: 'Section labels cannot contain a period, equals sign or backslash'
                    })
                }
                for (const field of Object.entries(section[1].fields).filter(z => z[0].includes('.') || z[0].includes('\\') || z[0].includes('='))) {
                    failures.push({
                        property: `sections.${section[0]}.fields.${field[0]}`,
                        reason: 'Field labels cannot contain a period, equals sign or backslash'
                    });
                }
            }
        }
        return failures.length ? { inputs: news, failures } : {};
    }
    // /**
    //  * Diff checks what impacts a hypothetical update will have on the resource's properties.
    //  *
    //  * @param id The ID of the resource to diff.
    //  * @param olds The old values of properties to diff.
    //  * @param news The new values of properties to diff.
    //  */
    // async diff(id: pulumi.ID, urn: pulumi.URN, olds: any, news: any): Promise<provider.DiffResult> { }
    // /**
    //  * Reads the current live state associated with a pulumi.  Enough state must be included in the inputs to uniquely
    //  * identify the resource; this is typically just the resource ID, but it may also include some properties.
    //  */
    // async read(id: pulumi.ID, urn: pulumi.URN, props?: any): Promise<provider.ReadResult> { }
    // /**
    //  * Call calls the indicated method.
    //  *
    //  * @param token The token of the method to call.
    //  * @param inputs The inputs to the method.
    //  */
    // async call(token: string, inputs: pulumi.Inputs): Promise<provider.InvokeResult> { }
    // /**
    //  * Invoke calls the indicated function.
    //  *
    //  * @param token The token of the function to call.
    //  * @param inputs The inputs to the function.
    //  */
    // async invoke(token: string, inputs: any): Promise<provider.InvokeResult> { }
}

// async function constructStaticPage(name: string, inputs: pulumi.Inputs,
//     options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {

//     // Create the component pulumi.
//     const staticPage = new StaticPage(name, inputs as StaticPageArgs, options);

//     // Return the component resource's URN and outputs as its state.
//     return {
//         urn: staticPage.urn,
//         state: {
//             bucket: staticPage.bucket,
//             websiteUrl: staticPage.websiteUrl,
//         },
//     };
// }

function getResourceTypeFromUrn(urn: pulumi.URN) {
    return ResourceTypes.find(z => urn.includes(`:${z}:`));
}
function getNameFromUrn(urn: pulumi.URN) {
    return last(split(urn, ':'))!;
}

function newUniqueName(prefix: string, randomSeed?: Buffer, randlen = 8, maxlen?: number, charset?: string): string {

    if (randlen <= 0) {
        randlen = 8
    }
    if (maxlen && maxlen > 0 && prefix.length + randlen > maxlen) {
        throw new Error(`name '${prefix}' plus ${randlen} random chars is longer than maximum length ${maxlen}`)
    }

    if (!charset) {
        charset = "0123456789abcdef";
    }

    let r: import("random").Random;
    if (!randomSeed || randomSeed.length === 0) {
        r = new Random(RNGFactory(randomBytes(randlen).toString('hex')))
    } else {
        const hash = createHash('sha256');
        hash.write(randomSeed);
        const seed = hash.digest('hex')
        r = new Random(RNGFactory(seed));
    }

    let randomSuffix = '';
    for (let i = 0; i < randlen; i++) {
        randomSuffix += charset[r.int(0, charset.length - 1)]
    }

    return prefix + randomSuffix;
}

interface CommonProperties { category: string; fields?: Record<string, Field>; sections?: Record<string, Section>;[key: string]: any };
interface Field {
    purpose?: FieldPurpose;
    type?: FieldAssignmentType;
    value: string;
}
interface Section {
    fields: Record<string, Field>;
}

function assignFields(fields: Record<string, Field>) {
    const assignments: FieldAssignment[] = [];
    for (const field of Object.entries(fields)) {
        assignments.push([field[0], field[1].purpose === 'PASSWORD' ? 'concealed' : 'text', field[1].value, field[1].purpose])
    }
    return assignments;
}
function assignSections(sections: Record<string, Section>) {

    const assignments: FieldAssignment[] = [];
    for (const section of Object.entries(sections)) {
        assignments.push(...assignFields(section[1].fields))
    }
    return assignments;

}