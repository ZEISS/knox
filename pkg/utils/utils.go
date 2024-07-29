package utils

import "github.com/google/uuid"

// PtrStr return a string of a string pointer.
func PtrStr(s *string) string {
	return *s
}

// PtrInt return a string of a string pointer.
func PtrInt(i *int, defaults ...int) int {
	if i != nil {
		return *i
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return 0
}

// IntPtr return a string pointer of a string.
func IntPtr(i int) *int {
	return &i
}

// PtrUUID return a string of a string pointer.
func PtrUUID(u *uuid.UUID) uuid.UUID {
	return *u
}

// UUIDPtr return a string pointer of a string.
func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

// StrPtr return a string pointer of a string.
func StrPtr(s string) *string {
	return &s
}
