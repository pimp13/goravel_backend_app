package tests

import (
	"github.com/goravel/framework/testing"

	"goravel_by_gin/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
