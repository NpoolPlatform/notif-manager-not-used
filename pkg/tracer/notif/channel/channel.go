package channel

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"
)

func trace(span trace1.Span, in *npool.ChannelReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("EventType.%v", index), in.GetEventType().String()),
		attribute.String(fmt.Sprintf("Channel.%v", index), in.GetChannel().String()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.ChannelReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.ChannelReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
