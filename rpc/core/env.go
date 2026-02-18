package core

import (
	"encoding/base64"
	"fmt"
	"time"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	standardEachScreen = 30
	maximumEachScreen     = 100

	//
	//
	EnrolDeadline = 5 * time.Second

	//
	//
	originSegmentVolume = 16 * 1024 * 1024 //
)

//
//

type Agreement interface {
	FetchStatus() sm.Status
	FetchRatifiers() (int64, []*kinds.Ratifier)
	FetchFinalLevel() int64
	FetchEpochStatusJSON() ([]byte, error)
	FetchEpochStatusBasicJSON() ([]byte, error)
}

type carrier interface {
	Observers() []string
	IsObserving() bool
	MemberDetails() p2p.MemberDetails
}

type nodes interface {
	AppendDurableNodes([]string) error
	AppendAbsoluteNodeIDXDatastore([]string) error
	AppendInternalNodeIDXDatastore([]string) error
	CallNodesAsync([]string) error
	Nodes() p2p.IDXNodeCollection
}

//
type alignHandler interface {
	WaitAlign() bool
}

//
//
//
type Context struct {
	//
	GatewayApplicationInquire   gateway.ApplicationLinkInquire
	GatewayApplicationTxpool gateway.ApplicationLinkTxpool

	//
	StatusDepot       sm.Depot
	LedgerDepot       sm.LedgerDepot
	ProofDepository     sm.ProofDepository
	AgreementStatus   Agreement
	AgreementHandler alignHandler
	TxpoolHandler   alignHandler
	P2PNodes         nodes
	P2PCarrier     carrier

	IsReplicaStyle bool

	//
	PublicKey       vault.PublicKey
	GeneratePaper       *kinds.OriginPaper //
	TransOrdinaler    transordinal.TransOrdinaler
	LedgerOrdinaler ordinaler.LedgerOrdinaler
	EventBus     *kinds.EventBus //
	Txpool      txpool.Txpool

	Tracer log.Tracer

	Settings cfg.RPCSettings

	//
	generateSegments []string
}

//

func certifyScreen(screenPointer *int, eachScreen, sumNumber int) (int, error) {
	if eachScreen < 1 {
		panic(fmt.Sprintf("REDACTED", eachScreen))
	}

	if screenPointer == nil { //
		return 1, nil
	}

	sections := ((sumNumber - 1) / eachScreen) + 1
	if sections == 0 {
		sections = 1 //
	}
	screen := *screenPointer
	if screen <= 0 || screen > sections {
		return 1, fmt.Errorf("REDACTED", sections, screen)
	}

	return screen, nil
}

func (env *Context) certifyEachScreen(eachScreenPointer *int) int {
	if eachScreenPointer == nil { //
		return standardEachScreen
	}

	eachScreen := *eachScreenPointer
	if eachScreen < 1 {
		return standardEachScreen
	} else if eachScreen > maximumEachScreen {
		return maximumEachScreen
	}
	return eachScreen
}

//
//
func (env *Context) InitOriginSegments() error {
	if env.generateSegments != nil {
		return nil
	}

	if env.GeneratePaper == nil {
		return nil
	}

	data, err := cometjson.Serialize(env.GeneratePaper)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i += originSegmentVolume {
		end := i + originSegmentVolume

		if end > len(data) {
			end = len(data)
		}

		env.generateSegments = append(env.generateSegments, base64.StdEncoding.EncodeToString(data[i:end]))
	}

	return nil
}

func certifyOmitNumber(screen, eachScreen int) int {
	omitNumber := (screen - 1) * eachScreen
	if omitNumber < 0 {
		return 0
	}

	return omitNumber
}

//
func (env *Context) fetchLevel(newestLevel int64, levelPointer *int64) (int64, error) {
	if levelPointer != nil {
		level := *levelPointer
		if level <= 0 {
			return 0, fmt.Errorf("REDACTED", level)
		}
		if level > newestLevel {
			return 0, fmt.Errorf("REDACTED",
				level, newestLevel)
		}
		root := env.LedgerDepot.Root()
		if level < root {
			return 0, fmt.Errorf("REDACTED",
				level, root)
		}
		return level, nil
	}
	return newestLevel, nil
}

func (env *Context) newestUnsubmittedLevel() int64 {
	memberIsAligning := env.AgreementHandler.WaitAlign()
	if memberIsAligning {
		return env.LedgerDepot.Level()
	}
	return env.LedgerDepot.Level() + 1
}
