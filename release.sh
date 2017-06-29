#!/bin/bash
# gox and ghr must be installed
# GITHUB_TOKEN environment variables must be set
ggallin release -u=edvakf --os="darwin linux windows" --arch="386 amd64"
