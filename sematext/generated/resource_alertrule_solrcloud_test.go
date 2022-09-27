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

// TestResourceAlertRuleSolrcloudLifecycle tests resource creation.
func TestResourceAlertRuleSolrcloudLifecycle(t *testing.T) {
	sematext.CommonTestLifecycleAlertRule(t, "sematext_alertrule_solrcloud", "SolrCloud")
}

// TestResourceAlertRuleSolrcloudUpdate tests for resource updates.
func TestResourceAlertRuleSolrcloudUpdate(t *testing.T) {
	sematext.CommonTestUpdateAlertRule(t, "sematext_alertrule_solrcloud", "SolrCloud")
}
