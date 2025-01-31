#!/bin/bash

REPO_ID="$(basename $(pwd))"
GROUP_ID="$(basename $(dirname $(pwd)))"
REPO_URL="https://github.com/${GROUP_ID}/${REPO_ID}"
echo "REPO_URL -> ${REPO_URL}"
#
# Generate a new draft release jq and gh
#
RELEASE_TAG="v$(jq -r .version codemeta.json)"
RELEASE_NOTES="$(jq -r .releaseNotes codemeta.json)"
read -r -p "Push release to GitHub with gh? (y/N) " YES_NO
if [ "$YES_NO" = "y" ]; then
	make save msg="prep for ${RELEASE_TAG}, ${RELEASE_NOTES}"
	# Now generate a draft releas
	gh release create "${RELEASE_TAG}" \
 		--draft \
		--notes="${RELEASE_NOTES}" \
		dist/*.zip 
	
	cat <<EOT

Now goto repo release and finalize draft.

	${REPO_URL}

EOT

fi
