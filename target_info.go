package offline_alloc

const (
	Sex = iota
	Age
	Area
	TargetNum
)

var AttrNum [TargetNum]int

func InitTargetInfo() {

	AttrNum[Sex] = 3

	AttrNum[Age] = 4

	AttrNum[Area] = 5

}
