---
page_title: "kestra_flow Data Source - terraform-provider-kestra"
subcategory: ""
description: |-
  Use this data source to access information about an existing Kestra Flow
---

# kestra_flow (Data Source)

Use this data source to access information about an existing Kestra Flow

## Example Usage

```terraform
data "kestra_flow" "example" {
  namespace = "company.team"
  id        = "my-flow"
}
```

## Schema

### Required

- `flow_id` (String) The flow id.
- `namespace` (String) The namespace.

### Read-Only

- `content` (String) The flow content as yaml.
- `id` (String) The ID of this resource.
- `revision` (Number) The flow revision.
- `tenant_id` (String) The tenant id.
