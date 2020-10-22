package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Medicine holds the schema definition for the Medicine entity.
type Medicine struct {
	ent.Schema
}

// Fields of the Medicine.
func (Medicine) Fields() []ent.Field {
	return []ent.Field{

		field.String("Medicine_Name").
			NotEmpty(),
	}
}

// Edges of the Medicine.
func (Medicine) Edges() []ent.Edge {
	return []ent.Edge{}

}
