package utils

import "github.com/google/uuid"

// PtrStr return a string of a string pointer.
func PtrStr(s *string) string {
	return *s
}

// PtrInt return a string of a string pointer.
func PtrInt(i *int) int {
	return *i
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
