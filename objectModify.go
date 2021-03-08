package burgeonsdk

func (this *BurgeonConnection) NewObjectModify() PostData {
	return PostData{Id: 3, Command: "ObjectModify", Params: make(map[string]interface{})}
}

//设置查询的表
func (this *PostData) ObjectModifySetTable(t string) {
	this.Params["table"] = t
}

//设置返回结果
//partial_update*boolean缺省值:true，表示仅修改传入的<column-name>对应的列
func (this *PostData) ObjectModifySetColumn(k string, v interface{}) {
	this.Params[k] = v
}
