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
}
