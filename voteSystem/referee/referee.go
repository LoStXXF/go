package referee

import (
	. "voteSystem/person"
)

// 候选人map，存放候选人地址，候选人ID=>候选人地址
var Candidates = make(map[uint]*Candidate)

/*加入候选map中，可以同时加入多个
 *参数:
 *	candidates:(...*Candidate):候选人地址，可以多个，也可以一个
 */
func JoinCandidate(candidates ...*Candidate) {

	//将候选人加入map容器中
	for _, value := range candidates {
		Candidates[value.Id] = value
	}
}

/*统计获胜者
 *返回值:
 *	Candidate:获胜候选人
 */
func Statistical() Candidate {
	var max uint = 0   //保存最大票数
	var index uint = 0 //保存最大票数的候选人的id
	for key, value := range Candidates {
		if max < value.Voters { //如果有票数大于max的，进行重新赋值
			max = value.Voters
			index = key //保存最大票数的id
		}
	}
	return *Candidates[index]
}
