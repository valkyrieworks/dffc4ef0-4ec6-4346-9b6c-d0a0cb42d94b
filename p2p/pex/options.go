package pex

import "time"

const (
	//
	requireLocationLimit = 1000

	//
	exportLocationCadence = time.Minute * 2

	//
	agedSegmentVolume = 64

	//
	agedSegmentNumber = 64

	//
	newSegmentVolume = 64

	//
	newSegmentNumber = 256

	//
	agedSegmentsEachCluster = 4

	//
	newSegmentsEachCluster = 32

	//
	maximumNewSegmentsEachLocation = 4

	//
	//
	countAbsentPeriods = 7

	//
	countAttempts = 3

	//
	maximumBreakdowns = 10 //

	//
	minimumFlawedPeriods = 7

	//
	fetchPreferenceFraction = 23

	//
	minimumFetchPreference = 32

	//
	//
	maximumFetchPreference = 250
)
