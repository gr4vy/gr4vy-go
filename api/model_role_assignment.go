/*
 * Gr4vy API
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.1.0-beta
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// RoleAssignment struct for RoleAssignment
type RoleAssignment struct {
	// The type of this resource. Always `role-assignment`.
	Type *string `json:"type,omitempty"`
	// The unique ID for this role assignment.
	Id *string `json:"id,omitempty"`
	Role *Role `json:"role,omitempty"`
	Assignee *RoleAssignmentAssignee `json:"assignee,omitempty"`
}

// NewRoleAssignment instantiates a new RoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignment() *RoleAssignment {
	this := RoleAssignment{}
	return &this
}

// NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentWithDefaults() *RoleAssignment {
	this := RoleAssignment{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleAssignment) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleAssignment) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RoleAssignment) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleAssignment) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleAssignment) GetRole() Role {
	if o == nil || o.Role == nil {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetRoleOk() (*Role, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleAssignment) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *RoleAssignment) SetRole(v Role) {
	o.Role = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *RoleAssignment) GetAssignee() RoleAssignmentAssignee {
	if o == nil || o.Assignee == nil {
		var ret RoleAssignmentAssignee
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetAssigneeOk() (*RoleAssignmentAssignee, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *RoleAssignment) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given RoleAssignmentAssignee and assigns it to the Assignee field.
func (o *RoleAssignment) SetAssignee(v RoleAssignmentAssignee) {
	o.Assignee = &v
}

func (o RoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignment struct {
	value *RoleAssignment
	isSet bool
}

func (v NullableRoleAssignment) Get() *RoleAssignment {
	return v.value
}

func (v *NullableRoleAssignment) Set(val *RoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignment(val *RoleAssignment) *NullableRoleAssignment {
	return &NullableRoleAssignment{value: val, isSet: true}
}

func (v NullableRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


