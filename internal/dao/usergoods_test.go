package dao

import (
	"context"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.Background()
	result, err := GetUserGoodsDao().PageUserGoodsByUserId(ctx, 1, 10, "1534069483103465474")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
