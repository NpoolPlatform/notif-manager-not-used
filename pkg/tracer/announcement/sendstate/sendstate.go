package sendstate

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
)

func trace(span trace1.Span, in *npool.SendStateReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("AnnouncementID.%v", index), in.GetAnnouncementID()),
		attribute.String(fmt.Sprintf("Channel.%v", index), in.GetChannel().String()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.SendStateReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Value", in.GetUserID().GetValue()),
		attribute.String("AnnouncementID.Op", in.GetAnnouncementID().GetOp()),
		attribute.String("AnnouncementID.Value", in.GetAnnouncementID().GetValue()),
		attribute.String("Channel.Op", in.GetChannel().GetOp()),
		attribute.Int("Channel.Value", int(in.GetChannel().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.SendStateReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
