package base

import (
	"encoding/base64"
	"fmt"
	"time"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	fallbackEveryScreen = 30
	maximumEveryScreen     = 100

	//
	//
	ListenDeadline = 5 * time.Second

	//
	//
	inaugurationSegmentExtent = 16 * 1024 * 1024 //
)

//
//

type Agreement interface {
	ObtainStatus() sm.Status
	ObtainAssessors() (int64, []*kinds.Assessor)
	ObtainFinalAltitude() int64
	ObtainIterationStatusJSN() ([]byte, error)
	ObtainIterationStatusPlainJSN() ([]byte, error)
}

type carrier interface {
	Observers() []string
	EqualsObserving() bool
	PeerDetails() p2p.PeerDetails
}

type nodes interface {
	AppendEnduringNodes([]string) error
	AppendAbsoluteNodeIDXDstore([]string) error
	AppendSecludedNodeIDXDstore([]string) error
	CallNodesAsyncronous([]string) error
	Nodes() p2p.IDXNodeAssign
}

//
type chronizeHandler interface {
	AwaitChronize() bool
}

//
//
//
type Context struct {
	//
	DelegateApplicationInquire   delegate.PlatformLinkInquire
	DelegateApplicationTxpool delegate.ApplicationLinkTxpool

	//
	StatusDepot       sm.Depot
	LedgerDepot       sm.LedgerDepot
	ProofHub     sm.ProofHub
	AgreementStatus   Agreement
	AgreementHandler chronizeHandler
	TxpoolHandler   chronizeHandler
	Peer2peerNodes         nodes
	Peer2peerCarrier     carrier

	//
	//
	EqualsAggregateStyle bool

	//
	PublicToken       security.PublicToken
	ProducePaper       *kinds.OriginPaper //
	TransferOrdinalizer    transferordinal.TransferOrdinalizer
	LedgerOrdinalizer ordinalizer.LedgerOrdinalizer
	IncidentChannel     *kinds.IncidentChannel //
	Txpool      txpooll.Txpool

	Tracer log.Tracer

	Settings cfg.RemoteSettings

	//
	produceSegments []string
}

//

func certifyScreen(screenReference *int, everyScreen, sumTally int) (int, error) {
	if everyScreen < 1 {
		panic(fmt.Sprintf("REDACTED", everyScreen))
	}

	if screenReference == nil { //
		return 1, nil
	}

	displays := ((sumTally - 1) / everyScreen) + 1
	if displays == 0 {
		displays = 1 //
	}
	screen := *screenReference
	if screen <= 0 || screen > displays {
		return 1, fmt.Errorf("REDACTED", displays, screen)
	}

	return screen, nil
}

func (env *Context) certifyEveryScreen(everyScreenReference *int) int {
	if everyScreenReference == nil { //
		return fallbackEveryScreen
	}

	everyScreen := *everyScreenReference
	if everyScreen < 1 {
		return fallbackEveryScreen
	} else if everyScreen > maximumEveryScreen {
		return maximumEveryScreen
	}
	return everyScreen
}

//
//
func (env *Context) InitializeInaugurationSegments() error {
	if env.produceSegments != nil {
		return nil
	}

	if env.ProducePaper == nil {
		return nil
	}

	data, err := strongmindjson.Serialize(env.ProducePaper)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i += inaugurationSegmentExtent {
		end := i + inaugurationSegmentExtent

		if end > len(data) {
			end = len(data)
		}

		env.produceSegments = append(env.produceSegments, base64.StdEncoding.EncodeToString(data[i:end]))
	}

	return nil
}

func certifyOmitTally(screen, everyScreen int) int {
	omitTally := (screen - 1) * everyScreen
	if omitTally < 0 {
		return 0
	}

	return omitTally
}

//
func (env *Context) obtainAltitude(newestAltitude int64, altitudeReference *int64) (int64, error) {
	if altitudeReference != nil {
		altitude := *altitudeReference
		if altitude <= 0 {
			return 0, fmt.Errorf("REDACTED", altitude)
		}
		if altitude > newestAltitude {
			return 0, fmt.Errorf("REDACTED",
				altitude, newestAltitude)
		}
		foundation := env.LedgerDepot.Foundation()
		if altitude < foundation {
			return 0, fmt.Errorf("REDACTED",
				altitude, foundation)
		}
		return altitude, nil
	}
	return newestAltitude, nil
}

func (env *Context) newestPendingAltitude() int64 {
	peerEqualsChronizing := env.AgreementHandler.AwaitChronize()
	if peerEqualsChronizing {
		return env.LedgerDepot.Altitude()
	}
	return env.LedgerDepot.Altitude() + 1
}
