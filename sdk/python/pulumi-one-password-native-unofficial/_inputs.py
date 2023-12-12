# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = [
    'FieldArgs',
    'PasswordRecipeArgs',
    'ReferenceArgs',
    'SectionArgs',
    'UrlArgs',
]

@pulumi.input_type
class FieldArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[str],
                 label: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['FieldType']] = None):
        pulumi.set(__self__, "value", value)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if type is None:
            type = 'STRING'
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['FieldType']]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['FieldType']]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class PasswordRecipeArgs:
    def __init__(__self__, *,
                 length: pulumi.Input[int],
                 digits: Optional[pulumi.Input[bool]] = None,
                 letters: Optional[pulumi.Input[bool]] = None,
                 symbols: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "length", length)
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if letters is not None:
            pulumi.set(__self__, "letters", letters)
        if symbols is not None:
            pulumi.set(__self__, "symbols", symbols)

    @property
    @pulumi.getter
    def length(self) -> pulumi.Input[int]:
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: pulumi.Input[int]):
        pulumi.set(self, "length", value)

    @property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "digits", value)

    @property
    @pulumi.getter
    def letters(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "letters")

    @letters.setter
    def letters(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "letters", value)

    @property
    @pulumi.getter
    def symbols(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "symbols")

    @symbols.setter
    def symbols(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "symbols", value)


@pulumi.input_type
class ReferenceArgs:
    def __init__(__self__, *,
                 item_id: pulumi.Input[str]):
        pulumi.set(__self__, "item_id", item_id)

    @property
    @pulumi.getter(name="itemId")
    def item_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "item_id")

    @item_id.setter
    def item_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "item_id", value)


@pulumi.input_type
class SectionArgs:
    def __init__(__self__, *,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input['ReferenceArgs']]]] = None):
        if attachments is not None:
            pulumi.set(__self__, "attachments", attachments)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if references is not None:
            pulumi.set(__self__, "references", references)

    @property
    @pulumi.getter
    def attachments(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]:
        return pulumi.get(self, "attachments")

    @attachments.setter
    def attachments(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]):
        pulumi.set(self, "attachments", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def references(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReferenceArgs']]]]:
        return pulumi.get(self, "references")

    @references.setter
    def references(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReferenceArgs']]]]):
        pulumi.set(self, "references", value)


@pulumi.input_type
class UrlArgs:
    def __init__(__self__, *,
                 href: pulumi.Input[str],
                 primary: pulumi.Input[bool],
                 label: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "href", href)
        if primary is None:
            primary = False
        pulumi.set(__self__, "primary", primary)
        if label is not None:
            pulumi.set(__self__, "label", label)

    @property
    @pulumi.getter
    def href(self) -> pulumi.Input[str]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: pulumi.Input[str]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def primary(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: pulumi.Input[bool]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)


