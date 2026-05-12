package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"modernc.org/sqlite"
	"zgo.at/zdb/drivers"
)

func init() {
	// Register custom functions for modernc/sqlite GLOBALLLY
	// percent_diff(curr, prev)
	sqlite.MustRegisterDeterministicScalarFunction("percent_diff", 2, func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
		if args[0] == nil || args[1] == nil {
			return nil, nil
		}
		curr := toFloatVal(args[0])
		prev := toFloatVal(args[1])
		if prev == 0 {
			return 0.0, nil
		}
		return ((curr - prev) / prev) * 100.0, nil
	})

	drivers.RegisterDriver(moderncDriver{})
}

func toFloatVal(v any) float64 {
	switch i := v.(type) {
	case int64:
		return float64(i)
	case float64:
		return i
	case int:
		return float64(i)
	case int32:
		return float64(i)
	case float32:
		return float64(i)
	default:
		return 0.0
	}
}

type moderncDriver struct{}

func (d moderncDriver) Name() string    { return "sqlite3" }
func (d moderncDriver) Dialect() string { return "sqlite" }

func (d moderncDriver) Connect(ctx context.Context, connect string, create bool) (*sql.DB, any, error) {
	if create {
		path := connect
		if strings.HasPrefix(path, "file:") {
			path = strings.TrimPrefix(path, "file:")
			if i := strings.IndexAny(path, "?#"); i > -1 {
				path = path[:i]
			}
		}
		dir := filepath.Dir(path)
		if dir != "." && dir != "" {
			_ = os.MkdirAll(dir, 0755)
		}
	}

	db, err := sql.Open("sqlite", connect)
	return db, nil, err
}

func (d moderncDriver) ErrUnique(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}

func (d moderncDriver) StartTest(t testing.TB, opt *drivers.TestOptions) context.Context {
	return context.Background()
}
