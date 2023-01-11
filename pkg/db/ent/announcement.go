// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	"github.com/google/uuid"
)

// Announcement is the model entity for the Announcement schema.
type Announcement struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Channels holds the value of the "channels" field.
	Channels []string `json:"channels,omitempty"`
	// EmailSend holds the value of the "email_send" field.
	EmailSend bool `json:"email_send,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Announcement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case announcement.FieldChannels:
			values[i] = new([]byte)
		case announcement.FieldEmailSend:
			values[i] = new(sql.NullBool)
		case announcement.FieldCreatedAt, announcement.FieldUpdatedAt, announcement.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case announcement.FieldTitle, announcement.FieldContent:
			values[i] = new(sql.NullString)
		case announcement.FieldID, announcement.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Announcement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Announcement fields.
func (a *Announcement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case announcement.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case announcement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = uint32(value.Int64)
			}
		case announcement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = uint32(value.Int64)
			}
		case announcement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = uint32(value.Int64)
			}
		case announcement.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				a.AppID = *value
			}
		case announcement.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case announcement.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case announcement.FieldChannels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field channels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Channels); err != nil {
					return fmt.Errorf("unmarshal field channels: %w", err)
				}
			}
		case announcement.FieldEmailSend:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_send", values[i])
			} else if value.Valid {
				a.EmailSend = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Announcement.
// Note that you need to call Announcement.Unwrap() before calling this method if this Announcement
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Announcement) Update() *AnnouncementUpdateOne {
	return (&AnnouncementClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Announcement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Announcement) Unwrap() *Announcement {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Announcement is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Announcement) String() string {
	var builder strings.Builder
	builder.WriteString("Announcement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AppID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(a.Content)
	builder.WriteString(", ")
	builder.WriteString("channels=")
	builder.WriteString(fmt.Sprintf("%v", a.Channels))
	builder.WriteString(", ")
	builder.WriteString("email_send=")
	builder.WriteString(fmt.Sprintf("%v", a.EmailSend))
	builder.WriteByte(')')
	return builder.String()
}

// Announcements is a parsable slice of Announcement.
type Announcements []*Announcement

func (a Announcements) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}