package job_test

import (
	"errors"
	"github.com/datakube/datakube/internal/test"
	"github.com/datakube/datakube/job"
	"github.com/datakube/datakube/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateJobNeededByTarget(t *testing.T) {
	jobStoreMock := new(test.JobStoreMock)

	job7DaysAgo := types.Job{
		Target: "targetWeekly",
		RunAt:  time.Now().AddDate(0, 0, -7),
	}

	jobYesterday := types.Job{
		Target: "targetDaily",
		RunAt:  time.Now().AddDate(0, 0, -1),
	}

	job1MonthAgo := types.Job{
		Target: "targetMonthly",
		RunAt:  time.Now().AddDate(0, -1, 0),
	}

	jobAlreadyQueued := types.Job{
		Target: "alreadyQueued",
		Status: types.STATUS_QUEUED,
	}

	jobInFuture := types.Job{
		Target: "inFuture",
		RunAt:  time.Now().AddDate(0, 0, 3),
	}

	jobStoreMock.On("GetLatestJobByTargetName", "targetWeekly").Return(job7DaysAgo, nil)
	jobStoreMock.On("GetLatestJobByTargetName", "targetDaily").Return(jobYesterday, nil)
	jobStoreMock.On("GetLatestJobByTargetName", "targetMonthly").Return(job1MonthAgo, nil)
	jobStoreMock.On("GetLatestJobByTargetName", "alreadyQueued").Return(jobAlreadyQueued, nil)
	jobStoreMock.On("GetLatestJobByTargetName", "noExistingJob").Return(types.Job{}, nil)
	jobStoreMock.On("GetLatestJobByTargetName", "inFuture").Return(jobInFuture, nil)
	jobStoreMock.On("GetLatestJobByTargetName", "notFound").Return(types.Job{}, errors.New("not found"))
	jobStoreMock.On("GetLatestJobByTargetName", "error").Return(types.Job{}, errors.New("error"))

	targetWeekly := types.Target{
		Name: "targetWeekly",
		Schedule: types.Schedule{
			Interval: "weekly",
		},
	}

	targetDaily := types.Target{
		Name: "targetDaily",
		Schedule: types.Schedule{
			Interval: "daily",
		},
	}

	targetMonthly := types.Target{
		Name: "targetMonthly",
		Schedule: types.Schedule{
			Interval: "monthly",
		},
	}

	targetQueued := types.Target{
		Name: "alreadyQueued",
		Schedule: types.Schedule{
			Interval: "weekly",
		},
	}

	noPreviousJob := types.Target{
		Name: "noExistingJob",
	}

	inFutureTarget := types.Target{
		Name: "inFuture",
		Schedule: types.Schedule{
			Interval: "weekly",
		},
	}

	targetNotFound := types.Target{
		Name: "notFound",
	}

	targetError := types.Target{
		Name: "error",
	}

	assert.True(t, job.ValidateJobNeededByTarget(targetWeekly, jobStoreMock))
	assert.True(t, job.ValidateJobNeededByTarget(targetDaily, jobStoreMock))
	assert.True(t, job.ValidateJobNeededByTarget(targetMonthly, jobStoreMock))
	assert.False(t, job.ValidateJobNeededByTarget(targetQueued, jobStoreMock))
	assert.False(t, job.ValidateJobNeededByTarget(targetQueued, jobStoreMock))
	assert.True(t, job.ValidateJobNeededByTarget(noPreviousJob, jobStoreMock))
	assert.False(t, job.ValidateJobNeededByTarget(inFutureTarget, jobStoreMock))
	assert.True(t, job.ValidateJobNeededByTarget(targetNotFound, jobStoreMock))
	assert.False(t, job.ValidateJobNeededByTarget(targetError, jobStoreMock))
	assert.True(t, jobStoreMock.AssertExpectations(t))
}
