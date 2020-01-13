---
layout: "sematext"
page_title: "Provider: Sematext"
description: |-
  The Sematext provider is used to interact with Sematext Cloud resources.
---

# Sematext Provider

The Sematext provider is used to interact with Sematext Cloud resources.

The provider allows you to manage your Sematext Cloud Apps.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Sematext Provider
provider "sematext" {
  region        = "${var.region}"  
  token        = "${var.token}"  
}

# Create a monitoring application within Sematext Cloud 
resource "sematext_app" "register_app_x" {
  # ...
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - This is the Sematext account token. It can also be
  sourced from the `SEMATEXT_TOKEN` environment variable. Token is required.
  
* `region` - This is the target Sematext Cloud region.
  Either 'US' or 'EU'. It can also be sourced from the `SEMATEXT_REGION` environment variable. 

