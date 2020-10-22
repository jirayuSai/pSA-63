package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Mmedicine holds the schema definition for the Mmedicine entity.
type Mmedicine struct {
	ent.Schema
}

// Fields of the Mmedicine.
func (Mmedicine) Fields() []ent.Field {
	return []ent.Field{
		field.String("Mmedicine_Name").
			Unique().
			NotEmpty(),
	}
}

// Edges of the Mmedicine.
func (Mmedicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prescriptions", Prescription.Type).StorageKey(edge.Column("Mmedicine_ID")),
	}
}
