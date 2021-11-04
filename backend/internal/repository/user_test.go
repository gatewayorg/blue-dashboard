package repository

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserImpl(t *testing.T) {
	tx := testRepo.Begin()
	ctx := testRepo.ContextWithTx(context.Background(), tx)
	defer tx.Rollback()

	err := testUser.Add(ctx, &model.User{
		Username: "member",
		Password: "123123",
		Name:     "member",
		RoleID:   1,
		Enable:   false,
		CreateAt: time.Now(),
	})
	assert.NoError(t, err)

	num, err := testUser.GetCountByRoleId(ctx, 1)
	assert.NoError(t, err)
	assert.True(t, num > 0)

	user, err := testUser.FindByUserName(ctx, "member")
	assert.NoError(t, err)

	user.Name = ""
	user.RoleID = 0
	user.Enable = true
	err = testUser.Save(ctx, user)
	assert.NoError(t, err)

	err = testUser.UpdatePwd(ctx, user.ID, "999999")
	assert.NoError(t, err)

	err = testUser.SetStatus(ctx, user.ID, false)
	assert.NoError(t, err)

	err = testUser.SelectRole(ctx, user.ID, 2)
	assert.NoError(t, err)

	users, err := testUser.GetAll(ctx, 1, 10)
	assert.NoError(t, err)
	for _, v := range users {
		t.Log(v)
	}
}
