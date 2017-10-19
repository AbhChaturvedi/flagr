//go:generate goqueryset -in flag.go

package entity

import (
	"bytes"
	"encoding/gob"

	"github.com/jinzhu/gorm"
)

// Flag is the unit of flags
// gen:qs
type Flag struct {
	gorm.Model
	Description string `sql:"type:text"`
	CreatedBy   string
	UpdatedBy   string
	Enabled     bool
	Segments    []Segment
	Variants    []Variant

	// Purely for evaluation
	FlagEvaluation FlagEvaluation `gorm:"-"`
}

// FlagEvaluation is a struct that holds the necessary info for evaluation
type FlagEvaluation struct {
	VariantsMap map[uint]*Variant
}

// Preload preloads the segments and variants into flags
func (f *Flag) Preload(db *gorm.DB) error {
	ss := []Segment{}
	segmentQuery := NewSegmentQuerySet(db)
	if err := segmentQuery.FlagIDEq(f.ID).OrderAscByRank().All(&ss); err != nil {
		return err
	}
	for i, s := range ss {
		if err := s.Preload(db); err != nil {
			return err
		}
		ss[i] = s
	}
	f.Segments = ss

	vs := []Variant{}
	variantQuery := NewVariantQuerySet(db)
	err := variantQuery.FlagIDEq(f.ID).OrderAscByID().All(&vs)
	if err != nil {
		return err
	}
	f.Variants = vs
	return nil
}

// PrepareEvaluation prepares the information for evaluation
func (f *Flag) PrepareEvaluation() error {
	f.FlagEvaluation = FlagEvaluation{
		VariantsMap: make(map[uint]*Variant),
	}
	for i := range f.Segments {
		if err := f.Segments[i].PrepareEvaluation(); err != nil {
			return err
		}
	}
	for i := range f.Variants {
		f.FlagEvaluation.VariantsMap[f.Variants[i].ID] = &f.Variants[i]
	}
	return nil
}

// Encode serialize the flag
func (f *Flag) Encode() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(f)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Decode de-serialize the flag
func (f *Flag) Decode(b []byte) error {
	var bb bytes.Buffer
	bb.Write(b)
	dec := gob.NewDecoder(&bb)
	err := dec.Decode(f)
	if err != nil {
		return err
	}
	return nil
}
