package burgeonsdk

func (this *BurgeonConnection) NewQuery() PostData {
	return PostData{Id: 1, Command: "Query", Params: make(map[string]interface{})}
}

//设置查询的表
func (this *PostData) QuerySetTable(t string) {
	this.Params["table"] = t
}

//设置返回结果
func (this *PostData) QuerySetResult(columns ...string) {
	this.Params["columns"] = columns
}

//设置查询调节  k 条件  v 值
func (this *PostData) QuerySetCondition(k string, v interface{}) {
	pps := make(map[string]interface{})
	pps["column"] = k
	pps["condition"] = v
	this.Params["params"] = pps
}

func (this *PostData) QuerySetStartRange(start, rg int) {
	this.Params["start"] = start
	this.Params["range"] = rg
}
