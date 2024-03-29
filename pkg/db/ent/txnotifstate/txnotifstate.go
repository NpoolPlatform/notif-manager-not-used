// Code generated by ent, DO NOT EDIT.

package txnotifstate

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the txnotifstate type in the database.
	Label = "tx_notif_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// FieldNotifState holds the string denoting the notif_state field in the database.
	FieldNotifState = "notif_state"
	// FieldTxType holds the string denoting the tx_type field in the database.
	FieldTxType = "tx_type"
	// Table holds the table name of the txnotifstate in the database.
	Table = "tx_notif_states"
)

// Columns holds all SQL columns for txnotifstate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTxID,
	FieldNotifState,
	FieldTxType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/notif-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultTxID holds the default value on creation for the "tx_id" field.
	DefaultTxID func() uuid.UUID
	// DefaultNotifState holds the default value on creation for the "notif_state" field.
	DefaultNotifState string
	// DefaultTxType holds the default value on creation for the "tx_type" field.
	DefaultTxType string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
