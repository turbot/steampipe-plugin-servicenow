package model

const SysDbObjectTableName = "sys_db_object"

type SysDbObjectListResult struct {
	Result []SysDbObject `json:"result"`
}

type SysDbObjectGetResult struct {
	Result SysDbObject `json:"result"`
}

type SysDbObject struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}
