// Copyright 2024 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package context

import (
	exprctx "github.com/pingcap/tidb/pkg/expression/context"
	infoschema "github.com/pingcap/tidb/pkg/infoschema/context"
	"github.com/pingcap/tidb/pkg/kv"
	tablelock "github.com/pingcap/tidb/pkg/lock/context"
	"github.com/pingcap/tidb/pkg/parser/model"
	"github.com/pingcap/tidb/pkg/sessionctx/variable"
	"github.com/pingcap/tidb/pkg/util"
	contextutil "github.com/pingcap/tidb/pkg/util/context"
)

// PlanContext is the context for building plan.
type PlanContext interface {
	exprctx.BuildContext
	contextutil.ValueStoreContext
	tablelock.TableLockReadContext
	// GetSessionVars gets the session variables.
	GetSessionVars() *variable.SessionVars
	// GetInfoSchema returns the current infoschema
	GetInfoSchema() infoschema.InfoSchemaMetaVersion
	// UpdateColStatsUsage updates the column stats usage.
	UpdateColStatsUsage(predicateColumns []model.TableItemID)
	// GetClient gets a kv.Client.
	GetClient() kv.Client
	// GetMPPClient gets a kv.MPPClient.
	GetMPPClient() kv.MPPClient
	// GetSessionManager gets the session manager.
	GetSessionManager() util.SessionManager
	// Txn returns the current transaction which is created before executing a statement.
	// The returned kv.Transaction is not nil, but it maybe pending or invalid.
	// If the active parameter is true, call this function will wait for the pending txn
	// to become valid.
	Txn(active bool) (kv.Transaction, error)
	// HasDirtyContent checks whether there's dirty update on the given table.
	HasDirtyContent(tid int64) bool
}