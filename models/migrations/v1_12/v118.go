// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_12 //nolint

import (
	"xorm.io/xorm"
)

func AddReviewCommitAndStale(x *xorm.Engine) error {
	type Review struct {
		CommitID string `xorm:"VARCHAR(40)"`
		Stale    bool   `xorm:"NOT NULL DEFAULT false"`
	}

	type ProtectedBranch struct {
		DismissStaleApprovals bool `xorm:"NOT NULL DEFAULT false"`
	}

	// Old reviews will have commit ID set to "" and not stale
	if err := x.Sync2(new(Review)); err != nil {
		return err
	}
	return x.Sync2(new(ProtectedBranch))
}
