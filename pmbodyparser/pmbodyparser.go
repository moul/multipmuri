package pmbodyparser

import (
	"fmt"
	"regexp"

	"moul.io/multipmuri"
)

type Kind string

const (
	Blocks      Kind = "blocks"
	DependsOn   Kind = "depends-on"
	Fixes       Kind = "fixes"
	Closes      Kind = "closes"
	Addresses   Kind = "addresses"
	RelatedWith Kind = "related-with"
	PartOf      Kind = "part-of"
	ParentOf    Kind = "parent-of"
)

type Relationship struct {
	Kind   Kind
	Target multipmuri.Entity
}

func (r Relationship) String() string {
	return fmt.Sprintf("%s %s", r.Kind, r.Target.Canonical())
}

// FIXME: add isDependent / isDepending helpers

type Relationships []Relationship

func ParseString(body string) (Relationships, []error) {
	return RelParseString(multipmuri.NewUnknownEntity(), body)
}

var (
	fixesRegex, _       = regexp.Compile(`(?im)^(fix|fixes)\s*[:= ]\s*([^\s,]+)$`)
	blocksRegex, _      = regexp.Compile(`(?im)^(block|blocks)\s*[:= ]\s*([^\s,]+)$`)
	closesRegex, _      = regexp.Compile(`(?im)^(close|closes)\s*[:= ]\s*([^\s,]+)$`)
	parentOfRegex, _    = regexp.Compile(`(?im)^(parent of|parent)\s*[:= ]\s*([^\s,]+)$`)
	partOfRegex, _      = regexp.Compile(`(?im)^(part of|part)\s*[:= ]\s*([^\s,]+)$`)
	relatedWithRegex, _ = regexp.Compile(`(?im)^(related|related with)\s*[:= ]\s*([^\s,]+)$`)
	addressesRegex, _   = regexp.Compile(`(?im)^(address|addresses)\s*[:= ]\s*([^\s,]+)$`)
	dependsOnRegex, _   = regexp.Compile(`(?im)^(depend|depends|depend on|depends on)\s*[:= ]\s*([^\s,]+)$`)
)

func RelParseString(context multipmuri.Entity, body string) (Relationships, []error) {
	relationships := Relationships{}
	errs := []error{}

	for kind, regex := range map[Kind]*regexp.Regexp{
		Fixes:       fixesRegex,
		Blocks:      blocksRegex,
		Closes:      closesRegex,
		DependsOn:   dependsOnRegex,
		ParentOf:    parentOfRegex,
		PartOf:      partOfRegex,
		RelatedWith: relatedWithRegex,
		Addresses:   addressesRegex,
	} {
		for _, match := range regex.FindAllStringSubmatch(body, -1) {
			decoded, err := context.RelDecodeString(match[len(match)-1])
			if err != nil {
				errs = append(errs, err)
				continue
			}
			relationships = append(
				relationships,
				Relationship{Kind: kind, Target: decoded},
			)
		}
	}

	return relationships, errs
}
