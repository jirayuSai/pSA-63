// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jirayuSai/app/ent/doctor"
	"github.com/jirayuSai/app/ent/mmedicine"
	"github.com/jirayuSai/app/ent/patient"
	"github.com/jirayuSai/app/ent/predicate"
	"github.com/jirayuSai/app/ent/prescription"
	"github.com/jirayuSai/app/ent/systemmember"
)

// PrescriptionUpdate is the builder for updating Prescription entities.
type PrescriptionUpdate struct {
	config
	hooks      []Hook
	mutation   *PrescriptionMutation
	predicates []predicate.Prescription
}

// Where adds a new predicate for the builder.
func (pu *PrescriptionUpdate) Where(ps ...predicate.Prescription) *PrescriptionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetDatetime sets the Datetime field.
func (pu *PrescriptionUpdate) SetDatetime(t time.Time) *PrescriptionUpdate {
	pu.mutation.SetDatetime(t)
	return pu
}

// SetPatientID sets the patient edge to Patient by id.
func (pu *PrescriptionUpdate) SetPatientID(id int) *PrescriptionUpdate {
	pu.mutation.SetPatientID(id)
	return pu
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillablePatientID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetPatientID(*id)
	}
	return pu
}

// SetPatient sets the patient edge to Patient.
func (pu *PrescriptionUpdate) SetPatient(p *Patient) *PrescriptionUpdate {
	return pu.SetPatientID(p.ID)
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (pu *PrescriptionUpdate) SetDoctorID(id int) *PrescriptionUpdate {
	pu.mutation.SetDoctorID(id)
	return pu
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableDoctorID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetDoctorID(*id)
	}
	return pu
}

// SetDoctor sets the doctor edge to Doctor.
func (pu *PrescriptionUpdate) SetDoctor(d *Doctor) *PrescriptionUpdate {
	return pu.SetDoctorID(d.ID)
}

// SetSystemmemberID sets the systemmember edge to Systemmember by id.
func (pu *PrescriptionUpdate) SetSystemmemberID(id int) *PrescriptionUpdate {
	pu.mutation.SetSystemmemberID(id)
	return pu
}

// SetNillableSystemmemberID sets the systemmember edge to Systemmember by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableSystemmemberID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetSystemmemberID(*id)
	}
	return pu
}

// SetSystemmember sets the systemmember edge to Systemmember.
func (pu *PrescriptionUpdate) SetSystemmember(s *Systemmember) *PrescriptionUpdate {
	return pu.SetSystemmemberID(s.ID)
}

// SetMmedicineID sets the mmedicine edge to Mmedicine by id.
func (pu *PrescriptionUpdate) SetMmedicineID(id int) *PrescriptionUpdate {
	pu.mutation.SetMmedicineID(id)
	return pu
}

// SetNillableMmedicineID sets the mmedicine edge to Mmedicine by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableMmedicineID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetMmedicineID(*id)
	}
	return pu
}

// SetMmedicine sets the mmedicine edge to Mmedicine.
func (pu *PrescriptionUpdate) SetMmedicine(m *Mmedicine) *PrescriptionUpdate {
	return pu.SetMmedicineID(m.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (pu *PrescriptionUpdate) Mutation() *PrescriptionMutation {
	return pu.mutation
}

// ClearPatient clears the patient edge to Patient.
func (pu *PrescriptionUpdate) ClearPatient() *PrescriptionUpdate {
	pu.mutation.ClearPatient()
	return pu
}

// ClearDoctor clears the doctor edge to Doctor.
func (pu *PrescriptionUpdate) ClearDoctor() *PrescriptionUpdate {
	pu.mutation.ClearDoctor()
	return pu
}

// ClearSystemmember clears the systemmember edge to Systemmember.
func (pu *PrescriptionUpdate) ClearSystemmember() *PrescriptionUpdate {
	pu.mutation.ClearSystemmember()
	return pu
}

// ClearMmedicine clears the mmedicine edge to Mmedicine.
func (pu *PrescriptionUpdate) ClearMmedicine() *PrescriptionUpdate {
	pu.mutation.ClearMmedicine()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PrescriptionUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrescriptionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrescriptionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prescription.Table,
			Columns: prescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prescription.FieldDatetime,
		})
	}
	if pu.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SystemmemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.SystemmemberTable,
			Columns: []string{prescription.SystemmemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SystemmemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.SystemmemberTable,
			Columns: []string{prescription.SystemmemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MmedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.MmedicineTable,
			Columns: []string{prescription.MmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mmedicine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MmedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.MmedicineTable,
			Columns: []string{prescription.MmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mmedicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescription.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PrescriptionUpdateOne is the builder for updating a single Prescription entity.
type PrescriptionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PrescriptionMutation
}

// SetDatetime sets the Datetime field.
func (puo *PrescriptionUpdateOne) SetDatetime(t time.Time) *PrescriptionUpdateOne {
	puo.mutation.SetDatetime(t)
	return puo
}

// SetPatientID sets the patient edge to Patient by id.
func (puo *PrescriptionUpdateOne) SetPatientID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetPatientID(id)
	return puo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillablePatientID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetPatientID(*id)
	}
	return puo
}

// SetPatient sets the patient edge to Patient.
func (puo *PrescriptionUpdateOne) SetPatient(p *Patient) *PrescriptionUpdateOne {
	return puo.SetPatientID(p.ID)
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (puo *PrescriptionUpdateOne) SetDoctorID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetDoctorID(id)
	return puo
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableDoctorID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetDoctorID(*id)
	}
	return puo
}

