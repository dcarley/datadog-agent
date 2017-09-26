// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017 Datadog, Inc.

package common

import (
	"fmt"
	"os"

	"github.com/DataDog/datadog-agent/pkg/config"
	"golang.org/x/sys/windows/svc/eventlog"
	log "github.com/cihub/seelog"
)

// PanicHandler handles how we log a crash causing panic
// output contains the full output (including stack traces)
func PanicHandler(output string) {
	// output contains the full output (including stack traces)
	err := config.SetupLogger(
		"error", config.Datadog.GetString("log_panic_file"),
		"", false, false, "")

	msg := fmt.Sprintf("Agent panicked (oh no!):\n\n%s\n", output)
	if err == nil {
		log.Error(msg)
		log.Flush()
	} else {
		fmt.Print(msg)
	}
	h, err := eventlog.Open("Datadog Agent")
	if nil != err {
		return
	}
	defer h.Close()

	h.Error(0xc0000001, output)

	os.Exit(1)
}
