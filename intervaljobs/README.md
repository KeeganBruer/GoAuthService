# Interval Jobs

Schedule re-occuring tasks on an interval schedule, i.e. every 30 minutes.

```golang
IntervalJobs := intervaljobs.NewIntervalJobManager(intervaljobs.ManagerConfigs{
    Interval: 1 * time.Minute, //check jobs every minute
})
IntervalJobs.AddJob(
    &intervaljobs.IntervalJobTime{
        Min: 10, //run every 10 min
    },
    func() {
        // ... tasks ...
    },
)
go IntervalJobs.Start()
```