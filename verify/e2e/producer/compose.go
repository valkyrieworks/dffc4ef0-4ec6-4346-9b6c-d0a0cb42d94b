package primary

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/signature381"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

var (
	//
	//
	simnetGroupings = map[string][]any{
		"REDACTED":      {"REDACTED", "REDACTED", "REDACTED"},
		"REDACTED": {0, 1000},
		"REDACTED": {
			map[string]string{},
			map[string]string{"REDACTED": "REDACTED", "REDACTED": "REDACTED", "REDACTED": "REDACTED"},
		},
		"REDACTED": {"REDACTED", "REDACTED"},
	}
	peerReleases = burdenedSelection{
		"REDACTED": 2,
	}

	//
	peerRepositories = consistentSelection{"REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	inetv6          = consistentSelection{false, true}
	//
	peerIfaceSchemes         = consistentSelection{"REDACTED", "REDACTED", "REDACTED", "REDACTED"} //
	peerPrivatevalueSchemes      = consistentSelection{"REDACTED", "REDACTED", "REDACTED"}
	peerLedgerAlignments            = consistentSelection{"REDACTED"} //
	peerLedgerChronizeAggregateStyle = consistentSelection{false, true}
	peerStatusAlignments            = consistentSelection{false, true}
	peerEndurePeriods      = consistentSelection{0, 1, 5}
	peerImagePeriods     = consistentSelection{0, 3}
	peerPreserveLedgers          = consistentSelection{
		0,
		2 * int(e2e.ProofLifespanAltitude),
		4 * int(e2e.ProofLifespanAltitude),
	}
	proof          = consistentSelection{0, 1, 10, 20, 200}
	ifaceLags        = consistentSelection{"REDACTED", "REDACTED", "REDACTED"}
	peerDisruptions = likelihoodGroupSelection{
		"REDACTED": 0.1,
		"REDACTED":      0.1,
		"REDACTED":       0.1,
		"REDACTED":    0.1,
		"REDACTED":    0.3,
	}
	agilePeerDisruptions = likelihoodGroupSelection{
		"REDACTED": 0.3,
	}
	ballotAdditionReviseAltitude = consistentSelection{int64(-1), int64(0), int64(1)} //
	ballotAdditionActivated      = burdenedSelection{true: 3, false: 1}
	ballotAdditionAltitudeDisplacement = consistentSelection{int64(0), int64(10), int64(100)}
	ballotAdditionExtent         = consistentSelection{uint(128), uint(512), uint(2048), uint(8192)} //
	tokenKind                   = consistentSelection{edwards25519.TokenKind, ellipticp256.TokenKind, signature381.TokenKind}
)

type composeSettings struct {
	arbitraryOrigin   *rand.Rand
	emissionPath    string
	variedEdition string
	titan   bool
}

//
func Compose(cfg *composeSettings) ([]e2e.Declaration, error) {
	modernizeEdition := "REDACTED"

	if cfg.variedEdition != "REDACTED" {
		var err error
		peerReleases, modernizeEdition, err = analyzeBurdenedReleases(cfg.variedEdition)
		if err != nil {
			return nil, err
		}
		if _, ok := peerReleases["REDACTED"]; ok {
			peerReleases["REDACTED"] = peerReleases["REDACTED"]
			delete(peerReleases, "REDACTED")
			if modernizeEdition == "REDACTED" {
				modernizeEdition = "REDACTED"
			}
		}
		if _, ok := peerReleases["REDACTED"]; ok {
			newestEdition, err := sourceDepotNewestDeliveryEdition(cfg.emissionPath)
			if err != nil {
				return nil, err
			}
			peerReleases[newestEdition] = peerReleases["REDACTED"]
			delete(peerReleases, "REDACTED")
			if modernizeEdition == "REDACTED" {
				modernizeEdition = newestEdition
			}
		}
	}
	fmt.Println("REDACTED")
	for ver, wt := range peerReleases {
		if ver == "REDACTED" {
			fmt.Printf("REDACTED", wt)
		} else {
			fmt.Printf("REDACTED", ver, wt)
		}
	}

	declarations := make([]e2e.Declaration, 0, len(simnetGroupings))
	for _, opt := range groupings(simnetGroupings) {
		declaration, err := composeSimnet(cfg.arbitraryOrigin, opt, modernizeEdition, cfg.titan)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, declaration)
	}
	return declarations, nil
}

