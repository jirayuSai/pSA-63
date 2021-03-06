// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jirayuSai/app/ent/doctor"
	"github.com/jirayuSai/app/ent/mmedicine"
	"github.com/jirayuSai/app/ent/patient"
	"github.com/jirayuSai/app/ent/prescription"
	"github.com/jirayuSai/app/ent/systemmember"
)

// Prescription is the model entity for the Prescription schema.
type Prescription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Datetime holds the value of the "Datetime" field.
	Datetime time.Time `json:"Datetime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PrescriptionQuery when eager-loading is set.
	Edges           PrescriptionEdges `json:"edges"`
	Doctor_ID       *int
	Mmedicine_ID    *int
	Patient_ID      *int
	Systemmember_ID *int
}

// PrescriptionEdges holds the relations/edges for other nodes in the graph.
type PrescriptionEdges struct {
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Doctor holds the value of the doctor edge.
	Doctor *Doctor
	// Systemmember holds the value of the systemmember edge.
	Systemmember *Systemmember
	// Mmedicine holds the value of the mmedicine edge.
	Mmedicine *Mmedicine
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[0] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) DoctorOrErr() (*Doctor, error) {
	if e.loadedTypes[1] {
		if e.Doctor == nil {
			// The edge doctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Doctor, nil
	}
	return nil, &NotLoadedError{edge: "doctor"}
}

// SystemmemberOrErr returns the Systemmember value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) SystemmemberOrErr() (*Systemmember, error) {
	if e.loadedTypes[2] {
		if e.Systemmember == nil {
			// The edge systemmember was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: systemmember.Label}
		}
		return e.Systemmember, nil
	}
	return nil, &NotLoadedError{edge: "systemmember"}
}

// MmedicineOrErr returns the Mmedicine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) MmedicineOrErr() (*Mmedicine, error) {
	if e.loadedTypes[3] {
		if e.Mmedicine == nil {
			// The edge mmedicine was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: mmedicine.Label}
		}
		return e.Mmedicine, nil
	}
	return nil, &NotLoadedError{edge: "mmedicine"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prescription) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // Datetime
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Prescription) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // Doctor_ID
		&sql.NullInt64{}, // Mmedicine_ID
		&sql.NullInt64{}, // Patient_ID
		&sql.NullInt64{}, // Systemmember_ID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prescription fields.
func (pr *Prescription) assignValues(values ...interface{}) error {
	if m, n := len(values), len(prescription.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Datetime", values[0])
	} else if value.Valid {
		pr.Datetime = value.Time
	}
	values = values[1:]
	if len(values) == len(prescription.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Doctor_ID", value)
		} else if value.Valid {
			pr.Doctor_ID = new(int)
			*pr.Doctor_ID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Mmedicine_ID", value)
		} else if value.Valid {
			pr.Mmedicine_ID = new(int)
			*pr.Mmedicine_ID = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Patient_ID", value)
		} else if value.Valid {
			pr.Patient_ID = new(int)
			*pr.Patient_ID = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Systemmember_ID", value)
		} else if value.Valid {
			pr.Systemmember_ID = new(int)
			*pr.Systemmember_ID = int(value.Int64)
		}
	}
	return nil
}

// QueryPatient queries the patient edge of the Prescription.
func (pr *Prescription) QueryPatient() *PatientQuery {
	return (&PrescriptionClient{config: pr.config}).QueryPatient(pr)
}

// QueryDoctor queries the doctor edge of the Prescription.
func (pr *Prescription) QueryDoctor() *DoctorQuery {
	return (&PrescriptionClient{config: pr.config}).QueryDoctor(pr)
}

// QuerySystemmember queries the systemmember edge of the Prescription.
func (pr *Prescription) QuerySystemmember() *SystemmemberQuery {
	return (&PrescriptionClient{config: pr.config}).QuerySystemmember(pr)
}

// QueryMmedicine queries the mmedicine edge of the Prescription.
func (pr *Prescription) QueryMmedicine() *MmedicineQuery {
	return (&PrescriptionClient{config: pr.config}).QueryMmedicine(pr)
}

// Update returns a builder for updating this Prescription.
// Note that, you need to call Prescription.Unwrap() before calling this method, if this Prescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prescription) Update() *PrescriptionUpdateOne {
	return (&PrescriptionClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Prescription) Unwrap() *Prescription {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prescription is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prescription) String() string {
	var builder strings.Builder
	builder.WriteString("Prescription(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", Datetime=")
	builder.WriteString(pr.Datetime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Prescriptions is a parsable slice of Prescription.
type Prescriptions []*Prescription

func (pr Prescriptions) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
