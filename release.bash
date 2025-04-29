#!/bin/bash

REPO_ID="$(basename $(pwd))"
GROUP_ID="$(basename $(dirname $(pwd)))"
REPO_URL="https://github.com/${GROUP_ID}/${REPO_ID}"
echo "REPO_URL -> ${REPO_URL}"

#
# Generate a new draft release jq and gh
#
RELEASE_TAG="v$(jq -r .version codemeta.json)"
RELEASE_NOTES="$(jq -r .releaseNotes codemeta.json | tr '\n' ' ')"
echo "tag: ${RELEASE_TAG}, notes:"
jq -r .releaseNotes codemeta.json >release_notes.tmp
cat release_notes.tmp

# Now we're ready to push things.
read -r -p "Push release to GitHub with gh? (y/N) " YES_NO
if [ "$YES_NO" = "y" ]; then
	make save msg="prep for ${RELEASE_TAG}, ${RELEASE_NOTES}"
	# Now generate a draft releas
	echo "Pushing release ${RELEASE_TAG} to GitHub"
	gh release create "${RELEASE_TAG}" \
 		--draft \
		-F release_notes.tmp \
		--generate-notes \
		dist/*.zip 
	
	cat <<EOT

Now goto repo release and finalize draft.

	${REPO_URL}/releases

EOT

fi
