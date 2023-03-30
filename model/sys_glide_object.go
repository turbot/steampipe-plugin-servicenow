package model

const SysGlideObjectTableName = "sys_glide_object"

type SysGlideObjectListResult struct {
	Result []SysGlideObject `json:"result"`
}

type SysGlideObjectGetResult struct {
	Result SysGlideObject `json:"result"`
}

type SysGlideObject struct {
	UseOriginalValue string `json:"use_original_value"`
	Visible          string `json:"visible"`
	SysModCount      string `json:"sys_mod_count"`
	Label            string `json:"label"`
	ScalarType       string `json:"scalar_type"`
	SysUpdatedOn     string `json:"sys_updated_on"`
	SysClassName     string `json:"sys_class_name"`
	SysID            string `json:"sys_id"`
	SysUpdateName    string `json:"sys_update_name"`
	SysUpdatedBy     string `json:"sys_updated_by"`
	SysCreatedOn     string `json:"sys_created_on"`
	Name             string `json:"name"`
	SysName          string `json:"sys_name"`
	Attributes       string `json:"attributes"`
	ClassName        string `json:"class_name"`
	ScalarLength     string `json:"scalar_length"`
	SysCreatedBy     string `json:"sys_created_by"`
	SysPolicy        string `json:"sys_policy"`
}
