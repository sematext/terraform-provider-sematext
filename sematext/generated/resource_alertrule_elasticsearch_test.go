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

// TestResourceAlertRuleElasticsearchLifecycle tests resource creation.
func TestResourceAlertRuleElasticsearchLifecycle(t *testing.T) {
	sematext.CommonTestLifecycleAlertRule(t, "sematext_alertrule_elasticsearch", "Elastic Search")
}

// TestResourceAlertRuleElasticsearchUpdate tests for resource updates.
func TestResourceAlertRuleElasticsearchUpdate(t *testing.T) {
	sematext.CommonTestUpdateAlertRule(t, "sematext_alertrule_elasticsearch", "Elastic Search")
}
