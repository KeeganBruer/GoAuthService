package Intervaljobs

import (
	"strconv"
	"strings"
	"time"
)

type IntervalJobTime struct {
	Day int
	Hr  int
	Min int
	Sec int
}
type IntervalJob struct {
	Handler  func()
	offset   *IntervalJobTime
	interval *IntervalJobTime
	timeStmp *IntervalJobTime
}
type ManagerConfigs struct {
	Interval time.Duration
}
type IntervalJobManager struct {
	Interval time.Duration
	Jobs     []*IntervalJob
}

func NewIntervalJobManager(cfg ManagerConfigs) *IntervalJobManager {
	manager := &IntervalJobManager{
		Jobs:     make([]*IntervalJob, 0),
		Interval: cfg.Interval,
	}

	return manager
}

type JobConfig struct {
	Offset         *IntervalJobTime
	Interval       *IntervalJobTime
	Handler        func()
	InstantTrigger bool
}

func (manager *IntervalJobManager) AddJob(cfg JobConfig) {
	startingOffset := cfg.Interval
	if cfg.InstantTrigger {
		startingOffset = &IntervalJobTime{}
	} else if cfg.Offset != nil {
		startingOffset = cfg.Offset
	}

	newJob := &IntervalJob{
		Handler:  cfg.Handler,
		offset:   startingOffset,
		interval: cfg.Interval,
	}
	manager.Jobs = append(manager.Jobs, newJob)
}
func (manager *IntervalJobManager) Start() {
	for i := range manager.Jobs {
		job := manager.Jobs[i]
		job.UpdateTimestamp(true)
	}

	delay := time.NewTicker(3 * time.Second)
	<-delay.C
	delay.Stop()

	manager.handleIntervalStep()

	ticker := time.NewTicker(manager.Interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		manager.handleIntervalStep()
	}
}

func (manager *IntervalJobManager) handleIntervalStep() {
	nowTime := time.Now()
	timeStmp := IntervalJobTime{
		Day: nowTime.YearDay(),
		Hr:  nowTime.Hour(),
		Min: nowTime.Minute(),
		Sec: nowTime.Second(),
	}
	//fmt.Println("== STEP:", timeStmp)
	for i := range manager.Jobs {
		job := manager.Jobs[i]
		if CompareTimestamps(job.timeStmp, &timeStmp) < 0 {
			continue
		}
		job.Handler()
		job.UpdateTimestamp()
	}
}
func (job *IntervalJob) UpdateTimestamp(first ...bool) {
	nowTime := time.Now()
	offset := job.interval
	if len(first) > 0 && first[0] {
		offset = job.offset
	}
	job.timeStmp = &IntervalJobTime{
		Day: nowTime.YearDay() + offset.Day,
		Hr:  nowTime.Hour() + offset.Hr,
		Min: nowTime.Minute() + offset.Min,
		Sec: nowTime.Second() + offset.Sec,
	}
}
func CompareTimestamps(timeStmp1 *IntervalJobTime, timeStmp2 *IntervalJobTime) int {
	time1 := timeStmp1.Sec
	time2 := timeStmp2.Sec

	time1 = time1 + (timeStmp1.Min * 60)
	time2 = time2 + (timeStmp2.Min * 60)

	time1 = time1 + (timeStmp1.Hr * 60 * 60)
	time2 = time2 + (timeStmp2.Hr * 60 * 60)

	time1 = time1 + (timeStmp1.Day * 24 * 60 * 60)
	time2 = time2 + (timeStmp2.Day * 24 * 60 * 60)

	return time2 - time1
}

func Spec(specKey string, specVal int) bool {
	meetsSpec := false
	fmStr := strings.Trim(specKey, " ")
	if fmStr == "" {
		meetsSpec = true
	} else if strings.HasPrefix(fmStr, "/") {
		fmStr = strings.Replace(fmStr, "/", "", 1)
		num, _ := strconv.Atoi(fmStr)
		if num != 0 && specVal%num == 0 {
			meetsSpec = true
		}
	}
	return meetsSpec
}
