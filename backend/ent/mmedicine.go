// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jirayuSai/app/ent/mmedicine"
)

// Mmedicine is the model entity for the Mmedicine schema.
type Mmedicine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MmedicineName holds the value of the "Mmedicine_Name" field.
	MmedicineName string `json:"Mmedicine_Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MmedicineQuery when eager-loading is set.
	Edges MmedicineEdges `json:"edges"`
}

// MmedicineEdges holds the relations/edges for other nodes in the graph.
type MmedicineEdges struct {
	// Prescriptions holds the value of the prescriptions edge.
	Prescriptions []*Prescription
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PrescriptionsOrErr returns the Prescriptions value or an error if the edge
// was not loaded in eager-loading.
func (e MmedicineEdges) PrescriptionsOrErr() ([]*Prescription, error) {
	if e.loadedTypes[0] {
		return e.Prescriptions, nil
	}
	return nil, &NotLoadedError{edge: "prescriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mmedicine) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Mmedicine_Name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mmedicine fields.
func (m *Mmedicine) assignValues(values ...interface{}) error {
	if m, n := len(values), len(mmedicine.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Mmedicine_Name", values[0])
	} else if value.Valid {
		m.MmedicineName = value.String
	}
	return nil
}

// QueryPrescriptions queries the prescriptions edge of the Mmedicine.
func (m *Mmedicine) QueryPrescriptions() *PrescriptionQuery {
	return (&MmedicineClient{config: m.config}).QueryPrescriptions(m)
}

// Update returns a builder for updating this Mmedicine.
// Note that, you need to call Mmedicine.Unwrap() before calling this method, if this Mmedicine
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mmedicine) Update() *MmedicineUpdateOne {
	return (&MmedicineClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Mmedicine) Unwrap() *Mmedicine {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mmedicine is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mmedicine) String() string {
	var builder strings.Builder
	builder.WriteString("Mmedicine(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", Mmedicine_Name=")
	builder.WriteString(m.MmedicineName)
	builder.WriteByte(')')
	return builder.String()
}

// Mmedicines is a parsable slice of Mmedicine.
type Mmedicines []*Mmedicine

func (m Mmedicines) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
