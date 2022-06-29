//go:build clickhouse
// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/vippsas/vipps-login-golang-migrate/database/clickhouse"
)
