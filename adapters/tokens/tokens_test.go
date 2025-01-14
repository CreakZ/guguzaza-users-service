package tokens

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJwtUtil(t *testing.T) {
	key := make([]byte, 30)
	if _, err := rand.Read(key); err != nil {
		t.Error(err)
	}

	util := NewJwtUtil(time.Second*34, key)

	userUuid := uuid.NewString()

	token, err := util.CreateJwt(context.Background(), userUuid)
	if err != nil {
		t.Error(err)
	}

	parsedUuid, err := util.ParseJwtClaims(context.Background(), token)
	if err != nil {
		t.Error(err)
	}

	if userUuid != parsedUuid {
		t.Errorf("created and parsed uuids are not the same:\n")
		t.Logf("initial uuid: %s\n", userUuid)
		t.Logf("parsed uuid: %s\n", parsedUuid)
	}
}
