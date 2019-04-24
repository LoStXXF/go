package person

type Voter struct {
	id        uint
	name      string     // 投票人姓名
	weight    uint       // 权重
	support   *Candidate //支持的人
	voted     bool
	entruster *Voter
}

var voterId uint = 1

/*创建选民
 *参数:
 *	_name(string):选民姓名
 *返回值:
 *	Voter:经过初始化的选民
 */
func CreateVoter(_name string) Voter {
	var tempVoter Voter
	tempVoter.id = voterId
	tempVoter.name = _name
	tempVoter.weight = 1
	tempVoter.support = nil
	tempVoter.voted = false
	tempVoter.entruster = nil
	voterId += 1
	return tempVoter
}

/*投票
 *参数:
 *	candidate(*Candidate):候选人的地址
 *绑定:
 *  voter(*Voter):选民调用者的地址
 *返回值:
 *	bool:true(成功):false(失败)
 */
func (voter *Voter) Vote(candidate *Candidate) bool {

	// 如果已经进行过投票，那么投票失败，返回false
	if voter.weight == 0 && voter.voted == true {
		return false
	}
	voter.support = candidate             //将候选人的地址保存到support中
	candidate.Voters += voter.weight      //将自己手中的票数加到候选人的票数中
	voter.weight = 0                      //将自己的票数归零
	voter.voted = true                    //设置投票状态为已投，即为true
	candidate.supporter[voter.id] = voter //将选民的地址保存到候选人的支持者容器中
	return true
}

/*委托他人投票
 *参数：
 *	other(*Voter):委托人的地址
 *绑定:
 *  voter(*Voter):自己的地址
 *返回值:
 *  bool:ture(成功):false(失败)
 */
func (voter *Voter) Entrust(other *Voter) bool {
	//检查是不是回环委托，也就是像你委托他，他委托别人，别人委托你
	for v := other.entruster; v != nil; v = v.entruster {
		if *v == *voter { //如果是，返回false，代表委托失败
			return false
		}
	}
	voter.entruster = other
	//如果委托人已经投过票，直接将你的票投给他投过的那个人
	if other.voted == true {
		voter.support = other.support //别人支持的人也就是你支持的人
		voter.Vote(voter.support)     //自己进行投票，实质也是自己在投票
	} else { //如果没有投过票，将自己的票给委托人
		other.weight += voter.weight //委托人没有投过票，就将自己的票数加到委托人的票数中
		voter.weight = 0             //自己的票数归零
		voter.voted = true           //把自己的状态设置为已投票
	}
	return true
}
