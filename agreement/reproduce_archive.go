package agreement

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	dbm "github.com/valkyrieworks/-db"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	ctsystem "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	subscriber = "REDACTED"
)

//
//

//
func RunReplayFile(config cfg.BaseConfig, csConfig *cfg.ConsensusConfig, console bool) {
	consensusState := newConsensusStateForReplay(config, csConfig)

	if err := consensusState.ReplayFile(csConfig.WalFile(), console); err != nil {
		ctsystem.Exit(fmt.Sprintf("REDACTED", err))
	}
}

//
func (cs *State) ReplayFile(file string, console bool) error {
	if cs.IsRunning() {
		return errors.New("REDACTED")
	}
	if cs.wal != nil {
		return errors.New("REDACTED")
	}

	cs.startForReplay()

	//

	ctx := context.Background()
	newStepSub, err := cs.eventBus.Subscribe(ctx, subscriber, kinds.EventQueryNewRoundStep)
	if err != nil {
		return fmt.Errorf("REDACTED", subscriber, kinds.EventQueryNewRoundStep)
	}
	defer func() {
		if err := cs.eventBus.Unsubscribe(ctx, subscriber, kinds.EventQueryNewRoundStep); err != nil {
			cs.Logger.Error("REDACTED", "REDACTED", err)
		}
	}()

	//
	fp, err := os.OpenFile(file, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}

	pb := newPlayback(file, fp, cs, cs.state.Copy())
	defer pb.fp.Close()

	var nextN int //
	var msg *TimedWALMessage
	for {
		if nextN == 0 && console {
			nextN = pb.replayConsoleLoop()
		}

		msg, err = pb.dec.Decode()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := pb.cs.readReplayMessage(msg, newStepSub); err != nil {
			return err
		}

		if nextN > 0 {
			nextN--
		}
		pb.count++
	}
}

//
//

type playback struct {
	cs *State

	fp    *os.File
	dec   *WALDecoder
	count int //

	//
	fileName     string   //
	genesisState sm.State //
}

func newPlayback(fileName string, fp *os.File, cs *State, genState sm.State) *playback {
	return &playback{
		cs:           cs,
		fp:           fp,
		fileName:     fileName,
		genesisState: genState,
		dec:          NewWALDecoder(fp),
	}
}

//
func (pb *playback) replayReset(count int, newStepSub kinds.Subscription) error {
	if err := pb.cs.Stop(); err != nil {
		return err
	}
	pb.cs.Wait()

	newCS := NewState(pb.cs.config, pb.genesisState.Copy(), pb.cs.blockExec,
		pb.cs.blockStore, pb.cs.txNotifier, pb.cs.evpool)
	newCS.SetEventBus(pb.cs.eventBus)
	newCS.startForReplay()

	if err := pb.fp.Close(); err != nil {
		return err
	}
	fp, err := os.OpenFile(pb.fileName, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}
	pb.fp = fp
	pb.dec = NewWALDecoder(fp)
	count = pb.count - count
	fmt.Printf("REDACTED", pb.count, count)
	pb.count = 0
	pb.cs = newCS
	var msg *TimedWALMessage
	for i := 0; i < count; i++ {
		msg, err = pb.dec.Decode()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		if err := pb.cs.readReplayMessage(msg, newStepSub); err != nil {
			return err
		}
		pb.count++
	}
	return nil
}

func (cs *State) startForReplay() {
	cs.Logger.Error("REDACTED")
	/*!
s
{
{
{
:
:
n
}
}
*/
}

