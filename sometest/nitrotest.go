package sometest

import (
	"flag"
	"github.com/spf13/nitro"
)

func someFunc() {
	var a [1024 * 1024 * 64]byte
	for aa := range a {
		a[aa] = 1
	}
}

func Analysis() {
	timer := nitro.Initialize()
	// 通过命令行参数开启分析
	flag.Parse()
	// 默认开启
	// nitro.AnalysisOn = true

	someFunc()
	timer.Step("step 1, write index")

	//otherFunc()
	timer.Step("step 2, batch insert")
}
