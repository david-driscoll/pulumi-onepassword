# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BankAccountArgs', 'BankAccount']

@pulumi.input_type
class BankAccountArgs:
    def __init__(__self__, *,
                 account_number: Optional[pulumi.Input[str]] = None,
                 bank_name: Optional[pulumi.Input[str]] = None,
                 branch_information: Optional[Any] = None,
                 iban: Optional[pulumi.Input[str]] = None,
                 name_on_account: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 routing_number: Optional[pulumi.Input[str]] = None,
                 swift: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BankAccount resource.
        """
        if account_number is not None:
            pulumi.set(__self__, "account_number", account_number)
        if bank_name is not None:
            pulumi.set(__self__, "bank_name", bank_name)
        if branch_information is not None:
            pulumi.set(__self__, "branch_information", branch_information)
        if iban is not None:
            pulumi.set(__self__, "iban", iban)
        if name_on_account is not None:
            pulumi.set(__self__, "name_on_account", name_on_account)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if pin is not None:
            pulumi.set(__self__, "pin", pin)
        if routing_number is not None:
            pulumi.set(__self__, "routing_number", routing_number)
        if swift is not None:
            pulumi.set(__self__, "swift", swift)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accountNumber")
    def account_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account_number")

    @account_number.setter
    def account_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_number", value)

    @property
    @pulumi.getter(name="bankName")
    def bank_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bank_name")

    @bank_name.setter
    def bank_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bank_name", value)

    @property
    @pulumi.getter(name="branchInformation")
    def branch_information(self) -> Optional[Any]:
        return pulumi.get(self, "branch_information")

    @branch_information.setter
    def branch_information(self, value: Optional[Any]):
        pulumi.set(self, "branch_information", value)

    @property
    @pulumi.getter
    def iban(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iban")

    @iban.setter
    def iban(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iban", value)

    @property
    @pulumi.getter(name="nameOnAccount")
    def name_on_account(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name_on_account")

    @name_on_account.setter
    def name_on_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_on_account", value)

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
    @pulumi.getter(name="routingNumber")
    def routing_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "routing_number")

    @routing_number.setter
    def routing_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_number", value)

    @property
    @pulumi.getter
    def swift(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "swift")

    @swift.setter
    def swift(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "swift", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class BankAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_number: Optional[pulumi.Input[str]] = None,
                 bank_name: Optional[pulumi.Input[str]] = None,
                 branch_information: Optional[Any] = None,
                 iban: Optional[pulumi.Input[str]] = None,
                 name_on_account: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 routing_number: Optional[pulumi.Input[str]] = None,
                 swift: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BankAccount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BankAccountArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BankAccount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BankAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BankAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_number: Optional[pulumi.Input[str]] = None,
                 bank_name: Optional[pulumi.Input[str]] = None,
                 branch_information: Optional[Any] = None,
                 iban: Optional[pulumi.Input[str]] = None,
                 name_on_account: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 routing_number: Optional[pulumi.Input[str]] = None,
                 swift: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = BankAccountArgs.__new__(BankAccountArgs)

            __props__.__dict__["account_number"] = account_number
            __props__.__dict__["bank_name"] = bank_name
            __props__.__dict__["branch_information"] = branch_information
            __props__.__dict__["iban"] = iban
            __props__.__dict__["name_on_account"] = name_on_account
            __props__.__dict__["notes"] = notes
            __props__.__dict__["pin"] = pin
            __props__.__dict__["routing_number"] = routing_number
            __props__.__dict__["swift"] = swift
            __props__.__dict__["type"] = type
        super(BankAccount, __self__).__init__(
            'onepassword:index:BankAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BankAccount':
        """
        Get an existing BankAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BankAccountArgs.__new__(BankAccountArgs)

        return BankAccount(resource_name, opts=opts, __props__=__props__)

