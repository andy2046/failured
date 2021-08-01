

# failured
`import "github.com/andy2046/failured"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package failured implements the Adaptive Accrual Failure Detector

This package implements [A New Adaptive Accrual Failure Detector for Dependable Distributed Systems](<a href="https://dl.acm.org/citation.cfm?id=1244129">https://dl.acm.org/citation.cfm?id=1244129</a>)

To use the failure detector, you need a heartbeat loop that will
call RegisterHeartbeat() at regular intervals.  At any point, you can call FailureProbability() which will
report the suspicion level since the last time RegisterHeartbeat() was called.




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Config](#Config)
* [type Detector](#Detector)
  * [func New(options ...Option) *Detector](#New)
  * [func (d *Detector) CheckFailure(now ...int64) bool](#Detector.CheckFailure)
  * [func (d *Detector) FailureProbability(now ...int64) float64](#Detector.FailureProbability)
  * [func (d *Detector) RegisterHeartbeat(now ...int64)](#Detector.RegisterHeartbeat)
* [type Option](#Option)


#### <a name="pkg-files">Package files</a>
[failured.go](/src/github.com/andy2046/failured/failured.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    // DefaultConfig is the default Detector Config.
    DefaultConfig = Config{
        WindowSize:       1000,
        Factor:           0.9,
        FailureThreshold: 0.5,
    }
)
```



## <a name="Config">type</a> [Config](/src/target/failured.go?s=1129:1275#L46)
``` go
type Config struct {
    // max size of inter-arrival time list
    WindowSize uint64
    // a scaling factor
    Factor float64

    FailureThreshold float64
}
```
Config used to init Detector.










## <a name="Detector">type</a> [Detector](/src/target/failured.go?s=752:1092#L29)
``` go
type Detector struct {
    // contains filtered or unexported fields
}
```
Detector is a failure detector.







### <a name="New">func</a> [New](/src/target/failured.go?s=1558:1595#L69)
``` go
func New(options ...Option) *Detector
```
New returns a new failure detector.





### <a name="Detector.CheckFailure">func</a> (\*Detector) [CheckFailure](/src/target/failured.go?s=3186:3236#L145)
``` go
func (d *Detector) CheckFailure(now ...int64) bool
```
CheckFailure returns true if `FailureProbability` is
equal to or higher than `FailureThreshold`.
`now` is the current time in Microseconds.




### <a name="Detector.FailureProbability">func</a> (\*Detector) [FailureProbability](/src/target/failured.go?s=2554:2613#L112)
``` go
func (d *Detector) FailureProbability(now ...int64) float64
```
FailureProbability calculates the suspicion level at time `now`
that the remote end has failed.
`now` is the current time in Milliseconds,
default to time.Now() in Milliseconds.




### <a name="Detector.RegisterHeartbeat">func</a> (\*Detector) [RegisterHeartbeat](/src/target/failured.go?s=1998:2048#L86)
``` go
func (d *Detector) RegisterHeartbeat(now ...int64)
```
RegisterHeartbeat registers a heartbeat at time `now`.
`now` is the time in Milliseconds at which the heartbeat was received,
default to time.Now() in Milliseconds.




## <a name="Option">type</a> [Option](/src/target/failured.go?s=1324:1352#L56)
``` go
type Option = func(*Config) error
```
Option applies config to Detector Config.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)