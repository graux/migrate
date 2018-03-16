// +build clickhouse

package main

import (
	_ "github.com/kshvakov/clickhouse"
	_ "github.com/graux/migrate/database/clickhouse"
)
