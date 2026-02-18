package main

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

	"github.com/valkyrieworks/vault/bls12381"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/release"
)

var (
	//
	//
	verifychainGroupings = map[string][]any{
		"REDACTED":      {"REDACTED", "REDACTED", "REDACTED"},
		"REDACTED": {0, 1000},
		"REDACTED": {
			map[string]string{},
			map[string]string{"REDACTED": "REDACTED", "REDACTED": "REDACTED", "REDACTED": "REDACTED"},
		},
		"REDACTED": {"REDACTED", "REDACTED"},
	}
	memberReleases = scaledOption{
		"REDACTED": 2,
	}

	//
	memberDatastores = consistentOption{"REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	ipv6          = consistentOption{false, true}
	//
	memberIfaceProtocols     = consistentOption{"REDACTED", "REDACTED", "REDACTED", "REDACTED"} //
	memberPrivatekeyProtocols  = consistentOption{"REDACTED", "REDACTED", "REDACTED"}
	memberLedgerAlignments        = consistentOption{"REDACTED"} //
	memberStatusAlignments        = consistentOption{false, true}
	memberEndurePeriods  = consistentOption{0, 1, 5}
	memberMirrorPeriods = consistentOption{0, 3}
	memberPreserveLedgers      = consistentOption{
		0,
		2 * int(e2e.ProofEraLevel),
		4 * int(e2e.ProofEraLevel),
	}
	proof          = consistentOption{0, 1, 10, 20, 200}
	ifaceLags        = consistentOption{"REDACTED", "REDACTED", "REDACTED"}
	memberVariations = likelihoodCollectionOption{
		"REDACTED": 0.1,
		"REDACTED":      0.1,
		"REDACTED":       0.1,
		"REDACTED":    0.1,
		"REDACTED":    0.3,
	}
	rapidMemberVariations = likelihoodCollectionOption{
		"REDACTED": 0.3,
	}
	ballotAdditionModifyLevel = consistentOption{int64(-1), int64(0), int64(1)} //
	ballotAdditionActivated      = scaledOption{true: 3, false: 1}
	ballotAdditionLevelDisplacement = consistentOption{int64(0), int64(10), int64(100)}
	ballotAdditionVolume         = consistentOption{uint(128), uint(512), uint(2048), uint(8192)} //
	keyKind                   = consistentOption{ed25519.KeyKind, secp256k1.KeyKind, bls12381.KeyKind}
)

type composeSettings struct {
	randomOrigin   *rand.Rand
	resultFolder    string
	multipleRelease string
	monitorstats   bool
}

//
func Compose(cfg *composeSettings) ([]e2e.Declaration, error) {
	enhanceRelease := "REDACTED"

	if cfg.multipleRelease != "REDACTED" {
		var err error
		memberReleases, enhanceRelease, err = analyzeScaledReleases(cfg.multipleRelease)
		if err != nil {
			return nil, err
		}
		if _, ok := memberReleases["REDACTED"]; ok {
			memberReleases["REDACTED"] = memberReleases["REDACTED"]
			delete(memberReleases, "REDACTED")
			if enhanceRelease == "REDACTED" {
				enhanceRelease = "REDACTED"
			}
		}
		if _, ok := memberReleases["REDACTED"]; ok {
			newestRelease, err := gitArchiveNewestDeliveryRelease(cfg.resultFolder)
			if err != nil {
				return nil, err
			}
			memberReleases[newestRelease] = memberReleases["REDACTED"]
			delete(memberReleases, "REDACTED")
			if enhanceRelease == "REDACTED" {
				enhanceRelease = newestRelease
			}
		}
	}
	fmt.Println("REDACTED")
	for ver, wt := range memberReleases {
		if ver == "REDACTED" {
			fmt.Printf("REDACTED", wt)
		} else {
			fmt.Printf("REDACTED", ver, wt)
		}
	}

	declarations := make([]e2e.Declaration, 0, len(verifychainGroupings))
	for _, opt := range groupings(verifychainGroupings) {
		declaration, err := composeVerifychain(cfg.randomOrigin, opt, enhanceRelease, cfg.monitorstats)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, declaration)
	}
	return declarations, nil
}

