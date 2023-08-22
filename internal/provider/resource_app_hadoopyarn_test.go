package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppHadoopyarn tests resource creation.
func TestAccResourceLifecycleAppHadoopyarn(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_hadoopyarn", "Hadoop-YARN")
}


// TestAccResourceUpdateAppHadoopyarn tests for resource updates.
func TestAccResourceUpdateAppHadoopyarn(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_hadoopyarn", "Hadoop-YARN")
}
