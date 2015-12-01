// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package kv

import (
	. "github.com/pingcap/check"
)

var _ = Suite(&testTxnSuite{})

type testTxnSuite struct {
}

func (s *testTxnSuite) SetUpTest(c *C) {
}

func (s *testTxnSuite) TearDownTest(c *C) {
}

func (s *testTxnSuite) TestBackOff(c *C) {
	mustBackOff(c, 1, 2)
	mustBackOff(c, 2, 4)
	mustBackOff(c, 3, 8)
	mustBackOff(c, 100000, 100)
}

func mustBackOff(c *C, cnt, sleep int) {
	c.Assert(BackOff(cnt), LessEqual, sleep*1000)
}
