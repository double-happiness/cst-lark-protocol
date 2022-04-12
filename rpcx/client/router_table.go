package client

import "cst-lark-protocal/rpcx/proto"

type RouterTable struct {
	SlotNode map[int]int
}

func newRouterTable() *RouterTable {
	return &RouterTable{}
}

func (s *RouterTable) transform(args *proto.SyncRouteTableResp) *RouterTable {
	// todo:将lark-seq返回结果转为map
	return &RouterTable{}
}
