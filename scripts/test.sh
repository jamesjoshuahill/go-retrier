#!/usr/bin/env bash
set -eu

ginkgo -r --randomizeAllSpecs --randomizeSuites --cover
