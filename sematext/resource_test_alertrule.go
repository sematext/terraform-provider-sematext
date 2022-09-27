package sematext

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

type AlertRuleTokenType struct {
	Names []string
}

type ResourceTestFixtureAlertRule interface {
	hydrate(resourceType string, appType string)
	toHCL() string
}

// ResourceTestFixtureAlertRule a common test fixture representing most AlertRule resources
type ResourceTestFixtureAlertRuleDefault struct {
	ResourceType   string
	ResourceName   string
	AlertRuleType  string
	Name           string
	StatePath      string
	PlanID         int
	DiscountCode   string
	AlertRuleToken AlertRuleTokenType
}

// ResourceTestFixtureAlertRuleAWS a test fixture representing a resource - AWS EBS, AWS EC2, AWS ELB
type ResourceTestFixtureAlertRuleAWS struct {
	ResourceType      string
	ResourceName      string
	AlertRuleType     string
	Name              string
	StatePath         string
	PlanID            int
	DiscountCode      string
	AlertRuleToken    AlertRuleTokenType
	AwsRegion         string
	AwsSecretKey      string
	AwsAccessKey      string
	AwsFetchFrequency string
}

func (rtf *ResourceTestFixtureAlertRuleDefault) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AlertRuleType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = stcloud.AssignPlanID(rtf.AlertRuleType)

	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.AlertRuleToken = AlertRuleTokenType{
		Names: []string{"Staging", "Development", "Production"},
	}

}

func (rtf *ResourceTestFixtureAlertRuleAWS) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AlertRuleType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = stcloud.AssignPlanID(rtf.AlertRuleType)
	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.AlertRuleToken = AlertRuleTokenType{
		Names: []string{"Staging", "Development", "Production"},
	}

	rtf.AwsRegion = os.Getenv("AWS_REGION")
	rtf.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	rtf.AwsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	rtf.AwsFetchFrequency = "FIVE_MINUTES"
}

