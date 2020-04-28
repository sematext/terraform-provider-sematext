# Sematext Redis Monitoring Resource

Creates a monitoring application within Sematext Cloud. 
Refer to [Refer to Sematext Provider for authentication detail](../index.md) 

## Example Usage

```hcl
# Configure the Sematext Provider
provider "sematext" {
  sematext_region = "US"
}

# Create a monitoring application
resource "sematext_monitor_redis" "mymonitor" {
  name = "my monitor name"
  billing_plan_id = 6
  discount_code = "<discount code>"
}
```

## Argument Reference

* `name` - List attributes that this resource exports.
* `billing_plan_id` - List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md) 
* `discount_code` - List attributes that this resource exports.



