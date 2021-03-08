package burgeonsdk

func (this *BurgeonConnection) NewProcessOrder() PostData {
	return PostData{Id: 4, Command: "ProcessOrder", Params: map[string]interface{}{
		"submit":    false,
		"id":        -1,
		"masterobj": make(map[string]interface{}),
		"detailobjs": map[string]interface{}{
			"tables":  make([]string, 0),
			"refobjs": make([]map[string]interface{}, 0),
		},
	}}
}

//存储过程单据是否提交
func (this *PostData) ProcessOrderIfSubmit(b bool) {
	this.Params["submit"] = b
}

//设置主表名称
func (this *PostData) ProcessOrderMasterObjSetTable(table string) {
	masterobj := this.Params["masterobj"].(map[string]interface{})
	masterobj["table"] = table
}

//设置主表字段
func (this *PostData) ProcessOrderMasterObjSetColumn(k string, v interface{}) {
	masterobj := this.Params["masterobj"].(map[string]interface{})
	masterobj[k] = v
}

//子表名称 可是复数  多表
func (this *PostData) ProcessOrderDetailObjsSetTables(t ...string) {
	detailobjs := this.Params["detailobjs"].(map[string]interface{})
	detailobjs["tables"] = t
}

//设置子表提交数据   表名  各个字段
func (this *PostData) ProcessOrderDetailObjsAddrefobjs(reftable string, list ...map[string]interface{}) {
	pdetailobjs := this.Params["detailobjs"].(map[string]interface{})
	refobjs := pdetailobjs["refobjs"].([]map[string]interface{})
	dobi := make([]map[string]interface{}, 0)
	dobi = append(dobi, refobjs...)
	ref := make(map[string]interface{})
	ref["table"] = reftable
	ref["addList"] = list
	dobi = append(dobi, ref)
	pdetailobjs["refobjs"] = dobi
}
