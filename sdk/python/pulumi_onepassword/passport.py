# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PassportArgs', 'Passport']

@pulumi.input_type
class PassportArgs:
    def __init__(__self__, *,
                 date_of_birth: Optional[pulumi.Input[str]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issuing_authority: Optional[pulumi.Input[str]] = None,
                 issuing_country: Optional[pulumi.Input[str]] = None,
                 nationality: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 place_of_birth: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Passport resource.
        """
        if date_of_birth is not None:
            pulumi.set(__self__, "date_of_birth", date_of_birth)
        if expiry_date is not None:
            pulumi.set(__self__, "expiry_date", expiry_date)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if gender is not None:
            pulumi.set(__self__, "gender", gender)
        if issued_on is not None:
            pulumi.set(__self__, "issued_on", issued_on)
        if issuing_authority is not None:
            pulumi.set(__self__, "issuing_authority", issuing_authority)
        if issuing_country is not None:
            pulumi.set(__self__, "issuing_country", issuing_country)
        if nationality is not None:
            pulumi.set(__self__, "nationality", nationality)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if number is not None:
            pulumi.set(__self__, "number", number)
        if place_of_birth is not None:
            pulumi.set(__self__, "place_of_birth", place_of_birth)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="dateOfBirth")
    def date_of_birth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "date_of_birth")

    @date_of_birth.setter
    def date_of_birth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_of_birth", value)

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiry_date")

    @expiry_date.setter
    def expiry_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry_date", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter
    def gender(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gender")

    @gender.setter
    def gender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gender", value)

    @property
    @pulumi.getter(name="issuedOn")
    def issued_on(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issued_on")

    @issued_on.setter
    def issued_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issued_on", value)

    @property
    @pulumi.getter(name="issuingAuthority")
    def issuing_authority(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuing_authority")

    @issuing_authority.setter
    def issuing_authority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuing_authority", value)

    @property
    @pulumi.getter(name="issuingCountry")
    def issuing_country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuing_country")

    @issuing_country.setter
    def issuing_country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuing_country", value)

    @property
    @pulumi.getter
    def nationality(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nationality")

    @nationality.setter
    def nationality(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nationality", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter(name="placeOfBirth")
    def place_of_birth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "place_of_birth")

    @place_of_birth.setter
    def place_of_birth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "place_of_birth", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Passport(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 date_of_birth: Optional[pulumi.Input[str]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issuing_authority: Optional[pulumi.Input[str]] = None,
                 issuing_country: Optional[pulumi.Input[str]] = None,
                 nationality: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 place_of_birth: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Passport resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PassportArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Passport resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PassportArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PassportArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 date_of_birth: Optional[pulumi.Input[str]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issuing_authority: Optional[pulumi.Input[str]] = None,
                 issuing_country: Optional[pulumi.Input[str]] = None,
                 nationality: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 place_of_birth: Optional[pulumi.Input[str]] = None,
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
            __props__ = PassportArgs.__new__(PassportArgs)

            __props__.__dict__["date_of_birth"] = date_of_birth
            __props__.__dict__["expiry_date"] = expiry_date
            __props__.__dict__["full_name"] = full_name
            __props__.__dict__["gender"] = gender
            __props__.__dict__["issued_on"] = issued_on
            __props__.__dict__["issuing_authority"] = issuing_authority
            __props__.__dict__["issuing_country"] = issuing_country
            __props__.__dict__["nationality"] = nationality
            __props__.__dict__["notes"] = notes
            __props__.__dict__["number"] = number
            __props__.__dict__["place_of_birth"] = place_of_birth
            __props__.__dict__["type"] = type
        super(Passport, __self__).__init__(
            'onepassword:index:Passport',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Passport':
        """
        Get an existing Passport resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PassportArgs.__new__(PassportArgs)

        return Passport(resource_name, opts=opts, __props__=__props__)

