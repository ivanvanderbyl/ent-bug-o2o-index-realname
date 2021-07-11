// Code generated by entc, DO NOT EDIT.

package travelrestrictionsnapshot

const (
	// Label holds the string label denoting the travelrestrictionsnapshot type in the database.
	Label = "travel_restriction_snapshot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"

	// EdgeNext holds the string denoting the next edge name in mutations.
	EdgeNext = "next"
	// EdgePrevious holds the string denoting the previous edge name in mutations.
	EdgePrevious = "previous"

	// Table holds the table name of the travelrestrictionsnapshot in the database.
	Table = "travel_restriction_snapshots"
	// NextTable is the table the holds the next relation/edge.
	NextTable = "travel_restriction_snapshots"
	// NextColumn is the table column denoting the next relation/edge.
	NextColumn = "travel_restriction_snapshot_previous"
	// PreviousTable is the table the holds the previous relation/edge.
	PreviousTable = "travel_restriction_snapshots"
	// PreviousColumn is the table column denoting the previous relation/edge.
	PreviousColumn = "travel_restriction_snapshot_previous"
)

// Columns holds all SQL columns for travelrestrictionsnapshot fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldData,
	FieldTimestamp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the TravelRestrictionSnapshot type.
var ForeignKeys = []string{
	"travel_restriction_snapshot_previous",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// VersionValidator is a validator for the "version" field. It is called by the builders before save.
	VersionValidator func(int64) error
)