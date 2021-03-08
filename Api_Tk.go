package burgeonsdk

func (bg *BurgeonConnection) AddNewTk(tkno string, card string, price int64, start, end, remark string) (string, error) {
	pts := bg.NewProcessOrder()
	pts.ProcessOrderMasterObjSetTable("C_V_VOUCHERS")
	pts.ProcessOrderIfSubmit(true)
	pts.ProcessOrderMasterObjSetColumn("VOUCHERS_NO", tkno)                            //券号
	pts.ProcessOrderMasterObjSetColumn("C_VIP_ID__CARDNO", card)                       //VIP卡号
	pts.ProcessOrderMasterObjSetColumn("VOU_TYPE", "VOU5")                             //优惠券类别(购物券)
	pts.ProcessOrderMasterObjSetColumn("IS_OFFLINE", "Y")                              //是否线下
	pts.ProcessOrderMasterObjSetColumn("AMT_DISCOUNT", price)                          //优惠金额
	pts.ProcessOrderMasterObjSetColumn("START_DATE", start)                            //开始日期
	pts.ProcessOrderMasterObjSetColumn("DELTYPE", "AND")                               //排除方式
	pts.ProcessOrderMasterObjSetColumn("VALID_DATE", end)                              //过期日期
	pts.ProcessOrderMasterObjSetColumn("C_CUSTOMER_ID__NAME", NewVipCardInfo.Customer) //经销商名
	pts.ProcessOrderMasterObjSetColumn("IS_ALLSTORE", "Y")                             //是否所有店铺
	pts.ProcessOrderMasterObjSetColumn("ISSHARE_PAYTYPE", "Y")                         //付款方式是否共用
	pts.ProcessOrderMasterObjSetColumn("IS_MDAMT", "N")                                //消费时允许修改金额
	pts.ProcessOrderMasterObjSetColumn("DESCRIPTION", remark)                          //备注
	return bg.Post(pts)
}
