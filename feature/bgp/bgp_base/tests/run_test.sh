#!/bin/bash

KNE_CONFIG_DIR="/usr/local/google/home/wenbli/gocode/src/github.com/openconfig/lemming/kne_config_files"

go test -v=1 -alsologtostderr -config ${KNE_CONFIG_DIR}/kne_config.yaml -testbed ${KNE_CONFIG_DIR}/testbed_gribi.textproto
