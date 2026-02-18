package settings_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/settings"
)

func VerifyStandardSettings(t *testing.T) {
	affirm := assert.New(t)

	//
	cfg := settings.StandardSettings()
	assert.NotNil(cfg.P2P)
	assert.NotNil(cfg.Txpool)
	assert.NotNil(cfg.Agreement)

	//
	cfg.AssignOrigin("REDACTED")
	cfg.Origin = "REDACTED"
	cfg.StoreRoute = "REDACTED"
	cfg.Txpool.JournalRoute = "REDACTED"

	assert.Equal("REDACTED", cfg.OriginEntry())
	assert.Equal("REDACTED", cfg.StoreFolder())
	assert.Equal("REDACTED", cfg.Txpool.JournalFolder())
}

func VerifySettingsCertifySimple(t *testing.T) {
	cfg := settings.StandardSettings()
	assert.NoError(t, cfg.CertifySimple())

	//
	cfg.Agreement.DeadlineNominate = -10 * time.Second
	assert.Error(t, cfg.CertifySimple())
	cfg.Agreement.DeadlineNominate = 3 * time.Second

	cfg.Agreement.GenerateEmptyLedgers = false
	cfg.Txpool.Kind = settings.TxpoolKindNoop
	assert.Error(t, cfg.CertifySimple())
}

func VerifyTLSSetup(t *testing.T) {
	affirm := assert.New(t)
	cfg := settings.StandardSettings()
	cfg.AssignOrigin("REDACTED")

	cfg.RPC.TLSTokenEntry = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.TokenEntry())
	cfg.RPC.TLSKeyEntry = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.KeyEntry())

	cfg.RPC.TLSTokenEntry = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.TokenEntry())
	cfg.RPC.TLSKeyEntry = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.KeyEntry())
}

func VerifyRootSettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyRootSettings()
	assert.NoError(t, cfg.CertifySimple())

	//
	cfg.TraceLayout = "REDACTED"
	assert.Error(t, cfg.CertifySimple())
}

func VerifyRPCSettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyRPCSettings()
	assert.NoError(t, cfg.CertifySimple())

	attributesToVerify := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}

	for _, fieldLabel := range attributesToVerify {
		reflect.ValueOf(cfg).Elem().FieldByName(fieldLabel).SetInt(-1)
		assert.Error(t, cfg.CertifySimple())
		reflect.ValueOf(cfg).Elem().FieldByName(fieldLabel).SetInt(0)
	}
}

func VerifyP2PSettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyP2PSettings()
	assert.NoError(t, cfg.CertifySimple())

	attributesToVerify := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}

	for _, fieldLabel := range attributesToVerify {
		reflect.ValueOf(cfg).Elem().FieldByName(fieldLabel).SetInt(-1)
		assert.Error(t, cfg.CertifySimple())
		reflect.ValueOf(cfg).Elem().FieldByName(fieldLabel).SetInt(0)
	}
}

func VerifyTxpoolSettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyTxpoolSettings()
	assert.NoError(t, cfg.CertifySimple())

	attributesToVerify := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}

	for _, fieldLabel := range attributesToVerify {
		reflect.ValueOf(cfg).Elem().FieldByName(fieldLabel).SetInt(-1)
		assert.Error(t, cfg.CertifySimple())
		reflect.ValueOf(cfg).Elem().FieldByName(fieldLabel).SetInt(0)
	}

	reflect.ValueOf(cfg).Elem().FieldByName("REDACTED").SetString("REDACTED")
	assert.Error(t, cfg.CertifySimple())
}

func VerifyStatusAlignSettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyStatusAlignSettings()
	require.NoError(t, cfg.CertifySimple())
}

func VerifyLedgerAlignSettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyLedgerAlignSettings()
	assert.NoError(t, cfg.CertifySimple())

	//
	cfg.Release = "REDACTED"
	assert.Error(t, cfg.CertifySimple())

	cfg.Release = "REDACTED"
	assert.Error(t, cfg.CertifySimple())
}

func Verifyagreementsettings_Verifybasic(t *testing.T) {
	//
	verifyscenarios := map[string]struct {
		adjust    func(*settings.AgreementSettings)
		anticipateErr bool
	}{
		"REDACTED":                       {func(c *settings.AgreementSettings) { c.DeadlineNominate = time.Second }, false},
		"REDACTED":              {func(c *settings.AgreementSettings) { c.DeadlineNominate = -1 }, true},
		"REDACTED":                  {func(c *settings.AgreementSettings) { c.DeadlineNominateVariance = time.Second }, false},
		"REDACTED":         {func(c *settings.AgreementSettings) { c.DeadlineNominateVariance = -1 }, true},
		"REDACTED":                       {func(c *settings.AgreementSettings) { c.DeadlinePreballot = time.Second }, false},
		"REDACTED":              {func(c *settings.AgreementSettings) { c.DeadlinePreballot = -1 }, true},
		"REDACTED":                  {func(c *settings.AgreementSettings) { c.DeadlinePreballotVariance = time.Second }, false},
		"REDACTED":         {func(c *settings.AgreementSettings) { c.DeadlinePreballotVariance = -1 }, true},
		"REDACTED":                     {func(c *settings.AgreementSettings) { c.DeadlinePreendorse = time.Second }, false},
		"REDACTED":            {func(c *settings.AgreementSettings) { c.DeadlinePreendorse = -1 }, true},
		"REDACTED":                {func(c *settings.AgreementSettings) { c.DeadlinePreendorseVariance = time.Second }, false},
		"REDACTED":       {func(c *settings.AgreementSettings) { c.DeadlinePreendorseVariance = -1 }, true},
		"REDACTED":                        {func(c *settings.AgreementSettings) { c.DeadlineEndorse = time.Second }, false},
		"REDACTED":               {func(c *settings.AgreementSettings) { c.DeadlineEndorse = -1 }, true},
		"REDACTED":              {func(c *settings.AgreementSettings) { c.NodeGossipPausePeriod = time.Second }, false},
		"REDACTED":     {func(c *settings.AgreementSettings) { c.NodeGossipPausePeriod = -1 }, true},
		"REDACTED":          {func(c *settings.AgreementSettings) { c.NodeInquireMaj23pausePeriod = time.Second }, false},
		"REDACTED": {func(c *settings.AgreementSettings) { c.NodeInquireMaj23pausePeriod = -1 }, true},
		"REDACTED":       {func(c *settings.AgreementSettings) { c.RepeatAttestInspectLevel = -1 }, true},
	}
	for note, tc := range verifyscenarios {
		//
		t.Run(note, func(t *testing.T) {
			cfg := settings.StandardAgreementSettings()
			tc.adjust(cfg)

			err := cfg.CertifySimple()
			if tc.anticipateErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyTelemetrySettingsCertifySimple(t *testing.T) {
	cfg := settings.VerifyTelemetrySettings()
	assert.NoError(t, cfg.CertifySimple())

	//
	cfg.MaximumAccessLinks = -1
	assert.Error(t, cfg.CertifySimple())
}
