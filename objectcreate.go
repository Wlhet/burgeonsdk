package burgeonsdk

func (this *BurgeonConnection) NewObjectCreate() PostData {
	return PostData{Id: 2, Command: "ObjectCreate", Params: make(map[string]interface{})}
}

//设置查询的表
func (this *PostData) ObjectCreateSetTable(t string) {
	this.Params["table"] = t
}

//设置返回结果
func (this *PostData) ObjectCreateSetColumn(k string, v interface{}) {
	this.Params[k] = v
}
