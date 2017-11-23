package toy4

import (
	"fmt"
	"time"
)

// Location
// func LoadLocation(name string) (*Location, error)

// Time
// time should be used to save and passby so use time.Time not *time.Time
// IsZero Unix time.Time.Unix() -> unix timestamp
// time.Time.UnixNano() -> nano second timestamp
// time.Parse time.ParseInLocation time.Time.Format // formats
func DemoTimeParse() {
	t, err := time.Parse("2006-01-02 15:04:05", "2017-11-22 15:39:40")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	t, err = time.ParseInLocation("2006-01-02 15:04:05", "2017-11-22 15:39:40", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	ft := "01/02 03:04:05PM '06"
	t, err = time.ParseInLocation(ft, "11/12 03:39:40PM '17", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}

// json time.RFC3339Nano
func DemoMarshalJsonTime() {
	t := time.Now()
	if y := t.Year(); y < 0 || y >= 10000 {
		fmt.Println("Time.MarshalJSON: year out of range [0, 9999]")
	}
	// "" for json
	fmt.Println(t.Format(`"2006-01-02 15:04:05"`))
}

// time.Round / Truncate

// Duration

// Timer
// type Timer struct {
// 	C <-chan Time
// 	r runtimeTimer
// }
// --> runtimeTimer
// type timer struct {
// 	i      int // heap index
// 	when   int64 -> timer wakes up at when, and then at when+period (period > 0 only)
// 	period int64
// 	f      func(interface{}, uintptr)
// 	arg    interface{}
// 	seq    uintptr
// }

func DemoSimTimeoutByTimeAfter() {
	c := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		c <- 1
	}()

	select {
	case <-c:
		fmt.Println("channel...")
	case <-time.After(time.Second * 2):
		close(c)
		fmt.Println("timeout...")
	}
}

// Ticker

// Weekday Month