// SetDoctor sets the doctor edge to Doctor.
func (puo *PrescriptionUpdateOne) SetDoctor(d *Doctor) *PrescriptionUpdateOne {
	return puo.SetDoctorID(d.ID)
}

// SetSystemmemberID sets the systemmember edge to Systemmember by id.
func (puo *PrescriptionUpdateOne) SetSystemmemberID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetSystemmemberID(id)
	return puo
}

// SetNillableSystemmemberID sets the systemmember edge to Systemmember by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableSystemmemberID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetSystemmemberID(*id)
	}
	return puo
}

// SetSystemmember sets the systemmember edge to Systemmember.
func (puo *PrescriptionUpdateOne) SetSystemmember(s *Systemmember) *PrescriptionUpdateOne {
	return puo.SetSystemmemberID(s.ID)
}

// SetMmedicineID sets the mmedicine edge to Mmedicine by id.
func (puo *PrescriptionUpdateOne) SetMmedicineID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetMmedicineID(id)
	return puo
}

// SetNillableMmedicineID sets the mmedicine edge to Mmedicine by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableMmedicineID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetMmedicineID(*id)
	}
	return puo
}

// SetMmedicine sets the mmedicine edge to Mmedicine.
func (puo *PrescriptionUpdateOne) SetMmedicine(m *Mmedicine) *PrescriptionUpdateOne {
	return puo.SetMmedicineID(m.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (puo *PrescriptionUpdateOne) Mutation() *PrescriptionMutation {
	return puo.mutation
}

// ClearPatient clears the patient edge to Patient.
func (puo *PrescriptionUpdateOne) ClearPatient() *PrescriptionUpdateOne {
	puo.mutation.ClearPatient()
	return puo
}

// ClearDoctor clears the doctor edge to Doctor.
func (puo *PrescriptionUpdateOne) ClearDoctor() *PrescriptionUpdateOne {
	puo.mutation.ClearDoctor()
	return puo
}

// ClearSystemmember clears the systemmember edge to Systemmember.
func (puo *PrescriptionUpdateOne) ClearSystemmember() *PrescriptionUpdateOne {
	puo.mutation.ClearSystemmember()
	return puo
}

// ClearMmedicine clears the mmedicine edge to Mmedicine.
func (puo *PrescriptionUpdateOne) ClearMmedicine() *PrescriptionUpdateOne {
	puo.mutation.ClearMmedicine()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PrescriptionUpdateOne) Save(ctx context.Context) (*Prescription, error) {

	var (
		err  error
		node *Prescription
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrescriptionUpdateOne) SaveX(ctx context.Context) *Prescription {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PrescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrescriptionUpdateOne) sqlSave(ctx context.Context) (pr *Prescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prescription.Table,
			Columns: prescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Prescription.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prescription.FieldDatetime,
		})
	}
	if puo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SystemmemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.SystemmemberTable,
			Columns: []string{prescription.SystemmemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SystemmemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.SystemmemberTable,
			Columns: []string{prescription.SystemmemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MmedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.MmedicineTable,
			Columns: []string{prescription.MmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mmedicine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MmedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.MmedicineTable,
			Columns: []string{prescription.MmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mmedicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Prescription{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescription.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
