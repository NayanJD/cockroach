// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package opgen

import (
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scop"
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scpb"
	"github.com/cockroachdb/cockroach/pkg/util/protoutil"
)

func init() {
	opRegistry.register((*scpb.SchemaParent)(nil),
		toPublic(
			scpb.Status_ABSENT,
			to(scpb.Status_PUBLIC,
				emit(func(this *scpb.SchemaParent) *scop.AddSchemaParent {
					return &scop.AddSchemaParent{
						Parent: *protoutil.Clone(this).(*scpb.SchemaParent),
					}
				}),
			),
		),
		toAbsent(
			scpb.Status_PUBLIC,
			to(scpb.Status_ABSENT,
				// TODO(postamar): remove revertibility constraint when possible
				revertible(false),
				emit(func(this *scpb.SchemaParent) *scop.RemoveSchemaParent {
					return &scop.RemoveSchemaParent{
						Parent: *protoutil.Clone(this).(*scpb.SchemaParent),
					}
				}),
			),
		),
	)
}
