package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"github.com/sematext/terraform-provider-sematext/sematext"
	"testing"
)

// TestAccResourceLifecycleAppHadoopmrv1 tests resource creation.
func TestAccResourceLifecycleAppHadoopmrv1(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_hadoopmrv1", "Hadoop-MRv1")
}


// TestAccResourceUpdateAppHadoopmrv1 tests for resource updates.
func TestAccResourceUpdateAppHadoopmrv1(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_hadoopmrv1", "Hadoop-MRv1")
}
