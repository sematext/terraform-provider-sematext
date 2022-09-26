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

// TestAccResourceLifecycleAppHadoopyarn tests resource creation.
func TestAccResourceLifecycleAppHadoopyarn(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_hadoopyarn", "Hadoop-YARN")
}

// TestAccResourceUpdateAppHadoopyarn tests for resource updates.
func TestAccResourceUpdateAppHadoopyarn(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_hadoopyarn", "Hadoop-YARN")
}
