#!/bin/bash

#
# Generate a new draft release using jq and gh
#
RELEASE_TAG="v$(jq -r .version codemeta.json)"
RELEASE_NOTES="$(jq -r .releaseNotes codemeta.json)"
make save msg="prep for ${RELEASE_TAG}, $RELEASE_NOTES}"
# Now generate a draft releas
gh release create "${RELEASE_TAG}" \
  --verify-tag --draft \
  --notes="${RELEASE_NOTES}" \
  dist/*.zip 
echo "Now goto repo release and finalize draft"
