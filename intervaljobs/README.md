# Interval Jobs

Schedule re-occuring tasks on an interval schedule, i.e. every 30 minutes.

```golang
IntervalJobs := intervaljobs.NewIntervalJobManager(intervaljobs.ManagerConfigs{
    Interval: 1 * time.Minute, //check jobs every minute
})
IntervalJobs.AddJob(intervaljobs.JobConfig{
    // run 5 minutes after startup, then every hour
    Interval: &intervaljobs.IntervalJobTime{
        Hr: 1,
    },
    Offset: &intervaljobs.IntervalJobTime{
        Min: 5, 
    },
    Handler: func() {
        // ... tasks ...
    },
})
go IntervalJobs.Start()
```