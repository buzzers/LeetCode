package problem_332

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	// 欧拉图 Hierholzer 算法

	// 构造 map [起点] --> [终点集合] 映射
	tickets2 := make(map[string][]string)
	for _, ticket := range tickets {
		tickets2[ticket[0]] = append(tickets2[ticket[0]], ticket[1])
	}
	// 排序
	for _, tickets := range tickets2 {
		sort.Strings(tickets)
	}

	// 从 JFK 开始搜索
	itinerary := findItinerary0("JFK", tickets2, nil)
	reverse(itinerary)

	return itinerary
}

func findItinerary0(src string, tickets map[string][]string, itinerary []string) []string {
	for {
		// 如果当前节点没有出边，结束递归
		tickets0 := tickets[src]
		if len(tickets0) == 0 {
			break
		}
		// 按字典序递归当前节点的出边
		dst := tickets0[0]
		tickets[src] = tickets0[1:]
		itinerary = findItinerary0(dst, tickets, itinerary)
	}
	// 将当前节点添加到解
	return append(itinerary, src)
}

func reverse(ss []string) {
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-i-1] = ss[len(ss)-i-1], ss[i]
	}
}
