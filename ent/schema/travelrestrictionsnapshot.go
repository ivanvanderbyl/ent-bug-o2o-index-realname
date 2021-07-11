package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TravelRestrictionSnapshot holds the schema definition for the TravelRestrictionSnapshot entity.
type TravelRestrictionSnapshot struct {
	ent.Schema
}

// Fields of the TravelRestrictionSnapshot.
func (TravelRestrictionSnapshot) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("version").Immutable().Positive(),
		field.JSON("data", map[string]ent.Value{}).Immutable(),
		field.Time("timestamp").Immutable(),
	}
}

// Edges of the TravelRestrictionSnapshot.
func (TravelRestrictionSnapshot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("previous", TravelRestrictionSnapshot.Type).
			Unique().
			From("next").
			Unique(),
	}
}
