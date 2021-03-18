#!/bin/bash

# Remove go-build* files in temp
rm -rf /tmp/go-build* 2>/dev/null 1>/dev/null

echo "- OS: $(head -n1 </etc/issue)"
echo "- $(go version)"
echo "- Current path: $(pwd)"
echo "- Current user: $(whoami)"
echo "- Time: $(date)"
echo "- To run linters                : ./.github/run-tests-lint.sh"
echo "- To run unit tests and coverage: ./.github/run-tests-coverage.sh"
echo "- To run full tests before merge: ./.github/run-tests-merge.sh"
