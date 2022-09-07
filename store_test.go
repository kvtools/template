package template

import (
	"context"
	"testing"
	"time"

	"github.com/kvtools/valkeyrie"
	"github.com/kvtools/valkeyrie/store"
	"github.com/kvtools/valkeyrie/testsuite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTimeout = 60 * time.Second

func TestRegister(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	kv, err := valkeyrie.NewStore(ctx, StoreName, []string{"TODO"}, nil)
	require.NoError(t, err)
	assert.NotNil(t, kv)

	assert.IsTypef(t, kv, new(Store), "Error registering and initializing the store")
}

func TestStore(t *testing.T) {
	// TODO remove the skip
	t.Skip("the real store implementation is required")

	kv := makeStore(t)
	lockKV := makeStore(t)
	ttlKV := makeStore(t)

	t.Cleanup(func() {
		testsuite.RunCleanup(t, kv)
	})

	testsuite.RunCleanup(t, kv)
	testsuite.RunTestCommon(t, kv)
	testsuite.RunTestAtomic(t, kv)
	testsuite.RunTestWatch(t, kv)
	testsuite.RunTestLock(t, kv)
	testsuite.RunTestLockTTL(t, kv, lockKV)
	testsuite.RunTestTTL(t, kv, ttlKV)
}

func makeStore(t *testing.T) store.Store {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config := &Config{
		Username: "example",
		Password: "example",
	}

	kv, err := New(ctx, []string{"TODO"}, config)
	require.NoErrorf(t, err, "cannot create store")

	return kv
}
