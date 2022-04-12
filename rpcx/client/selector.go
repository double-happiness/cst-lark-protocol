package client

import (
	"context"
	"net/url"
	"strconv"
)

// 自定义selector
type SelectByUser struct { // 实现Selector接口
	/**
	mpNode	节点对应addr
	key		nodeId
	val		addr
	*/
	server map[int]string
	nodeId int
}

// 路由选择服务节点算法
func (s *SelectByUser) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) (node string) {
	// 使用给定节点路由
	if serviceMethod != "SyncRouteTable" {
		node = s.server[s.nodeId]
	} else {
		// 返回随机节点
		for _, v := range s.server {
			node = v
		}
	}
	return
}

// 更新新的服务节点算法
// watch的时候调用 UpdateServer
func (s *SelectByUser) UpdateServer(servers map[string]string) {
	/**
	metadata
	key		tcp@10.42.0.187:19099
	value	node=nodeId
	*/
	tmp := make(map[int]string)
	for k, metadata := range servers {
		if v, err := url.ParseQuery(metadata); err == nil {
			ww := v.Get("node")
			if ww != "" {
				if nodeId, errAtoi := strconv.Atoi(ww); errAtoi == nil {
					tmp[nodeId] = k
				}
			}
		}
	}
	s.server = tmp
	tmp = nil
}
