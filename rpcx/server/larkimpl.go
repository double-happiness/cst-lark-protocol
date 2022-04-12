package server

import (
	"context"
	"cst-lark-protocal/rpcx/proto"
)

type LarkImpl struct {
}

func (s *LarkImpl) SyncRouteTable(ctx context.Context, args *proto.SyncRouteTableReq, reply *proto.SyncRouteTableResp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = proto.SyncRouteTableResp{}

	return nil
}

// NextSeq is server rpc method as defined
func (s *LarkImpl) NextSeq(ctx context.Context, args *proto.NextSeqReq, reply *proto.NextSeqResp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = proto.NextSeqResp{}

	return nil
}

// Segment is server rpc method as defined
func (s *LarkImpl) Segment(ctx context.Context, args *proto.SegmentReq, reply *proto.SegmentResp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = proto.SegmentResp{}

	return nil
}
