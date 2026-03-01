package pex

import "time"

const (
	//
	requireLocatorLimit = 1000

	//
	exportLocatorDuration = time.Minute * 2

	//
	agedSegmentExtent = 64

	//
	agedSegmentTotal = 64

	//
	freshSegmentExtent = 64

	//
	freshSegmentTotal = 256

	//
	agedSegmentsEveryCohort = 4

	//
	freshSegmentsEveryCohort = 32

	//
	maximumFreshSegmentsEveryLocator = 4

	//
	//
	countAbsentEpochs = 7

	//
	countAttempts = 3

	//
	maximumMishaps = 10 //

	//
	minimumFlawedEpochs = 7

	//
	fetchPreferenceRatio = 23

	//
	minimumFetchPreference = 32

	//
	//
	maximumFetchPreference = 250
)
