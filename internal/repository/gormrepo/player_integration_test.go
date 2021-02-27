// +build integration

package gormrepo_test

import (
	"testing"

	"github.com/remnv/go-boiler/internal/repository/gormrepo"
	"github.com/remnv/go-boiler/internal/storage"
	"github.com/remnv/go-boiler/internal/test"
	"github.com/stretchr/testify/require"
)

func TestPlayerGorm_Add(t *testing.T) {
	t.Run("ShouldAddPlayer", func(t *testing.T) {
		// INIT
		tx := storage.MySqlDbConn(&dbName)
		defer cleanDB(t, tx)

		pl := test.FakePlayer(t, nil)

		playerRepo := gormrepo.NewPlayer(tx)

		// CODE UNDER TEST
		addedPl, err := playerRepo.Add(pl)
		require.NoError(t, err)

		// EXPECTATION
		require.NotNil(t, addedPl)
		require.NotNil(t, addedPl.ID)
		require.Equal(t, pl.UserId, addedPl.UserId)
		require.Equal(t, pl.Name, addedPl.Name)
		require.Equal(t, pl.Level, addedPl.Level)
		require.Equal(t, pl.Job, addedPl.Job)
	})
}
