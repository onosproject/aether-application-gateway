// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package models

// Application defines model for Application.
type Application struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Applications - temporary
var Applications = []Application{
	{ID: "1", Name: "app-1"},
	{ID: "2", Name: "app-2"},
}
