package casbin

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2/util"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authz"
	"net/http"
	"strings"
)

type PermissionService struct {
	enforcer *EnforcerProvider
}

func NewPermissionService(enforcer *EnforcerProvider) *PermissionService {
	return &PermissionService{
		enforcer: enforcer,
	}
}

var _ authz.PermissionManagementService = (*PermissionService)(nil)

func (p *PermissionService) AddGrant(ctx context.Context, resource authz.Resource, action authz.Action, subject authz.Subject, tenantID string, effect authz.Effect) error {
	enforcer, err := p.enforcer.Get(ctx)
	if err != nil {
		return err
	}

	eff, err := mapToEffect(effect)
	if err != nil {
		return err
	}
	tenantID = authz.NormalizeTenantId(ctx, tenantID)
	_, err = enforcer.AddPolicy(subject.GetIdentity(), resource.GetNamespace(), resource.GetIdentity(), action.GetIdentity(), tenantID, eff)
	if err != nil {
		return err
	}
	return nil
}

func (p *PermissionService) ListAcl(ctx context.Context, subjects ...authz.Subject) ([]authz.PermissionBean, error) {
	//list
	enforcer, err := p.enforcer.Get(ctx)
	if err != nil {
		return nil, err
	}
	policies := enforcer.GetPolicy()
	var ret []authz.PermissionBean
	for _, policy := range policies {
		//tenant has been filtered by LoadFilteredPolicy
		if len(subjects) > 0 {
			for _, subject := range subjects {
				if util.KeyMatch(policy[0], subject.GetIdentity()) {
					ret = append(ret, authz.NewPermissionBean(authz.NewEntityResource(policy[1], policy[2]),
						authz.ActionStr(policy[3]),
						authz.SubjectStr(policy[0]),
						policy[4], mapToAuthEffect(policy[5]),
					))
				}
			}
		} else {
			ret = append(ret, authz.NewPermissionBean(authz.NewEntityResource(policy[1], policy[2]),
				authz.ActionStr(policy[3]),
				authz.SubjectStr(policy[0]),
				policy[4], mapToAuthEffect(policy[5]),
			))
		}
	}
	return ret, nil
}

func (p *PermissionService) UpdateGrant(ctx context.Context, subject authz.Subject, acl []authz.UpdateSubjectPermission) error {
	enforcer, err := p.enforcer.Get(ctx)
	if err != nil {
		return err
	}
	_, err = enforcer.RemoveFilteredPolicy(0, subject.GetIdentity())
	if err != nil {
		return err
	}

	rules := make([][]string, len(acl))
	for i, permission := range acl {
		eff, err := mapToEffect(permission.Effect)
		if err != nil {
			return err
		}
		rules[i] = []string{subject.GetIdentity(),
			permission.Resource.GetNamespace(),
			permission.Resource.GetIdentity(),
			permission.Action.GetIdentity(),
			authz.NormalizeTenantId(ctx, permission.TenantID),
			eff}
	}
	_, err = enforcer.AddPolicies(rules)
	if err != nil {
		return err
	}
	return nil
}

func (p *PermissionService) RemoveGrant(ctx context.Context, resource authz.Resource, action authz.Action, subject authz.Subject, tenantID string, effects []authz.Effect) error {
	enforcer, err := p.enforcer.Get(ctx)
	if err != nil {
		return err
	}
	var effectStr []string
	if len(effects) == 0 {
		effects = []authz.Effect{authz.EffectGrant, authz.EffectForbidden, authz.EffectUnknown}
	}
	for _, eff := range effects {
		e, err := mapToEffect(eff)
		if err != nil {
			return err
		}
		effectStr = append(effectStr, e)
	}

	_, err = enforcer.RemoveFilteredPolicy(0, subject.GetIdentity(), resource.GetNamespace(), resource.GetIdentity(), action.GetIdentity(), authz.NormalizeTenantId(ctx, tenantID), strings.Join(effectStr, ","))
	if err != nil {
		return err
	}
	return nil
}

func mapToEffect(effect authz.Effect) (string, error) {
	eff := "allow"
	if effect == authz.EffectGrant {
		eff = "allow"
	} else if effect == authz.EffectForbidden {
		eff = "deny"
	} else {
		return "", errors.New(http.StatusBadRequest, "EFFECT_INVALID", fmt.Sprintf("effect should be one of %s,%s", "grant", "forbidden"))
	}
	return eff, nil
}

func mapToAuthEffect(eff string) authz.Effect {
	if eff == "allow" {
		return authz.EffectGrant
	} else if eff == "deny" {
		return authz.EffectForbidden
	}
	return authz.EffectUnknown
}

var PermissionProviderSet = wire.NewSet(
	NewPermissionService,
	wire.Bind(new(authz.PermissionManagementService), new(*PermissionService)),
	wire.Bind(new(authz.PermissionChecker), new(*PermissionService)),
)
