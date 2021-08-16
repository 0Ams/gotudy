#!/bin/bash

function cleanTestCache() {
    go clean -testcache;
}

function test_() {
    go test ./algorithm/...;
}

case $1 in
    clean) cleanTestCache;
    ;;
esac

test_;