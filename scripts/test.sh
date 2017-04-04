#!/usr/bin/env bash
set -eu

ginkgo -r -p --randomizeAllSpecs --randomizeSuites --cover
