# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MembershipArgs', 'Membership']

@pulumi.input_type
class MembershipArgs:
    def __init__(__self__, *,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 member_id: Optional[pulumi.Input[str]] = None,
                 member_name: Optional[pulumi.Input[str]] = None,
                 member_since: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 telephone: Optional[pulumi.Input[str]] = None,
                 website: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Membership resource.
        """
        if expiry_date is not None:
            pulumi.set(__self__, "expiry_date", expiry_date)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if member_id is not None:
            pulumi.set(__self__, "member_id", member_id)
        if member_name is not None:
            pulumi.set(__self__, "member_name", member_name)
        if member_since is not None:
            pulumi.set(__self__, "member_since", member_since)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if pin is not None:
            pulumi.set(__self__, "pin", pin)
        if telephone is not None:
            pulumi.set(__self__, "telephone", telephone)
        if website is not None:
            pulumi.set(__self__, "website", website)

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiry_date")

    @expiry_date.setter
    def expiry_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry_date", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="memberId")
    def member_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member_id")

    @member_id.setter
    def member_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_id", value)

    @property
    @pulumi.getter(name="memberName")
    def member_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member_name")

    @member_name.setter
    def member_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_name", value)

    @property
    @pulumi.getter(name="memberSince")
    def member_since(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member_since")

    @member_since.setter
    def member_since(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_since", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def pin(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pin")

    @pin.setter
    def pin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pin", value)

    @property
    @pulumi.getter
    def telephone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "telephone")

    @telephone.setter
    def telephone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "telephone", value)

    @property
    @pulumi.getter
    def website(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "website")

    @website.setter
    def website(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website", value)


class Membership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 member_id: Optional[pulumi.Input[str]] = None,
                 member_name: Optional[pulumi.Input[str]] = None,
                 member_since: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 telephone: Optional[pulumi.Input[str]] = None,
                 website: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Membership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MembershipArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Membership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 member_id: Optional[pulumi.Input[str]] = None,
                 member_name: Optional[pulumi.Input[str]] = None,
                 member_since: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 telephone: Optional[pulumi.Input[str]] = None,
                 website: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MembershipArgs.__new__(MembershipArgs)

            __props__.__dict__["expiry_date"] = expiry_date
            __props__.__dict__["group"] = group
            __props__.__dict__["member_id"] = member_id
            __props__.__dict__["member_name"] = member_name
            __props__.__dict__["member_since"] = member_since
            __props__.__dict__["notes"] = notes
            __props__.__dict__["pin"] = pin
            __props__.__dict__["telephone"] = telephone
            __props__.__dict__["website"] = website
        super(Membership, __self__).__init__(
            'onepassword:index:Membership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Membership':
        """
        Get an existing Membership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MembershipArgs.__new__(MembershipArgs)

        return Membership(resource_name, opts=opts, __props__=__props__)

