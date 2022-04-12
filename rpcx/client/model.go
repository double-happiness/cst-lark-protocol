package client

import "cst-lark-protocal/rpcx/proto"

type SegmentResp struct {
	BizTag string `json:"biz_tag"`
	Start  int64  `json:"start"` // 开始seq
	End    int64  `json:"end"`   // 结束seq
}

func newSegmentResp() *SegmentResp {
	return &SegmentResp{}
}

func (s *SegmentResp) transform(args *proto.SegmentResp) *SegmentResp {
	return &SegmentResp{
		BizTag: args.BizTag,
		Start:  args.Start,
		End:    args.End,
	}
}
