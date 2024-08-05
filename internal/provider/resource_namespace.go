package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func resourceNamespace() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a Kestra Namespace.",

		CreateContext: resourceNamespaceCreate,
		ReadContext:   resourceNamespaceRead,
		UpdateContext: resourceNamespaceUpdate,
		DeleteContext: resourceNamespaceDelete,
		Schema: map[string]*schema.Schema{
			"tenant_id": {
				Description: "The tenant id.",
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
			},
			"namespace_id": {
				Description: "The namespace.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "The namespace friendly description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"variables": {
				Description:      "The namespace variables in yaml string.",
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: isYamlEquals,
			},
			"plugin_defaults": {
				Description:      "The namespace plugin defaults in yaml string.",
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: isYamlEquals,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNamespaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	body, err := namespaceSchemaToApi(d)
	if err != nil {
		return diag.FromErr(err)
	}

	tenantId := c.TenantId

	r, reqErr := c.request("POST", fmt.Sprintf("%s/namespaces", apiRoot(tenantId)), body)
	if reqErr != nil {
		return diag.FromErr(reqErr.Err)
	}

	errs := namespaceApiToSchema(r.(map[string]interface{}), d, c)
	if errs != nil {
		return errs
	}

	return diags
}

func resourceNamespaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	namespaceId := d.Id()
	tenantId := c.TenantId

	r, reqErr := c.request("GET", fmt.Sprintf("%s/namespaces/%s", apiRoot(tenantId), namespaceId), nil)
	if reqErr != nil {
		if reqErr.StatusCode == http.StatusNotFound {
			d.SetId("")
			return diags
		}

		return diag.FromErr(reqErr.Err)
	}

	errs := namespaceApiToSchema(r.(map[string]interface{}), d, c)
	if errs != nil {
		return errs
	}

	return diags
}

func resourceNamespaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	if d.HasChanges("description", "variables", "plugin_defaults") {
		body, err := namespaceSchemaToApi(d)
		if err != nil {
			return diag.FromErr(err)
		}

		namespaceId := d.Id()
		tenantId := c.TenantId

		r, reqErr := c.request("PUT", fmt.Sprintf("%s/namespaces/%s", apiRoot(tenantId), namespaceId), body)
		if reqErr != nil {
			return diag.FromErr(reqErr.Err)
		}

		errs := namespaceApiToSchema(r.(map[string]interface{}), d, c)
		if errs != nil {
			return errs
		}

		return diags
	} else {
		return resourceNamespaceRead(ctx, d, meta)

	}
}

func resourceNamespaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	namespaceId := d.Id()
	tenantId := c.TenantId

	_, reqErr := c.request("DELETE", fmt.Sprintf("%s/namespaces/%s", apiRoot(tenantId), namespaceId), nil)
	if reqErr != nil {
		return diag.FromErr(reqErr.Err)
	}

	d.SetId("")

	return diags
}
