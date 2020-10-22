package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{

		field.String("Patient_Name").
			Unique().
			NotEmpty(),
		field.String("Gender"),
		field.Int("Patient_Phone").
			Positive(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prescriptions", Prescription.Type).StorageKey(edge.Column("Patient_ID")),
	}
}
