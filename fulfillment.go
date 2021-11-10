// fulfillment.go
package main

import (
        "log"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fulfillment() *schema.Resource {
        return &schema.Resource{
                Create: fulfillmentCreate,
                Read:   fulfillmentRead,
                Update: fulfillmentUpdate,
                Delete: fulfillmentDelete,

                Schema: map[string]*schema.Schema{
                        "uuid_count": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func fulfillmentCreate(d *schema.ResourceData, m interface{}) error {
        ecosystem := d.Get("ecosystem").(string)
        universe := d.Get("universe").(string)
        port := d.Get("port").(string)

        return fulfillmentRead(d, m)
}

func fulfillmentRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func fulfillmentUpdate(d *schema.ResourceData, m interface{}) error {
        return fulfillmentRead(d, m)
}

func fulfillmentDelete(d *schema.ResourceData, m interface{}) error {
        d.SetId("")
        return nil
}
