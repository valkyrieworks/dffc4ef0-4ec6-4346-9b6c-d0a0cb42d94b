package statereplication

import (
	"context"
	"fmt"
	"strings"
	"time"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/utils/log"
	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/minimal"
	clientorigin "github.com/valkyrieworks/minimal/origin"
	httpclient "github.com/valkyrieworks/minimal/origin/rest"
	clientrpc "github.com/valkyrieworks/minimal/rpc"
	clientdb "github.com/valkyrieworks/minimal/depot/db"
	ctstatus "github.com/valkyrieworks/schema/consensuscore/status"
	httpendpoint "github.com/valkyrieworks/rpc/requester/rest"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

//

//
//
type StateProvider interface {
	//
	AppHash(ctx context.Context, height uint64) ([]byte, error)
	//
	Commit(ctx context.Context, height uint64) (*kinds.Commit, error)
	//
	State(ctx context.Context, height uint64) (sm.State, error)
}

//
type lightClientStateProvider struct {
	ctsync.Mutex //
	lc            *minimal.Client
	version       ctstatus.Version
	initialHeight int64
	providers     map[clientorigin.Provider]string
}

//
func NewLightClientStateProvider(
	ctx context.Context,
	chainID string,
	version ctstatus.Version,
	initialHeight int64,
	servers []string,
	trustOptions minimal.TrustOptions,
	logger log.Logger,
) (StateProvider, error) {
	if len(servers) < 2 {
		return nil, fmt.Errorf("REDACTED", len(servers))
	}

	providers := make([]clientorigin.Provider, 0, len(servers))
	providerRemotes := make(map[clientorigin.Provider]string)
	for _, server := range servers {
		client, err := rpcClient(server)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		provider := httpclient.NewWithClient(chainID, client)
		providers = append(providers, provider)
		//
		//
		providerRemotes[provider] = server
	}

	lc, err := minimal.NewClient(ctx, chainID, trustOptions, providers[0], providers[1:],
		clientdb.New(dbm.NewMemDB(), "REDACTED"), minimal.Logger(logger), minimal.MaxRetryAttempts(5))
	if err != nil {
		return nil, err
	}
	return &lightClientStateProvider{
		lc:            lc,
		version:       version,
		initialHeight: initialHeight,
		providers:     providerRemotes,
	}, nil
}

//
func (s *lightClientStateProvider) AppHash(ctx context.Context, height uint64) ([]byte, error) {
	s.Lock()
	defer s.Unlock()

	//
	header, err := s.lc.VerifyLightBlockAtHeight(ctx, int64(height+1), time.Now())
	if err != nil {
		return nil, err
	}
	//
	//
	//
	//
	//
	//
	//
	//
	_, err = s.lc.VerifyLightBlockAtHeight(ctx, int64(height+2), time.Now())
	if err != nil {
		return nil, err
	}
	return header.AppHash, nil
}

//
func (s *lightClientStateProvider) Commit(ctx context.Context, height uint64) (*kinds.Commit, error) {
	s.Lock()
	defer s.Unlock()
	header, err := s.lc.VerifyLightBlockAtHeight(ctx, int64(height), time.Now())
	if err != nil {
		return nil, err
	}
	return header.Commit, nil
}

//
func (s *lightClientStateProvider) State(ctx context.Context, height uint64) (sm.State, error) {
	s.Lock()
	defer s.Unlock()

	state := sm.State{
		ChainID:       s.lc.ChainID(),
		Version:       s.version,
		InitialHeight: s.initialHeight,
	}
	if state.InitialHeight == 0 {
		state.InitialHeight = 1
	}

	//
	//
	//
	//
	//
	//
	//
	//
	lastLightBlock, err := s.lc.VerifyLightBlockAtHeight(ctx, int64(height), time.Now())
	if err != nil {
		return sm.State{}, err
	}
	currentLightBlock, err := s.lc.VerifyLightBlockAtHeight(ctx, int64(height+1), time.Now())
	if err != nil {
		return sm.State{}, err
	}
	nextLightBlock, err := s.lc.VerifyLightBlockAtHeight(ctx, int64(height+2), time.Now())
	if err != nil {
		return sm.State{}, err
	}

	state.Version = ctstatus.Version{
		Consensus: currentLightBlock.Version,
		Software:  release.TMCoreSemVer,
	}
	state.LastBlockHeight = lastLightBlock.Height
	state.LastBlockTime = lastLightBlock.Time
	state.LastBlockID = lastLightBlock.Commit.BlockID
	state.AppHash = currentLightBlock.AppHash
	state.LastResultsHash = currentLightBlock.LastResultsHash
	state.LastValidators = lastLightBlock.ValidatorSet
	state.Validators = currentLightBlock.ValidatorSet
	state.NextValidators = nextLightBlock.ValidatorSet
	state.LastHeightValidatorsChanged = nextLightBlock.Height

	//
	primaryURL, ok := s.providers[s.lc.Primary()]
	if !ok || primaryURL == "REDACTED" {
		return sm.State{}, fmt.Errorf("REDACTED")
	}
	primaryRPC, err := rpcClient(primaryURL)
	if err != nil {
		return sm.State{}, fmt.Errorf("REDACTED", err)
	}
	rpcclient := clientrpc.NewClient(primaryRPC, s.lc)
	result, err := rpcclient.ConsensusParams(ctx, &currentLightBlock.Height)
	if err != nil {
		return sm.State{}, fmt.Errorf("REDACTED",
			nextLightBlock.Height, err)
	}
	state.ConsensusParams = result.ConsensusParams
	state.LastHeightConsensusParamsChanged = currentLightBlock.Height

	return state, nil
}

//
func rpcClient(server string) (*httpendpoint.HTTP, error) {
	if !strings.Contains(server, "REDACTED") {
		server = "REDACTED" + server
	}
	c, err := httpendpoint.New(server, "REDACTED")
	if err != nil {
		return nil, err
	}
	return c, nil
}
