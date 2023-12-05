# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 connect_host: Optional[pulumi.Input[str]] = None,
                 connect_token: Optional[pulumi.Input[str]] = None,
                 service_account_token: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        if connect_host is not None:
            pulumi.set(__self__, "connect_host", connect_host)
        if connect_token is not None:
            pulumi.set(__self__, "connect_token", connect_token)
        if service_account_token is not None:
            pulumi.set(__self__, "service_account_token", service_account_token)
        if vault is not None:
            pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter(name="connectHost")
    def connect_host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connect_host")

    @connect_host.setter
    def connect_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connect_host", value)

    @property
    @pulumi.getter(name="connectToken")
    def connect_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connect_token")

    @connect_token.setter
    def connect_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connect_token", value)

    @property
    @pulumi.getter(name="serviceAccountToken")
    def service_account_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_account_token")

    @service_account_token.setter
    def service_account_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_token", value)

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


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connect_host: Optional[pulumi.Input[str]] = None,
                 connect_token: Optional[pulumi.Input[str]] = None,
                 service_account_token: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a One-password-native-unoffical resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a One-password-native-unoffical resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connect_host: Optional[pulumi.Input[str]] = None,
                 connect_token: Optional[pulumi.Input[str]] = None,
                 service_account_token: Optional[pulumi.Input[str]] = None,
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
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["connect_host"] = connect_host
            __props__.__dict__["connect_token"] = connect_token
            __props__.__dict__["service_account_token"] = service_account_token
            __props__.__dict__["vault"] = vault
        super(Provider, __self__).__init__(
            'one-password-native-unoffical',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="connectHost")
    def connect_host(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "connect_host")

    @property
    @pulumi.getter(name="connectToken")
    def connect_token(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "connect_token")

    @property
    @pulumi.getter(name="serviceAccountToken")
    def service_account_token(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "service_account_token")

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Output[Optional[str]]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

