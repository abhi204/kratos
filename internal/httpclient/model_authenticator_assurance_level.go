/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.0-alpha.4.pre.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// AuthenticatorAssuranceLevel The authenticator assurance level can be one of \"aal1\", \"aal2\", or \"aal3\". A higher number means that it is harder for an attacker to compromise the account.  Generally, \"aal1\" implies that one authentication factor was used while AAL2 implies that two factors (e.g. password + TOTP) have been used.  To learn more about these levels please head over to: https://www.ory.sh/kratos/docs/concepts/credentials
type AuthenticatorAssuranceLevel string

// List of authenticatorAssuranceLevel
const (
	AUTHENTICATORASSURANCELEVEL_AAL0 AuthenticatorAssuranceLevel = "aal0"
	AUTHENTICATORASSURANCELEVEL_AAL1 AuthenticatorAssuranceLevel = "aal1"
	AUTHENTICATORASSURANCELEVEL_AAL2 AuthenticatorAssuranceLevel = "aal2"
	AUTHENTICATORASSURANCELEVEL_AAL3 AuthenticatorAssuranceLevel = "aal3"
)

func (v *AuthenticatorAssuranceLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticatorAssuranceLevel(value)
	for _, existing := range []AuthenticatorAssuranceLevel{"aal0", "aal1", "aal2", "aal3"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticatorAssuranceLevel", value)
}

// Ptr returns reference to authenticatorAssuranceLevel value
func (v AuthenticatorAssuranceLevel) Ptr() *AuthenticatorAssuranceLevel {
	return &v
}

type NullableAuthenticatorAssuranceLevel struct {
	value *AuthenticatorAssuranceLevel
	isSet bool
}

func (v NullableAuthenticatorAssuranceLevel) Get() *AuthenticatorAssuranceLevel {
	return v.value
}

func (v *NullableAuthenticatorAssuranceLevel) Set(val *AuthenticatorAssuranceLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorAssuranceLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorAssuranceLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorAssuranceLevel(val *AuthenticatorAssuranceLevel) *NullableAuthenticatorAssuranceLevel {
	return &NullableAuthenticatorAssuranceLevel{value: val, isSet: true}
}

func (v NullableAuthenticatorAssuranceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorAssuranceLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
