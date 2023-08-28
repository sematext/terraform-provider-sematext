package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceTestFixtureAWS a test fixture representing a resource - AWS EBS, AWS EC2, AWS ELB
type ResourceTestFixtureAWS struct {
	ResourceType      string
	ResourceName      string
	AppType           string
	Name              string
	StatePath         string
	PlanID            int
	DiscountCode      string
	AppToken          AppTokenType
	AwsRegion         string
	AwsSecretKey      string
	AwsAccessKey      string
	AwsFetchFrequency string
}

func (rtf *ResourceTestFixtureAWS) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AppType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = stcloud.AssignPlanID(rtf.AppType)
	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.AppToken = AppTokenType{
		Names: []string{"Staging", "Development", "Production"},
	}

	rtf.AwsRegion = os.Getenv("AWS_REGION")
	rtf.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	rtf.AwsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	rtf.AwsFetchFrequency = "FIVE_MINUTES"
}

func (rtf *ResourceTestFixtureAWS) toHCL() string {

	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		billing_plan_id = %d
		discount_code = "%s"
		apptoken {
			names = %s
		}
		aws_region = "%s"
		aws_secret_key = "%s"
		aws_access_key = "%s"
		aws_fetch_frequency = "%s"
	}
	`,
		rtf.ResourceType,
		rtf.ResourceName,
		rtf.Name,
		rtf.PlanID,
		rtf.DiscountCode,
		arrayLiteralString(rtf.AppToken.Names),
		rtf.AwsRegion,
		rtf.AwsSecretKey,
		rtf.AwsAccessKey,
		rtf.AwsFetchFrequency,
	)

}

var providerProtoV6ProviderFactoriesAWS = map[string]func() (tfprotov6.ProviderServer, error){
	"sematext_cloud": providerserver.NewProtocol6WithError(New("test")()),
}

func providerPreCheckTestAWS(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	token := os.Getenv("SEMATEXT_API_KEY")
	if !IsValidUUID(token) {
		t.Fatal("ERROR : SEMATEXT_API_KEY environment not set correctly")
	}

	region := os.Getenv("SEMATEXT_REGION")
	if !stcloud.IsValidSematextRegion(region) {
		t.Fatal("ERROR : SEMATEXT_REGION environment not set correctly")
	}

	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		t.Fatal("ERROR : AWS_ACCESS_KEY_ID must be set for acceptance tests")
	}

	if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		t.Fatal("ERROR : AWS_SECRET_ACCESS_KEY must be set for acceptance tests")
	}

	if os.Getenv("AWS_REGION") == "" {
		t.Fatal("ERROR : AWS_REGION must be set for acceptance tests")
	}

}

func TestAccResourceAppAWS(t *testing.T, resourceType string, appType string) {

	testname := resourceType + ".test"

	fixture0 := &ResourceTestFixtureAWS{}
	fixture0.hydrate(resourceType, appType)

	fixture1 := fixture0

	fixture1.Name = fixture0.Name + "_2"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { providerPreCheckTestAWS(t) },
		ProtoV6ProviderFactories: providerProtoV6ProviderFactoriesAWS,
		Steps: []resource.TestStep{
			// Create and Read testing
			{

				Config: fixture0.toHCL(),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testname, "name", fixture0.ResourceName),
					resource.TestCheckResourceAttr(testname, "aws_region", fixture0.AwsRegion),
					resource.TestCheckResourceAttr(testname, "aws_secret_key", fixture0.AwsSecretKey),
					resource.TestCheckResourceAttr(testname, "aws_access_key", fixture0.AwsAccessKey),
					resource.TestCheckResourceAttr(testname, "aws_fetch_frequency", fixture0.AwsFetchFrequency),
					resource.TestCheckResourceAttr(testname, "billing_plan_id", strconv.Itoa(fixture0.PlanID)),
					resource.TestCheckResourceAttr(testname, "discount_code", fixture0.DiscountCode),
				),
			},
			// ImportState testing
			{
				ResourceName:      testname,
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				//ImportStateVerifyIgnore: []string{"configurable_attribute", "defaulted"},
			},
			// Update and Read testing
			{
				Config: fixture1.toHCL(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testname, "name", fixture1.Name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
