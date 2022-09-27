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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/sematext/terraform-provider-sematext/sematext/generated"
)

var testProviders map[string]*schema.Provider
var testProvider schema.Provider
var testProviderFunc func() *schema.Provider

type NotificationHookTokenType struct {
	Names []string
}

type ResourceTestFixtureNotificationHook interface {
	hydrate(resourceType string, appType string)
	toHCL() string
}

// ResourceTestFixtureNotificationHook a common test fixture representing most NotificationHook resources
type ResourceTestFixtureNotificationHookDefault struct {
	ResourceType          string
	ResourceName          string
	NotificationHookType  string
	Name                  string
	StatePath             string
	PlanID                int
	DiscountCode          string
	NotificationHookToken NotificationHookTokenType
}

// ResourceTestFixtureNotificationHookAWS a test fixture representing a resource - AWS EBS, AWS EC2, AWS ELB
type ResourceTestFixtureNotificationHookAWS struct {
	ResourceType          string
	ResourceName          string
	NotificationHookType  string
	Name                  string
	StatePath             string
	PlanID                int
	DiscountCode          string
	NotificationHookToken NotificationHookTokenType
	AwsRegion             string
	AwsSecretKey          string
	AwsAccessKey          string
	AwsFetchFrequency     string
}

func (rtf *ResourceTestFixtureNotificationHookDefault) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.NotificationHookType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = stcloud.AssignPlanID(rtf.NotificationHookType)

	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.NotificationHookToken = NotificationHookTokenType{
		Names: []string{"Staging", "Development", "Production"},
	}

}

func (rtf *ResourceTestFixtureNotificationHookAWS) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.NotificationHookType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = stcloud.AssignPlanID(rtf.NotificationHookType)
	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.NotificationHookToken = NotificationHookTokenType{
		Names: []string{"Staging", "Development", "Production"},
	}

	rtf.AwsRegion = os.Getenv("AWS_REGION")
	rtf.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	rtf.AwsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	rtf.AwsFetchFrequency = "FIVE_MINUTES"
}

func (rtf *ResourceTestFixtureNotificationHookDefault) toHCL() string {

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
		arrayLiteralString(rtf.NotificationHookToken.Names),
	)

}

func (rtf *ResourceTestFixtureNotificationHookAWS) toHCL() string {

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
		arrayLiteralString(rtf.NotificationHookToken.Names),
		rtf.AwsRegion,
		rtf.AwsSecretKey,
		rtf.AwsAccessKey,
		rtf.AwsFetchFrequency,
	)

}

// CommonNotificationHookBasicTest is a common test of resource creation.
func CommonTestLifecycleNotificationHook(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonNotificationHookBasicTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf := ResourceTestFixtureNotificationHookAWS{}
		rtf.hydrate(resourceType, appType)

		fixture := rtf.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmNotificationHookDestructionAWS(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmNotificationHookCreationAWS(rtf),
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

		rtf := ResourceTestFixtureNotificationHookDefault{}
		rtf.hydrate(resourceType, appType)
		fixture := rtf.toHCL()

		fmt.Println(fixture)

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmNotificationHookDestructionDefault(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmNotificationHookCreationDefault(rtf),
						resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					),
				},
			},
		})
	}
}

// CommonNotificationHookUpdateTest tests for resource updates.
func CommonTestUpdateNotificationHook(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonNotificationHookUpdateTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {
	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf0 := ResourceTestFixtureNotificationHookAWS{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0

		fixture1 := rtf1.toHCL()

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmNotificationHookDestructionAWS(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmNotificationHookCreationAWS(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmNotificationHookCreationAWS(rtf1),
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

		rtf0 := ResourceTestFixtureNotificationHookDefault{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0
		fixture1 := rtf1.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmNotificationHookDestructionDefault(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmNotificationHookCreationDefault(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmNotificationHookCreationDefault(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					),
				},
			},
		})
	}

}

// ConfirmNotificationHookCreationDefault checks the NotificationHook ID exists in both state and stcloud.
func ConfirmNotificationHookCreationDefault(rtf ResourceTestFixtureNotificationHookDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmNotificationHookCreationDefault func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.NotificationHook
		var rs *terraform.ResourceState
		var appResponse stcloud.NotificationHookResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		/*
			fmt.Println("---------------------------------------")
			fmt.Println("rtf.StatePath")
			fmt.Println(rtf.StatePath)
			spew.Dump(s.RootModule().Resources)
			fmt.Println("---------------------------------------")
		*/
		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmNotificationHookCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if appResponse, _, err = client.NotificationHooksAPI.GetUsingGET(context.Background(), id); err != nil {
			return fmt.Errorf("ConfirmNotificationHookCreation : Error in checking NotificationHook %s, %s", rtf.StatePath, err)
		}

		if app, err = extractNotificationHook(&appResponse); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmNotificationHookCreation : Error in checking NotificationHook %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmNotificationHookCreationAWS checks the NotificationHook ID exists in both state and stcloud.
func ConfirmNotificationHookCreationAWS(rtf ResourceTestFixtureNotificationHookAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmNotificationHookCreationAWS func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.NotificationHook
		var rs *terraform.ResourceState
		var appResponse stcloud.NotificationHookResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmNotificationHookCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if appResponse, _, err = client.NotificationHooksAPI.GetUsingGET(context.Background(), id); err != nil {
			return fmt.Errorf("ConfirmNotificationHookCreationAWS : Error in checking NotificationHook %s, %s", rtf.StatePath, err)
		}

		if app, err = extractNotificationHook(&appResponse); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmNotificationHookCreation : Error in checking NotificationHook %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmNotificationHookDestructionDefault checks the NotificationHook ID has been removed from state and the API has marked the app as DELETED.
func ConfirmNotificationHookDestructionDefault(rtf ResourceTestFixtureNotificationHookDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmNotificationHookDestructionDefault func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.NotificationHook
		var appResponse stcloud.NotificationHookResponse

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

			if appResponse, _, err = client.NotificationHooksAPI.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmNotificationHookDestructionDefault : Failed to retrieve Sematext Cloud NotificationHook %s, %s", rtf.StatePath, err)
			}

			if app, err = extractNotificationHook(&appResponse); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmNotificationHookDestructionDefault : Missing app in genericAPIResponse checking Sematext Cloud NotificationHook %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmNotificationHookDestructionDefault : Unexpected status in checking Sematext Cloud NotificationHook %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}

// ConfirmNotificationHookDestructionAWS checks the NotificationHook ID exists in both state and API and is marked as DELETED.
func ConfirmNotificationHookDestructionAWS(rtf ResourceTestFixtureNotificationHookAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmNotificationHookDestructionDefaultAWS func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.NotificationHook
		var appResponse stcloud.NotificationHookResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO Consider shift to template and make check tighter after MVP
				continue
			}
			if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
				return err
			}

			if appResponse, _, err = client.NotificationHooksAPI.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmNotificationHookDestructionDefaultAWS : Failed to retrieve Sematext Cloud NotificationHook %s, %s", rtf.StatePath, err)
			}

			if app, err = extractNotificationHook(&appResponse); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmNotificationHookDestructionDefaultAWS : Missing app in genericAPIResponse checking Sematext Cloud NotificationHook %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmNotificationHookDestructionDefaultAWS : Unexpected status in checking Sematext Cloud NotificationHook %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}

func testSetProviderConfiguration(t *testing.T) {

	p := generated.Provider()

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
