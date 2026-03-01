package verify

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
)

func RestoreVerifyOrigin(verifyAlias string) *settings.Settings {
	return RestoreVerifyOriginUsingSuccessionUUID(verifyAlias, "REDACTED")
}

func RestoreVerifyOriginUsingSuccessionUUID(verifyAlias string, successionUUID string) *settings.Settings {
	//
	originPath, err := os.MkdirTemp("REDACTED", fmt.Sprintf("REDACTED", successionUUID, verifyAlias))
	if err != nil {
		panic(err)
	}

	settings.AssureOrigin(originPath)

	foundationSettings := settings.FallbackFoundationSettings()
	inaugurationRecordRoute := filepath.Join(originPath, foundationSettings.Inauguration)
	privateTokenRecordRoute := filepath.Join(originPath, foundationSettings.PrivateAssessorToken)
	privateStatusRecordRoute := filepath.Join(originPath, foundationSettings.PrivateAssessorStatus)

	if !strongos.RecordPresent(inaugurationRecordRoute) {
		if successionUUID == "REDACTED" {
			successionUUID = FallbackVerifySuccessionUUID
		}
		verifyInauguration := fmt.Sprintf(verifyInaugurationTextformat, successionUUID)
		strongos.ShouldRecordRecord(inaugurationRecordRoute, []byte(verifyInauguration), 0o644)
	}
	//
	strongos.ShouldRecordRecord(privateTokenRecordRoute, []byte(verifyPrivateAssessorToken), 0o644)
	strongos.ShouldRecordRecord(privateStatusRecordRoute, []byte(verifyPrivateAssessorStatus), 0o644)

	settings := settings.VerifySettings().AssignOrigin(originPath)
	return settings
}

var verifyInaugurationTextformat = `REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED{
REDACTED{
REDACTED,
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED,
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED[
REDACTED"
REDACTED]
REDACTED,
REDACTED{
REDACTED"
REDACTED,
REDACTED}
REDACTED,
REDACTED[
REDACTED{
REDACTED{
REDACTED,
REDACTED"
REDACTED,
REDACTED,
REDACTED"
REDACTED}
REDACTED,
REDACTED"
REDACTED`

var verifyPrivateAssessorToken = `REDACTED{
REDACTED,
REDACTED{
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED,
REDACTED"
REDACTED}
REDACTED`

var verifyPrivateAssessorStatus = `REDACTED{
REDACTED,
REDACTED,
REDACTED0
REDACTED`
