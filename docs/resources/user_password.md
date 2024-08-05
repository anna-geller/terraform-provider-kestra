---
page_title: "kestra_user_password Resource - terraform-provider-kestra"
subcategory: ""
description: |-
  Manages a Kestra User Basic Auth Password.
---

# kestra_user_password (Resource)

Manages a Kestra User Basic Auth Password.

## Example Usage

```terraform
resource "kestra_user_password" "example" {
  user_id  = "4by6NvSLcPXFhCj8nwbZOM"
  password = "my-random-password"
}
```

## Schema

### Required

- `password` (String, Sensitive) The user password.
- `user_id` (String) The user id.

### Read-Only

- `id` (String) The ID of this resource.
