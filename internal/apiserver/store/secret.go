// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// SecretStore defines the secret storage interface.
type SecretStore interface {
	Create(secret *v1.Secret, opts metav1.CreateOptions) error
	Update(secret *v1.Secret, opts metav1.UpdateOptions) error
	Delete(username, secretID string, opts metav1.DeleteOptions) error
	DeleteCollection(username string, secretIDs []string, opts metav1.DeleteOptions) error
	Get(username, secretID string, opts metav1.GetOptions) (*v1.Secret, error)
	List(username string, opts metav1.ListOptions) (*v1.SecretList, error)
}
