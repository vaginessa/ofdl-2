// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ofdl/ofdl/ent/message"
	"github.com/ofdl/ofdl/ent/messagemedia"
)

// MessageMedia is the model entity for the MessageMedia schema.
type MessageMedia struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MessageID holds the value of the "message_id" field.
	MessageID int `json:"message_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Full holds the value of the "full" field.
	Full string `json:"full,omitempty"`
	// DownloadedAt holds the value of the "downloaded_at" field.
	DownloadedAt time.Time `json:"downloaded_at,omitempty"`
	// StashID holds the value of the "stash_id" field.
	StashID string `json:"stash_id,omitempty"`
	// OrganizedAt holds the value of the "organized_at" field.
	OrganizedAt time.Time `json:"organized_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageMediaQuery when eager-loading is set.
	Edges        MessageMediaEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MessageMediaEdges holds the relations/edges for other nodes in the graph.
type MessageMediaEdges struct {
	// Message holds the value of the message edge.
	Message *Message `json:"message,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MessageOrErr returns the Message value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageMediaEdges) MessageOrErr() (*Message, error) {
	if e.loadedTypes[0] {
		if e.Message == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: message.Label}
		}
		return e.Message, nil
	}
	return nil, &NotLoadedError{edge: "message"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageMedia) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagemedia.FieldID, messagemedia.FieldMessageID:
			values[i] = new(sql.NullInt64)
		case messagemedia.FieldType, messagemedia.FieldFull, messagemedia.FieldStashID:
			values[i] = new(sql.NullString)
		case messagemedia.FieldDownloadedAt, messagemedia.FieldOrganizedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageMedia fields.
func (mm *MessageMedia) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagemedia.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mm.ID = int(value.Int64)
		case messagemedia.FieldMessageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field message_id", values[i])
			} else if value.Valid {
				mm.MessageID = int(value.Int64)
			}
		case messagemedia.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mm.Type = value.String
			}
		case messagemedia.FieldFull:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full", values[i])
			} else if value.Valid {
				mm.Full = value.String
			}
		case messagemedia.FieldDownloadedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field downloaded_at", values[i])
			} else if value.Valid {
				mm.DownloadedAt = value.Time
			}
		case messagemedia.FieldStashID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stash_id", values[i])
			} else if value.Valid {
				mm.StashID = value.String
			}
		case messagemedia.FieldOrganizedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field organized_at", values[i])
			} else if value.Valid {
				mm.OrganizedAt = value.Time
			}
		default:
			mm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MessageMedia.
// This includes values selected through modifiers, order, etc.
func (mm *MessageMedia) Value(name string) (ent.Value, error) {
	return mm.selectValues.Get(name)
}

// QueryMessage queries the "message" edge of the MessageMedia entity.
func (mm *MessageMedia) QueryMessage() *MessageQuery {
	return NewMessageMediaClient(mm.config).QueryMessage(mm)
}

// Update returns a builder for updating this MessageMedia.
// Note that you need to call MessageMedia.Unwrap() before calling this method if this MessageMedia
// was returned from a transaction, and the transaction was committed or rolled back.
func (mm *MessageMedia) Update() *MessageMediaUpdateOne {
	return NewMessageMediaClient(mm.config).UpdateOne(mm)
}

// Unwrap unwraps the MessageMedia entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mm *MessageMedia) Unwrap() *MessageMedia {
	_tx, ok := mm.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageMedia is not a transactional entity")
	}
	mm.config.driver = _tx.drv
	return mm
}

// String implements the fmt.Stringer.
func (mm *MessageMedia) String() string {
	var builder strings.Builder
	builder.WriteString("MessageMedia(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mm.ID))
	builder.WriteString("message_id=")
	builder.WriteString(fmt.Sprintf("%v", mm.MessageID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(mm.Type)
	builder.WriteString(", ")
	builder.WriteString("full=")
	builder.WriteString(mm.Full)
	builder.WriteString(", ")
	builder.WriteString("downloaded_at=")
	builder.WriteString(mm.DownloadedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("stash_id=")
	builder.WriteString(mm.StashID)
	builder.WriteString(", ")
	builder.WriteString("organized_at=")
	builder.WriteString(mm.OrganizedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MessageMediaSlice is a parsable slice of MessageMedia.
type MessageMediaSlice []*MessageMedia
