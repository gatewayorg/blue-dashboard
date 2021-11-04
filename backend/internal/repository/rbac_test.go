package repository

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/pkg/fields"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRbacImpl_Rule(t *testing.T) {
	tx := testRepo.Begin()
	ctx := testRepo.ContextWithTx(context.Background(), tx)
	defer tx.Rollback()

	err := testRbac.AddRule(ctx, &model.Rule{
		Service: "dashboard.rbac.PublicRole",
		Method:  "AddRule",
		Detail:  "add rule",
	})
	assert.NoError(t, err)

	num, err := testRbac.GetRuleCount(ctx)
	assert.NoError(t, err)
	assert.True(t, num > 0)

	rules, err := testRbac.GetRule(ctx, 1, 10)
	assert.NoError(t, err)
	for _, v := range rules {
		t.Log(v)
	}

	err = testRbac.SetRuleDetail(ctx, rules[0].ID, "test1111")
	assert.NoError(t, err)
}

func TestRbacImpl_Role(t *testing.T) {
	tx := testRepo.Begin()
	ctx := testRepo.ContextWithTx(context.Background(), tx)
	defer tx.Rollback()

	id, err := testRbac.AddRole(ctx, &model.Role{
		Name:     "admin",
		Detail:   "superman",
		Enable:   false,
		RuleIDs:  fields.Uint64s{},
		CreateAt: time.Now(),
	})
	t.Log(id)
	assert.NoError(t, err)

	roles, err := testRbac.GetRoleAll(ctx)
	assert.NoError(t, err)
	t.Log(roles)

	roles, err = testRbac.GetRole(ctx, 1, 10)
	assert.NoError(t, err)

	num, err := testRbac.GetRuleCount(ctx)
	assert.NoError(t, err)
	assert.True(t, num > 0)

	r, err := testRbac.FindRoleById(ctx, roles[0].ID)
	assert.NoError(t, err)

	r.Detail = "super"
	r.Name = "aaaa"
	r.Enable = true
	r.RuleIDs = fields.Uint64s{1, 2, 3}
	err = testRbac.SaveRole(ctx, r)
	assert.NoError(t, err)

	err = testRbac.SelectRole(ctx, r.ID, []uint64{1, 2, 3, 4})
	assert.NoError(t, err)

	err = testRbac.SetRoleStatus(ctx, r.ID, false)
	assert.NoError(t, err)
}

func TestRbacImpl_FindRoleById(t *testing.T) {
	role, err := testRbac.FindRoleById(context.Background(), 1111)
	assert.NoError(t, err)
	t.Log(role)
}
