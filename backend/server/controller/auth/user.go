package auth

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/env"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
	"github.com/temphia/temphia/backend/server/lib/apiutils"
)

// https://blog.cloudflare.com/opaque-oblivious-passwords/

func (c *Controller) logInUser(tenantId, userOrEmail, password string, ctx *gin.Context) error {

	if tenantId == "" {
		return easyerr.NotFound()
	}

	var user *entities.User
	var err error

	if strings.Contains(userOrEmail, "@") {
		user, err = c.coredb.GetUserByEmail(tenantId, userOrEmail)
	} else {
		user, err = c.coredb.GetUserByID(tenantId, userOrEmail)
	}
	if err != nil {
		return err
	}

	userGrp, err := c.coredb.GetUserGroup(tenantId, user.GroupID)
	if err != nil {
		return err
	}

	lenv := &env.AuthFlow{
		User:            user,
		UserGroup:       userGrp,
		ChallengeType:   "password",
		ChallengeData:   password,
		Stage:           "",
		UserClaim:       nil,
		SiteClaim:       nil,
		FlowSideEffects: env.FlowSideEffects{},
	}

	err = c.fencer.AuthFlowCheck(tenantId, lenv)
	if err != nil {
		return err
	}

	_claim := claim.NewUserLogged(tenantId, user.UserId, user.GroupID)

	// generate token
	token, err := c.signer.Sign(tenantId, _claim)
	if err != nil {
		return err
	}

	resp := &vmodels.LoginResponse{
		Token:    token,
		Message:  "",
		Console:  "",
		Plugs:    nil,
		UserData: vmodels.UserData{},
	}

	apiutils.WriteJSON(ctx, resp, nil)
	return nil
}

func (c *Controller) userChallenge(lenv *env.AuthFlow) error {

	if lenv.ChallengeData == lenv.User.Password {
		return nil
	}

	return errors.New("user/password incorrect")
}

func (c *Controller) ForgotPassword()         {}
func (c *Controller) ForgotPasswordComplete() {}
func (c *Controller) ChangePassword()         {}
