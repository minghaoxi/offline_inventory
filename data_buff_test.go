package offline_alloc

import (
	"fmt"
	"testing"
)

func MockAdOrderItem() *AdOrderItem {
	obj := NewAdOrderItem()

	/*sex_arr := make([]int, 1)
	sex_arr[0] = 1
	obj.M_TargetAttr[Sex] = sex_arr

	age_arr := make([]int, 2)
	age_arr[0] = 2
	age_arr[1] = 3
	obj.M_TargetAttr[Age] = age_arr*/

	area_arr := make([]int, 1)
	area_arr[0] = 4
	obj.M_TargetAttr[Area] = area_arr

	return obj

}

func MockAdOrderItem2() *AdOrderItem {
	obj := NewAdOrderItem()

	/*sex_arr := make([]int, 1)
	sex_arr[0] = 1
	obj.M_TargetAttr[Sex] = sex_arr*/

	age_arr := make([]int, 2)
	age_arr[0] = 1
	age_arr[1] = 3
	obj.M_TargetAttr[Age] = age_arr

	/*area_arr := make([]int, 1)
	area_arr[0] = 4
	obj.M_TargetAttr[Area] = area_arr*/

	return obj

}

func TestMergeTargetAttr(t *testing.T) {

	InitTargetInfo()

	data_buff := NewDataBuff()

	order := MockAdOrderItem()

	order2 := MockAdOrderItem2()

	data_buff.AddOrderItem(order)

	data_buff.AddOrderItem(order2)

	for i := 0; i < TargetNum; i++ {

		data_buff.MergeTargetAttr(i)

	}

	data_buff.AllMergeTargetAttr()

	//fmt.Println(data_buff.M_MergeAttr2TargetAttr)

	//fmt.Println(data_buff.M_MergeAttr2TargetAttr)

	//fmt.Println(data_buff.M_FlatMergeAttr2AllMergeAttr)

	fmt.Println(data_buff.M_AllMergeAttr2FlatMergeAttr)

	fmt.Println(order)

	fmt.Println(order2)

}