//
func composeSimnet(r *rand.Rand, opt map[string]any, modernizeEdition string, titan bool) (e2e.Declaration, error) {
	declaration := e2e.Declaration{
		IDXPrv6:             inetv6.Select(r).(bool),
		IfaceScheme:     peerIfaceSchemes.Select(r).(string),
		PrimaryAltitude:    int64(opt["REDACTED"].(int)),
		PrimaryStatus:     opt["REDACTED"].(map[string]string),
		Assessors:       &map[string]int64{},
		AssessorRevisions: map[string]map[string]int64{},
		TokenKind:          tokenKind.Select(r).(string),
		Proof:         proof.Select(r).(int),
		Peers:            map[string]*e2e.DeclarationPeer{},
		ModernizeEdition:   modernizeEdition,
		Titan:       titan,
	}

	switch ifaceLags.Select(r).(string) {
	case "REDACTED":
	case "REDACTED":
		declaration.ArrangeNominationDeferral = 100 * time.Millisecond
		declaration.HandleNominationDeferral = 100 * time.Millisecond
		declaration.BallotAdditionDeferral = 20 * time.Millisecond
		declaration.CulminateLedgerDeferral = 200 * time.Millisecond
	case "REDACTED":
		declaration.ArrangeNominationDeferral = 200 * time.Millisecond
		declaration.HandleNominationDeferral = 200 * time.Millisecond
		declaration.InspectTransferDeferral = 20 * time.Millisecond
		declaration.BallotAdditionDeferral = 100 * time.Millisecond
		declaration.CulminateLedgerDeferral = 500 * time.Millisecond
	}
	declaration.BallotAdditionsReviseAltitude = ballotAdditionReviseAltitude.Select(r).(int64)
	if declaration.BallotAdditionsReviseAltitude == 1 {
		declaration.BallotAdditionsReviseAltitude = declaration.PrimaryAltitude + ballotAdditionAltitudeDisplacement.Select(r).(int64)
	}
	if ballotAdditionActivated.Select(r).(bool) {
		foundationAltitude := max(declaration.BallotAdditionsReviseAltitude+1, declaration.PrimaryAltitude)
		declaration.BallotAdditionsActivateAltitude = foundationAltitude + ballotAdditionAltitudeDisplacement.Select(r).(int64)
	}

	declaration.BallotAdditionExtent = ballotAdditionExtent.Select(r).(uint)

	var countOrigins, countAssessors, countComplete, countAgileCustomers int
	switch opt["REDACTED"].(string) {
	case "REDACTED":
		countAssessors = 1
	case "REDACTED":
		countAssessors = 4
	case "REDACTED":
		//
		countOrigins = r.Intn(2)
		countAgileCustomers = r.Intn(3)
		countAssessors = 4 + r.Intn(4)
		countComplete = r.Intn(4)
	default:
		return declaration, fmt.Errorf("REDACTED", opt["REDACTED"])
	}

	//
	for i := 1; i <= countOrigins; i++ {
		declaration.Peers[fmt.Sprintf("REDACTED", i)] = composePeer(
			r, e2e.StyleGerm, 0, false)
	}

	//
	//
	//
	followingInitiateLocated := declaration.PrimaryAltitude + 5
	assembly := countAssessors*2/3 + 1
	for i := 1; i <= countAssessors; i++ {
		initiateLocated := int64(0)
		if i > assembly {
			initiateLocated = followingInitiateLocated
			followingInitiateLocated += 5
		}
		alias := fmt.Sprintf("REDACTED", i)
		declaration.Peers[alias] = composePeer(
			r, e2e.StyleAssessor, initiateLocated, i <= 2)

		if initiateLocated == 0 {
			(*declaration.Assessors)[alias] = int64(30 + r.Intn(71))
		} else {
			declaration.AssessorRevisions[fmt.Sprint(initiateLocated+5)] = map[string]int64{
				alias: int64(30 + r.Intn(71)),
			}
		}
	}

	//
	switch opt["REDACTED"].(string) {
	case "REDACTED":
	case "REDACTED":
		declaration.AssessorRevisions["REDACTED"] = *declaration.Assessors
		declaration.Assessors = &map[string]int64{}
	default:
		return declaration, fmt.Errorf("REDACTED", opt["REDACTED"])
	}

	//
	for i := 1; i <= countComplete; i++ {
		initiateLocated := int64(0)
		if r.Float64() >= 0.5 {
			initiateLocated = followingInitiateLocated
			followingInitiateLocated += 5
		}
		declaration.Peers[fmt.Sprintf("REDACTED", i)] = composePeer(
			r, e2e.StyleComplete, initiateLocated, false)
	}

	//
	//
	//
	var germIdentifiers, nodeIdentifiers, agileSuppliers []string
	for alias, peer := range declaration.Peers {
		if peer.Style == string(e2e.StyleGerm) {
			germIdentifiers = append(germIdentifiers, alias)
		} else {
			//
			//
			if (peer.InitiateLocated == 0 || peer.InitiateLocated == declaration.PrimaryAltitude) && peer.PreserveLedgers == 0 {
				agileSuppliers = append(agileSuppliers, alias)
			}
			nodeIdentifiers = append(nodeIdentifiers, alias)
		}
	}

	for _, alias := range germIdentifiers {
		for _, anotherAlias := range germIdentifiers {
			if alias != anotherAlias {
				declaration.Peers[alias].Origins = append(declaration.Peers[alias].Origins, anotherAlias)
			}
		}
	}

	sort.Slice(nodeIdentifiers, func(i, j int) bool {
		idxAlias, jthAlias := nodeIdentifiers[i], nodeIdentifiers[j]
		switch {
		case declaration.Peers[idxAlias].InitiateLocated < declaration.Peers[jthAlias].InitiateLocated:
			return true
		case declaration.Peers[idxAlias].InitiateLocated > declaration.Peers[jthAlias].InitiateLocated:
			return false
		default:
			return strings.Compare(idxAlias, jthAlias) == -1
		}
	})
	for i, alias := range nodeIdentifiers {
		if len(germIdentifiers) > 0 && (i == 0 || r.Float64() >= 0.5) {
			declaration.Peers[alias].Origins = consistentGroupSelection(germIdentifiers).Select(r)
		} else if i > 0 {
			declaration.Peers[alias].EnduringNodes = consistentGroupSelection(nodeIdentifiers[:i]).Select(r)
		}
	}

	//
	for i := 1; i <= countAgileCustomers; i++ {
		initiateLocated := declaration.PrimaryAltitude + 5
		declaration.Peers[fmt.Sprintf("REDACTED", i)] = composeAgilePeer(
			r, initiateLocated+(5*int64(i)), agileSuppliers,
		)
	}

	return declaration, nil
}

