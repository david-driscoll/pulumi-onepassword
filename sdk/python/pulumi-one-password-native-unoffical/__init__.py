# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .api_credential_item import *
from .bank_account_item import *
from .credit_card_item import *
from .crypto_wallet_item import *
from .database_item import *
from .document_item import *
from .driver_license_item import *
from .email_account_item import *
from .get_api_credential import *
from .get_attachment import *
from .get_bank_account import *
from .get_credit_card import *
from .get_crypto_wallet import *
from .get_database import *
from .get_document import *
from .get_driver_license import *
from .get_email_account import *
from .get_identity import *
from .get_item import *
from .get_login import *
from .get_medical_record import *
from .get_membership import *
from .get_outdoor_license import *
from .get_passport import *
from .get_password import *
from .get_reward_program import *
from .get_secret_reference import *
from .get_secure_note import *
from .get_server import *
from .get_social_security_number import *
from .get_software_license import *
from .get_ssh_key import *
from .get_vault import *
from .get_wireless_router import *
from .identity_item import *
from .item import *
from .login_item import *
from .medical_record_item import *
from .membership_item import *
from .outdoor_license_item import *
from .passport_item import *
from .password_item import *
from .provider import *
from .reward_program_item import *
from .secure_note_item import *
from .server_item import *
from .social_security_number_item import *
from .software_license_item import *
from .ssh_key_item import *
from .wireless_router_item import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi-one-password-native-unoffical.bankaccount as __bankaccount
    bankaccount = __bankaccount
    import pulumi-one-password-native-unoffical.creditcard as __creditcard
    creditcard = __creditcard
    import pulumi-one-password-native-unoffical.cryptowallet as __cryptowallet
    cryptowallet = __cryptowallet
    import pulumi-one-password-native-unoffical.emailaccount as __emailaccount
    emailaccount = __emailaccount
    import pulumi-one-password-native-unoffical.identity as __identity
    identity = __identity
    import pulumi-one-password-native-unoffical.medicalrecord as __medicalrecord
    medicalrecord = __medicalrecord
    import pulumi-one-password-native-unoffical.rewardprogram as __rewardprogram
    rewardprogram = __rewardprogram
    import pulumi-one-password-native-unoffical.server as __server
    server = __server
    import pulumi-one-password-native-unoffical.softwarelicense as __softwarelicense
    softwarelicense = __softwarelicense
else:
    bankaccount = _utilities.lazy_import('pulumi-one-password-native-unoffical.bankaccount')
    creditcard = _utilities.lazy_import('pulumi-one-password-native-unoffical.creditcard')
    cryptowallet = _utilities.lazy_import('pulumi-one-password-native-unoffical.cryptowallet')
    emailaccount = _utilities.lazy_import('pulumi-one-password-native-unoffical.emailaccount')
    identity = _utilities.lazy_import('pulumi-one-password-native-unoffical.identity')
    medicalrecord = _utilities.lazy_import('pulumi-one-password-native-unoffical.medicalrecord')
    rewardprogram = _utilities.lazy_import('pulumi-one-password-native-unoffical.rewardprogram')
    server = _utilities.lazy_import('pulumi-one-password-native-unoffical.server')
    softwarelicense = _utilities.lazy_import('pulumi-one-password-native-unoffical.softwarelicense')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "one-password-native-unoffical",
  "mod": "index",
  "fqn": "pulumi-one-password-native-unoffical",
  "classes": {
   "one-password-native-unoffical:index:APICredentialItem": "APICredentialItem",
   "one-password-native-unoffical:index:BankAccountItem": "BankAccountItem",
   "one-password-native-unoffical:index:CreditCardItem": "CreditCardItem",
   "one-password-native-unoffical:index:CryptoWalletItem": "CryptoWalletItem",
   "one-password-native-unoffical:index:DatabaseItem": "DatabaseItem",
   "one-password-native-unoffical:index:DocumentItem": "DocumentItem",
   "one-password-native-unoffical:index:DriverLicenseItem": "DriverLicenseItem",
   "one-password-native-unoffical:index:EmailAccountItem": "EmailAccountItem",
   "one-password-native-unoffical:index:IdentityItem": "IdentityItem",
   "one-password-native-unoffical:index:Item": "Item",
   "one-password-native-unoffical:index:LoginItem": "LoginItem",
   "one-password-native-unoffical:index:MedicalRecordItem": "MedicalRecordItem",
   "one-password-native-unoffical:index:MembershipItem": "MembershipItem",
   "one-password-native-unoffical:index:OutdoorLicenseItem": "OutdoorLicenseItem",
   "one-password-native-unoffical:index:PassportItem": "PassportItem",
   "one-password-native-unoffical:index:PasswordItem": "PasswordItem",
   "one-password-native-unoffical:index:RewardProgramItem": "RewardProgramItem",
   "one-password-native-unoffical:index:SSHKeyItem": "SSHKeyItem",
   "one-password-native-unoffical:index:SecureNoteItem": "SecureNoteItem",
   "one-password-native-unoffical:index:ServerItem": "ServerItem",
   "one-password-native-unoffical:index:SocialSecurityNumberItem": "SocialSecurityNumberItem",
   "one-password-native-unoffical:index:SoftwareLicenseItem": "SoftwareLicenseItem",
   "one-password-native-unoffical:index:WirelessRouterItem": "WirelessRouterItem"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "one-password-native-unoffical",
  "token": "pulumi:providers:one-password-native-unoffical",
  "fqn": "pulumi-one-password-native-unoffical",
  "class": "Provider"
 }
]
"""
)
