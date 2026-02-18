package directives

import (
	"context"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	ifacetypes "github.com/valkyrieworks/atci/kinds"
	ctconfig "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/validation"
	recordfakes "github.com/valkyrieworks/status/locator/simulations"
	"github.com/valkyrieworks/status/simulations"
	xactfakes "github.com/valkyrieworks/status/txlocator/simulations"
	"github.com/valkyrieworks/kinds"
)

const (
	height int64 = 10
	base   int64 = 2
)

func setupReIndexEventCmd() *cobra.Command {
	reIndexEventCmd := &cobra.Command{
		Use: ReIndexEventCmd.Use,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	_ = reIndexEventCmd.ExecuteContext(context.Background())

	return reIndexEventCmd
}

func TestReIndexEventCheckHeight(t *testing.T) {
	mockBlockStore := &simulations.BlockStore{}
	mockBlockStore.
		On("REDACTED").Return(base).
		On("REDACTED").Return(height)

	testCases := []struct {
		startHeight int64
		endHeight   int64
		validHeight bool
	}{
		{0, 0, true},
		{0, base, true},
		{0, base - 1, false},
		{0, height, true},
		{0, height + 1, true},
		{0, 0, true},
		{base - 1, 0, false},
		{base, 0, true},
		{base, base, true},
		{base, base - 1, false},
		{base, height, true},
		{base, height + 1, true},
		{height, 0, true},
		{height, base, false},
		{height, height - 1, false},
		{height, height, true},
		{height, height + 1, true},
		{height + 1, 0, false},
	}

	for _, tc := range testCases {
		startHeight = tc.startHeight
		endHeight = tc.endHeight

		err := checkValidHeight(mockBlockStore)
		if tc.validHeight {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func TestLoadEventSink(t *testing.T) {
	testCases := []struct {
		sinks   string
		connURL string
		loadErr bool
	}{
		{"REDACTED", "REDACTED", true},
		{"REDACTED", "REDACTED", true},
		{"REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", true}, //
		//
		{"REDACTED", "REDACTED", true},
	}

	for idx, tc := range testCases {
		cfg := ctconfig.TestConfig()
		cfg.TxIndex.Indexer = tc.sinks
		cfg.TxIndex.PsqlConn = tc.connURL
		_, _, err := loadEventSinks(cfg, validation.DefaultTestChainID)
		if tc.loadErr {
			require.Error(t, err, idx)
		} else {
			require.NoError(t, err, idx)
		}
	}
}

func TestLoadBlockStore(t *testing.T) {
	cfg := ctconfig.TestConfig()
	cfg.DBPath = t.TempDir()
	_, _, err := loadStateAndBlockStore(cfg)
	require.Error(t, err)

	_, err = dbm.NewDB("REDACTED", dbm.GoLevelDBBackend, cfg.DBDir())
	require.NoError(t, err)

	//
	_, err = dbm.NewDB("REDACTED", dbm.GoLevelDBBackend, cfg.DBDir())
	require.NoError(t, err)

	bs, ss, err := loadStateAndBlockStore(cfg)
	require.NoError(t, err)
	require.NotNil(t, bs)
	require.NotNil(t, ss)
}

func TestReIndexEvent(t *testing.T) {
	mockBlockStore := &simulations.BlockStore{}
	mockStateStore := &simulations.Store{}
	mockBlockIndexer := &recordfakes.BlockIndexer{}
	mockTxIndexer := &xactfakes.TxIndexer{}

	mockBlockStore.
		On("REDACTED").Return(base).
		On("REDACTED").Return(height).
		On("REDACTED", base).Return(nil).Once().
		On("REDACTED", base).Return(&kinds.Block{Data: kinds.Data{Txs: kinds.Txs{make(kinds.Tx, 1)}}}).
		On("REDACTED", height).Return(&kinds.Block{Data: kinds.Data{Txs: kinds.Txs{make(kinds.Tx, 1)}}})

	abciResp := &ifacetypes.ResponseFinalizeBlock{
		TxResults: []*ifacetypes.ExecTxResult{
			{Code: 1},
		},
	}

	mockBlockIndexer.
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(errors.New("REDACTED")).Once().
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)

	mockTxIndexer.
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(errors.New("REDACTED")).Once().
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)

	mockStateStore.
		On("REDACTED", base).Return(nil, errors.New("REDACTED")).Once().
		On("REDACTED", base).Return(abciResp, nil).
		On("REDACTED", height).Return(abciResp, nil)

	testCases := []struct {
		startHeight int64
		endHeight   int64
		reIndexErr  bool
	}{
		{base, height, true}, //
		{base, height, true}, //
		{base, height, true}, //
		{base, height, true}, //
		{base, base, false},
		{height, height, false},
	}

	for _, tc := range testCases {
		args := eventReIndexArgs{
			startHeight:  tc.startHeight,
			endHeight:    tc.endHeight,
			blockIndexer: mockBlockIndexer,
			txIndexer:    mockTxIndexer,
			blockStore:   mockBlockStore,
			stateStore:   mockStateStore,
		}

		err := eventReIndex(setupReIndexEventCmd(), args)
		if tc.reIndexErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
