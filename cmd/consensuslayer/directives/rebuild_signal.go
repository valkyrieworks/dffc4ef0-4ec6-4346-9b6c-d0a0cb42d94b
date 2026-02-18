package directives

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/-db"

	ifacetypes "github.com/valkyrieworks/atci/kinds"
	ctconfig "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/loadmeter"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/locator"
	blocklocator "github.com/valkyrieworks/status/locator/record/kv"
	"github.com/valkyrieworks/status/locator/drain/sqldb"
	"github.com/valkyrieworks/status/txlocator"
	"github.com/valkyrieworks/status/txlocator/kv"
	"github.com/valkyrieworks/kinds"
)

const (
	reindexFailed = "REDACTED"
)

var (
	ErrHeightNotAvailable = errors.New("REDACTED")
	ErrInvalidRequest     = errors.New("REDACTED")
)

//
var ReIndexEventCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Long: `
REDACTED,
REDACTED 
REDACTED 
REDACTED 
REDACTEDt
REDACTED.

REDACTEDu
REDACTED.
REDACTED`,
	Example: `
REDACTEDt
REDACTED2
REDACTED0
REDACTED0
REDACTED`,
	Run: func(cmd *cobra.Command, args []string) {
		bs, ss, err := loadStateAndBlockStore(config)
		if err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		state, err := ss.Load()
		if err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		if err := checkValidHeight(bs); err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		bi, ti, err := loadEventSinks(config, status.ChainID)
		if err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		riArgs := eventReIndexArgs{
			startHeight:  startHeight,
			endHeight:    endHeight,
			blockIndexer: bi,
			txIndexer:    ti,
			blockStore:   bs,
			stateStore:   ss,
		}
		if err := eventReIndex(cmd, riArgs); err != nil {
			panic(fmt.Errorf("REDACTED", reindexFailed, err))
		}

		fmt.Println("REDACTED")
	},
}

var (
	startHeight int64
	endHeight   int64
)

func init() {
	ReIndexEventCmd.Flags().Int64Var(&startHeight, "REDACTED", 0, "REDACTED")
	ReIndexEventCmd.Flags().Int64Var(&endHeight, "REDACTED", 0, "REDACTED")
}

func loadEventSinks(cfg *ctconfig.Config, chainID string) (locator.BlockIndexer, txlocator.TxIndexer, error) {
	switch strings.ToLower(cfg.TxIndex.Indexer) {
	case "REDACTED":
		return nil, nil, errors.New("REDACTED")
	case "REDACTED":
		conn := cfg.TxIndex.PsqlConn
		if conn == "REDACTED" {
			return nil, nil, errors.New("REDACTED")
		}
		es, err := sqldb.NewEventSink(conn, chainID)
		if err != nil {
			return nil, nil, err
		}
		return es.BlockIndexer(), es.TxIndexer(), nil
	case "REDACTED":
		store, err := dbm.NewDB("REDACTED", dbm.BackendType(cfg.DBBackend), cfg.DBDir())
		if err != nil {
			return nil, nil, err
		}

		txIndexer := kv.NewTxIndex(store)
		blockIndexer := blocklocator.New(dbm.NewPrefixDB(store, []byte("REDACTED")))
		return blockIndexer, txIndexer, nil
	default:
		return nil, nil, fmt.Errorf("REDACTED", cfg.TxIndex.Indexer)
	}
}

type eventReIndexArgs struct {
	startHeight  int64
	endHeight    int64
	blockIndexer locator.BlockIndexer
	txIndexer    txlocator.TxIndexer
	blockStore   status.BlockStore
	stateStore   status.Store
}

func eventReIndex(cmd *cobra.Command, args eventReIndexArgs) error {
	var bar loadmeter.Bar
	bar.NewOption(args.startHeight-1, args.endHeight)

	fmt.Println("REDACTED")
	defer bar.Finish()
	for height := args.startHeight; height <= args.endHeight; height++ {
		select {
		case <-cmd.Context().Done():
			return fmt.Errorf("REDACTED", height, cmd.Context().Err())
		default:
			block := args.blockStore.LoadBlock(height)
			if block == nil {
				return fmt.Errorf("REDACTED", height)
			}

			resp, err := args.stateStore.LoadFinalizeBlockResponse(height)
			if err != nil {
				return fmt.Errorf("REDACTED", height)
			}

			e := kinds.EventDataNewBlockEvents{
				Height: height,
				Events: resp.Events,
			}

			numTxs := len(resp.TxResults)

			var batch *txlocator.Batch
			if numTxs > 0 {
				batch = txlocator.NewBatch(int64(numTxs))

				for idx, txResult := range resp.TxResults {
					tr := ifacetypes.TxResult{
						Height: height,
						Index:  uint32(idx),
						Tx:     block.Txs[idx],
						Result: *txResult,
					}

					if err = batch.Add(&tr); err != nil {
						return fmt.Errorf("REDACTED", err)
					}
				}

				if err := args.txIndexer.AddBatch(batch); err != nil {
					return fmt.Errorf("REDACTED", height, err)
				}
			}

			if err := args.blockIndexer.Index(e); err != nil {
				return fmt.Errorf("REDACTED", height, err)
			}
		}

		bar.Play(height)
	}

	return nil
}

func checkValidHeight(bs status.BlockStore) error {
	base := bs.Base()

	if startHeight == 0 {
		startHeight = base
		fmt.Printf("REDACTED", base)
	}

	if startHeight < base {
		return fmt.Errorf("REDACTED",
			ErrHeightNotAvailable, startHeight, base)
	}

	height := bs.Height()

	if startHeight > height {
		return fmt.Errorf(
			"REDACTED", ErrHeightNotAvailable, startHeight, height)
	}

	if endHeight == 0 || endHeight > height {
		endHeight = height
		fmt.Printf("REDACTED", height)
	}

	if endHeight < base {
		return fmt.Errorf(
			"REDACTED", ErrHeightNotAvailable, endHeight, base)
	}

	if endHeight < startHeight {
		return fmt.Errorf(
			"REDACTED",
			ErrInvalidRequest, startHeight, endHeight)
	}

	return nil
}
