package announcement

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
)

func trace(span trace1.Span, in *npool.AnnouncementReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("Title.%v", index), in.GetTitle()),
		attribute.String(fmt.Sprintf("Content.%v", index), in.GetContent()),
		attribute.Int(fmt.Sprintf("EndAt.%v", index), int(in.GetEndAt())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AnnouncementReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("Channel.Op", in.GetChannel().GetOp()),
		attribute.String("Channel.Value", channel.NotifChannel(in.GetChannel().GetValue()).String()),
		attribute.String("EndAt.Op", in.GetEndAt().GetOp()),
		attribute.Int("EndAt.Value", int(in.GetEndAt().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AnnouncementReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
