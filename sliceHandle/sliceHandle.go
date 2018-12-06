package sliceHandle

/*
函数功能：删除或添加切片元素
参数说明：
	slice：切片，需要进行删除的切片
	index：要删除的下标，从这个下标开始
	n：删除的个数
	adds：要添加的int数
返回值：
	删除后的切片([]int)，是否成功(bool)
*/
func addOrDelS(slice []int, index int, n int, adds ...int) ([]int, bool) {
	if len(slice) == 0 || index > index+n || index < 0 {
		return slice, false
	} else {
		var temp = make([]int, len(slice[:index]))
		copy(temp, slice)
		for _, v := range adds {
			temp = append(temp, v)
		}
		fmt.Println(temp, slice)
		if index+n > len(slice) {
			return append(temp, slice[len(slice):]...), true
		} else {
			return append(temp, slice[index+n:]...), true
		}
	}
}

