// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizations_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func init() {
	acctest.RegisterServiceErrorCheckFunc(names.OrganizationsServiceID, testAccErrorCheckSkip)
}

func testAccErrorCheckSkip(t *testing.T) resource.ErrorCheckFunc {
	return acctest.ErrorCheckSkipMessagesContaining(t,
		"MASTER_ACCOUNT_NOT_GOVCLOUD_ENABLED",
	)
}

const (
	organizationIDRegexPattern = `o-[0-9a-z]{10}`
)

func TestAccOrganizations_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Organization": {
			acctest.CtBasic:                     testAccOrganization_basic,
			acctest.CtDisappears:                testAccOrganization_disappears,
			"AwsServiceAccessPrincipals":        testAccOrganization_serviceAccessPrincipals,
			"EnabledPolicyTypes":                testAccOrganization_EnabledPolicyTypes,
			"FeatureSet_Basic":                  testAccOrganization_FeatureSet,
			"FeatureSet_Update":                 testAccOrganization_FeatureSetUpdate,
			"FeatureSet_ForcesNew":              testAccOrganization_FeatureSetForcesNew,
			"DataSource_basic":                  testAccOrganizationDataSource_basic,
			"DataSource_memberAccount":          testAccOrganizationDataSource_memberAccount,
			"DataSource_delegatedAdministrator": testAccOrganizationDataSource_delegatedAdministrator,
			"Identity":                          testAccOrganizationsOrganization_IdentitySerial,
		},
		"Account": {
			acctest.CtBasic:   testAccAccount_basic,
			"CloseOnDeletion": testAccAccount_CloseOnDeletion,
			"ParentId":        testAccAccount_ParentID,
			"tags":            testAccAccount_Tags,
			"GovCloud":        testAccAccount_govCloud,
			"AccountUpdate":   testAccAccount_AccountUpdate,
			"Identity":        testAccOrganizationsAccount_IdentitySerial,
		},
		"OrganizationalUnit": {
			acctest.CtBasic:                      testAccOrganizationalUnit_basic,
			acctest.CtDisappears:                 testAccOrganizationalUnit_disappears,
			"update":                             testAccOrganizationalUnit_update,
			"tags":                               testAccOrganizationsOrganizationalUnit_tagsSerial,
			"DataSource_basic":                   testAccOrganizationalUnitDataSource_basic,
			"DescendantOUsDataSource_basic":      testAccOrganizationalUnitDescendantOUsDataSource_basic,
			"ChildAccountsDataSource_basic":      testAccOrganizationalUnitChildAccountsDataSource_basic,
			"DescendantAccountsDataSource_basic": testAccOrganizationalUnitDescendantAccountsDataSource_basic,
			"PluralDataSource_basic":             testAccOrganizationalUnitsDataSource_basic,
			"Identity":                           testAccOrganizationsOrganizationalUnit_IdentitySerial,
		},
		"Policy": {
			acctest.CtBasic:          testAccPolicy_basic,
			"concurrent":             testAccPolicy_concurrent,
			"Description":            testAccPolicy_description,
			"tags":                   testAccOrganizationsPolicy_tagsSerial,
			"SkipDestroy":            testAccPolicy_skipDestroy,
			acctest.CtDisappears:     testAccPolicy_disappears,
			"Type_AI_OPT_OUT":        testAccPolicy_type_AI_OPT_OUT,
			"Type_Backup":            testAccPolicy_type_Backup,
			"Type_SCP":               testAccPolicy_type_SCP,
			"Type_Tag":               testAccPolicy_type_Tag,
			"ImportAwsManagedPolicy": testAccPolicy_importManagedPolicy,
			"Identity":               testAccOrganizationsPolicy_IdentitySerial,
		},
		"PolicyAttachment": {
			"Account":            testAccPolicyAttachment_Account,
			"OrganizationalUnit": testAccPolicyAttachment_OrganizationalUnit,
			"Root":               testAccPolicyAttachment_Root,
			"SkipDestroy":        testAccPolicyAttachment_skipDestroy,
			acctest.CtDisappears: testAccPolicyAttachment_disappears,
			"Identity":           testAccOrganizationsPolicyAttachment_IdentitySerial,
		},
		"PolicyDataSource": {
			"UnattachedPolicy": testAccPolicyDataSource_UnattachedPolicy,
		},
		"ResourcePolicy": {
			acctest.CtBasic:      testAccResourcePolicy_basic,
			acctest.CtDisappears: testAccResourcePolicy_disappears,
			"tags":               testAccOrganizationsResourcePolicy_tagsSerial,
			"Identity":           testAccOrganizationsResourcePolicy_IdentitySerial,
		},
		"DelegatedAdministrator": {
			acctest.CtBasic:      testAccDelegatedAdministrator_basic,
			acctest.CtDisappears: testAccDelegatedAdministrator_disappears,
			"Identity":           testAccOrganizationsDelegatedAdministrator_IdentitySerial,
		},
		"DelegatedAdministrators": {
			acctest.CtBasic: testAccDelegatedAdministratorsDataSource_basic,
		},
		"DelegatedServices": {
			acctest.CtBasic: testAccDelegatedServicesDataSource_basic,
			"multiple":      testAccDelegatedServicesDataSource_multiple,
		},
		"ResourceTags": {
			acctest.CtBasic: testAccResourceTagsDataSource_basic,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
