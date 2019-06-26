// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeUrlMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeUrlMapCreate,
		Read:   resourceComputeUrlMapRead,
		Update: resourceComputeUrlMapUpdate,
		Delete: resourceComputeUrlMapDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeUrlMapImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"default_service": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     computeUrlMapHostRuleSchema(),
				// Default schema.HashSchema is used.
			},
			"path_matcher": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_service": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"path_rule": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"paths": {
										Type:     schema.TypeSet,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set: schema.HashString,
									},
									"service": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: compareSelfLinkOrResourceName,
									},
								},
							},
						},
					},
				},
			},
			"test": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"map_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func computeUrlMapHostRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hosts": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"path_matcher": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeUrlMapCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	defaultServiceProp, err := expandComputeUrlMapDefaultService(d.Get("default_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_service"); !isEmptyValue(reflect.ValueOf(defaultServiceProp)) && (ok || !reflect.DeepEqual(v, defaultServiceProp)) {
		obj["defaultService"] = defaultServiceProp
	}
	descriptionProp, err := expandComputeUrlMapDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hostRulesProp, err := expandComputeUrlMapHost_rule(d.Get("host_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_rule"); !isEmptyValue(reflect.ValueOf(hostRulesProp)) && (ok || !reflect.DeepEqual(v, hostRulesProp)) {
		obj["hostRules"] = hostRulesProp
	}
	fingerprintProp, err := expandComputeUrlMapFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	nameProp, err := expandComputeUrlMapName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pathMatchersProp, err := expandComputeUrlMapPath_matcher(d.Get("path_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("path_matcher"); !isEmptyValue(reflect.ValueOf(pathMatchersProp)) && (ok || !reflect.DeepEqual(v, pathMatchersProp)) {
		obj["pathMatchers"] = pathMatchersProp
	}
	testsProp, err := expandComputeUrlMapTest(d.Get("test"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("test"); !isEmptyValue(reflect.ValueOf(testsProp)) && (ok || !reflect.DeepEqual(v, testsProp)) {
		obj["tests"] = testsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new UrlMap: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating UrlMap: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating UrlMap",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create UrlMap: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating UrlMap %q: %#v", d.Id(), res)

	return resourceComputeUrlMapRead(d, meta)
}

func resourceComputeUrlMapRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeUrlMap %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeUrlMapCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("default_service", flattenComputeUrlMapDefaultService(res["defaultService"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("description", flattenComputeUrlMapDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("host_rule", flattenComputeUrlMapHost_rule(res["hostRules"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("map_id", flattenComputeUrlMapMap_id(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeUrlMapFingerprint(res["fingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("name", flattenComputeUrlMapName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("path_matcher", flattenComputeUrlMapPath_matcher(res["pathMatchers"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("test", flattenComputeUrlMapTest(res["tests"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}

	return nil
}

func resourceComputeUrlMapUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	defaultServiceProp, err := expandComputeUrlMapDefaultService(d.Get("default_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultServiceProp)) {
		obj["defaultService"] = defaultServiceProp
	}
	descriptionProp, err := expandComputeUrlMapDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hostRulesProp, err := expandComputeUrlMapHost_rule(d.Get("host_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_rule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostRulesProp)) {
		obj["hostRules"] = hostRulesProp
	}
	fingerprintProp, err := expandComputeUrlMapFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	nameProp, err := expandComputeUrlMapName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pathMatchersProp, err := expandComputeUrlMapPath_matcher(d.Get("path_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("path_matcher"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pathMatchersProp)) {
		obj["pathMatchers"] = pathMatchersProp
	}
	testsProp, err := expandComputeUrlMapTest(d.Get("test"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("test"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, testsProp)) {
		obj["tests"] = testsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating UrlMap %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating UrlMap %q: %s", d.Id(), err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Updating UrlMap",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeUrlMapRead(d, meta)
}

func resourceComputeUrlMapDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting UrlMap %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "UrlMap")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting UrlMap",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting UrlMap %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeUrlMapImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/urlMaps/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeUrlMapCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapDefaultService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeUrlMapDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapHost_rule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeUrlMapHostRuleSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"description":  flattenComputeUrlMapHost_ruleDescription(original["description"], d),
			"hosts":        flattenComputeUrlMapHost_ruleHosts(original["hosts"], d),
			"path_matcher": flattenComputeUrlMapHost_rulePathMatcher(original["pathMatcher"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapHost_ruleDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapHost_ruleHosts(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeUrlMapHost_rulePathMatcher(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapMap_id(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeUrlMapFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapPath_matcher(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"default_service": flattenComputeUrlMapPath_matcherDefaultService(original["defaultService"], d),
			"description":     flattenComputeUrlMapPath_matcherDescription(original["description"], d),
			"name":            flattenComputeUrlMapPath_matcherName(original["name"], d),
			"path_rule":       flattenComputeUrlMapPath_matcherPath_rule(original["pathRules"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapPath_matcherDefaultService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeUrlMapPath_matcherDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapPath_matcherName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapPath_matcherPath_rule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"paths":   flattenComputeUrlMapPath_matcherPath_rulePaths(original["paths"], d),
			"service": flattenComputeUrlMapPath_matcherPath_ruleService(original["service"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapPath_matcherPath_rulePaths(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeUrlMapPath_matcherPath_ruleService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeUrlMapTest(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"description": flattenComputeUrlMapTestDescription(original["description"], d),
			"host":        flattenComputeUrlMapTestHost(original["host"], d),
			"path":        flattenComputeUrlMapTestPath(original["path"], d),
			"service":     flattenComputeUrlMapTestService(original["service"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapTestDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTestHost(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTestPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTestService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

// ResourceRef only supports 1 type and UrlMap has references to a BackendBucket or BackendService. Just read the self_link string
// instead of extracting the name and making a self_link out of it.
func expandComputeUrlMapDefaultService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapHost_rule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeUrlMapHost_ruleDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedHosts, err := expandComputeUrlMapHost_ruleHosts(original["hosts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHosts); val.IsValid() && !isEmptyValue(val) {
			transformed["hosts"] = transformedHosts
		}

		transformedPathMatcher, err := expandComputeUrlMapHost_rulePathMatcher(original["path_matcher"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPathMatcher); val.IsValid() && !isEmptyValue(val) {
			transformed["pathMatcher"] = transformedPathMatcher
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapHost_ruleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapHost_ruleHosts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeUrlMapHost_rulePathMatcher(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPath_matcher(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDefaultService, err := expandComputeUrlMapPath_matcherDefaultService(original["default_service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDefaultService); val.IsValid() && !isEmptyValue(val) {
			transformed["defaultService"] = transformedDefaultService
		}

		transformedDescription, err := expandComputeUrlMapPath_matcherDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedName, err := expandComputeUrlMapPath_matcherName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedPath_rule, err := expandComputeUrlMapPath_matcherPath_rule(original["path_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath_rule); val.IsValid() && !isEmptyValue(val) {
			transformed["pathRules"] = transformedPath_rule
		}

		req = append(req, transformed)
	}
	return req, nil
}

// ResourceRef only supports 1 type and UrlMap has references to a BackendBucket or BackendService. Just read the self_link string
// instead of extracting the name and making a self_link out of it.
func expandComputeUrlMapPath_matcherDefaultService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPath_matcherDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPath_matcherName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPath_matcherPath_rule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPaths, err := expandComputeUrlMapPath_matcherPath_rulePaths(original["paths"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPaths); val.IsValid() && !isEmptyValue(val) {
			transformed["paths"] = transformedPaths
		}

		transformedService, err := expandComputeUrlMapPath_matcherPath_ruleService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapPath_matcherPath_rulePaths(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

// ResourceRef only supports 1 type and UrlMap has references to a BackendBucket or BackendService. Just read the self_link string
// instead of extracting the name and making a self_link out of it.
func expandComputeUrlMapPath_matcherPath_ruleService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeUrlMapTestDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedHost, err := expandComputeUrlMapTestHost(original["host"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
			transformed["host"] = transformedHost
		}

		transformedPath, err := expandComputeUrlMapTestPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandComputeUrlMapTestService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapTestDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTestHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

// ResourceRef only supports 1 type and UrlMap has references to a BackendBucket or BackendService. Just read the self_link string
// instead of extracting the name and making a self_link out of it.
func expandComputeUrlMapTestService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
