// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/ivanvanderbyl/entbug/ent/schema"
	"github.com/ivanvanderbyl/entbug/ent/travelrestrictionsnapshot"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	travelrestrictionsnapshotFields := schema.TravelRestrictionSnapshot{}.Fields()
	_ = travelrestrictionsnapshotFields
	// travelrestrictionsnapshotDescVersion is the schema descriptor for version field.
	travelrestrictionsnapshotDescVersion := travelrestrictionsnapshotFields[0].Descriptor()
	// travelrestrictionsnapshot.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	travelrestrictionsnapshot.VersionValidator = travelrestrictionsnapshotDescVersion.Validators[0].(func(int64) error)
}
