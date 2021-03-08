package burgeonsdk

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const (
	EmptyString = ""
	ApiErrCode  = -1
)

var NewVipCardInfo *VipCardConfig

func InitVipCardConfig(viptype, customer, store, area, validdate string) {
	NewVipCardInfo = &VipCardConfig{
		VipType:      viptype,
		Customer:     customer,
		Store:        store,
		IntegralArea: area,
		ValidDate:    validdate,
	}
}

//获取Vip当前积分
func (bg *BurgeonConnection) GetVipIntegralByCardNO(card string) (int64, error) {
	pst := bg.NewQuery()
	pst.QuerySetTable("C_VIP")
	pst.QuerySetCondition("CARDNO", card)
	pst.QuerySetResult("INTEGRAL")
	s, err := bg.Post(pst)
	if err != nil {
		return -1, err
	} else {
		return gjson.Get(s, "0.rows.0.0").Int(), nil
	}
}

//获取Vip等级
func (bg *BurgeonConnection) GetVipLevelByCardNO(card string) (int64, error) {
	pst := bg.NewQuery()
	pst.QuerySetTable("C_VIP")
	pst.QuerySetCondition("CARDNO", card)
	pst.QuerySetResult("VIP_LEVEL")
	s, err := bg.Post(pst)
	if err != nil {
		return ApiErrCode, err
	} else {
		return gjson.Get(s, "0.rows.0.0").Int(), nil
	}
}

//获取VIP等级名称
func (bg *BurgeonConnection) GetVipLevelNameByCard(card string) (string, error) {
	pst := bg.NewQuery()
	pst.QuerySetTable("C_VIP")
	pst.QuerySetCondition("CARDNO", card)
	pst.QuerySetResult("C_VIPTYPE_ID;NAME")
	s, err := bg.Post(pst)
	if err != nil {
		return EmptyString, err
	} else {
		name := gjson.Get(s, "0.rows.0.0").String()
		return name, nil
	}
}

func (bg *BurgeonConnection) GetVipLevelUpMRetail(card string) (string, error) {
	pst := bg.NewQuery()
	pst.QuerySetTable("C_VIP")
	pst.QuerySetCondition("CARDNO", card)
	pst.QuerySetResult("M_RETAIL_ID;NAME")
	s, err := bg.Post(pst)
	fmt.Println(s)
	if err != nil {
		return EmptyString, err
	} else {
		name := gjson.Get(s, "0.rows.0.0").String()
		return name, nil
	}
}

//获取Vip消费记录/返回消费金额,积分,日期[[2165,44,20210304]]
func (bg *BurgeonConnection) GetVipRecordsOfConsumptionByCard(card string) ([][]int64, error) {
	pst := bg.NewQuery()
	pst.QuerySetTable("FA_VIPINTEGRAL_FTP")
	pst.QuerySetCondition("C_VIP_ID;CARDNO", card)
	pst.QuerySetResult("AMT_ACTUAL", "INTEGRAL", "CHANGDATE")
	s, err := bg.Post(pst)
	result := make([][]int64, 0)
	if err != nil {
		return result, err
	} else {
		array := gjson.Get(s, "0.rows").Array()
		for _, v := range array {
			vv := v.Array()
			result = append(result, []int64{vv[0].Int(), vv[1].Int(), vv[2].Int()})
		}
		return result, nil
	}
}

//获新增VIP
func (bg *BurgeonConnection) CreateVip(card, mobil, name, birthday, sex, remark string) (string, error) {
	switch sex {
	case "M":
	case "W":
	case "男":
		sex = "M"
	case "女":
		sex = "W"
	default:
		sex = "N"
	}
	pst := bg.NewObjectCreate()
	pst.ObjectCreateSetTable("C_V_ADDVIP")
	pst.ObjectCreateSetColumn("C_VIPTYPE_ID__NAME", NewVipCardInfo.VipType)
	pst.ObjectCreateSetColumn("CARDNO", card)
	pst.ObjectCreateSetColumn("C_CUSTOMER_ID__NAME", NewVipCardInfo.Customer)
	pst.ObjectCreateSetColumn("C_STORE_ID__NAME", NewVipCardInfo.Store)
	pst.ObjectCreateSetColumn("MOBIL", mobil)
	pst.ObjectCreateSetColumn("SEX", sex)
	pst.ObjectCreateSetColumn("VIPNAME", name)
	pst.ObjectCreateSetColumn("BIRTHDAY", birthday)
	pst.ObjectCreateSetColumn("C_INTEGRALAREA_ID__NAME", NewVipCardInfo.IntegralArea)
	pst.ObjectCreateSetColumn("OPENID", " ")
	pst.ObjectCreateSetColumn("DESCRIPTION", remark)
	return bg.Post(pst)
}

func (bg *BurgeonConnection) ChangeVipBirthday(card string, birthday string) (string, error) {
	pts := bg.NewObjectModify()
	pts.ObjectModifySetTable("c_vip")
	pts.ObjectModifySetColumn("ak", card)
	pts.ObjectModifySetColumn("birthday", birthday)
	pts.ObjectModifySetColumn("partial_update", true) //partial_update*boolean缺省值:true，表示仅修改传入的<column-name>对应的列
	return bg.Post(pts)
}

func (bg *BurgeonConnection) ChangeVipName(card string, name string) (string, error) {
	pts := bg.NewObjectModify()
	pts.ObjectModifySetTable("c_vip")
	pts.ObjectModifySetColumn("ak", card)
	pts.ObjectModifySetColumn("vipname", name)
	pts.ObjectModifySetColumn("partial_update", true) //partial_update*boolean缺省值:true，表示仅修改传入的<column-name>对应的列
	return bg.Post(pts)
}

func (bg *BurgeonConnection) ChangeVipSex(card string, sex string) (string, error) {
	pts := bg.NewObjectModify()
	pts.ObjectModifySetTable("c_vip")
	pts.ObjectModifySetColumn("ak", card)
	pts.ObjectModifySetColumn("sex", sex)
	pts.ObjectModifySetColumn("partial_update", true) //partial_update*boolean缺省值:true，表示仅修改传入的<column-name>对应的列
	return bg.Post(pts)
}

//新增VIP积分调整单
func (bg *BurgeonConnection) AddVipIntegral(card, billdate, remark string, integral int64) (string, error) {
	pts := bg.NewProcessOrder()
	pts.ProcessOrderMasterObjSetTable("C_VIPINTEGRALADJ")
	pts.ProcessOrderIfSubmit(true)
	pts.ProcessOrderMasterObjSetColumn("BILLDATE", billdate)
	pts.ProcessOrderMasterObjSetColumn("ADJTYPE", 1)
	pts.ProcessOrderDetailObjsSetTables("C_VIPINTEGRALADJITEM")
	pts.ProcessOrderDetailObjsAddrefobjs("C_VIPINTEGRALADJITEM",
		map[string]interface{}{
			"C_VIP_ID__CARDNO": card,
			"INTEGRALADJ":      integral,
			"DESCRIPTION":      remark,
		})
	return bg.Post(pts)
}
