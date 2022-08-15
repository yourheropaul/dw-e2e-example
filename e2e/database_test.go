package e2e

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ItCanGenerateANewFakeArticle(t *testing.T) {
	a := newArticle()
	j, err := json.Marshal(a)
	require.Nil(t, err)
	t.Logf(string(j))
}
