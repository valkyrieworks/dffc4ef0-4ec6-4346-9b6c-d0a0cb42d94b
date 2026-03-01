package autosave

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

func generateVerifyCohortUsingHeaderExtentThreshold(t *testing.T, headerExtentThreshold int64) *Cluster {
	verifyUUID := commitrand.Str(12)
	verifyPath := "REDACTED" + verifyUUID
	err := strongos.AssurePath(verifyPath, 0o700)
	require.NoError(t, err, "REDACTED")

	headerRoute := verifyPath + "REDACTED"
	g, err := InitiateCluster(headerRoute, ClusterLeadingExtentThreshold(headerExtentThreshold))
	require.NoError(t, err, "REDACTED")
	require.NotEqual(t, nil, g, "REDACTED")

	return g
}

func obliterateVerifyCohort(t *testing.T, g *Cluster) {
	g.Shutdown()

	err := os.RemoveAll(g.Dir)
	require.NoError(t, err, "REDACTED")
}

func affirmCohortDetails(t *testing.T, gDetails CohortDetails, minimumPosition, maximumPosition int, sumExtent, headerExtent int64) {
	assert.Equal(t, minimumPosition, gDetails.MinimumOrdinal)
	assert.Equal(t, maximumPosition, gDetails.MaximumOrdinal)
	assert.Equal(t, sumExtent, gDetails.SumExtent)
	assert.Equal(t, headerExtent, gDetails.HeaderExtent)
}

func VerifyInspectHeaderExtentThreshold(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 1000*1000)

	//
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 0, 0, 0)

	//
	for i := 0; i < 999; i++ {
		err := g.PersistRow(commitrand.Str(999))
		require.NoError(t, err, "REDACTED")
	}
	err := g.PurgeAlsoChronize()
	require.NoError(t, err)
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 0, 999000, 999000)

	//
	g.inspectHeaderExtentThreshold()
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 0, 999000, 999000)

	//
	err = g.PersistRow(commitrand.Str(999))
	require.NoError(t, err, "REDACTED")
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)

	//
	g.inspectHeaderExtentThreshold()
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 1, 1000000, 0)

	//
	err = g.PersistRow(commitrand.Str(999))
	require.NoError(t, err, "REDACTED")
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)

	//
	g.inspectHeaderExtentThreshold()
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 1, 1001000, 1000)

	//
	for i := 0; i < 999; i++ {
		err = g.PersistRow(commitrand.Str(999))
		require.NoError(t, err, "REDACTED")
	}
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 1, 2000000, 1000000)

	//
	g.inspectHeaderExtentThreshold()
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 2, 2000000, 0)

	//
	_, err = g.Leading.Record([]byte(commitrand.Str(999) + "REDACTED"))
	require.NoError(t, err, "REDACTED")
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 2, 2001000, 1000)

	//
	g.inspectHeaderExtentThreshold()
	affirmCohortDetails(t, g.FetchCohortDetails(), 0, 2, 2001000, 1000)

	//
	obliterateVerifyCohort(t, g)
}

func VerifyPivotRecord(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 0)

	//
	//
	sourcePath, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		if err := os.Chdir(sourcePath); err != nil {
			t.Error(err)
		}
	}()

	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	err = os.Chdir(dir)
	require.NoError(t, err)

	require.True(t, filepath.IsAbs(g.Leading.Route))
	require.True(t, filepath.IsAbs(g.Dir))

	//
	err = g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)
	g.PivotRecord()
	err = g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)

	//
	section1, err := os.ReadFile(g.Leading.Route + "REDACTED")
	assert.NoError(t, err, "REDACTED")
	if string(section1) != "REDACTED" {
		t.Errorf("REDACTED", string(section1))
	}

	//
	section2, err := os.ReadFile(g.Leading.Route)
	assert.NoError(t, err, "REDACTED")
	if string(section2) != "REDACTED" {
		t.Errorf("REDACTED", string(section2))
	}

	//
	records, err := os.ReadDir("REDACTED")
	require.NoError(t, err)
	assert.Empty(t, records)

	//
	obliterateVerifyCohort(t, g)
}

func VerifyPersist(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 0)

	recorded := []byte("REDACTED")
	_, err := g.Record(recorded)
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)

	fetch := make([]byte, len(recorded))
	gr, err := g.FreshFetcher(0)
	require.NoError(t, err, "REDACTED")

	_, err = gr.Obtain(fetch)
	assert.NoError(t, err, "REDACTED")
	assert.Equal(t, recorded, fetch)

	//
	obliterateVerifyCohort(t, g)
}

//
//
func VerifyCohortFetcherFetch(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 0)

	instructor := []byte("REDACTED")
	_, err := g.Record(instructor)
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)
	g.PivotRecord()
	reconstructed := []byte("REDACTED")
	_, err = g.Record(reconstructed)
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)

	sumRecordedMagnitude := len(instructor) + len(reconstructed)
	fetch := make([]byte, sumRecordedMagnitude)
	gr, err := g.FreshFetcher(0)
	require.NoError(t, err, "REDACTED")

	n, err := gr.Obtain(fetch)
	assert.NoError(t, err, "REDACTED")
	assert.Equal(t, sumRecordedMagnitude, n, "REDACTED")
	instructorAdditionReconstructed := instructor
	instructorAdditionReconstructed = append(instructorAdditionReconstructed, reconstructed...)
	assert.Equal(t, instructorAdditionReconstructed, fetch)

	//
	obliterateVerifyCohort(t, g)
}

//
//
func VerifyCohortFetcherFetch2(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 0)

	instructor := []byte("REDACTED")
	_, err := g.Record(instructor)
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)
	g.PivotRecord()
	reconstructed := []byte("REDACTED")
	reconstructedFragment := []byte("REDACTED")
	_, err = g.Record(reconstructedFragment) //
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)

	sumMagnitude := len(instructor) + len(reconstructed)
	fetch := make([]byte, sumMagnitude)
	gr, err := g.FreshFetcher(0)
	require.NoError(t, err, "REDACTED")

	//
	n, err := gr.Obtain(fetch)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, len(instructor)+len(reconstructedFragment), n, "REDACTED")

	//
	n, err = gr.Obtain([]byte("REDACTED"))
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 0, n)

	//
	obliterateVerifyCohort(t, g)
}

func VerifyMinimumPosition(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 0)

	assert.Zero(t, g.MinimumOrdinal(), "REDACTED")

	//
	obliterateVerifyCohort(t, g)
}

func VerifyMaximumPosition(t *testing.T) {
	g := generateVerifyCohortUsingHeaderExtentThreshold(t, 0)

	assert.Zero(t, g.MaximumOrdinal(), "REDACTED")

	err := g.PersistRow("REDACTED")
	require.NoError(t, err)
	err = g.PurgeAlsoChronize()
	require.NoError(t, err)
	g.PivotRecord()

	assert.Equal(t, 1, g.MaximumOrdinal(), "REDACTED")

	//
	obliterateVerifyCohort(t, g)
}
