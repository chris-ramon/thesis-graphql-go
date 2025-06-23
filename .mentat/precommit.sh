#!/bin/bash

echo "Running markdown linter on README.md..."
markdownlint README.md || true

echo "Running spell checker on README.md..."
cspell README.md || true

# Since this is just a documentation repository with no actual code,
# there's no need for traditional testing, building, or linting.
echo "No code to build or test in this repository."
