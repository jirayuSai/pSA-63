package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Prescription holds the schema definition for the Prescription entity.
type Prescription struct {
	ent.Schema
}

// Fields of the Prescription.
func (Prescription) Fields() []ent.Field {
	return []ent.Field{

		field.Time("Datetime"),
	}
}

// Edges of the Prescription.
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", Patient.Type).
			Ref("prescriptions").
			Unique(),

		edge.From("doctor", Doctor.Type).
			Ref("prescriptions").
			Unique(),

		edge.From("systemmember", Systemmember.Type).
			Ref("prescriptions").
			Unique(),

		edge.From("mmedicine", Mmedicine.Type).
			Ref("prescriptions").
			Unique(),
	}
}
