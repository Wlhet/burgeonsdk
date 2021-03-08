package burgeonsdk

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func StringToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//2020-09-26 15:06:23.000
func GetTimeStringMillisecond() string {
	return time.Now().Format("2006-01-02 15:04:05.000")

}

//2020-09-26 15:06:23
func GetTimeStringSecond() string {
	tmp := time.Now().Format("2006-01-02 15:04:05")
	return tmp
}

func GetTimeStringDay() string {
	tmp := time.Now().Format("20060102")
	return tmp
}

//当月最后一天
func GetNowLastStringDay() string {
	da := time.Now()                              //当前时间
	nextMonth := da.AddDate(0, 1, 0)              //月份加一
	LastDay := nextMonth.AddDate(0, 0, -da.Day()) //减去当前的日数,就是本月最后一天
	return LastDay.Format("20060102")
}

//当月第一一天
func GetNowFirstStringDay() string {
	da := time.Now()                          //当前时间          //月份加一
	FirstDay := da.AddDate(0, 0, -da.Day()+1) //减去当前的日数,就是本月最后一天
	return FirstDay.Format("20060102")
}

//微信获取的生日1999-7-1 需要转换为19990707此类型
func HandleBirthday(b string) string {
	if len(b) < 8 {
		return "19900101"
	}
	rs := strings.Split(b, "-")
	year := rs[0]
	month := rs[1]
	day := rs[2]
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	return year + month + day
}

// Substr 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length || end < 0 {
		return ""
	}

	if end > length {
		return string(rs[start:])
	}
	return string(rs[start:end])
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(s ...string) string {
	sort.Strings(s)
	h := sha1.New()
	h.Write([]byte(strings.Join(s, "")))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(s ...string) string {
	sort.Strings(s)
	h := md5.New()
	h.Write([]byte(strings.Join(s, "")))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

//随机生成字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