//
//
//
//
func composePeer(
	r *rand.Rand, style e2e.Style, initiateLocated int64, compelRepository bool,
) *e2e.DeclarationPeer {
	peer := e2e.DeclarationPeer{
		Edition:               peerReleases.Select(r).(string),
		Style:                  string(style),
		InitiateLocated:               initiateLocated,
		Repository:              peerRepositories.Select(r).(string),
		PrivatevalueScheme:       peerPrivatevalueSchemes.Select(r).(string),
		LedgerChronizeEdition:      peerLedgerAlignments.Select(r).(string),
		LedgerChronizeAggregateStyle: peerLedgerChronizeAggregateStyle.Select(r).(bool),
		StatusChronize:             peerStatusAlignments.Select(r).(bool) && initiateLocated > 0,
		EndureDuration:       referenceUint64n(uint64(peerEndurePeriods.Select(r).(int))),
		ImageDuration:      uint64(peerImagePeriods.Select(r).(int)),
		PreserveLedgers:          uint64(peerPreserveLedgers.Select(r).(int)),
		Disrupt:               peerDisruptions.Select(r),
	}

	//
	//
	if compelRepository {
		peer.PreserveLedgers = 0
		peer.ImageDuration = 3
	}

	//
	//
	if peer.EndureDuration != nil && *peer.EndureDuration == 0 && peer.PreserveLedgers > 0 {
		if r.Float64() > 0.5 {
			peer.PreserveLedgers = 0
		} else {
			peer.EndureDuration = referenceUint64n(peer.PreserveLedgers)
		}
	}

	//
	//
	if peer.PreserveLedgers > 0 {
		if peer.EndureDuration != nil && peer.PreserveLedgers < *peer.EndureDuration {
			peer.PreserveLedgers = *peer.EndureDuration
		}
		if peer.PreserveLedgers < peer.ImageDuration {
			peer.PreserveLedgers = peer.ImageDuration
		}
	}

	return &peer
}

