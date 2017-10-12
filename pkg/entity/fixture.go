package entity

import (
	"github.com/checkr/flagr/swagger_gen/models"
	"github.com/jinzhu/gorm"
)

// GenFixtureFlag is a fixture
func GenFixtureFlag() Flag {
	f := Flag{
		Model:       gorm.Model{ID: 100},
		Description: "",
		Segments:    []Segment{GenFixtureSegment()},
		Variants: []Variant{
			Variant{
				Model:  gorm.Model{ID: 300},
				FlagID: 100,
				Key:    "control",
				Attachment: map[string]string{
					"value": "123",
				},
			},
			Variant{
				Model:  gorm.Model{ID: 301},
				FlagID: 100,
				Key:    "treatment",
				Attachment: map[string]string{
					"value": "321",
				},
			},
		},
	}
	f.PrepareEvaluation()
	return f
}

// GenFixtureSegment is a fixture
func GenFixtureSegment() Segment {
	s := Segment{
		Model:          gorm.Model{ID: 200},
		FlagID:         100,
		Description:    "",
		Rank:           0,
		RolloutPercent: 100,
		Constraints: []Constraint{
			Constraint{
				Model:     gorm.Model{ID: 500},
				SegmentID: 200,
				Property:  "dl_state",
				Operator:  models.ConstraintOperatorEQ,
				Value:     `"CA"`,
			},
		},
		Distributions: []Distribution{
			Distribution{
				Model:      gorm.Model{ID: 400},
				SegmentID:  200,
				VariantID:  300,
				VariantKey: "control",
				Percent:    50,
			},
			Distribution{
				Model:      gorm.Model{ID: 401},
				SegmentID:  200,
				VariantID:  301,
				VariantKey: "treatment",
				Percent:    50,
			},
		},
	}
	s.PrepareEvaluation()
	return s
}
