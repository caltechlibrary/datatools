#!/bin/bash
#

WORKING_BRANCH=$(git branch | grep '* ' | cut -d\  -f 2)
if [ "$WORKING_BRANCH" = "gh-pages" ]; then
    git commit -am "publishing to gh-pages branch"
    git push origin gh-pages
else
    echo "You're in $WORKING_BRANCH branch"
    echo "You need to pull in changes to the gh-pages branch to publish"
    read -p "Pull into gh-pages and publish? Y/N " YES_NO
    if [ "$YES_NO" = "Y" ] || [ "$YES_NO" = "y" ]; then
        echo "Committing and pushing to $WORKING_BRANCH"
        git commit -am "commiting to $WORKING_BANCH";
        git push origin "$WORKING_BRANCH";
        echo "Changing branchs from $WORKING_BRANCH to gh-pages";
        git checkout gh-pages
        echo "Merging changes from origin gh-pages"
        git pull origin gh-pages
        git commit -am "merging origin gh-pages"
        echo "Pulling changes from $WORKING_BRANCH info gh-pages"
        git pull origin "$WORKING_BRANCH"
        echo "Merging changes from $WORKING_BRANCH"
        git commit -am "merging $WORKING_BRANCH with gh-pages"
        echo "Pushing changes up and publishing"
        git push origin gh-pages
        echo "Changing back to your working branch $WORKING_BRANCH"
        git checkout "$WORKING_BRANCH"
    fi
fi