//
func (pb *playback) replayConsoleLoop() int {
	for {
		fmt.Printf("REDACTED")
		bufReader := bufio.NewReader(os.Stdin)
		line, more, err := bufReader.ReadLine()
		if more {
			ctsystem.Exit("REDACTED")
		} else if err != nil {
			ctsystem.Exit(err.Error())
		}

		tokens := strings.Split(string(line), "REDACTED")
		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {
		case "REDACTED":
			//
			//

			if len(tokens) == 1 {
				return 0
			}
			i, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("REDACTED")
			} else {
				return i
			}

		case "REDACTED":
			//
			//

			//
			//

			ctx := context.Background()
			//

			newStepSub, err := pb.cs.eventBus.Subscribe(ctx, subscriber, kinds.EventQueryNewRoundStep)
			if err != nil {
				ctsystem.Exit(fmt.Sprintf("REDACTED", subscriber, kinds.EventQueryNewRoundStep))
			}
			defer func() {
				if err := pb.cs.eventBus.Unsubscribe(ctx, subscriber, kinds.EventQueryNewRoundStep); err != nil {
					pb.cs.Logger.Error("REDACTED", "REDACTED", err)
				}
			}()

			if len(tokens) == 1 {
				if err := pb.replayReset(1, newStepSub); err != nil {
					pb.cs.Logger.Error("REDACTED", "REDACTED", err)
				}
			} else {
				i, err := strconv.Atoi(tokens[1])
				if err != nil {
					fmt.Println("REDACTED")
				} else if i > pb.count {
					fmt.Printf("REDACTED", pb.count)
				} else if err := pb.replayReset(i, newStepSub); err != nil {
					pb.cs.Logger.Error("REDACTED", "REDACTED", err)
				}
			}

		case "REDACTED":
			//
			//
			//

			rs := pb.cs.RoundState
			if len(tokens) == 1 {
				fmt.Println(rs)
			} else {
				switch tokens[1] {
				case "REDACTED":
					fmt.Printf("REDACTED", rs.Height, rs.Round, rs.Step)
				case "REDACTED":
					fmt.Println(rs.Validators)
				case "REDACTED":
					fmt.Println(rs.Proposal)
				case "REDACTED":
					fmt.Printf("REDACTED", rs.ProposalBlockParts.StringShort(), rs.ProposalBlock.StringShort())
				case "REDACTED":
					fmt.Println(rs.LockedRound)
				case "REDACTED":
					fmt.Printf("REDACTED", rs.LockedBlockParts.StringShort(), rs.LockedBlock.StringShort())
				case "REDACTED":
					fmt.Println(rs.Votes.StringIndented("REDACTED"))

				default:
					fmt.Println("REDACTED", tokens[1])
				}
			}
		case "REDACTED":
			fmt.Println(pb.count)
		}
	}
}

//

//
func newConsensusStateForReplay(config cfg.BaseConfig, csConfig *cfg.ConsensusConfig) *State {
	dbType := dbm.BackendType(config.DBBackend)
	//
	blockStoreDB, err := dbm.NewDB("REDACTED", dbType, config.DBDir())
	if err != nil {
		ctsystem.Exit(err.Error())
	}
	blockStore := depot.NewBlockStore(blockStoreDB)

	//
	stateDB, err := dbm.NewDB("REDACTED", dbType, config.DBDir())
	if err != nil {
		ctsystem.Exit(err.Error())
	}
	stateStore := sm.NewStore(stateDB, sm.StoreOptions{
		DiscardABCIResponses: false,
	})
	gdoc, err := sm.MakeGenesisDocFromFile(config.GenesisFile())
	if err != nil {
		ctsystem.Exit(err.Error())
	}
	state, err := sm.MakeGenesisState(gdoc)
	if err != nil {
		ctsystem.Exit(err.Error())
	}

	//
	clientCreator := gateway.DefaultClientCreator(config.ProxyApp, config.ABCI, config.DBDir())
	proxyApp := gateway.NewAppConns(clientCreator, gateway.NopMetrics())
	err = proxyApp.Start()
	if err != nil {
		ctsystem.Exit(fmt.Sprintf("REDACTED", err))
	}

	eventBus := kinds.NewEventBus()
	if err := eventBus.Start(); err != nil {
		ctsystem.Exit(fmt.Sprintf("REDACTED", err))
	}

	handshaker := NewHandshaker(stateStore, state, blockStore, gdoc)
	handshaker.SetEventBus(eventBus)
	err = handshaker.Handshake(proxyApp)
	if err != nil {
		ctsystem.Exit(fmt.Sprintf("REDACTED", err))
	}

	mempool, evpool := emptyMempool{}, sm.EmptyEvidencePool{}
	blockExec := sm.NewBlockExecutor(stateStore, log.TestingLogger(), proxyApp.Consensus(), mempool, evpool, blockStore)

	consensusState := NewState(csConfig, state.Copy(), blockExec,
		blockStore, mempool, evpool)

	consensusState.SetEventBus(eventBus)
	return consensusState
}
