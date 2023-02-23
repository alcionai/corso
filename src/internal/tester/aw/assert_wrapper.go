package aw

import (
	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// this file is subfoldered to avoid cyclic dependencies

// ---------------------------------------------------------------------------
// require wrappers
// ---------------------------------------------------------------------------

// MustErr wraps require.Error()
func MustErr(t require.TestingT, err error, etc ...any) {
	require.Error(t, err, addClues(err, etc))
}

// MustNoErr wraps require.NoError()
func MustNoErr(t require.TestingT, err error, etc ...any) {
	require.NoError(t, err, addClues(err, etc))
}

// MustErrIs wraps require.ErrorIs()
func MustErrIs(t require.TestingT, err, target error, etc ...any) {
	require.ErrorIs(t, err, target, addClues(err, etc))
}

// MustNotErrIs wraps require.NotErrorIs()
func MustNotErrIs(t require.TestingT, err, target error, etc ...any) {
	require.NotErrorIs(t, err, target, addClues(err, etc))
}

// MustErrAs wraps require.ErrorAs()
func MustErrAs(t require.TestingT, err, target error, etc ...any) {
	require.ErrorAs(t, err, target, addClues(err, etc))
}

// ---------------------------------------------------------------------------
// assert wrappers
// ---------------------------------------------------------------------------

// Err wraps assert.Error()
func Err(t assert.TestingT, err error, etc ...any) bool {
	return assert.Error(t, err, addClues(err, etc))
}

// NoErr wraps assert.NoError()
func NoErr(t assert.TestingT, err error, etc ...any) bool {
	return assert.NoError(t, err, addClues(err, etc))
}

// ErrIs wraps assert.ErrorIs()
func ErrIs(t assert.TestingT, err, target error, etc ...any) bool {
	return assert.ErrorIs(t, err, target, addClues(err, etc))
}

// NotErrIs wraps assert.NotErrorIs()
func NotErrIs(t assert.TestingT, err, target error, etc ...any) bool {
	return assert.NotErrorIs(t, err, target, addClues(err, etc))
}

// ErrAs wraps assert.ErrorAs()
func ErrAs(t assert.TestingT, err, target error, etc ...any) bool {
	return assert.ErrorAs(t, err, target, addClues(err, etc))
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func addClues(err error, a []any) []any {
	if err == nil {
		return a
	}

	if a == nil {
		a = make([]any, 0)
	}

	return append(a, clues.InErr(err).Slice()...)
}
