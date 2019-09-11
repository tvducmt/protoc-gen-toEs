package test

import (
	"encoding/json"
	fmt "fmt"
	"testing"

	"git.zapa.cloud/merchant-tools/helper/proto"
	proto1 "git.zapa.cloud/merchant-tools/helper/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func buildProto3(createTime *timestamp.Timestamp, toDate *proto1.Date,
	pmcName string, discountAmount int64, transCode string, transChargeStatus TransactionMessage3_TransChargeStatus,
	merchantName, email, phone, address, name string, itemCount int32) *TransactionMessage3 {

	userInfoEmbedded := &TransactionMessage3_UserInfo{
		Email:   email,
		Phone:   phone,
		Address: address,
	}
	companyInfoEmbedded := &TransactionMessage3_CompanyInfo{
		Name: name,
	}

	goodProto3 := &TransactionMessage3{
		CreateTime:        createTime,
		ToDate:            toDate,
		PmcName:           pmcName,
		DiscountAmount:    discountAmount,
		TransCode:         transCode,
		TransChargeStatus: transChargeStatus,
		MerchantName:      merchantName,
		UserInfo:          userInfoEmbedded,
		CompanyInfo:       companyInfoEmbedded,
		ItemCount:         itemCount,
	}
	return goodProto3
}

func TestGoodProto3(t *testing.T) {

	createTime := &timestamp.Timestamp{Seconds: 60}
	toDate := &proto.Date{Year: 2019, Month: 9, Day: 23}
	// transChargeStatus := 12
	// userInfo := &TransactionMessage3_UserInfo{
	// 	Email: "tvduc@gmail.com",
	// 	Phone: "0364859473",
	// }
	// companyInfo := &TransactionMessage3_CompanyInfo{
	// 	UserInf: &TransactionMessage3_UserInfo{
	// 		Email: "ducttan@",
	// 	},
	// }

	// var err error
	goodProto3 := buildProto3(createTime, toDate, "pmcNameVal", 10, "transcodeVal", TransactionMessage3_TCS_PENDING, "merchantNameVal", "tvduc@gmail.com", "0363637773", "dia chi", "nameVal", 32)
	esMap := map[string]interface{}{}
	goodProto3.GetEsMap(&esMap)
	src, _ := json.Marshal(esMap)
	fmt.Println(string(src))
	// if err != nil {
	// 	t.Fatalf("unexpected fail in validator: %v", err)
	// }
}
