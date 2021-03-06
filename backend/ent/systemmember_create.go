// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jirayuSai/app/ent/prescription"
	"github.com/jirayuSai/app/ent/systemmember"
)

// SystemmemberCreate is the builder for creating a Systemmember entity.
type SystemmemberCreate struct {
	config
	mutation *SystemmemberMutation
	hooks    []Hook
}

// SetSystemmemberName sets the Systemmember_Name field.
func (sc *SystemmemberCreate) SetSystemmemberName(s string) *SystemmemberCreate {
	sc.mutation.SetSystemmemberName(s)
	return sc
}

// SetPassword sets the Password field.
func (sc *SystemmemberCreate) SetPassword(s string) *SystemmemberCreate {
	sc.mutation.SetPassword(s)
	return sc
}

// AddPrescriptionIDs adds the prescriptions edge to Prescription by ids.
func (sc *SystemmemberCreate) AddPrescriptionIDs(ids ...int) *SystemmemberCreate {
	sc.mutation.AddPrescriptionIDs(ids...)
	return sc
}

// AddPrescriptions adds the prescriptions edges to Prescription.
func (sc *SystemmemberCreate) AddPrescriptions(p ...*Prescription) *SystemmemberCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPrescriptionIDs(ids...)
}

// Mutation returns the SystemmemberMutation object of the builder.
func (sc *SystemmemberCreate) Mutation() *SystemmemberMutation {
	return sc.mutation
}

// Save creates the Systemmember in the database.
func (sc *SystemmemberCreate) Save(ctx context.Context) (*Systemmember, error) {
	if _, ok := sc.mutation.SystemmemberName(); !ok {
		return nil, &ValidationError{Name: "Systemmember_Name", err: errors.New("ent: missing required field \"Systemmember_Name\"")}
	}
	if v, ok := sc.mutation.SystemmemberName(); ok {
		if err := systemmember.SystemmemberNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Systemmember_Name", err: fmt.Errorf("ent: validator failed for field \"Systemmember_Name\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "Password", err: errors.New("ent: missing required field \"Password\"")}
	}
	if v, ok := sc.mutation.Password(); ok {
		if err := systemmember.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "Password", err: fmt.Errorf("ent: validator failed for field \"Password\": %w", err)}
		}
	}
	var (
		err  error
		node *Systemmember
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemmemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SystemmemberCreate) SaveX(ctx context.Context) *Systemmember {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SystemmemberCreate) sqlSave(ctx context.Context) (*Systemmember, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *SystemmemberCreate) createSpec() (*Systemmember, *sqlgraph.CreateSpec) {
	var (
		s     = &Systemmember{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: systemmember.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: systemmember.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.SystemmemberName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemmember.FieldSystemmemberName,
		})
		s.SystemmemberName = value
	}
	if value, ok := sc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemmember.FieldPassword,
		})
		s.Password = value
	}
	if nodes := sc.mutation.PrescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   systemmember.PrescriptionsTable,
			Columns: []string{systemmember.PrescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
