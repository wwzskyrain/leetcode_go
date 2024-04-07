package Throne_Inheritance

// ThroneInheritance 1600. 王位继承顺序
type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		king:  kingName,
		edges: map[string][]string{},
		dead:  map[string]bool{},
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.edges[parentName] = append(this.edges[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var preOrder func(s string)
	var ans []string
	preOrder = func(name string) {
		if !this.dead[name] {
			ans = append(ans, name)
		}
		for _, childName := range this.edges[name] {
			preOrder(childName)
		}
	}
	preOrder(this.king)
	return ans
}
