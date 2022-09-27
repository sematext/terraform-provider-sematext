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

// TestResourceAlertRuleNodejsLifecycle tests resource creation.
func TestResourceAlertRuleNodejsLifecycle(t *testing.T) {
	sematext.CommonTestLifecycleAlertRule(t, "sematext_alertrule_nodejs", "Node.js")
}

// TestResourceAlertRuleNodejsUpdate tests for resource updates.
func TestResourceAlertRuleNodejsUpdate(t *testing.T) {
	sematext.CommonTestUpdateAlertRule(t, "sematext_alertrule_nodejs", "Node.js")
}
