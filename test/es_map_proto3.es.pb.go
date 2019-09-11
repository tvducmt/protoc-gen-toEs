// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: es_map_proto3.proto

package test

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "git.zapa.cloud/merchant-tools/helper/proto"
	_ "github.com/tvducmt/protoc-gen-toEs/protobuf"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_golang_glog "github.com/golang/glog"
	git_zapa_cloud_merchant_tools_helper_proto "git.zapa.cloud/merchant-tools/helper/proto"
	reflect "reflect"
	time "time"
	flag "flag"
	github_com_golang_protobuf_ptypes_timestamp "github.com/golang/protobuf/ptypes/timestamp"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func checkNull(field interface{}) bool {
	zero := reflect.Zero(reflect.TypeOf(field)).Interface()
	if reflect.DeepEqual(field, zero) {
		return true
	}
	return false
}
func checkTimestampType(field interface{}) (*github_com_golang_protobuf_ptypes_timestamp.Timestamp, bool) {
	if ts, ok := field.(*github_com_golang_protobuf_ptypes_timestamp.Timestamp); ok {
		return ts, true
	}
	return nil, false
}
func checkDateType(field interface{}) (*git_zapa_cloud_merchant_tools_helper_proto.Date, bool) {
	if date, ok := field.(*git_zapa_cloud_merchant_tools_helper_proto.Date); ok {
		return date, true
	}
	return nil, false
}
func makeKeyMap(m *map[string]interface{}, key string) *map[string]interface{} {
	if t, ok := (*m)[key]; ok {
		if t, ok := t.(*map[string]interface{}); ok {
			return t
		}
	}
	t := &map[string]interface{}{}
	(*m)[key] = t
	return t
}
func (this *TransactionMessage3) GetEsMap(esMap *map[string]interface{}) {
	flag.Parse()
	if !checkNull(this.CreateTime) {
		if ts, ok := checkTimestampType(this.CreateTime); ok {
			if ts != nil {
				tm, err := github_com_golang_protobuf_ptypes.Timestamp(ts)
				if err != nil {
					github_com_golang_glog.Errorln(err)
				} else {
					(*esMap)["reqTimdestamp"] = tm.UnixNano() / int64(time.Millisecond)
				}
			}
		}
	}
	if !checkNull(this.ToDate) {
		if date, ok := checkDateType(this.ToDate); ok {
			if date != nil {
				tm := git_zapa_cloud_merchant_tools_helper_proto.DateToTimeSearch(date)
				(*esMap)["reqDate"] = tm.UnixNano() / int64(time.Millisecond)
			}
		}
	}
	if !checkNull(this.PmcName) {
		(*esMap)["pmcName"] = this.PmcName
	}
	if !checkNull(this.DiscountAmount) {
		(*esMap)["disCountAmount"] = this.DiscountAmount
	}
	if !checkNull(this.TransCode) {
		(*esMap)["transErrCode"] = this.TransCode
	}
	if !checkNull(this.TransChargeStatus) {
		(*esMap)["transChargeStatus"] = this.TransChargeStatus
	}
	if !checkNull(this.MerchantName) {
		(*esMap)["merchantName"] = this.MerchantName
	}
	if !checkNull(this.UserInfo) {
		this.GetUserInfo().GetEsMap(makeKeyMap(esMap, "userInfo"))
	}
	if !checkNull(this.CompanyInfo) {
		this.GetCompanyInfo().GetEsMap(makeKeyMap(esMap, "companyInfo"))
	}
	if !checkNull(this.ItemCount) {
		(*esMap)["itemCount"] = this.ItemCount
	}
}
func (this *TransactionMessage3_UserInfo) GetEsMap(esMap *map[string]interface{}) {
	flag.Parse()
	if !checkNull(this.Email) {
		(*esMap)["email"] = this.Email
	}
	if !checkNull(this.Phone) {
		(*esMap)["phoneNumber"] = this.Phone
	}
}
func (this *TransactionMessage3_CompanyInfo) GetEsMap(esMap *map[string]interface{}) {
	flag.Parse()
	if !checkNull(this.UserInf) {
		this.GetUserInf().GetEsMap(makeKeyMap(esMap, "userINFFo"))
	}
}