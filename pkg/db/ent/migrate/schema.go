// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnnouncementsColumns holds the columns for the "announcements" table.
	AnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "lang_id", Type: field.TypeUUID, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "channels", Type: field.TypeJSON, Nullable: true},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "type", Type: field.TypeString, Nullable: true, Default: "DefaultAnnouncementType"},
	}
	// AnnouncementsTable holds the schema information for the "announcements" table.
	AnnouncementsTable = &schema.Table{
		Name:       "announcements",
		Columns:    AnnouncementsColumns,
		PrimaryKey: []*schema.Column{AnnouncementsColumns[0]},
	}
	// NotifsColumns holds the columns for the "notifs" table.
	NotifsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "already_read", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "lang_id", Type: field.TypeUUID, Nullable: true},
		{Name: "event_type", Type: field.TypeString, Nullable: true, Default: "DefaultEventType"},
		{Name: "use_template", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "channels", Type: field.TypeJSON, Nullable: true},
		{Name: "email_send", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// NotifsTable holds the schema information for the "notifs" table.
	NotifsTable = &schema.Table{
		Name:       "notifs",
		Columns:    NotifsColumns,
		PrimaryKey: []*schema.Column{NotifsColumns[0]},
	}
	// ReadAnnouncementsColumns holds the columns for the "read_announcements" table.
	ReadAnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "announcement_id", Type: field.TypeUUID, Nullable: true},
	}
	// ReadAnnouncementsTable holds the schema information for the "read_announcements" table.
	ReadAnnouncementsTable = &schema.Table{
		Name:       "read_announcements",
		Columns:    ReadAnnouncementsColumns,
		PrimaryKey: []*schema.Column{ReadAnnouncementsColumns[0]},
	}
	// SendAnnouncementsColumns holds the columns for the "send_announcements" table.
	SendAnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "announcement_id", Type: field.TypeUUID, Nullable: true},
		{Name: "channel", Type: field.TypeString, Nullable: true, Default: "DefaultChannel"},
	}
	// SendAnnouncementsTable holds the schema information for the "send_announcements" table.
	SendAnnouncementsTable = &schema.Table{
		Name:       "send_announcements",
		Columns:    SendAnnouncementsColumns,
		PrimaryKey: []*schema.Column{SendAnnouncementsColumns[0]},
	}
	// UserAnnouncementsColumns holds the columns for the "user_announcements" table.
	UserAnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "announcement_id", Type: field.TypeUUID, Nullable: true},
	}
	// UserAnnouncementsTable holds the schema information for the "user_announcements" table.
	UserAnnouncementsTable = &schema.Table{
		Name:       "user_announcements",
		Columns:    UserAnnouncementsColumns,
		PrimaryKey: []*schema.Column{UserAnnouncementsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnnouncementsTable,
		NotifsTable,
		ReadAnnouncementsTable,
		SendAnnouncementsTable,
		UserAnnouncementsTable,
	}
)

func init() {
}
