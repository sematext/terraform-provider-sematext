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

// TestResourceAlertRuleHadoopmrv1Lifecycle tests resource creation.
func TestResourceAlertRuleHadoopmrv1Lifecycle(t *testing.T) {
	sematext.CommonTestLifecycleAlertRule(t, "sematext_alertrule_hadoopmrv1", "Hadoop-MRv1")
}

// TestResourceAlertRuleHadoopmrv1Update tests for resource updates.
func TestResourceAlertRuleHadoopmrv1Update(t *testing.T) {
	sematext.CommonTestUpdateAlertRule(t, "sematext_alertrule_hadoopmrv1", "Hadoop-MRv1")
}
