package verify

import "testing"

func TestValidate(t *testing.T) {
	type AddWorkOrder struct {
		Order          string `json:"order" validate:"required" msg:"缺少工单号参数"`
		CustomerNumber int    `json:"customer" validate:"required,gt=10" msg:"客服编号必须大于10"`
	}

	req := AddWorkOrder{
		Order: "123",
	}

	if err := Validate(req); err != nil {
		t.Fatal(TranslateErr2MsgTag(err))
	}
}

func TestValidateGt(t *testing.T) {
	type AddWorkOrder struct {
		Order          string `json:"order" validate:"required" msg:"缺少工单号参数"`
		CustomerNumber int    `json:"customer" validate:"required,gt=10" msg:"客服编号必须大于10"`
	}

	req := AddWorkOrder{
		Order:          "123",
		CustomerNumber: 15,
	}

	if err := Validate(req); err != nil {
		t.Fatal(TranslateErr2MsgTag(err))
	}
}
