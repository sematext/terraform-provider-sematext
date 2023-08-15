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

// TestAccResourceLifecycleAppClickhouse tests resource creation.
func TestAccResourceLifecycleAppClickhouse(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_clickhouse", "ClickHouse")
}


// TestAccResourceUpdateAppClickhouse tests for resource updates.
func TestAccResourceUpdateAppClickhouse(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_clickhouse", "ClickHouse")
}
