package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{

		field.String("Doctor_Name").
			NotEmpty(),
		field.String("Doctor_Email").
			Unique(),
	}
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prescriptions", Prescription.Type).StorageKey(edge.Column("Doctor_ID")),
	}
}
