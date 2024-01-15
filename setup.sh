#!/usr/bin/env bash
set -e

yes | goenv install --skip-existing

go install golang.org/x/tools/gopls@latest