func composeAgilePeer(r *rand.Rand, initiateLocated int64, suppliers []string) *e2e.DeclarationPeer {
	return &e2e.DeclarationPeer{
		Style:            string(e2e.StyleAgile),
		Edition:         peerReleases.Select(r).(string),
		InitiateLocated:         initiateLocated,
		Repository:        peerRepositories.Select(r).(string),
		EndureDuration: referenceUint64n(0),
		EnduringNodes: suppliers,
		Disrupt:         agilePeerDisruptions.Select(r),
	}
}

func referenceUint64n(i uint64) *uint64 {
	return &i
}

//
//
//
//
//
//
func analyzeBurdenedReleases(s string) (burdenedSelection, string, error) {
	wc := make(burdenedSelection)
	var finalEdition string

	listings := strings.Split(strings.TrimSpace(s), "REDACTED")

	for _, record := range listings {
		fragments := strings.Split(strings.TrimSpace(record), "REDACTED")

		var ver string
		switch len(fragments) {
		case 2:
			//
			ver = strings.TrimSpace(
				strings.Join([]string{"REDACTED", fragments[0]}, "REDACTED"),
			)
		case 3:
			//
			ver = strings.TrimSpace(
				strings.Join([]string{fragments[0], fragments[1]}, "REDACTED"),
			)
		default:
			return nil, "REDACTED", fmt.Errorf(
				"REDACTED",
				record,
			)
		}

		loadTxt := strings.TrimSpace(fragments[len(fragments)-1])
		load, err := strconv.Atoi(loadTxt)
		if err != nil {
			return nil, "REDACTED", fmt.Errorf("REDACTED", loadTxt, err)
		}

		if load < 1 {
			return nil, "REDACTED", errors.New("REDACTED")
		}

		wc[ver] = uint(load)
		finalEdition = ver
	}

	return wc, finalEdition, nil
}

//
//
//
func sourceDepotNewestDeliveryEdition(sourceDepotPath string) (string, error) {
	choices := &git.PlainOpenOptions{
		DetectDotGit: true,
	}
	r, err := git.PlainOpenWithOptions(sourceDepotPath, choices)
	if err != nil {
		return "REDACTED", err
	}
	labels := make([]string, 0)
	labelEntities, err := r.TagObjects()
	if err != nil {
		return "REDACTED", err
	}
	err = labelEntities.ForEach(func(labelEntity *object.Tag) error {
		labels = append(labels, labelEntity.Name)
		return nil
	})
	if err != nil {
		return "REDACTED", err
	}
	return locateNewestDeliveryLabel(edition.TEMPBaseSemaphoreEdtn, labels)
}

func locateNewestDeliveryLabel(foundationEdtn string, labels []string) (string, error) {
	foundationSemaphoreEdtn, err := semver.NewVersion(strings.Split(foundationEdtn, "REDACTED")[0])
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", foundationEdtn, err)
	}
	componentEdtn := fmt.Sprintf("REDACTED", foundationSemaphoreEdtn.Major(), foundationSemaphoreEdtn.Minor())
	//
	//
	componentTxt := "REDACTED" + componentEdtn
	edtnConnection, err := semver.NewConstraint(componentTxt)
	if err != nil {
		return "REDACTED", err
	}
	var newestEdtn *semver.Version
	for _, tag := range labels {
		if !strings.HasPrefix(tag, "REDACTED") {
			continue
		}
		currentEdtn, err := semver.NewVersion(tag)
		//
		if err != nil {
			continue
		}
		//
		if len(currentEdtn.Prerelease()) != 0 {
			continue
		}
		//
		if !edtnConnection.Check(currentEdtn) {
			continue
		}
		if newestEdtn == nil || currentEdtn.GreaterThan(newestEdtn) {
			newestEdtn = currentEdtn
		}
	}
	//
	//
	if newestEdtn == nil {
		return "REDACTED", nil
	}
	//
	//
	vs := newestEdtn.String()
	if !strings.HasPrefix(vs, "REDACTED") {
		return "REDACTED" + vs, nil
	}
	return vs, nil
}
