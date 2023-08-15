// Code generated by ent, DO NOT EDIT.

package message

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSubscriptionID holds the string denoting the subscription_id field in the database.
	FieldSubscriptionID = "subscription_id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldPostedAt holds the string denoting the posted_at field in the database.
	FieldPostedAt = "posted_at"
	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"
	// EdgeSubscription holds the string denoting the subscription edge name in mutations.
	EdgeSubscription = "subscription"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// MediaTable is the table that holds the media relation/edge.
	MediaTable = "message_media"
	// MediaInverseTable is the table name for the MessageMedia entity.
	// It exists in this package in order to avoid circular dependency with the "messagemedia" package.
	MediaInverseTable = "message_media"
	// MediaColumn is the table column denoting the media relation/edge.
	MediaColumn = "message_id"
	// SubscriptionTable is the table that holds the subscription relation/edge.
	SubscriptionTable = "messages"
	// SubscriptionInverseTable is the table name for the Subscription entity.
	// It exists in this package in order to avoid circular dependency with the "subscription" package.
	SubscriptionInverseTable = "subscriptions"
	// SubscriptionColumn is the table column denoting the subscription relation/edge.
	SubscriptionColumn = "subscription_id"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldSubscriptionID,
	FieldText,
	FieldPostedAt,
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

// OrderOption defines the ordering options for the Message queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySubscriptionID orders the results by the subscription_id field.
func BySubscriptionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionID, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByPostedAt orders the results by the posted_at field.
func ByPostedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostedAt, opts...).ToFunc()
}

// ByMediaCount orders the results by media count.
func ByMediaCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMediaStep(), opts...)
	}
}

// ByMedia orders the results by media terms.
func ByMedia(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediaStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubscriptionField orders the results by subscription field.
func BySubscriptionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionStep(), sql.OrderByField(field, opts...))
	}
}
func newMediaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MediaTable, MediaColumn),
	)
}
func newSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubscriptionTable, SubscriptionColumn),
	)
}
