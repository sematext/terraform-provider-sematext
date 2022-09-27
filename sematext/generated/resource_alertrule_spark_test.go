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

// TestResourceAlertRuleSparkLifecycle tests resource creation.
func TestResourceAlertRuleSparkLifecycle(t *testing.T) {
	sematext.CommonTestLifecycleAlertRule(t, "sematext_alertrule_spark", "Spark")
}

// TestResourceAlertRuleSparkUpdate tests for resource updates.
func TestResourceAlertRuleSparkUpdate(t *testing.T) {
	sematext.CommonTestUpdateAlertRule(t, "sematext_alertrule_spark", "Spark")
}
