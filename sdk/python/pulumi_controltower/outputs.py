# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'AwsAccountSso',
]

@pulumi.output_type
class AwsAccountSso(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "firstName":
            suggest = "first_name"
        elif key == "lastName":
            suggest = "last_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AwsAccountSso. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AwsAccountSso.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AwsAccountSso.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 email: str,
                 first_name: str,
                 last_name: str):
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "first_name", first_name)
        pulumi.set(__self__, "last_name", last_name)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> str:
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> str:
        return pulumi.get(self, "last_name")