//
func composeVerifychain(r *rand.Rand, opt map[string]any, enhanceRelease string, monitorstats bool) (e2e.Declaration, error) {
	declaration := e2e.Declaration{
		IDXIpv6:             ipv6.Select(r).(bool),
		IfaceProtocol:     memberIfaceProtocols.Select(r).(string),
		PrimaryLevel:    int64(opt["REDACTED"].(int)),
		PrimaryStatus:     opt["REDACTED"].(map[string]string),
		Ratifiers:       &map[string]int64{},
		RatifierRefreshes: map[string]map[string]int64{},
		KeyKind:          keyKind.Select(r).(string),
		Proof:         proof.Select(r).(int),
		Instances:            map[string]*e2e.DeclarationMember{},
		EnhanceRelease:   enhanceRelease,
		Monitorstats:       monitorstats,
	}

	switch ifaceLags.Select(r).(string) {
	case "REDACTED":
	case "REDACTED":
		declaration.ArrangeNominationDeferral = 100 * time.Millisecond
		declaration.HandleNominationDeferral = 100 * time.Millisecond
		declaration.BallotAdditionDeferral = 20 * time.Millisecond
		declaration.CompleteLedgerDeferral = 200 * time.Millisecond
	case "REDACTED":
		declaration.ArrangeNominationDeferral = 200 * time.Millisecond
		declaration.HandleNominationDeferral = 200 * time.Millisecond
		declaration.InspectTransferDeferral = 20 * time.Millisecond
		declaration.BallotAdditionDeferral = 100 * time.Millisecond
		declaration.CompleteLedgerDeferral = 500 * time.Millisecond
	}
	declaration.BallotPluginsModifyLevel = ballotAdditionModifyLevel.Select(r).(int64)
	if declaration.BallotPluginsModifyLevel == 1 {
		declaration.BallotPluginsModifyLevel = declaration.PrimaryLevel + ballotAdditionLevelDisplacement.Select(r).(int64)
	}
	if ballotAdditionActivated.Select(r).(bool) {
		rootLevel := max(declaration.BallotPluginsModifyLevel+1, declaration.PrimaryLevel)
		declaration.BallotPluginsActivateLevel = rootLevel + ballotAdditionLevelDisplacement.Select(r).(int64)
	}

	declaration.BallotAdditionVolume = ballotAdditionVolume.Select(r).(uint)

	var countOrigins, countRatifiers, countEntireties, countRapidAgents int
	switch opt["REDACTED"].(string) {
	case "REDACTED":
		countRatifiers = 1
	case "REDACTED":
		countRatifiers = 4
	case "REDACTED":
		//
		countOrigins = r.Intn(2)
		countRapidAgents = r.Intn(3)
		countRatifiers = 4 + r.Intn(4)
		countEntireties = r.Intn(4)
	default:
		return declaration, fmt.Errorf("REDACTED", opt["REDACTED"])
	}

	//
	for i := 1; i <= countOrigins; i++ {
		declaration.Instances[fmt.Sprintf("REDACTED", i)] = composeMember(
			r, e2e.StyleOrigin, 0, false)
	}

	//
	//
	//
	followingBeginAt := declaration.PrimaryLevel + 5
	assembly := countRatifiers*2/3 + 1
	for i := 1; i <= countRatifiers; i++ {
		beginAt := int64(0)
		if i > assembly {
			beginAt = followingBeginAt
			followingBeginAt += 5
		}
		label := fmt.Sprintf("REDACTED", i)
		declaration.Instances[label] = composeMember(
			r, e2e.StyleRatifier, beginAt, i <= 2)

		if beginAt == 0 {
			(*declaration.Ratifiers)[label] = int64(30 + r.Intn(71))
		} else {
			declaration.RatifierRefreshes[fmt.Sprint(beginAt+5)] = map[string]int64{
				label: int64(30 + r.Intn(71)),
			}
		}
	}

	//
	switch opt["REDACTED"].(string) {
	case "REDACTED":
	case "REDACTED":
		declaration.RatifierRefreshes["REDACTED"] = *declaration.Ratifiers
		declaration.Ratifiers = &map[string]int64{}
	default:
		return declaration, fmt.Errorf("REDACTED", opt["REDACTED"])
	}

	//
	for i := 1; i <= countEntireties; i++ {
		beginAt := int64(0)
		if r.Float64() >= 0.5 {
			beginAt = followingBeginAt
			followingBeginAt += 5
		}
		declaration.Instances[fmt.Sprintf("REDACTED", i)] = composeMember(
			r, e2e.StyleComplete, beginAt, false)
	}

	//
	//
	//
	var originLabels, nodeLabels, rapidSources []string
	for label, member := range declaration.Instances {
		if member.Style == string(e2e.StyleOrigin) {
			originLabels = append(originLabels, label)
		} else {
			//
			//
			if (member.BeginAt == 0 || member.BeginAt == declaration.PrimaryLevel) && member.PreserveLedgers == 0 {
				rapidSources = append(rapidSources, label)
			}
			nodeLabels = append(nodeLabels, label)
		}
	}

	for _, label := range originLabels {
		for _, anotherLabel := range originLabels {
			if label != anotherLabel {
				declaration.Instances[label].Origins = append(declaration.Instances[label].Origins, anotherLabel)
			}
		}
	}

	sort.Slice(nodeLabels, func(i, j int) bool {
		idxLabel, idx2Label := nodeLabels[i], nodeLabels[j]
		switch {
		case declaration.Instances[idxLabel].BeginAt < declaration.Instances[idx2Label].BeginAt:
			return true
		case declaration.Instances[idxLabel].BeginAt > declaration.Instances[idx2Label].BeginAt:
			return false
		default:
			return strings.Compare(idxLabel, idx2Label) == -1
		}
	})
	for i, label := range nodeLabels {
		if len(originLabels) > 0 && (i == 0 || r.Float64() >= 0.5) {
			declaration.Instances[label].Origins = consistentCollectionOption(originLabels).Select(r)
		} else if i > 0 {
			declaration.Instances[label].DurableNodes = consistentCollectionOption(nodeLabels[:i]).Select(r)
		}
	}

	//
	for i := 1; i <= countRapidAgents; i++ {
		beginAt := declaration.PrimaryLevel + 5
		declaration.Instances[fmt.Sprintf("REDACTED", i)] = composeRapidMember(
			r, beginAt+(5*int64(i)), rapidSources,
		)
	}

	return declaration, nil
}

