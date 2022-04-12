package client

import (
	"context"
	"cst-lark-protocal/rpcx/proto"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	nacosclient "github.com/rpcxio/rpcx-nacos/client"
	"github.com/smallnest/rpcx/client"
	"log"
)

type Client struct {
	LarkClient proto.LarkClient
}

func NewRpcxClient(clientConfig constant.ClientConfig, serverConfig []constant.ServerConfig) (larkRpcClient *proto.LarkClient) {
	opt := client.DefaultOption
	register, err := nacosclient.NewNacosDiscovery("", "", "", clientConfig, serverConfig)
	if err != nil {
		log.Panicln(err)
	}
	xClient := client.NewXClient("", client.Failover, client.SelectByUser, register, opt)
	return proto.NewLarkClient(xClient)
}

/**
SyncRouteTable	直接将路由表以map[slot]node形式返回
*/
func (c *Client) SyncRouteTable(ctx context.Context) (resp map[int]int, err error) {
	c.LarkClient.SyncRouteTable(ctx, nil)
	req := &proto.SyncRouteTableReq{}
	reply, err := c.LarkClient.SyncRouteTable(ctx, req)
	if err != nil {
		log.Println("SyncRouteTable err: ", err)
		return
	}

	return newRouterTable().transform(reply).SlotNode, nil
}

/**
NextSeq	获取指定biz_tag的下一个seqId
*/
func (c *Client) NextSeq(ctx context.Context, bizTag string) (seqId int64, err error) {
	req := &proto.NextSeqReq{BizTag: bizTag}
	resp, err := c.LarkClient.NextSeq(ctx, req)
	if err != nil {
		log.Println("SyncRouteTable err: ", err)
		return
	}
	return resp.SeqId, nil
}

/**
Segment	获取指定biz_tag的固定数量的id
*/
func (c *Client) Segment(ctx context.Context, bizTag string, count int32) (segmentResp *SegmentResp, err error) {
	req := &proto.SegmentReq{
		BizTag: bizTag,
		Count:  count,
	}
	resp, err := c.LarkClient.Segment(ctx, req)
	if err != nil {
		log.Println("SyncRouteTable err: ", err)
		return
	}
	return newSegmentResp().transform(resp), nil
}
