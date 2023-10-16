# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import emailaccount as _emailaccount
from . import outputs
from ._enums import *

__all__ = [
    'GetEmailAccountResult',
    'AwaitableGetEmailAccountResult',
    'get_email_account',
    'get_email_account_output',
]

@pulumi.output_type
class GetEmailAccountResult:
    def __init__(__self__, auth_method=None, category=None, contact_information=None, fields=None, notes=None, password=None, port_number=None, sections=None, security=None, server=None, smtp=None, tags=None, title=None, type=None, username=None, uuid=None, vault=None):
        if auth_method and not isinstance(auth_method, str):
            raise TypeError("Expected argument 'auth_method' to be a str")
        pulumi.set(__self__, "auth_method", auth_method)
        if category and not isinstance(category, dict):
            raise TypeError("Expected argument 'category' to be a dict")
        pulumi.set(__self__, "category", category)
        if contact_information and not isinstance(contact_information, dict):
            raise TypeError("Expected argument 'contact_information' to be a dict")
        pulumi.set(__self__, "contact_information", contact_information)
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if port_number and not isinstance(port_number, str):
            raise TypeError("Expected argument 'port_number' to be a str")
        pulumi.set(__self__, "port_number", port_number)
        if sections and not isinstance(sections, dict):
            raise TypeError("Expected argument 'sections' to be a dict")
        pulumi.set(__self__, "sections", sections)
        if security and not isinstance(security, str):
            raise TypeError("Expected argument 'security' to be a str")
        pulumi.set(__self__, "security", security)
        if server and not isinstance(server, str):
            raise TypeError("Expected argument 'server' to be a str")
        pulumi.set(__self__, "server", server)
        if smtp and not isinstance(smtp, dict):
            raise TypeError("Expected argument 'smtp' to be a dict")
        pulumi.set(__self__, "smtp", smtp)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vault and not isinstance(vault, str):
            raise TypeError("Expected argument 'vault' to be a str")
        pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> Optional[str]:
        return pulumi.get(self, "auth_method")

    @property
    @pulumi.getter
    def category(self) -> str:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="contactInformation")
    def contact_information(self) -> Optional['_emailaccount.outputs.ContactInformationSection']:
        return pulumi.get(self, "contact_information")

    @property
    @pulumi.getter
    def fields(self) -> Mapping[str, 'outputs.GetField']:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="portNumber")
    def port_number(self) -> Optional[str]:
        return pulumi.get(self, "port_number")

    @property
    @pulumi.getter
    def sections(self) -> Mapping[str, 'outputs.GetSection']:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def security(self) -> Optional[str]:
        return pulumi.get(self, "security")

    @property
    @pulumi.getter
    def server(self) -> Optional[str]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def smtp(self) -> Optional['_emailaccount.outputs.SmtpSection']:
        return pulumi.get(self, "smtp")

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
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vault(self) -> str:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")


class AwaitableGetEmailAccountResult(GetEmailAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEmailAccountResult(
            auth_method=self.auth_method,
            category=self.category,
            contact_information=self.contact_information,
            fields=self.fields,
            notes=self.notes,
            password=self.password,
            port_number=self.port_number,
            sections=self.sections,
            security=self.security,
            server=self.server,
            smtp=self.smtp,
            tags=self.tags,
            title=self.title,
            type=self.type,
            username=self.username,
            uuid=self.uuid,
            vault=self.vault)


def get_email_account(title: Optional[str] = None,
                      uuid: Optional[str] = None,
                      vault: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEmailAccountResult:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    __args__ = dict()
    __args__['title'] = title
    __args__['uuid'] = uuid
    __args__['vault'] = vault
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('onepassword:index:GetEmailAccount', __args__, opts=opts, typ=GetEmailAccountResult).value

    return AwaitableGetEmailAccountResult(
        auth_method=__ret__.auth_method,
        category=__ret__.category,
        contact_information=__ret__.contact_information,
        fields=__ret__.fields,
        notes=__ret__.notes,
        password=__ret__.password,
        port_number=__ret__.port_number,
        sections=__ret__.sections,
        security=__ret__.security,
        server=__ret__.server,
        smtp=__ret__.smtp,
        tags=__ret__.tags,
        title=__ret__.title,
        type=__ret__.type,
        username=__ret__.username,
        uuid=__ret__.uuid,
        vault=__ret__.vault)


@_utilities.lift_output_func(get_email_account)
def get_email_account_output(title: Optional[pulumi.Input[Optional[str]]] = None,
                             uuid: Optional[pulumi.Input[Optional[str]]] = None,
                             vault: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEmailAccountResult]:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    ...
