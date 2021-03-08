package burgeonsdk

func (this *BurgeonConnection) NewExecuteSql() PostData {
	return PostData{Id: 5, Command: "ExecuteSQL", Params: map[string]interface{}{
		"name":   "",
		"values": make([]string, 0),
		"result": make([]string, 0),
	}}
}

func (this *PostData) ExecuteSqlSetName(name string) {
	this.Params["name"] = name
}

func (this *PostData) ExecuteSqlSetValues(value string) {
	values := this.Params["values"].([]string)
	values = append(values, value)
}
