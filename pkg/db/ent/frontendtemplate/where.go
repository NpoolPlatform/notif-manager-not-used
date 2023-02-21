// Code generated by ent, DO NOT EDIT.

package frontendtemplate

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// LangID applies equality check predicate on the "lang_id" field. It's identical to LangIDEQ.
func LangID(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// UsedFor applies equality check predicate on the "used_for" field. It's identical to UsedForEQ.
func UsedFor(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// LangIDEQ applies the EQ predicate on the "lang_id" field.
func LangIDEQ(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// LangIDNEQ applies the NEQ predicate on the "lang_id" field.
func LangIDNEQ(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLangID), v))
	})
}

// LangIDIn applies the In predicate on the "lang_id" field.
func LangIDIn(vs ...uuid.UUID) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLangID), v...))
	})
}

// LangIDNotIn applies the NotIn predicate on the "lang_id" field.
func LangIDNotIn(vs ...uuid.UUID) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLangID), v...))
	})
}

// LangIDGT applies the GT predicate on the "lang_id" field.
func LangIDGT(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLangID), v))
	})
}

// LangIDGTE applies the GTE predicate on the "lang_id" field.
func LangIDGTE(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLangID), v))
	})
}

// LangIDLT applies the LT predicate on the "lang_id" field.
func LangIDLT(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLangID), v))
	})
}

// LangIDLTE applies the LTE predicate on the "lang_id" field.
func LangIDLTE(v uuid.UUID) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLangID), v))
	})
}

// UsedForEQ applies the EQ predicate on the "used_for" field.
func UsedForEQ(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// UsedForNEQ applies the NEQ predicate on the "used_for" field.
func UsedForNEQ(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedFor), v))
	})
}

// UsedForIn applies the In predicate on the "used_for" field.
func UsedForIn(vs ...string) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsedFor), v...))
	})
}

// UsedForNotIn applies the NotIn predicate on the "used_for" field.
func UsedForNotIn(vs ...string) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsedFor), v...))
	})
}

// UsedForGT applies the GT predicate on the "used_for" field.
func UsedForGT(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedFor), v))
	})
}

// UsedForGTE applies the GTE predicate on the "used_for" field.
func UsedForGTE(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedFor), v))
	})
}

// UsedForLT applies the LT predicate on the "used_for" field.
func UsedForLT(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedFor), v))
	})
}

// UsedForLTE applies the LTE predicate on the "used_for" field.
func UsedForLTE(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedFor), v))
	})
}

// UsedForContains applies the Contains predicate on the "used_for" field.
func UsedForContains(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsedFor), v))
	})
}

// UsedForHasPrefix applies the HasPrefix predicate on the "used_for" field.
func UsedForHasPrefix(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsedFor), v))
	})
}

// UsedForHasSuffix applies the HasSuffix predicate on the "used_for" field.
func UsedForHasSuffix(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsedFor), v))
	})
}

// UsedForIsNil applies the IsNil predicate on the "used_for" field.
func UsedForIsNil() predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsedFor)))
	})
}

// UsedForNotNil applies the NotNil predicate on the "used_for" field.
func UsedForNotNil() predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsedFor)))
	})
}

// UsedForEqualFold applies the EqualFold predicate on the "used_for" field.
func UsedForEqualFold(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsedFor), v))
	})
}

// UsedForContainsFold applies the ContainsFold predicate on the "used_for" field.
func UsedForContainsFold(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsedFor), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.FrontendTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContent)))
	})
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContent)))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FrontendTemplate) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FrontendTemplate) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
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
func Not(p predicate.FrontendTemplate) predicate.FrontendTemplate {
	return predicate.FrontendTemplate(func(s *sql.Selector) {
		p(s.Not())
	})
}