func (rtf *ResourceTestFixtureAlertRuleDefault) toHCL() string {

	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		billing_plan_id = %d
		discount_code = "%s"
		apptoken {
			names = %s
		}
	}
	`,
		rtf.ResourceType,
		rtf.ResourceName,
		rtf.Name,
		rtf.PlanID,
		rtf.DiscountCode,
		arrayLiteralString(rtf.AlertRuleToken.Names),
	)

}

func (rtf *ResourceTestFixtureAlertRuleAWS) toHCL() string {

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
		arrayLiteralString(rtf.AlertRuleToken.Names),
		rtf.AwsRegion,
		rtf.AwsSecretKey,
		rtf.AwsAccessKey,
		rtf.AwsFetchFrequency,
	)

}

// CommonAlertRuleBasicTest is a common test of resource creation.
func ResourceTestLifecycleAlertRule(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonAlertRuleBasicTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf := ResourceTestFixtureAlertRuleAWS{}
		rtf.hydrate(resourceType, appType)

		fixture := rtf.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmDestructionAWSAlertRule(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmCreationAWSAlertRule(rtf),
						resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					),
				},
			},
		})

	default:

		fmt.Println("---------------------------------------")
		fmt.Println("default identified")
		fmt.Println("---------------------------------------")

		rtf := ResourceTestFixtureAlertRuleDefault{}
		rtf.hydrate(resourceType, appType)
		fixture := rtf.toHCL()

		fmt.Println(fixture)

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmDestructionDefaultAlertRule(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmCreationDefaultAlertRule(rtf),
						resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					),
				},
			},
		})
	}
}

// CommonTestUpdateAlertRule tests for resource updates.
func ResourceTestUpdateAlertRule(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonAlertRuleUpdateTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {
	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf0 := ResourceTestFixtureAlertRuleAWS{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0

		fixture1 := rtf1.toHCL()

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmDestructionAWSAlertRule(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmCreationAWSAlertRule(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmCreationAWSAlertRule(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					),
				},
			},
		})

	default:

		fmt.Println("---------------------------------------")
		fmt.Println("default identified")
		fmt.Println("---------------------------------------")

		rtf0 := ResourceTestFixtureAlertRuleDefault{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0
		fixture1 := rtf1.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmDestructionDefaultAlertRule(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmCreationDefaultAlertRule(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmCreationDefaultAlertRule(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					),
				},
			},
		})
	}

}

// ConfirmCreationDefaultAlertRule checks the AlertRule ID exists in both state and stcloud.
func ConfirmCreationDefaultAlertRule(rtf ResourceTestFixtureAlertRuleDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmCreationDefaultAlertRule func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.AlertRule
		var rs *terraform.ResourceState
		var appResponse stcloud.AlertRuleResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		/*
			fmt.Println("---------------------------------------")
			fmt.Println("rtf.StatePath")
			fmt.Println(rtf.StatePath)
			spew.Dump(s.RootModule().Resources)
			fmt.Println("---------------------------------------")
		*/
		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmCreationDefaultAlertRule : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if appResponse, _, err = client.AlertsApi.GetUsingGET(context.Background(), id); err != nil { // TODO need  GET-by-id
			return fmt.Errorf("ConfirmCreationDefaultAlertRule : Error in checking AlertRule %s, %s", rtf.StatePath, err)
		}

		if app, err = extractAlertRule(appResponse); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmCreationDefaultAlertRule : Error in checking AlertRule %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmCreationAWSAlertRule checks the AlertRule ID exists in both state and stcloud.
func ConfirmCreationAWSAlertRule(rtf ResourceTestFixtureAlertRuleAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmCreationAWSAlertRule func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.AlertRule
		var rs *terraform.ResourceState
		var appResponse stcloud.AlertRuleResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmCreationAWSAlertRule : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if appResponse, _, err = client.AlertsApi.GetUsingGET(context.Background(), id); err != nil { // TODO - need a get by id
			return fmt.Errorf("ConfirmCreationAWSAlertRule : Error in checking AlertRule %s, %s", rtf.StatePath, err)
		}

		if app, err = extractAlertRule(appResponse); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmCreationAWSAlertRule : Error in checking AlertRule %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmDestructionDefaultAlertRule checks the AlertRule ID has been removed from state and the API has marked the app as DELETED.
func ConfirmDestructionDefaultAlertRule(rtf ResourceTestFixtureAlertRuleDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmAlertRuleDestructionDefault func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var alertRule *stcloud.AlertRule
		var alertRuleResponse stcloud.AlertRuleResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		//spew.Dump(s.RootModule().Resources)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") {
				continue
			}
			if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
				return err
			}

			time.Sleep(time.Second * 10) // TODO - workaround for cache latency on SC API, no longer required, confirm and obsolete after MVP.

			if alertRuleResponse, _, err = client.AlertsApi.GetUsingGET(context.Background(), id); err != nil { // TODO need a get by id
				return fmt.Errorf("ConfirmDestructionDefaultAlertRule : Failed to retrieve Sematext Cloud AlertRule %s, %s", rtf.StatePath, err)
			}

			if alertRule, err = extractAlertRule(alertRuleResponse); err != nil {
				return err
			}

			if alertRule == nil {
				return fmt.Errorf("ConfirmDestructionDefaultAlertRule : Missing AlertRule in genericAPIResponse checking Sematext Cloud AlertRule %s", rtf.StatePath)
			}

			if alertRuleResponse.Status != "DELETED" {
				return fmt.Errorf("ConfirmDestructionDefaultAlertRule : Unexpected status in checking Sematext Cloud AlertRule %s : Expected DELETED, got %s ", rtf.StatePath, alertRuleResponse.Status)
			}

		}

		return nil
	}
}

// ConfirmDestructionAWSAlertRule checks the AlertRule ID exists in both state and API and is marked as DELETED.
func ConfirmDestructionAWSAlertRule(rtf ResourceTestFixtureAlertRuleAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmDestructionAWSAlertRule func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.AlertRule
		var appResponse stcloud.AlertRuleResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO Consider shift to template and make check tighter after MVP
				continue
			}
			if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
				return err
			}

			if appResponse, _, err = client.AlertsApi.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmDestructionAWSAlertRule : Failed to retrieve Sematext Cloud AlertRule %s, %s", rtf.StatePath, err)
			}

			if app, err = extractAlertRule(appResponse); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmDestructionAWSAlertRule : Missing app in genericAPIResponse checking Sematext Cloud AlertRule %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmDestructionAWSAlertRule : Unexpected status in checking Sematext Cloud AlertRule %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}

func testSetProviderConfiguration(t *testing.T) {

	p := Provider()

	config := map[string]interface{}{
		"sematext_region": "US",
	}
	err := p.Configure(context.Background(), terraform.NewResourceConfigRaw(config))
	if err != nil {
		t.Fatal(err)
	}
}

func testCheckProviderFixture(t *testing.T) {

	token := os.Getenv("SEMATEXT_API_KEY")
	if !IsValidUUID(token) {
		t.Fatal("ERROR : SEMATEXT_API_KEY environment not set correctly")
	}

	region := os.Getenv("SEMATEXT_REGION")
	if !IsValidSematextRegion(region) {
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
