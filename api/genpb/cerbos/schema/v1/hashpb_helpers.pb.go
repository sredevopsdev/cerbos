// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.2.0

package schemav1

import (
	protowire "google.golang.org/protobuf/encoding/protowire"
	hash "hash"
)

func cerbos_schema_v1_Schema_hashpb_sum(m *Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.schema.v1.Schema.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.schema.v1.Schema.definition"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.Definition))

	}
}

func cerbos_schema_v1_ValidationError_hashpb_sum(m *ValidationError, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.schema.v1.ValidationError.path"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Path))

	}
	if _, ok := ignore["cerbos.schema.v1.ValidationError.message"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Message))

	}
	if _, ok := ignore["cerbos.schema.v1.ValidationError.source"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Source)))

	}
}
