package model

const SysDictionaryTableName = "sys_dictionary"

type SysDictionaryListResult struct {
	Result []SysDictionary `json:"result"`
}

type SysDictionaryGetResult struct {
	Result SysDictionary `json:"result"`
}

type InternalType struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}
type SysDictionary struct {
	Element      string       `json:"element"`
	InternalType InternalType `json:"internal_type"`
	// Calculation          string `json:"calculation"`
	// DynamicRefQual       string `json:"dynamic_ref_qual"`
	// ChoiceField          string `json:"choice_field"`
	// FunctionField        string `json:"function_field"`
	// SysUpdatedOn         string `json:"sys_updated_on"`
	// SpellCheck           string `json:"spell_check"`
	// ReferenceCascadeRule string `json:"reference_cascade_rule"`
	// // Reference              string       `json:"reference"`
	// SysUpdatedBy           string       `json:"sys_updated_by"`
	// ReadOnly               string       `json:"read_only"`
	// SysCreatedOn           string       `json:"sys_created_on"`
	// ArrayDenormalized      string       `json:"array_denormalized"`
	// ElementReference       string       `json:"element_reference"`
	// SysName                string       `json:"sys_name"`
	// ReferenceKey           string       `json:"reference_key"`
	// ReferenceQualCondition string       `json:"reference_qual_condition"`
	// XMLView                string       `json:"xml_view"`
	// Dependent              string       `json:"dependent"`
	// SysCreatedBy           string       `json:"sys_created_by"`
	// MaxLength              string       `json:"max_length"`
	// UseDependentField      string       `json:"use_dependent_field"`
	// DeleteRoles            string       `json:"delete_roles"`
	// VirtualType            string       `json:"virtual_type"`
	// Active                 string       `json:"active"`
	// ChoiceTable            string       `json:"choice_table"`
	// ForeignDatabase        string       `json:"foreign_database"`
	// SysUpdateName          string       `json:"sys_update_name"`
	// Unique                 string       `json:"unique"`
	// Name                   string       `json:"name"`
	// DependentOnField       string       `json:"dependent_on_field"`
	// DynamicCreation        string       `json:"dynamic_creation"`
	// Primary                string       `json:"primary"`
	// SysPolicy              string       `json:"sys_policy"`
	// NextElement            string       `json:"next_element"`
	// Virtual                string       `json:"virtual"`
	// Widget                 string       `json:"widget"`
	// UseDynamicDefault      string       `json:"use_dynamic_default"`
	// Sizeclass              string       `json:"sizeclass"`
	// Mandatory              string       `json:"mandatory"`
	// SysClassName           string       `json:"sys_class_name"`
	// DynamicDefaultValue    string       `json:"dynamic_default_value"`
	// SysID                  string       `json:"sys_id"`
	// WriteRoles             string       `json:"write_roles"`
	// Array                  string       `json:"array"`
	// Audit                  string       `json:"audit"`
	// ReadRoles              string       `json:"read_roles"`
	// CreateRoles            string       `json:"create_roles"`
	// DynamicCreationScript  string       `json:"dynamic_creation_script"`
	// Defaultsort            string       `json:"defaultsort"`
	// ColumnLabel            string       `json:"column_label"`
	// Comments               string       `json:"comments"`
	// UseReferenceQualifier  string       `json:"use_reference_qualifier"`
	// Display                string       `json:"display"`
	// ReferenceFloats        string       `json:"reference_floats"`
	// SysModCount            string       `json:"sys_mod_count"`
	// DefaultValue           string       `json:"default_value"`
	// Staged                 string       `json:"staged"`
	// ReferenceType          string       `json:"reference_type"`
	// Formula                string       `json:"formula"`
	// Attributes             string       `json:"attributes"`
	// Choice                 string       `json:"choice"`
	// ReferenceQual          string       `json:"reference_qual"`
	// TableReference         string       `json:"table_reference"`
	// TextIndex              string       `json:"text_index"`
	// FunctionDefinition     string       `json:"function_definition"`
}
