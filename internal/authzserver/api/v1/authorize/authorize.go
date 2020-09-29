// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package authorize implements the authorize handlers.
package authorize

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"

	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/authzserver/authorization"
	"github.com/marmotedu/iam/internal/authzserver/store"
	"github.com/marmotedu/iam/internal/pkg/code"
)

// Authorize returns whether a request is allow or deny to access a resource and do some action
// under specified condition.
func Authorize(c *gin.Context) {
	var r ladon.Request
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	auth := authorization.NewAuthorizer(store.NewAuthorization())
	if r.Context == nil {
		r.Context = ladon.Context{}
	}
	r.Context["username"] = c.GetHeader("username")
	rsp := auth.Authorize(&r)

	core.WriteResponse(c, nil, rsp)
}
