package test

import (
	"encoding/json"
	reflect "reflect"
	"testing"

	"git.zapa.cloud/merchant-tools/helper/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func TestTimestampType(t *testing.T) {
	tables := []struct {
		input  *timestamp.Timestamp
		output string
	}{
		{&timestamp.Timestamp{Seconds: 60}, `{"reqTimdestamp":60000}`},
		{&timestamp.Timestamp{Seconds: 20}, `{"reqTimdestamp":20000}`},
		{nil, `{}`},
	}
	for _, table := range tables {
		esMap := map[string]interface{}{}
		transaction := &TransactionMessage3{
			CreateTime: table.input,
		}
		transaction.GetEsMap(&esMap)
		src, err := json.Marshal(esMap)
		if err != nil {
			t.Error(err)
		}
		// fmt.Println(string(src))
		got := string(src)
		want := table.output
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	}
}

func TestDateType(t *testing.T) {
	tables := []struct {
		input  *proto.Date
		output string
	}{
		{&proto.Date{Year: 2019, Month: 9, Day: 23}, `{"reqDate":1569171600000}`},
		{nil, `{}`},
	}
	for _, table := range tables {
		esMap := map[string]interface{}{}
		transaction := &TransactionMessage3{
			ToDate: table.input,
		}
		transaction.GetEsMap(&esMap)
		src, err := json.Marshal(esMap)
		if err != nil {
			t.Error(err)
		}
		// fmt.Println(string(src))
		got := string(src)
		want := table.output
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	}
}

func TestInt64Type(t *testing.T) {
	esMap := map[string]interface{}{}
	transaction := &TransactionMessage3{
		DiscountAmount: 12,
	}
	transaction.GetEsMap(&esMap)
	src, err := json.Marshal(esMap)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got := string(src)
	want := `{"disCountAmount":12}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	esMap1 := map[string]interface{}{}
	transaction1 := &TransactionMessage3{
		DiscountAmount: 0,
	}
	transaction1.GetEsMap(&esMap1)
	src1, err := json.Marshal(esMap1)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got1 := string(src1)
	want1 := `{}`
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
}

func TestInt32Type(t *testing.T) {
	esMap := map[string]interface{}{}
	transaction := &TransactionMessage3{
		ItemCount: 12,
	}
	transaction.GetEsMap(&esMap)
	src, err := json.Marshal(esMap)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got := string(src)
	want := `{"itemCount":12}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	esMap1 := map[string]interface{}{}
	transaction1 := &TransactionMessage3{
		ItemCount: 0,
	}
	transaction1.GetEsMap(&esMap1)
	src1, err := json.Marshal(esMap1)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got1 := string(src1)
	want1 := `{}`
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
}

func TestStringType(t *testing.T) {
	esMap := map[string]interface{}{}
	transaction := &TransactionMessage3{
		MerchantName: "Tiki merchant",
	}
	transaction.GetEsMap(&esMap)
	src, err := json.Marshal(esMap)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got := string(src)
	want := `{"merchantName":"Tiki merchant"}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	esMap1 := map[string]interface{}{}
	transaction1 := &TransactionMessage3{
		MerchantName: "",
	}
	transaction1.GetEsMap(&esMap1)
	src1, err := json.Marshal(esMap1)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got1 := string(src1)
	want1 := `{}`
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
}

func TestEnumType(t *testing.T) {
	esMap := map[string]interface{}{}
	transaction := &TransactionMessage3{
		TransChargeStatus: TransactionMessage3_TCS_PENDING,
	}
	transaction.GetEsMap(&esMap)
	src, err := json.Marshal(esMap)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got := string(src)
	want := `{"transChargeStatus":3}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	esMap1 := map[string]interface{}{}
	transaction1 := &TransactionMessage3{
		TransChargeStatus: 0,
	}
	transaction1.GetEsMap(&esMap1)
	src1, err := json.Marshal(esMap1)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got1 := string(src1)
	want1 := `{}`
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
}

func TestStructType(t *testing.T) {
	esMap := map[string]interface{}{}
	transaction := &TransactionMessage3{
		UserInfo: &TransactionMessage3_UserInfo{
			Email:   "tvduc@gmail.com",
			Phone:   "03672816663",
			Address: "Dia chi",
		},
	}
	transaction.GetEsMap(&esMap)
	src, err := json.Marshal(esMap)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got := string(src)
	want := `{"userInfo":{"email":"tvduc@gmail.com","phoneNumber":"03672816663"}}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	esMap1 := map[string]interface{}{}
	transaction1 := &TransactionMessage3{
		UserInfo: nil,
	}
	transaction1.GetEsMap(&esMap1)
	src1, err := json.Marshal(esMap1)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got1 := string(src1)
	want1 := `{}`
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}

	esMap2 := map[string]interface{}{}
	transaction2 := &TransactionMessage3{
		UserInfo: &TransactionMessage3_UserInfo{
			Email:   "",
			Phone:   "",
			Address: "Dia chi",
		},
	}
	transaction2.GetEsMap(&esMap2)
	src2, err := json.Marshal(esMap2)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got2 := string(src2)
	want2 := `{"userInfo":{}}`
	if !reflect.DeepEqual(want2, got2) {
		t.Fatalf("expected: %v, got: %v", want2, got2)
	}

	esMap3 := map[string]interface{}{}
	transaction3 := &TransactionMessage3{
		UserInfo: &TransactionMessage3_UserInfo{
			Birthday: &proto.Date{Year: 1997, Month: 3, Day: 3},
		},
	}
	transaction3.GetEsMap(&esMap3)
	src3, err := json.Marshal(esMap3)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(src))
	got3 := string(src3)
	want3 := `{"userInfo":{"birthday":857322000000}}`
	if !reflect.DeepEqual(want3, got3) {
		t.Fatalf("expected: %v, got: %v", want3, got3)
	}
}

func TestAllType(t *testing.T) {
	esMap := map[string]interface{}{}
	trans := TransactionMessage3{
		CreateTime:        &timestamp.Timestamp{Seconds: 60},
		ToDate:            &proto.Date{Year: 2019, Month: 9, Day: 23},
		TransChargeStatus: 12,
		UserInfo: &TransactionMessage3_UserInfo{
			Email: "tvduc@gmail.com",
			Phone: "0364859473",
		},
	}
	trans.GetEsMap(&esMap)
	src, _ := json.Marshal(esMap)
	// fmt.Println(string(src))
	got := string(src)
	want := `{"reqDate":1569171600000,"reqTimdestamp":60000,"transChargeStatus":12,"userInfo":{"email":"tvduc@gmail.com","phoneNumber":"0364859473"}}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

}

// func buildProto3(createTime *timestamp.Timestamp, toDate *proto1.Date,
// 	discountAmount int64, transChargeStatus TransactionMessage3_TransChargeStatus,
// 	merchantName, email, phone, address string, itemCount int32) *TransactionMessage3 {

// 	userInfoEmbedded := &TransactionMessage3_UserInfo{
// 		Email:   email,
// 		Phone:   phone,
// 		Address: address,
// 	}
// 	if (&TransactionMessage3_UserInfo{}) == userInfoEmbedded {
// 		fmt.Println("is zero value")
// 	}

// 	goodProto3 := &TransactionMessage3{
// 		CreateTime:        createTime,
// 		ToDate:            toDate,
// 		DiscountAmount:    discountAmount,
// 		TransChargeStatus: transChargeStatus,
// 		MerchantName:      merchantName,
// 		UserInfo:          userInfoEmbedded,
// 		ItemCount:         itemCount,
// 	}
// 	return goodProto3
// }
