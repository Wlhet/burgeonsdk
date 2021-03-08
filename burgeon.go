package burgeonsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
)

var NewVipCardInfo *VipCardConfig
var BaseConnection *BurgeonConnection

type BurgeonConnection struct {
	AppId  string
	AppKey string
	ApiUrl string
}

type VipCardConfig struct {
	VipType      string //"ECCO会员卡"  会员卡类型
	Customer     string //"爱步" 经销商
	Store        string //"ECCO公司仓" 所属仓
	ValidDate    string ////"20301231" 过期时间
	IntegralArea string //"ECCO区域" 积分区域
	VipBrandName string //Vip品牌名
	TkEndDate    string //新用户送的券的有效期
}

type PostData struct {
	Id      int                    `json:"id"`
	Command string                 `json:"command"`
	Params  map[string]interface{} `json:"params"`
}

//向外单独提供
func NewBurgeonConnection(AppId, AppKey, ApiUrl string) *BurgeonConnection {
	return &BurgeonConnection{
		AppId:  AppId,
		AppKey: AppKey,
		ApiUrl: ApiUrl,
	}
}

//用于初始化包内默认连接对象
func InitBurgeonConnection(AppId, AppKey, ApiUrl string) {
	BaseConnection = &BurgeonConnection{
		AppId:  AppId,
		AppKey: AppKey,
		ApiUrl: ApiUrl,
	}
}

func (bg *BurgeonConnection) InitVipCardConfig(viptype, customer, store, area, validdate string) {
	NewVipCardInfo = &VipCardConfig{
		VipType:      viptype,
		Customer:     customer,
		Store:        store,
		IntegralArea: area,
		ValidDate:    validdate,
	}
}

//Burgeon--------------------------------------------------------------------------
//APP_KEY_2 + tmp + app_ser_md5s
//将密码MD5后和账号时间戳按照 账号 时间戳 密码MD5 拼接 返回
func (bg *BurgeonConnection) GetSign() (string, string) {
	tmp := bg.GetTimeMillisecondString()
	md5key := StringToMd5(bg.AppKey)
	return tmp, StringToMd5(bg.AppId + tmp + md5key)
}

//2020-09-26 15:06:23.000
func (bg *BurgeonConnection) GetTimeMillisecondString() string {
	tmp := time.Now().Format("2006-01-02 15:04:05.000")
	return tmp
}

//封装了请求提交函数  只需传入 transactions即可  返回json字符串
func (bg *BurgeonConnection) Post(comm ...PostData) (string, error) {
	data_json_byte, _ := json.Marshal(&comm)
	data_json_str := string(data_json_byte)
	tmp, sign := bg.GetSign()
	req := httplib.Post(bg.ApiUrl)
	req.Param("sip_appkey", bg.AppId)
	req.Param("sip_timestamp", tmp)
	req.Param("sip_sign", sign)
	req.Param("transactions", data_json_str)
	fmt.Println(data_json_str)
	response, err := req.String()
	if err != nil {
		return "", err
	}
	code := gjson.Get(response, "0.code").Int()
	errmsg := gjson.Get(response, "0.message").String()
	if code == 0 {
		return response, nil
	} else {
		return errmsg, errors.New(errmsg)
	}
}
