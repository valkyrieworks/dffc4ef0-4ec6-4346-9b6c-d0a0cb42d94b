package autofile

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cometos "github.com/valkyrieworks/utils/os"
	engineseed "github.com/valkyrieworks/utils/random"
)

func instantiateVerifyClusterWithFrontVolumeCeiling(t *testing.T, frontVolumeCeiling int64) *Cluster {
	verifyUID := engineseed.Str(12)
	verifyFolder := "REDACTED" + verifyUID
	err := cometos.AssureFolder(verifyFolder, 0o700)
	require.NoError(t, err, "REDACTED")

	frontRoute := verifyFolder + "REDACTED"
	g, err := AccessCluster(frontRoute, ClusterFrontVolumeCeiling(frontVolumeCeiling))
	require.NoError(t, err, "REDACTED")
	require.NotEqual(t, nil, g, "REDACTED")

	return g
}

func obliterateVerifyCluster(t *testing.T, g *Cluster) {
	g.End()

	err := os.RemoveAll(g.Dir)
	require.NoError(t, err, "REDACTED")
}

func affirmClusterDetails(t *testing.T, gDetails ClusterDetails, minimumOrdinal, maximumOrdinal int, sumVolume, frontVolume int64) {
	assert.Equal(t, minimumOrdinal, gDetails.MinimumOrdinal)
	assert.Equal(t, maximumOrdinal, gDetails.MaximumOrdinal)
	assert.Equal(t, sumVolume, gDetails.SumVolume)
	assert.Equal(t, frontVolume, gDetails.FrontVolume)
}

func VerifyInspectFrontVolumeCeiling(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 1000*1000)

	//
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 0, 0, 0)

	//
	for i := 0; i < 999; i++ {
		err := g.RecordRow(engineseed.Str(999))
		require.NoError(t, err, "REDACTED")
	}
	err := g.PurgeAndAlign()
	require.NoError(t, err)
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 0, 999000, 999000)

	//
	g.inspectFrontVolumeCeiling()
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 0, 999000, 999000)

	//
	err = g.RecordRow(engineseed.Str(999))
	require.NoError(t, err, "REDACTED")
	err = g.PurgeAndAlign()
	require.NoError(t, err)

	//
	g.inspectFrontVolumeCeiling()
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 1, 1000000, 0)

	//
	err = g.RecordRow(engineseed.Str(999))
	require.NoError(t, err, "REDACTED")
	err = g.PurgeAndAlign()
	require.NoError(t, err)

	//
	g.inspectFrontVolumeCeiling()
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 1, 1001000, 1000)

	//
	for i := 0; i < 999; i++ {
		err = g.RecordRow(engineseed.Str(999))
		require.NoError(t, err, "REDACTED")
	}
	err = g.PurgeAndAlign()
	require.NoError(t, err)
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 1, 2000000, 1000000)

	//
	g.inspectFrontVolumeCeiling()
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 2, 2000000, 0)

	//
	_, err = g.Front.Record([]byte(engineseed.Str(999) + "REDACTED"))
	require.NoError(t, err, "REDACTED")
	err = g.PurgeAndAlign()
	require.NoError(t, err)
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 2, 2001000, 1000)

	//
	g.inspectFrontVolumeCeiling()
	affirmClusterDetails(t, g.ReaderClusterDetails(), 0, 2, 2001000, 1000)

	//
	obliterateVerifyCluster(t, g)
}

func VerifySpinEntry(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 0)

	//
	//
	origFolder, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		if err := os.Chdir(origFolder); err != nil {
			t.Error(err)
		}
	}()

	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	err = os.Chdir(dir)
	require.NoError(t, err)

	require.True(t, filepath.IsAbs(g.Front.Route))
	require.True(t, filepath.IsAbs(g.Dir))

	//
	err = g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)
	g.SpinEntry()
	err = g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)

	//
	body1, err := os.ReadFile(g.Front.Route + "REDACTED")
	assert.NoError(t, err, "REDACTED")
	if string(body1) != "REDACTED" {
		t.Errorf("REDACTED", string(body1))
	}

	//
	body2, err := os.ReadFile(g.Front.Route)
	assert.NoError(t, err, "REDACTED")
	if string(body2) != "REDACTED" {
		t.Errorf("REDACTED", string(body2))
	}

	//
	entries, err := os.ReadDir("REDACTED")
	require.NoError(t, err)
	assert.Empty(t, entries)

	//
	obliterateVerifyCluster(t, g)
}

func VerifyRecord(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 0)

	inscribed := []byte("REDACTED")
	_, err := g.Record(inscribed)
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)

	reader := make([]byte, len(inscribed))
	gr, err := g.NewScanner(0)
	require.NoError(t, err, "REDACTED")

	_, err = gr.Scan(reader)
	assert.NoError(t, err, "REDACTED")
	assert.Equal(t, inscribed, reader)

	//
	obliterateVerifyCluster(t, g)
}

//
//
func VerifyClusterScannerReader(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 0)

	tutor := []byte("REDACTED")
	_, err := g.Record(tutor)
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)
	g.SpinEntry()
	composite := []byte("REDACTED")
	_, err = g.Record(composite)
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)

	sumInscribedExtent := len(tutor) + len(composite)
	reader := make([]byte, sumInscribedExtent)
	gr, err := g.NewScanner(0)
	require.NoError(t, err, "REDACTED")

	n, err := gr.Scan(reader)
	assert.NoError(t, err, "REDACTED")
	assert.Equal(t, sumInscribedExtent, n, "REDACTED")
	tutorAdditionComposite := tutor
	tutorAdditionComposite = append(tutorAdditionComposite, composite...)
	assert.Equal(t, tutorAdditionComposite, reader)

	//
	obliterateVerifyCluster(t, g)
}

//
//
func VerifyClusterScannerReader2(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 0)

	tutor := []byte("REDACTED")
	_, err := g.Record(tutor)
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)
	g.SpinEntry()
	composite := []byte("REDACTED")
	compositeSection := []byte("REDACTED")
	_, err = g.Record(compositeSection) //
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)

	sumExtent := len(tutor) + len(composite)
	reader := make([]byte, sumExtent)
	gr, err := g.NewScanner(0)
	require.NoError(t, err, "REDACTED")

	//
	n, err := gr.Scan(reader)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, len(tutor)+len(compositeSection), n, "REDACTED")

	//
	n, err = gr.Scan([]byte("REDACTED"))
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 0, n)

	//
	obliterateVerifyCluster(t, g)
}

func VerifyMinimumOrdinal(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 0)

	assert.Zero(t, g.MinimumOrdinal(), "REDACTED")

	//
	obliterateVerifyCluster(t, g)
}

func VerifyMaximumOrdinal(t *testing.T) {
	g := instantiateVerifyClusterWithFrontVolumeCeiling(t, 0)

	assert.Zero(t, g.MaximumOrdinal(), "REDACTED")

	err := g.RecordRow("REDACTED")
	require.NoError(t, err)
	err = g.PurgeAndAlign()
	require.NoError(t, err)
	g.SpinEntry()

	assert.Equal(t, 1, g.MaximumOrdinal(), "REDACTED")

	//
	obliterateVerifyCluster(t, g)
}
