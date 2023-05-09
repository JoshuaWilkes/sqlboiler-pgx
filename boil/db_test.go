package boil

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetSetDB(t *testing.T) {
	t.Parallel()

	SetDB(&PGXPoolContextExecutor{&pgxpool.Pool{}})

	if GetDB() == nil {
		t.Errorf("Expected GetDB to return a database handle, got nil")
	}
}
