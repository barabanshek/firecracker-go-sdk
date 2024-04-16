// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotLoadParams snapshot load params
// swagger:model SnapshotLoadParams
type SnapshotLoadParams struct {

	// Snapshot type to load.
	SnapshotType string `json:"snapshot_type,omitempty"`

	// Enable support for incremental (diff) snapshots by tracking dirty guest pages.
	EnableDiffSnapshots bool `json:"enable_diff_snapshots,omitempty"`

	// Path to the file that contains the guest memory to be loaded.
	// Required: true
	MemFilePath *string `json:"mem_file_path"`

	// When set to true, the vm is also resumed if the snapshot load is successful.
	ResumeVM bool `json:"resume_vm,omitempty"`

	// Path to the file that contains the microVM state to be loaded.
	// Required: true
	SnapshotPath *string `json:"snapshot_path"`
}

// Validate validates this snapshot load params
func (m *SnapshotLoadParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotLoadParams) validateMemFilePath(formats strfmt.Registry) error {

	if err := validate.Required("mem_file_path", "body", m.MemFilePath); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotLoadParams) validateSnapshotPath(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_path", "body", m.SnapshotPath); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotLoadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotLoadParams) UnmarshalBinary(b []byte) error {
	var res SnapshotLoadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
