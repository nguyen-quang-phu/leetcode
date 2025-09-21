package leetcode

// @leet start
import "slices"

type Router struct {
	size    int
	packets [][3]int
	ids     map[[3]int]struct{}
	dest    map[int][]int
}

func Constructor(memoryLimit int) Router {
	r := Router{
		size: memoryLimit,
		ids:  make(map[[3]int]struct{}),
		dest: make(map[int][]int),
	}
	return r
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	if _, ok := this.ids[[3]int{source, destination, timestamp}]; ok {
		return false
	}
	if len(this.packets) == this.size {
		p := this.packets[0]
		this.packets = this.packets[1:]
		delete(this.ids, p)
		this.dest[p[1]] = this.dest[p[1]][1:]
	}
	this.packets = append(this.packets, [3]int{source, destination, timestamp})
	this.ids[[3]int{source, destination, timestamp}] = struct{}{}
	this.dest[destination] = append(this.dest[destination], timestamp)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.packets) == 0 {
		return []int{}
	}
	p := this.packets[0]
	this.packets = this.packets[1:]
	delete(this.ids, p)
	this.dest[p[1]] = this.dest[p[1]][1:]
	return []int{p[0], p[1], p[2]}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	times := this.dest[destination]
	left, _ := slices.BinarySearch(times, startTime)
	right, _ := slices.BinarySearch(times, endTime+1)
	return right - left
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
// @leet end

// Keynold
