#!/bin/bash

#
# Generate a new draft release using Make and gh
#
make
make website
make release
RELEASE_TAG="v$(jq -r .version codemeta.json)"
RELEASE_NOTES="$(jq .releaseNotes codemeta.json)"
# Now generate a draft releas
gh release create "${RELEASE_TAG}" \
  --verify-tag --draft \
  --notes="${RELEASE_NOTES}" \
  dist/*.zip 
echo "Now goto repo release and finalize draft"
