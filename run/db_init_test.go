package run

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestInitDb(t *testing.T) {
	_, err := InitMysql("119.45.34.108:3306", "root", "benang2024", "ipr")
	err = errors.Cause(err)
	if err != nil {
		fmt.Println(err.Error())
	}
}
