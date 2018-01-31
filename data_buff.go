package offline_alloc

import (
	"strconv"
)

type DataBuff struct {
	M_OrderItems []*AdOrderItem

	M_TargetAttr2MergeAttr [TargetNum][]int

	M_MergeAttr2TargetAttr [TargetNum][][]int

	M_NumOfMergeAttr [TargetNum]int

	M_FlatMergeAttr2AllMergeAttr []int

	M_AllMergeAttr2FlatMergeAttr [][]int
}

func NewDataBuff() *DataBuff {
	obj := new(DataBuff)
	obj.Initialize()
	return obj
}

func (this *DataBuff) Initialize() {
	//this.TargetAttr2MergeAttr = make([]uint32, 0, 100000)

	//this.MergeTargetAttr() = make([][]uint32, 0, 100000)
	for i := 0; i < TargetNum; i++ {
		this.M_MergeAttr2TargetAttr[i] = make([][]int, 0, 1000)

		this.M_TargetAttr2MergeAttr[i] = make([]int, AttrNum[i], AttrNum[i])

	}

	this.M_AllMergeAttr2FlatMergeAttr = make([][]int, 0, 1000)

}

func (this *DataBuff) AddOrderItem(order *AdOrderItem) {
	this.M_OrderItems = append(this.M_OrderItems, order)
}

func (this *DataBuff) AddOrderItems(orders []*AdOrderItem) {
	this.M_OrderItems = append(this.M_OrderItems, orders...)
}

func (this *DataBuff) MergeTargetAttr(target_type int) {

	attr2orders := make([]string, AttrNum[target_type], AttrNum[target_type])

	for i := 0; i < len(attr2orders); i++ {
		attr2orders[i] = "_"
	}

	for i := 0; i < len(this.M_OrderItems); i++ {
		for j := 0; j < len(this.M_OrderItems[i].M_TargetAttr[target_type]); j++ {
			attr2orders[this.M_OrderItems[i].M_TargetAttr[target_type][j]] += strconv.Itoa(i)
			attr2orders[this.M_OrderItems[i].M_TargetAttr[target_type][j]] += "_"

		}

	}

	orders_new_attr_map := make(map[string]int)

	for attr := 0; attr < AttrNum[target_type]; attr++ {

		new_attr_idx := 0

		if val, ok := orders_new_attr_map[attr2orders[attr]]; ok {
			//do something here
			//orders2attrs[attr2order2[attr]] = append(orders2attrs[attr2order2[attr]], attr)
			new_attr_idx = val

			this.M_MergeAttr2TargetAttr[target_type][new_attr_idx] = append(this.M_MergeAttr2TargetAttr[target_type][new_attr_idx], attr)

		} else {

			attr_slice := make([]int, 0, 100)
			attr_slice = append(attr_slice, attr)

			new_attr_idx = len(orders_new_attr_map)

			orders_new_attr_map[attr2orders[attr]] = new_attr_idx

			//this.MergeAttr2TargetAttr[target_type] = make([][]int, 0, 1000)
			this.M_MergeAttr2TargetAttr[target_type] = append(this.M_MergeAttr2TargetAttr[target_type], attr_slice)

		}

		this.M_TargetAttr2MergeAttr[target_type][attr] = new_attr_idx

	}

	for i := 0; i < len(this.M_OrderItems); i++ {

		this.M_OrderItems[i].MergeAttr(len(this.M_MergeAttr2TargetAttr[target_type]), target_type, this.M_TargetAttr2MergeAttr[target_type])

	}

	this.M_NumOfMergeAttr[target_type] = len(this.M_MergeAttr2TargetAttr[target_type])

}

func (this *DataBuff) AllMergeTargetAttr() {

	total_all_merge_attrs := this.M_NumOfMergeAttr[Sex] * this.M_NumOfMergeAttr[Age] * this.M_NumOfMergeAttr[Area]

	this.M_FlatMergeAttr2AllMergeAttr = make([]int, total_all_merge_attrs)

	all_merge_attr2orders := make([]string, total_all_merge_attrs)

	for i := 0; i < len(all_merge_attr2orders); i++ {
		all_merge_attr2orders[i] = "_"
	}

	order2mergeAttrs := make([][]int, len(this.M_OrderItems))

	for i := 0; i < len(order2mergeAttrs); i++ {
		order2mergeAttrs[i] = make([]int, 0, 1000)
	}

	for i := 0; i < len(this.M_OrderItems); i++ {

		this.MakeAllMergeDimensionOrderMap(i, order2mergeAttrs, all_merge_attr2orders)

	}

	orders_all_merge_attr_map := make(map[string]int)

	for attr := 0; attr < total_all_merge_attrs; attr++ {

		all_merge_index := 0

		if val, ok := orders_all_merge_attr_map[all_merge_attr2orders[attr]]; ok {

			all_merge_index = val

			this.M_AllMergeAttr2FlatMergeAttr[all_merge_index] = append(this.M_AllMergeAttr2FlatMergeAttr[all_merge_index], attr)

		} else {

			attr_slice := make([]int, 0, 1000)

			attr_slice = append(attr_slice, attr)

			all_merge_index = len(orders_all_merge_attr_map)

			orders_all_merge_attr_map[all_merge_attr2orders[attr]] = all_merge_index

			this.M_AllMergeAttr2FlatMergeAttr = append(this.M_AllMergeAttr2FlatMergeAttr, attr_slice)

		}

		this.M_FlatMergeAttr2AllMergeAttr[attr] = all_merge_index

	}

	for i := 0; i < len(this.M_OrderItems); i++ {

		this.M_OrderItems[i].AllMergeAttr(len(this.M_AllMergeAttr2FlatMergeAttr), order2mergeAttrs[i], this.M_FlatMergeAttr2AllMergeAttr)

	}

}

func (this *DataBuff) MakeAllMergeDimensionOrderMap(order_index int, order2mergeAttrs [][]int, all_merge_attr2orders []string) {

	order := this.M_OrderItems[order_index]

	target_sex := order.M_MergeTargetAttr[Sex]

	target_age := order.M_MergeTargetAttr[Age]

	target_area := order.M_MergeTargetAttr[Area]

	order2mergeAttrs[order_index] = make([]int, len(target_sex)*len(target_age)*len(target_area))

	pos := 0

	for sex_index := 0; sex_index < len(target_sex); sex_index++ {

		age_base_idx := target_sex[sex_index] * this.M_NumOfMergeAttr[Age]

		for age_index := 0; age_index < len(target_age); age_index++ {

			area_index_base := (age_base_idx + target_age[age_index]) * this.M_NumOfMergeAttr[Area]

			for area_index := 0; area_index < len(target_area); area_index++ {

				tmp_index := (area_index_base + target_area[area_index])

				order2mergeAttrs[order_index][pos] = tmp_index

				all_merge_attr2orders[tmp_index] += strconv.Itoa(order_index)

				all_merge_attr2orders[tmp_index] += "_"

				pos++

			}

		}

	}

}
