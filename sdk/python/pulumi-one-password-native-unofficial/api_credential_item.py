# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['APICredentialItemArgs', 'APICredentialItem']

@pulumi.input_type
class APICredentialItemArgs:
    def __init__(__self__, *,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 credential: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input['ReferenceArgs']]]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input['UrlArgs']]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a APICredentialItem resource.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        if attachments is not None:
            pulumi.set(__self__, "attachments", attachments)
        if category is not None:
            pulumi.set(__self__, "category", 'API Credential')
        if credential is not None:
            pulumi.set(__self__, "credential", credential)
        if expires is not None:
            pulumi.set(__self__, "expires", expires)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if references is not None:
            pulumi.set(__self__, "references", references)
        if sections is not None:
            pulumi.set(__self__, "sections", sections)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)
        if vault is not None:
            pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter
    def attachments(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]:
        return pulumi.get(self, "attachments")

    @attachments.setter
    def attachments(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]):
        pulumi.set(self, "attachments", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        The category of the vault the item is in.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def credential(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "credential")

    @credential.setter
    def credential(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credential", value)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def references(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReferenceArgs']]]]:
        return pulumi.get(self, "references")

    @references.setter
    def references(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReferenceArgs']]]]):
        pulumi.set(self, "references", value)

    @property
    @pulumi.getter
    def sections(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]:
        return pulumi.get(self, "sections")

    @sections.setter
    def sections(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]):
        pulumi.set(self, "sections", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UrlArgs']]]]:
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UrlArgs']]]]):
        pulumi.set(self, "urls", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_from", value)

    @property
    @pulumi.getter
    def vault(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @vault.setter
    def vault(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault", value)


@pulumi.input_type
class _APICredentialItemState:
    def __init__(__self__, *,
                 vault: pulumi.Input[str]):
        """
        Input properties used for looking up and filtering APICredentialItem resources.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Input[str]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @vault.setter
    def vault(self, value: pulumi.Input[str]):
        pulumi.set(self, "vault", value)


class APICredentialItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 credential: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReferenceArgs']]]]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlArgs']]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a APICredentialItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[APICredentialItemArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a APICredentialItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param APICredentialItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(APICredentialItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 credential: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReferenceArgs']]]]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlArgs']]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = APICredentialItemArgs.__new__(APICredentialItemArgs)

            __props__.__dict__["attachments"] = attachments
            __props__.__dict__["category"] = 'API Credential'
            __props__.__dict__["credential"] = None if credential is None else pulumi.Output.secret(credential)
            __props__.__dict__["expires"] = expires
            __props__.__dict__["fields"] = fields
            __props__.__dict__["filename"] = filename
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["notes"] = notes
            __props__.__dict__["references"] = references
            __props__.__dict__["sections"] = sections
            __props__.__dict__["tags"] = tags
            __props__.__dict__["title"] = title
            __props__.__dict__["type"] = type
            __props__.__dict__["urls"] = urls
            __props__.__dict__["username"] = username
            __props__.__dict__["valid_from"] = valid_from
            __props__.__dict__["vault"] = vault
            __props__.__dict__["id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["attachments", "credential", "fields", "sections"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(APICredentialItem, __self__).__init__(
            'one-password-native-unofficial:index:APICredentialItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vault: Optional[pulumi.Input[str]] = None) -> 'APICredentialItem':
        """
        Get an existing APICredentialItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _APICredentialItemState.__new__(_APICredentialItemState)

        __props__.__dict__["vault"] = vault
        __props__.__dict__["attachments"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["credential"] = None
        __props__.__dict__["expires"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["filename"] = None
        __props__.__dict__["hostname"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["references"] = None
        __props__.__dict__["sections"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["urls"] = None
        __props__.__dict__["username"] = None
        __props__.__dict__["valid_from"] = None
        return APICredentialItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Mapping[str, 'outputs.OutputAttachment']]:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def credential(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "credential")

    @property
    @pulumi.getter
    def expires(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Mapping[str, 'outputs.OutputField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def references(self) -> pulumi.Output[Sequence['outputs.OutputReference']]:
        return pulumi.get(self, "references")

    @property
    @pulumi.getter
    def sections(self) -> pulumi.Output[Mapping[str, 'outputs.OutputSection']]:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        The title of the item.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def urls(self) -> pulumi.Output[Optional[Sequence['outputs.OutputUrl']]]:
        return pulumi.get(self, "urls")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "vault")

