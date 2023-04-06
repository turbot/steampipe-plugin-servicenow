package model

const SysDocumentationTableName = "sys_documentation"

type SysDocumentationListResult struct {
	Result []SysDocumentation `json:"result"`
}

type SysDocumentationGetResult struct {
	Result SysDocumentation `json:"result"`
}

type SysDocumentation struct {
	Hint    string `json:"hint"`
	Label   string `json:"label"`
	Element string `json:"element"`
}
