# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from . import server as _server
from ._enums import *

__all__ = [
    'GetServerResult',
    'AwaitableGetServerResult',
    'get_server',
    'get_server_output',
]

@pulumi.output_type
class GetServerResult:
    def __init__(__self__, admin_console=None, attachments=None, category=None, fields=None, hosting_provider=None, id=None, notes=None, password=None, references=None, sections=None, tags=None, title=None, url=None, urls=None, username=None, vault=None):
        if admin_console and not isinstance(admin_console, dict):
            raise TypeError("Expected argument 'admin_console' to be a dict")
        pulumi.set(__self__, "admin_console", admin_console)
        if attachments and not isinstance(attachments, dict):
            raise TypeError("Expected argument 'attachments' to be a dict")
        pulumi.set(__self__, "attachments", attachments)
        if category and not isinstance(category, dict):
            raise TypeError("Expected argument 'category' to be a dict")
        pulumi.set(__self__, "category", category)
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if hosting_provider and not isinstance(hosting_provider, dict):
            raise TypeError("Expected argument 'hosting_provider' to be a dict")
        pulumi.set(__self__, "hosting_provider", hosting_provider)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if references and not isinstance(references, list):
            raise TypeError("Expected argument 'references' to be a list")
        pulumi.set(__self__, "references", references)
        if sections and not isinstance(sections, dict):
            raise TypeError("Expected argument 'sections' to be a dict")
        pulumi.set(__self__, "sections", sections)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if urls and not isinstance(urls, list):
            raise TypeError("Expected argument 'urls' to be a list")
        pulumi.set(__self__, "urls", urls)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if vault and not isinstance(vault, dict):
            raise TypeError("Expected argument 'vault' to be a dict")
        pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter(name="adminConsole")
    def admin_console(self) -> Optional['_server.outputs.AdminConsoleSection']:
        return pulumi.get(self, "admin_console")

    @property
    @pulumi.getter
    def attachments(self) -> Mapping[str, 'outputs.OutputAttachment']:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def category(self) -> str:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def fields(self) -> Mapping[str, 'outputs.OutputField']:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="hostingProvider")
    def hosting_provider(self) -> Optional['_server.outputs.HostingProviderSection']:
        return pulumi.get(self, "hosting_provider")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def references(self) -> Sequence['outputs.OutputReference']:
        return pulumi.get(self, "references")

    @property
    @pulumi.getter
    def sections(self) -> Mapping[str, 'outputs.OutputSection']:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The title of the item.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def urls(self) -> Optional[Sequence['outputs.OutputUrl']]:
        return pulumi.get(self, "urls")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vault(self) -> Mapping[str, str]:
        return pulumi.get(self, "vault")


class AwaitableGetServerResult(GetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerResult(
            admin_console=self.admin_console,
            attachments=self.attachments,
            category=self.category,
            fields=self.fields,
            hosting_provider=self.hosting_provider,
            id=self.id,
            notes=self.notes,
            password=self.password,
            references=self.references,
            sections=self.sections,
            tags=self.tags,
            title=self.title,
            url=self.url,
            urls=self.urls,
            username=self.username,
            vault=self.vault)


def get_server(id: Optional[str] = None,
               title: Optional[str] = None,
               vault: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str vault: The UUID of the vault the item is in.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['title'] = title
    __args__['vault'] = vault
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('one-password-native-unofficial:index:GetServer', __args__, opts=opts, typ=GetServerResult).value

    return AwaitableGetServerResult(
        admin_console=__ret__.admin_console,
        attachments=__ret__.attachments,
        category=__ret__.category,
        fields=__ret__.fields,
        hosting_provider=__ret__.hosting_provider,
        id=__ret__.id,
        notes=__ret__.notes,
        password=__ret__.password,
        references=__ret__.references,
        sections=__ret__.sections,
        tags=__ret__.tags,
        title=__ret__.title,
        url=__ret__.url,
        urls=__ret__.urls,
        username=__ret__.username,
        vault=__ret__.vault)


@_utilities.lift_output_func(get_server)
def get_server_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                      title: Optional[pulumi.Input[Optional[str]]] = None,
                      vault: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str vault: The UUID of the vault the item is in.
    """
    ...
