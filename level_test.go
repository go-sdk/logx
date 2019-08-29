package logx

import (
	"fmt"
	"testing"
)

func TestLevel_String(t *testing.T) {
	fmt.Println(FatalLevel)
	fmt.Println(ErrorLevel)
	fmt.Println(WarnLevel)
	fmt.Println(InfoLevel)
	fmt.Println(DebugLevel)
}
