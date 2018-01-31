package offline_alloc

import (
//"fmt"
)

type AdOrderItem struct {
	M_TargetAttr         [TargetNum][]int
	M_MergeTargetAttr    [TargetNum][]int
	M_AllMergeTargetAttr []int
}

func NewAdOrderItem() *AdOrderItem {
	obj := new(AdOrderItem)

	obj.M_AllMergeTargetAttr = make([]int, 0, 1000)
	for i := 0; i < TargetNum; i++ {
		obj.M_TargetAttr[i] = make([]int, 0, 1000)
		obj.M_MergeTargetAttr[i] = make([]int, 0, 1000)
	}

	return obj

}

func (this *AdOrderItem) MergeAttr(merge_size int, target_type int, targetAttr2mergeAttr []int) {

	exist_map := make(map[int]interface{})

	for i := 0; i < len(this.M_TargetAttr[target_type]); i++ {

		attr := this.M_TargetAttr[target_type][i]

		new_attr_idx := targetAttr2mergeAttr[attr]

		if _, ok := exist_map[new_attr_idx]; ok {

		} else {
			exist_map[new_attr_idx] = new(interface{})
			this.M_MergeTargetAttr[target_type] = append(this.M_MergeTargetAttr[target_type], new_attr_idx)
		}

	}

	if len(this.M_TargetAttr[target_type]) == 0 {

		for i := 0; i < merge_size; i++ {
			this.M_MergeTargetAttr[target_type] = append(this.M_MergeTargetAttr[target_type], i)
		}
	}

}

func (this *AdOrderItem) AllMergeAttr(all_merge_size int, order_flat_merge_attr []int, flat_merge2all_merge []int) {

	exist_map := make(map[int]interface{})

	for i := 0; i < len(order_flat_merge_attr); i++ {

		all_merge_index := flat_merge2all_merge[order_flat_merge_attr[i]]

		if _, ok := exist_map[all_merge_index]; ok {

		} else {

			exist_map[all_merge_index] = new(interface{})

			this.M_AllMergeTargetAttr = append(this.M_AllMergeTargetAttr, all_merge_index)

		}

	}

	if len(order_flat_merge_attr) == 0 {
		for i := 0; i < all_merge_size; i++ {
			this.M_AllMergeTargetAttr = append(this.M_AllMergeTargetAttr, i)
		}
	}
}
