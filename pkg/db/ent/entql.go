// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notif"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/readannouncement"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   announcement.Table,
			Columns: announcement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: announcement.FieldID,
			},
		},
		Type: "Announcement",
		Fields: map[string]*sqlgraph.FieldSpec{
			announcement.FieldCreatedAt: {Type: field.TypeUint32, Column: announcement.FieldCreatedAt},
			announcement.FieldUpdatedAt: {Type: field.TypeUint32, Column: announcement.FieldUpdatedAt},
			announcement.FieldDeletedAt: {Type: field.TypeUint32, Column: announcement.FieldDeletedAt},
			announcement.FieldAppID:     {Type: field.TypeUUID, Column: announcement.FieldAppID},
			announcement.FieldTitle:     {Type: field.TypeString, Column: announcement.FieldTitle},
			announcement.FieldContent:   {Type: field.TypeString, Column: announcement.FieldContent},
			announcement.FieldChannels:  {Type: field.TypeJSON, Column: announcement.FieldChannels},
			announcement.FieldEmailSend: {Type: field.TypeBool, Column: announcement.FieldEmailSend},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   notif.Table,
			Columns: notif.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notif.FieldID,
			},
		},
		Type: "Notif",
		Fields: map[string]*sqlgraph.FieldSpec{
			notif.FieldCreatedAt:   {Type: field.TypeUint32, Column: notif.FieldCreatedAt},
			notif.FieldUpdatedAt:   {Type: field.TypeUint32, Column: notif.FieldUpdatedAt},
			notif.FieldDeletedAt:   {Type: field.TypeUint32, Column: notif.FieldDeletedAt},
			notif.FieldAppID:       {Type: field.TypeUUID, Column: notif.FieldAppID},
			notif.FieldUserID:      {Type: field.TypeUUID, Column: notif.FieldUserID},
			notif.FieldAlreadyRead: {Type: field.TypeBool, Column: notif.FieldAlreadyRead},
			notif.FieldLangID:      {Type: field.TypeUUID, Column: notif.FieldLangID},
			notif.FieldEventType:   {Type: field.TypeString, Column: notif.FieldEventType},
			notif.FieldUseTemplate: {Type: field.TypeBool, Column: notif.FieldUseTemplate},
			notif.FieldTitle:       {Type: field.TypeString, Column: notif.FieldTitle},
			notif.FieldContent:     {Type: field.TypeString, Column: notif.FieldContent},
			notif.FieldChannels:    {Type: field.TypeJSON, Column: notif.FieldChannels},
			notif.FieldEmailSend:   {Type: field.TypeBool, Column: notif.FieldEmailSend},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   readannouncement.Table,
			Columns: readannouncement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: readannouncement.FieldID,
			},
		},
		Type: "ReadAnnouncement",
		Fields: map[string]*sqlgraph.FieldSpec{
			readannouncement.FieldCreatedAt:      {Type: field.TypeUint32, Column: readannouncement.FieldCreatedAt},
			readannouncement.FieldUpdatedAt:      {Type: field.TypeUint32, Column: readannouncement.FieldUpdatedAt},
			readannouncement.FieldDeletedAt:      {Type: field.TypeUint32, Column: readannouncement.FieldDeletedAt},
			readannouncement.FieldAppID:          {Type: field.TypeUUID, Column: readannouncement.FieldAppID},
			readannouncement.FieldUserID:         {Type: field.TypeUUID, Column: readannouncement.FieldUserID},
			readannouncement.FieldAnnouncementID: {Type: field.TypeUUID, Column: readannouncement.FieldAnnouncementID},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *AnnouncementQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AnnouncementQuery builder.
func (aq *AnnouncementQuery) Filter() *AnnouncementFilter {
	return &AnnouncementFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *AnnouncementMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AnnouncementMutation builder.
func (m *AnnouncementMutation) Filter() *AnnouncementFilter {
	return &AnnouncementFilter{config: m.config, predicateAdder: m}
}

// AnnouncementFilter provides a generic filtering capability at runtime for AnnouncementQuery.
type AnnouncementFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AnnouncementFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AnnouncementFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(announcement.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AnnouncementFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(announcement.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AnnouncementFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(announcement.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AnnouncementFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(announcement.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AnnouncementFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(announcement.FieldAppID))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *AnnouncementFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(announcement.FieldTitle))
}

// WhereContent applies the entql string predicate on the content field.
func (f *AnnouncementFilter) WhereContent(p entql.StringP) {
	f.Where(p.Field(announcement.FieldContent))
}

// WhereChannels applies the entql json.RawMessage predicate on the channels field.
func (f *AnnouncementFilter) WhereChannels(p entql.BytesP) {
	f.Where(p.Field(announcement.FieldChannels))
}

// WhereEmailSend applies the entql bool predicate on the email_send field.
func (f *AnnouncementFilter) WhereEmailSend(p entql.BoolP) {
	f.Where(p.Field(announcement.FieldEmailSend))
}

// addPredicate implements the predicateAdder interface.
func (nq *NotifQuery) addPredicate(pred func(s *sql.Selector)) {
	nq.predicates = append(nq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the NotifQuery builder.
func (nq *NotifQuery) Filter() *NotifFilter {
	return &NotifFilter{config: nq.config, predicateAdder: nq}
}

// addPredicate implements the predicateAdder interface.
func (m *NotifMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the NotifMutation builder.
func (m *NotifMutation) Filter() *NotifFilter {
	return &NotifFilter{config: m.config, predicateAdder: m}
}

// NotifFilter provides a generic filtering capability at runtime for NotifQuery.
type NotifFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *NotifFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *NotifFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(notif.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *NotifFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(notif.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *NotifFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(notif.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *NotifFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(notif.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *NotifFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(notif.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *NotifFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(notif.FieldUserID))
}

// WhereAlreadyRead applies the entql bool predicate on the already_read field.
func (f *NotifFilter) WhereAlreadyRead(p entql.BoolP) {
	f.Where(p.Field(notif.FieldAlreadyRead))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *NotifFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(notif.FieldLangID))
}

// WhereEventType applies the entql string predicate on the event_type field.
func (f *NotifFilter) WhereEventType(p entql.StringP) {
	f.Where(p.Field(notif.FieldEventType))
}

// WhereUseTemplate applies the entql bool predicate on the use_template field.
func (f *NotifFilter) WhereUseTemplate(p entql.BoolP) {
	f.Where(p.Field(notif.FieldUseTemplate))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *NotifFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(notif.FieldTitle))
}

// WhereContent applies the entql string predicate on the content field.
func (f *NotifFilter) WhereContent(p entql.StringP) {
	f.Where(p.Field(notif.FieldContent))
}

// WhereChannels applies the entql json.RawMessage predicate on the channels field.
func (f *NotifFilter) WhereChannels(p entql.BytesP) {
	f.Where(p.Field(notif.FieldChannels))
}

// WhereEmailSend applies the entql bool predicate on the email_send field.
func (f *NotifFilter) WhereEmailSend(p entql.BoolP) {
	f.Where(p.Field(notif.FieldEmailSend))
}

// addPredicate implements the predicateAdder interface.
func (raq *ReadAnnouncementQuery) addPredicate(pred func(s *sql.Selector)) {
	raq.predicates = append(raq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ReadAnnouncementQuery builder.
func (raq *ReadAnnouncementQuery) Filter() *ReadAnnouncementFilter {
	return &ReadAnnouncementFilter{config: raq.config, predicateAdder: raq}
}

// addPredicate implements the predicateAdder interface.
func (m *ReadAnnouncementMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ReadAnnouncementMutation builder.
func (m *ReadAnnouncementMutation) Filter() *ReadAnnouncementFilter {
	return &ReadAnnouncementFilter{config: m.config, predicateAdder: m}
}

// ReadAnnouncementFilter provides a generic filtering capability at runtime for ReadAnnouncementQuery.
type ReadAnnouncementFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ReadAnnouncementFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ReadAnnouncementFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(readannouncement.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ReadAnnouncementFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(readannouncement.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ReadAnnouncementFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(readannouncement.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ReadAnnouncementFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(readannouncement.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ReadAnnouncementFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(readannouncement.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *ReadAnnouncementFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(readannouncement.FieldUserID))
}

// WhereAnnouncementID applies the entql [16]byte predicate on the announcement_id field.
func (f *ReadAnnouncementFilter) WhereAnnouncementID(p entql.ValueP) {
	f.Where(p.Field(readannouncement.FieldAnnouncementID))
}