//
//
//
//
func composeMember(
	r *rand.Rand, style e2e.Style, beginAt int64, compelRepository bool,
) *e2e.DeclarationMember {
	member := e2e.DeclarationMember{
		Release:          memberReleases.Select(r).(string),
		Style:             string(style),
		BeginAt:          beginAt,
		Datastore:         memberDatastores.Select(r).(string),
		PrivatekeyProtocol:  memberPrivatekeyProtocols.Select(r).(string),
		LedgerAlignRelease: memberLedgerAlignments.Select(r).(string),
		StatusAlign:        memberStatusAlignments.Select(r).(bool) && beginAt > 0,
		EndureCadence:  pointerUint64(uint64(memberEndurePeriods.Select(r).(int))),
		MirrorCadence: uint64(memberMirrorPeriods.Select(r).(int)),
		PreserveLedgers:     uint64(memberPreserveLedgers.Select(r).(int)),
		Vary:          memberVariations.Select(r),
	}

	//
	//
	if compelRepository {
		member.PreserveLedgers = 0
		member.MirrorCadence = 3
	}

	//
	//
	if member.EndureCadence != nil && *member.EndureCadence == 0 && member.PreserveLedgers > 0 {
		if r.Float64() > 0.5 {
			member.PreserveLedgers = 0
		} else {
			member.EndureCadence = pointerUint64(member.PreserveLedgers)
		}
	}

	//
	//
	if member.PreserveLedgers > 0 {
		if member.EndureCadence != nil && member.PreserveLedgers < *member.EndureCadence {
			member.PreserveLedgers = *member.EndureCadence
		}
		if member.PreserveLedgers < member.MirrorCadence {
			member.PreserveLedgers = member.MirrorCadence
		}
	}

	return &member
}

