type RandomizedSet struct {
	data   []int
	access map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:   make([]int, 0, 10),
		access: make(map[int]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.access[val]; ok {
		return false
	} else {
		this.data = append(this.data, val)
		this.access[val] = len(this.data) - 1
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.access[val]; ok {
		delete(this.access, val)
		if pos != len(this.data)-1 {
			this.access[this.data[len(this.data)-1]] = pos
			this.data[pos] = this.data[len(this.data)-1]
		}
		this.data = this.data[:len(this.data)-1]
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.IntN(len(this.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */