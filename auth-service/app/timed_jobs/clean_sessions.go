package timedjobs

import (
	"fmt"
	"time"
)

func CleanSessions() {
	fmt.Println("Task executed at:", time.Now())

}
