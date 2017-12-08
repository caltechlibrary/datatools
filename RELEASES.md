
# Release Process and tags

## Preleases and production releases

This describes a simple release process organized around semantic versioned tags.

When a new release is ready it should be tag as v0.XX.XX-pre (where XX is a number) and published as a 'pre-release' on
Github. Send a note to the development group so someone other than the proposer of the release will be able to 
indepentently evaluate the release changes.  At a minimum they should run `bash test_cmds.bash` and walk through some 
of the documentation examples. If the release is verified as ready then a new release will be cut with a tag like
v0.XX.XX (there the XX are the same as in v0.XX.XX-pre). E.g. v0.0.10-pre would become v0.0.10 on success.

If the release fails verification, bugs need to be fixed and a new release proposed create after the fixes. 
The "patch" number persion should be incremented in v0.XX.XX-pre (e.g. v0.0.10-pre would be followed by 
v0.0.11-pre indicating that patches). 

NOTE: This means there their can be skips in the production patch numbers between release. E.g. v.0.0.9 might
be followed by v0.0.10-pre, v0.0.11-pre, v0.0.12-pre before a v0.0.12 appears as a production release.

Production and pre-releases should include Zip files of the compiled cli to be tested by `bash test_cmds.bash`.

## Dev releases

Dev release may happend from time to time as needed. They should always end in a '-dev' version number (e.g. v0.0.10-dev). 
They normally should not have any pre-compiled binaries to avoid confusion. They should be flagged as draft (pre-release)
on Github.

## Making a release

1. Set the version number in PACKAGE.go (where PACKAGE is the name of the package, e.g. dataset is the name of the dataset
package so you'd change the version number in dataset.go).
2. Run `make clean`
3. Run `make test` and make sure they pass, if some fail document it if you plan to release it (e.g. GSheet integration not tested because...)
4. Run `make release`

You are now ready to go to Github and create a release. If you are uploading compiled versions upload the zip files in the _dist_
folder.

