// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Arthur1/economical-graphql-server-demo/ent/artist"
	"github.com/Arthur1/economical-graphql-server-demo/ent/song"
)

// ArtistCreate is the builder for creating a Artist entity.
type ArtistCreate struct {
	config
	mutation *ArtistMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *ArtistCreate) SetName(s string) *ArtistCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetID sets the "id" field.
func (ac *ArtistCreate) SetID(i int) *ArtistCreate {
	ac.mutation.SetID(i)
	return ac
}

// AddSongIDs adds the "songs" edge to the Song entity by IDs.
func (ac *ArtistCreate) AddSongIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddSongIDs(ids...)
	return ac
}

// AddSongs adds the "songs" edges to the Song entity.
func (ac *ArtistCreate) AddSongs(s ...*Song) *ArtistCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddSongIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (ac *ArtistCreate) Mutation() *ArtistMutation {
	return ac.mutation
}

// Save creates the Artist in the database.
func (ac *ArtistCreate) Save(ctx context.Context) (*Artist, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtistCreate) SaveX(ctx context.Context) *Artist {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtistCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtistCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtistCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Artist.name"`)}
	}
	return nil
}

func (ac *ArtistCreate) sqlSave(ctx context.Context) (*Artist, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArtistCreate) createSpec() (*Artist, *sqlgraph.CreateSpec) {
	var (
		_node = &Artist{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(artist.Table, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := ac.mutation.SongsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.SongsTable,
			Columns: []string{artist.SongsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(song.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArtistCreateBulk is the builder for creating many Artist entities in bulk.
type ArtistCreateBulk struct {
	config
	builders []*ArtistCreate
}

// Save creates the Artist entities in the database.
func (acb *ArtistCreateBulk) Save(ctx context.Context) ([]*Artist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artist, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtistCreateBulk) SaveX(ctx context.Context) []*Artist {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtistCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtistCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
