//go:build unit

package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeRegistrationEmailSuffixWhitelist(t *testing.T) {
	got, err := NormalizeRegistrationEmailSuffixWhitelist([]string{"example.com", "@EXAMPLE.COM", " @foo.bar "})
	require.NoError(t, err)
	require.Equal(t, []string{"@example.com", "@foo.bar"}, got)
}

func TestNormalizeRegistrationEmailSuffixWhitelist_Invalid(t *testing.T) {
	_, err := NormalizeRegistrationEmailSuffixWhitelist([]string{"@invalid_domain"})
	require.Error(t, err)
}

func TestParseRegistrationEmailSuffixWhitelist(t *testing.T) {
	got := ParseRegistrationEmailSuffixWhitelist(`["example.com","@foo.bar","@invalid_domain"]`)
	require.Equal(t, []string{"@example.com", "@foo.bar"}, got)
}

func TestIsRegistrationEmailSuffixAllowed(t *testing.T) {
	require.True(t, IsRegistrationEmailSuffixAllowed("user@example.com", []string{"@example.com"}))
	require.False(t, IsRegistrationEmailSuffixAllowed("user@sub.example.com", []string{"@example.com"}))
	require.True(t, IsRegistrationEmailSuffixAllowed("user@any.com", []string{}))
}

func TestNormalizeRegistrationEmailSuffixWhitelist_Wildcard(t *testing.T) {
	got, err := NormalizeRegistrationEmailSuffixWhitelist([]string{"@*.edu.cn", "*.Example.com"})
	require.NoError(t, err)
	require.Equal(t, []string{"@*.edu.cn", "@*.example.com"}, got)

	_, err = NormalizeRegistrationEmailSuffixWhitelist([]string{"*.com"})
	require.Error(t, err)

	_, err = NormalizeRegistrationEmailSuffixWhitelist([]string{"foo.*.com"})
	require.Error(t, err)
}

func TestIsRegistrationEmailSuffixAllowed_Wildcard(t *testing.T) {
	wl := []string{"@*.edu.cn"}
	require.True(t, IsRegistrationEmailSuffixAllowed("user@pku.edu.cn", wl))
	require.True(t, IsRegistrationEmailSuffixAllowed("user@mail.pku.edu.cn", wl))
	require.False(t, IsRegistrationEmailSuffixAllowed("user@edu.cn", wl))
	require.False(t, IsRegistrationEmailSuffixAllowed("user@fakeedu.cn", wl))
	require.False(t, IsRegistrationEmailSuffixAllowed("user@other.com", wl))
	require.True(t, IsRegistrationEmailSuffixAllowed("user@edu.cn", []string{"@*.edu.cn", "@edu.cn"}))
}
