#!/bin/sh
set -e

echo "test mode: $TEST_MODE"

if [ "$TEST_MODE" = "true" ]; then
    echo "test mode"
else
    ./cgibinftw
fi