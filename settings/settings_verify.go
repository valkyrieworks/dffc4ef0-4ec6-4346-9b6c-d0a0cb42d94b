package param_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
)

func VerifyFallbackSettings(t *testing.T) {
	affirm := assert.New(t)

	//
	cfg := settings.FallbackSettings()
	assert.NotNil(cfg.P2P)
	assert.NotNil(cfg.Txpool)
	assert.NotNil(cfg.Agreement)

	//
	cfg.AssignOrigin("REDACTED")
	cfg.Inauguration = "REDACTED"
	cfg.DatastoreRoute = "REDACTED"
	cfg.Txpool.JournalRoute = "REDACTED"

	assert.Equal("REDACTED", cfg.InaugurationRecord())
	assert.Equal("REDACTED", cfg.DatastorePath())
	assert.Equal("REDACTED", cfg.Txpool.JournalPath())
}

func VerifySettingsCertifyFundamental(t *testing.T) {
	cfg := settings.FallbackSettings()
	assert.NoError(t, cfg.CertifyFundamental())

	//
	cfg.Agreement.DeadlineNominate = -10 * time.Second
	assert.Error(t, cfg.CertifyFundamental())
	cfg.Agreement.DeadlineNominate = 3 * time.Second

	cfg.Agreement.GenerateVoidLedgers = false
	cfg.Txpool.Kind = settings.TxpoolKindNooperation
	assert.Error(t, cfg.CertifyFundamental())
}

func VerifyTransportsecSetup(t *testing.T) {
	affirm := assert.New(t)
	cfg := settings.FallbackSettings()
	cfg.AssignOrigin("REDACTED")

	cfg.RPC.TransportsecLicenseRecord = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.LicenseRecord())
	cfg.RPC.TransportsecTokenRecord = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.TokenRecord())

	cfg.RPC.TransportsecLicenseRecord = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.LicenseRecord())
	cfg.RPC.TransportsecTokenRecord = "REDACTED"
	assert.Equal("REDACTED", cfg.RPC.TokenRecord())
}

func VerifyFoundationSettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyFoundationSettings()
	assert.NoError(t, cfg.CertifyFundamental())

	//
	cfg.RecordLayout = "REDACTED"
	assert.Error(t, cfg.CertifyFundamental())
}

func VerifyRemoteSettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyRemoteSettings()
	assert.NoError(t, cfg.CertifyFundamental())

	areasTowardVerify := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}

	for _, attributeAlias := range areasTowardVerify {
		reflect.ValueOf(cfg).Elem().FieldByName(attributeAlias).SetInt(-1)
		assert.Error(t, cfg.CertifyFundamental())
		reflect.ValueOf(cfg).Elem().FieldByName(attributeAlias).SetInt(0)
	}
}

func VerifyPeer2peerSettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyPeer2peerSettings()
	assert.NoError(t, cfg.CertifyFundamental())

	areasTowardVerify := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}

	for _, attributeAlias := range areasTowardVerify {
		reflect.ValueOf(cfg).Elem().FieldByName(attributeAlias).SetInt(-1)
		assert.Error(t, cfg.CertifyFundamental())
		reflect.ValueOf(cfg).Elem().FieldByName(attributeAlias).SetInt(0)
	}
}

func VerifyTxpoolSettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyTxpoolSettings()
	assert.NoError(t, cfg.CertifyFundamental())

	areasTowardVerify := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}

	for _, attributeAlias := range areasTowardVerify {
		reflect.ValueOf(cfg).Elem().FieldByName(attributeAlias).SetInt(-1)
		assert.Error(t, cfg.CertifyFundamental())
		reflect.ValueOf(cfg).Elem().FieldByName(attributeAlias).SetInt(0)
	}

	reflect.ValueOf(cfg).Elem().FieldByName("REDACTED").SetString("REDACTED")
	assert.Error(t, cfg.CertifyFundamental())
}

func VerifyStatusChronizeSettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyStatusChronizeSettings()
	require.NoError(t, cfg.CertifyFundamental())
}

func VerifyLedgerChronizeSettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyLedgerChronizeSettings()
	assert.NoError(t, cfg.CertifyFundamental())

	//
	cfg.Edition = "REDACTED"
	assert.Error(t, cfg.CertifyFundamental())

	cfg.Edition = "REDACTED"
	assert.Error(t, cfg.CertifyFundamental())
}

func Mockagreementsettings_Certifyfundamental(t *testing.T) {
	//
	verifycases := map[string]struct {
		alter    func(*settings.AgreementSettings)
		anticipateFault bool
	}{
		"REDACTED":                       {func(c *settings.AgreementSettings) { c.DeadlineNominate = time.Second }, false},
		"REDACTED":              {func(c *settings.AgreementSettings) { c.DeadlineNominate = -1 }, true},
		"REDACTED":                  {func(c *settings.AgreementSettings) { c.DeadlineNominateVariation = time.Second }, false},
		"REDACTED":         {func(c *settings.AgreementSettings) { c.DeadlineNominateVariation = -1 }, true},
		"REDACTED":                       {func(c *settings.AgreementSettings) { c.DeadlinePreballot = time.Second }, false},
		"REDACTED":              {func(c *settings.AgreementSettings) { c.DeadlinePreballot = -1 }, true},
		"REDACTED":                  {func(c *settings.AgreementSettings) { c.DeadlinePreballotVariation = time.Second }, false},
		"REDACTED":         {func(c *settings.AgreementSettings) { c.DeadlinePreballotVariation = -1 }, true},
		"REDACTED":                     {func(c *settings.AgreementSettings) { c.DeadlinePreendorse = time.Second }, false},
		"REDACTED":            {func(c *settings.AgreementSettings) { c.DeadlinePreendorse = -1 }, true},
		"REDACTED":                {func(c *settings.AgreementSettings) { c.DeadlinePreendorseVariation = time.Second }, false},
		"REDACTED":       {func(c *settings.AgreementSettings) { c.DeadlinePreendorseVariation = -1 }, true},
		"REDACTED":                        {func(c *settings.AgreementSettings) { c.DeadlineEndorse = time.Second }, false},
		"REDACTED":               {func(c *settings.AgreementSettings) { c.DeadlineEndorse = -1 }, true},
		"REDACTED":              {func(c *settings.AgreementSettings) { c.NodeMulticastSnoozeInterval = time.Second }, false},
		"REDACTED":     {func(c *settings.AgreementSettings) { c.NodeMulticastSnoozeInterval = -1 }, true},
		"REDACTED":          {func(c *settings.AgreementSettings) { c.NodeInquireMajor23dormantInterval = time.Second }, false},
		"REDACTED": {func(c *settings.AgreementSettings) { c.NodeInquireMajor23dormantInterval = -1 }, true},
		"REDACTED":       {func(c *settings.AgreementSettings) { c.DuplicateAttestInspectAltitude = -1 }, true},
	}
	for description, tc := range verifycases {
		//
		t.Run(description, func(t *testing.T) {
			cfg := settings.FallbackAgreementSettings()
			tc.alter(cfg)

			err := cfg.CertifyFundamental()
			if tc.anticipateFault {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyTelemetrySettingsCertifyFundamental(t *testing.T) {
	cfg := settings.VerifyTelemetrySettings()
	assert.NoError(t, cfg.CertifyFundamental())

	//
	cfg.MaximumInitiateLinks = -1
	assert.Error(t, cfg.CertifyFundamental())
}
