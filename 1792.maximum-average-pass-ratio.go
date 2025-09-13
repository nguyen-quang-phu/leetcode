package leetcode

import "container/heap"

// @leet start

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    mh := minHeap{}
    for _, v := range classes {
        mh = append(mh, myClasses{
            info: v,
            gain: (float64(v[0] + 1)) / (float64(v[1] + 1)) - float64(v[0]) / float64(v[1]),
        })
    }

    heap.Init(&mh)
    for i := 0; i < extraStudents; i++ {
        mh[0].info[0] ++
        mh[0].info[1] ++
        mh[0].gain = (float64(mh[0].info[0] + 1)) / (float64(mh[0].info[1] + 1)) -
        float64(mh[0].info[0]) / float64(mh[0].info[1])
        heap.Fix(&mh, 0)
    }

    var enum float64 = 0
    for _, v := range mh {
        enum += float64(v.info[0]) / float64(v.info[1])
    }

    return enum / float64(len(mh))
}

type myClasses struct {
    info []int
    gain float64
}

type minHeap []myClasses

func (m minHeap) Len() int {
    return len(m)
}

func (m minHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m minHeap) Less(i, j int) bool {
    return m[i].gain > m[j].gain
}

func (m *minHeap) Push(x any) {
    *m = append(*m, x.(myClasses))
}

func (m *minHeap) Pop() any {
    x := (*m)[len(*m) - 1]
    *m = (*m)[0:len(*m) - 1]
    return x
}
// @leet end

// Keynold
