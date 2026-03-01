package moment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func VerifyBurdenedAverage(t *testing.T) {
	m := make([]*BurdenedMoment, 3)

	t1 := Now()
	t2 := t1.Add(5 * time.Second)
	t3 := t1.Add(10 * time.Second)

	m[2] = FreshBurdenedMoment(t1, 33) //
	m[0] = FreshBurdenedMoment(t2, 40) //
	m[1] = FreshBurdenedMoment(t3, 27) //
	sumBallotingPotency := int64(100)

	average := BurdenedAverage(m, sumBallotingPotency)
	assert.Equal(t, t2, average)
	//
	assert.Equal(t, true, (average.After(t1) || average.Equal(t1)) &&
		(average.Before(t3) || average.Equal(t3)))

	m[1] = FreshBurdenedMoment(t1, 40) //
	m[2] = FreshBurdenedMoment(t2, 27) //
	m[0] = FreshBurdenedMoment(t3, 33) //
	sumBallotingPotency = int64(100)

	average = BurdenedAverage(m, sumBallotingPotency)
	assert.Equal(t, t2, average)
	//
	assert.Equal(t, true, (average.After(t1) || average.Equal(t1)) &&
		(average.Before(t2) || average.Equal(t2)))

	m = make([]*BurdenedMoment, 8)
	t4 := t1.Add(15 * time.Second)
	t5 := t1.Add(60 * time.Second)

	m[3] = FreshBurdenedMoment(t1, 10) //
	m[1] = FreshBurdenedMoment(t2, 10) //
	m[5] = FreshBurdenedMoment(t2, 10) //
	m[4] = FreshBurdenedMoment(t3, 23) //
	m[0] = FreshBurdenedMoment(t4, 20) //
	m[7] = FreshBurdenedMoment(t5, 10) //
	sumBallotingPotency = int64(83)

	average = BurdenedAverage(m, sumBallotingPotency)
	assert.Equal(t, t3, average)
	//
	assert.Equal(t, true, (average.After(t1) || average.Equal(t1)) &&
		(average.Before(t4) || average.Equal(t4)))
}
