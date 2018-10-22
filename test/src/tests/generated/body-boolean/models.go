package booleangroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "test/src/tests/generated/body-boolean"

// BoolModel ...
type BoolModel struct {
	autorest.Response `json:"-"`
	Value             *bool `json:"value,omitempty"`
}

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
