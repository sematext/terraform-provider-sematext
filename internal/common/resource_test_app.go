package sematext

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/*
var testProviders map[string]*schema.Provider
var testProvider schema.Provider
var testProviderFunc func() *schema.Provider

type ResourceTestFixture interface {
	hydrate(resourceType string, appType string)
	toHCL() string
}

// ResourceTestFixture a common test fixture representing most resources
type ResourceTestFixtureDefault struct {
	ResourceType string
	ResourceName string
	AppType      string
	Name         string
	StatePath    string
	PlanID       int
	DiscountCode string
	AppToken     AppTokenType
}

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

*/

func TestAccAppResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccAppResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "one"),                  //@TODO switch to SC App
					resource.TestCheckResourceAttr("scaffolding_example.test", "defaulted", "example value when not configured"), //@TODO switch to SC App
					resource.TestCheckResourceAttr("scaffolding_example.test", "id", "example-id"),                               //@TODO switch to SC App
				),
			},
			// ImportState testing
			{
				ResourceName:      "scaffolding_example.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"configurable_attribute", "defaulted"},
			},
			// Update and Read testing
			{
				Config: testAccAppResourceConfig("two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "two"), //@TODO switch to SC App
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccAppResourceConfig(configurableAttribute string) string {
	return fmt.Sprintf(`
resource "scaffolding_example" "test" {
  configurable_attribute = %[1]q
}
`, configurableAttribute)
}

/*

func (rtf *ResourceTestFixtureDefault) hydrate(resourceType string, appType string) {

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

func (rtf *ResourceTestFixtureDefault) toHCL() string {

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
		arrayLiteralString(rtf.AppToken.Names),
	)

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

// ResourceTestLifecycleApp is a common test of resource creation.
func ResourceTestLifecycleApp(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonBasicTestApp Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf := ResourceTestFixtureAWS{}
		rtf.hydrate(resourceType, appType)

		fixture := rtf.toHCL()

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmResourceDestructionAWSApp(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmResourceCreationAWSApp(rtf),
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

		rtf := ResourceTestFixtureDefault{}
		rtf.hydrate(resourceType, appType)
		fixture := rtf.toHCL()

		fmt.Println(fixture)

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmResourceDestructionDefaultApp(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmResourceCreationDefaultApp(rtf),
						resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					),
				},
			},
		})
	}
}

// ResourceTestUpdateApp tests for resource updates.
func ResourceTestUpdateApp(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorUpdateTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {
	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf0 := ResourceTestFixtureAWS{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0

		fixture1 := rtf1.toHCL()

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmResourceDestructionAWSApp(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmResourceCreationAWSApp(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmResourceCreationAWSApp(rtf1),
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

		rtf0 := ResourceTestFixtureDefault{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0
		fixture1 := rtf1.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testSetProviderConfiguration(t) },
			Providers:    testProviders,
			CheckDestroy: ConfirmResourceDestructionDefaultApp(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmResourceCreationDefaultApp(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmResourceCreationDefaultApp(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					),
				},
			},
		})
	}

}

// ConfirmResourceCreationDefaultApp checks the App ID exists in both state and stcloud.
func ConfirmResourceCreationDefaultApp(rtf ResourceTestFixtureDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmResourceCreationDefaultApp func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.App
		var rs *terraform.ResourceState
		var appResponse stcloud.AppResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmMonitorCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if appResponse, _, err = client.AppsApi.GetUsingGET(context.Background(), id); err != nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s, %s", rtf.StatePath, err)
		}

		if app, err = extractApp(&appResponse); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmResourceCreationAWSApp checks the App ID exists in both state and stcloud.
func ConfirmResourceCreationAWSApp(rtf ResourceTestFixtureAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmResourceCreationAWSApp func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.App
		var rs *terraform.ResourceState
		var appResponse stcloud.AppResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmMonitorCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if appResponse, _, err = client.AppsApi.GetUsingGET(context.Background(), id); err != nil {
			return fmt.Errorf("ConfirmMonitorCreationAWS : Error in checking monitor %s, %s", rtf.StatePath, err)
		}

		if app, err = extractApp(&appResponse); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmResourceDestructionDefaultApp checks the App ID has been removed from state and the API has marked the app as DELETED.
func ConfirmResourceDestructionDefaultApp(rtf ResourceTestFixtureDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmResourceDestructionDefaultApp func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.App
		var appResponse stcloud.AppResponse

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

			if appResponse, _, err = client.AppsApi.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefault : Failed to pull app in checking monitor %s, %s", rtf.StatePath, err)
			}

			if app, err = extractApp(&appResponse); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefault : Missing app in genericAPIResponse checking monitor %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmMonitorDestructionDefault : Unexpected status in checking monitor %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}

// ConfirmResourceDestructionAWSApp checks the App ID exists in both state and API and is marked as DELETED.
func ConfirmResourceDestructionAWSApp(rtf ResourceTestFixtureAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmResourceDestructionAWSApp func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.App
		var appResponse stcloud.AppResponse

		client := testProvider.Meta().(*stcloud.APIClient)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO Consider shift to template and make check tighter after MVP
				continue
			}
			if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
				return err
			}

			if appResponse, _, err = client.AppsApi.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefaultAWS : Failed to pull app in checking monitor %s, %s", rtf.StatePath, err)
			}

			if app, err = extractApp(&appResponse); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefaultAWS : Missing app in genericAPIResponse checking monitor %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmMonitorDestructionDefaultAWS : Unexpected status in checking monitor %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}

func testSetProviderConfiguration(t *testing.T) {

	p := schema.Provider{}

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

*/