func composeRapidMember(r *rand.Rand, beginAt int64, sources []string) *e2e.DeclarationMember {
	return &e2e.DeclarationMember{
		Style:            string(e2e.StyleRapid),
		Release:         memberReleases.Select(r).(string),
		BeginAt:         beginAt,
		Datastore:        memberDatastores.Select(r).(string),
		EndureCadence: pointerUint64(0),
		DurableNodes: sources,
		Vary:         rapidMemberVariations.Select(r),
	}
}

func pointerUint64(i uint64) *uint64 {
	return &i
}

//
//
//
//
//
//
func analyzeScaledReleases(s string) (scaledOption, string, error) {
	wc := make(scaledOption)
	var finalRelease string

	records := strings.Split(strings.TrimSpace(s), "REDACTED")

	for _, entry := range records {
		segments := strings.Split(strings.TrimSpace(entry), "REDACTED")

		var ver string
		switch len(segments) {
		case 2:
			//
			ver = strings.TrimSpace(
				strings.Join([]string{"REDACTED", segments[0]}, "REDACTED"),
			)
		case 3:
			//
			ver = strings.TrimSpace(
				strings.Join([]string{segments[0], segments[1]}, "REDACTED"),
			)
		default:
			return nil, "REDACTED", fmt.Errorf(
				"REDACTED",
				entry,
			)
		}

		magnitudeStr := strings.TrimSpace(segments[len(segments)-1])
		magnitude, err := strconv.Atoi(magnitudeStr)
		if err != nil {
			return nil, "REDACTED", fmt.Errorf("REDACTED", magnitudeStr, err)
		}

		if magnitude < 1 {
			return nil, "REDACTED", errors.New("REDACTED")
		}

		wc[ver] = uint(magnitude)
		finalRelease = ver
	}

	return wc, finalRelease, nil
}

//
//
//
func gitArchiveNewestDeliveryRelease(gitArchiveFolder string) (string, error) {
	opts := &git.PlainOpenOptions{
		DetectDotGit: true,
	}
	r, err := git.PlainOpenWithOptions(gitArchiveFolder, opts)
	if err != nil {
		return "REDACTED", err
	}
	labels := make([]string, 0)
	labelItems, err := r.TagObjects()
	if err != nil {
		return "REDACTED", err
	}
	err = labelItems.ForEach(func(labelItem *object.Tag) error {
		labels = append(labels, labelItem.Name)
		return nil
	})
	if err != nil {
		return "REDACTED", err
	}
	return locateNewestDeliveryLabel(release.TMCoreSemaphoreRev, labels)
}

func locateNewestDeliveryLabel(rootRev string, labels []string) (string, error) {
	rootSemaphoreRev, err := semver.NewVersion(strings.Split(rootRev, "REDACTED")[0])
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", rootRev, err)
	}
	componentRev := fmt.Sprintf("REDACTED", rootSemaphoreRev.Major(), rootSemaphoreRev.Minor())
	//
	//
	componentStr := "REDACTED" + componentRev
	revConnect, err := semver.NewConstraint(componentStr)
	if err != nil {
		return "REDACTED", err
	}
	var newestRev *semver.Version
	for _, tag := range labels {
		if !strings.HasPrefix(tag, "REDACTED") {
			continue
		}
		currentRev, err := semver.NewVersion(tag)
		//
		if err != nil {
			continue
		}
		//
		if len(currentRev.Prerelease()) != 0 {
			continue
		}
		//
		if !revConnect.Check(currentRev) {
			continue
		}
		if newestRev == nil || currentRev.GreaterThan(newestRev) {
			newestRev = currentRev
		}
	}
	//
	//
	if newestRev == nil {
		return "REDACTED", nil
	}
	//
	//
	vs := newestRev.String()
	if !strings.HasPrefix(vs, "REDACTED") {
		return "REDACTED" + vs, nil
	}
	return vs, nil
}
