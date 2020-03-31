---
layout: "sematext"
page_title: "Sematext: sematext_data00"
description: |-
  TODO Placeholder.
---

# sematext_data00

Use this data source to retrieve the collaborators for a given repository.

## Example Usage

```hcl
data "seamatext_data00" "test" {
  owner      = "example_owner"
  repository = "example_repository"
}
```

## Arguments Reference

 * `owner` - (Required) The organization that owns the repository.
 
 * `repository` - (Required) The name of the repository.
 
 * `affiliation` - (Optional) Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
 
## Attributes Reference

 * `attr00` - An Array of whatever.  Each `whatever` block consists of the fields documented below.
 
___
 
The `att00` block consists of:

* `something` - something.

