package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/sematext/terraform-provider-sematext/sematext"
	"testing"
)

// TestResourceAlertRuleAwsec2Lifecycle tests resource creation.
func TestResourceAlertRuleAwsec2Lifecycle(t *testing.T) {
	sematext.CommonTestLifecycleAlertRule(t, "sematext_alertrule_awsec2", "AWS EC2")
}

// TestResourceAlertRuleAwsec2Update tests for resource updates.
func TestResourceAlertRuleAwsec2Update(t *testing.T) {
	sematext.CommonTestUpdateAlertRule(t, "sematext_alertrule_awsec2", "AWS EC2")
}
