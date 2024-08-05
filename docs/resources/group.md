---
page_title: "kestra_group Resource - terraform-provider-kestra"
subcategory: ""
description: |-
  Manages a Kestra Group.
---

# kestra_group (Resource)

Manages a Kestra Group.

## Example Usage

```terraform
resource "kestra_group" "example" {
  namespace   = "company.team"
  name        = "Friendly name"
  description = "Friendly description"
}
```

## Schema

### Required

- `name` (String) The group name.

### Optional

- `description` (String) The group description.
- `namespace` (String) The linked namespace.

### Read-Only

- `id` (String) The ID of this resource.
- `tenant_id` (String) The tenant id.

## Import

Import is supported using the following syntax:

```shell
terraform import kestra_group.example {{group_id}}
```
