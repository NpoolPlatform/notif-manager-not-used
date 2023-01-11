// Code generated by ent, DO NOT EDIT.

package notif

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// AlreadyRead applies equality check predicate on the "already_read" field. It's identical to AlreadyReadEQ.
func AlreadyRead(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlreadyRead), v))
	})
}

// LangID applies equality check predicate on the "lang_id" field. It's identical to LangIDEQ.
func LangID(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// EventType applies equality check predicate on the "event_type" field. It's identical to EventTypeEQ.
func EventType(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventType), v))
	})
}

// UseTemplate applies equality check predicate on the "use_template" field. It's identical to UseTemplateEQ.
func UseTemplate(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseTemplate), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// EmailSend applies equality check predicate on the "email_send" field. It's identical to EmailSendEQ.
func EmailSend(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailSend), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// AlreadyReadEQ applies the EQ predicate on the "already_read" field.
func AlreadyReadEQ(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlreadyRead), v))
	})
}

// AlreadyReadNEQ applies the NEQ predicate on the "already_read" field.
func AlreadyReadNEQ(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAlreadyRead), v))
	})
}

// AlreadyReadIsNil applies the IsNil predicate on the "already_read" field.
func AlreadyReadIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAlreadyRead)))
	})
}

// AlreadyReadNotNil applies the NotNil predicate on the "already_read" field.
func AlreadyReadNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAlreadyRead)))
	})
}

// LangIDEQ applies the EQ predicate on the "lang_id" field.
func LangIDEQ(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// LangIDNEQ applies the NEQ predicate on the "lang_id" field.
func LangIDNEQ(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLangID), v))
	})
}

// LangIDIn applies the In predicate on the "lang_id" field.
func LangIDIn(vs ...uuid.UUID) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLangID), v...))
	})
}

// LangIDNotIn applies the NotIn predicate on the "lang_id" field.
func LangIDNotIn(vs ...uuid.UUID) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLangID), v...))
	})
}

// LangIDGT applies the GT predicate on the "lang_id" field.
func LangIDGT(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLangID), v))
	})
}

// LangIDGTE applies the GTE predicate on the "lang_id" field.
func LangIDGTE(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLangID), v))
	})
}

// LangIDLT applies the LT predicate on the "lang_id" field.
func LangIDLT(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLangID), v))
	})
}

// LangIDLTE applies the LTE predicate on the "lang_id" field.
func LangIDLTE(v uuid.UUID) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLangID), v))
	})
}

// LangIDIsNil applies the IsNil predicate on the "lang_id" field.
func LangIDIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLangID)))
	})
}

// LangIDNotNil applies the NotNil predicate on the "lang_id" field.
func LangIDNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLangID)))
	})
}

// EventTypeEQ applies the EQ predicate on the "event_type" field.
func EventTypeEQ(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventType), v))
	})
}

// EventTypeNEQ applies the NEQ predicate on the "event_type" field.
func EventTypeNEQ(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventType), v))
	})
}

// EventTypeIn applies the In predicate on the "event_type" field.
func EventTypeIn(vs ...string) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEventType), v...))
	})
}

// EventTypeNotIn applies the NotIn predicate on the "event_type" field.
func EventTypeNotIn(vs ...string) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEventType), v...))
	})
}

// EventTypeGT applies the GT predicate on the "event_type" field.
func EventTypeGT(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventType), v))
	})
}

// EventTypeGTE applies the GTE predicate on the "event_type" field.
func EventTypeGTE(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventType), v))
	})
}

// EventTypeLT applies the LT predicate on the "event_type" field.
func EventTypeLT(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventType), v))
	})
}

// EventTypeLTE applies the LTE predicate on the "event_type" field.
func EventTypeLTE(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventType), v))
	})
}

// EventTypeContains applies the Contains predicate on the "event_type" field.
func EventTypeContains(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventType), v))
	})
}

// EventTypeHasPrefix applies the HasPrefix predicate on the "event_type" field.
func EventTypeHasPrefix(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventType), v))
	})
}

// EventTypeHasSuffix applies the HasSuffix predicate on the "event_type" field.
func EventTypeHasSuffix(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventType), v))
	})
}

// EventTypeIsNil applies the IsNil predicate on the "event_type" field.
func EventTypeIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEventType)))
	})
}

// EventTypeNotNil applies the NotNil predicate on the "event_type" field.
func EventTypeNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEventType)))
	})
}

// EventTypeEqualFold applies the EqualFold predicate on the "event_type" field.
func EventTypeEqualFold(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventType), v))
	})
}

// EventTypeContainsFold applies the ContainsFold predicate on the "event_type" field.
func EventTypeContainsFold(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventType), v))
	})
}

// UseTemplateEQ applies the EQ predicate on the "use_template" field.
func UseTemplateEQ(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseTemplate), v))
	})
}

// UseTemplateNEQ applies the NEQ predicate on the "use_template" field.
func UseTemplateNEQ(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUseTemplate), v))
	})
}

// UseTemplateIsNil applies the IsNil predicate on the "use_template" field.
func UseTemplateIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUseTemplate)))
	})
}

// UseTemplateNotNil applies the NotNil predicate on the "use_template" field.
func UseTemplateNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUseTemplate)))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Notif {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContent)))
	})
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContent)))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// ChannelsIsNil applies the IsNil predicate on the "channels" field.
func ChannelsIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChannels)))
	})
}

// ChannelsNotNil applies the NotNil predicate on the "channels" field.
func ChannelsNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChannels)))
	})
}

// EmailSendEQ applies the EQ predicate on the "email_send" field.
func EmailSendEQ(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailSend), v))
	})
}

// EmailSendNEQ applies the NEQ predicate on the "email_send" field.
func EmailSendNEQ(v bool) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmailSend), v))
	})
}

// EmailSendIsNil applies the IsNil predicate on the "email_send" field.
func EmailSendIsNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmailSend)))
	})
}

// EmailSendNotNil applies the NotNil predicate on the "email_send" field.
func EmailSendNotNil() predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmailSend)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notif) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notif) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Notif) predicate.Notif {
	return predicate.Notif(func(s *sql.Selector) {
		p(s.Not())
	})
}
