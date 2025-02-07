# Operator Release

This document describes the release process for the operator

## Release script

The `release.sh` script will generate/update the Helm chart files and the OperatorHub.io bundle.

Ensure the following software is installed before running the release script:

- `yq`
- `operator-sdk`

Run the release script:

1. Run the `release.sh` script from the root of the mondoo-operator repo with the previous version of the operator as the first parameter and the new version of the operator as the second parameter (without any leading 'v' in the version string). For example:

```bash
$ ./release.sh 0.2.0 0.2.1
```

### Helm Chart and Operator bundle

Mondoo Operator helm chart has been auto-generated using the [helmify](https://github.com/arttor/helmify) tool via the `release.sh` script. The CI uses [chart-releaser-action](https://github.com/helm/chart-releaser-action) to self host the charts using GitHub pages.

The following steps need to be followed to release Helm chart.

#### Helm Chart Release Workflow

Helm chart release action is executed against release tags. It checks each chart in the charts folder, and whenever there's a new chart version, creates a corresponding GitHub release named for the chart version, adds Helm chart artifacts to the release, and creates or updates an index.yaml file with metadata about those releases, which is then hosted on GitHub Pages.

### Committing the release

After running the `release.sh` script, you can create a pull request containing the changes made to the repo (mainly the files under ./charts and ./config. Once the pull request has merged, you need to tag the release.

1. git checkout main
2. git pull
3. git tag v0.2.0
4. git push origin v0.2.0

## OLM Bundle for OperatorHub.io

### Steps

The release.sh script generates a `/bundle` directory with the required manifests and metadata.

1. Create a PR to https://github.com/k8s-operatorhub/community-operators

- fork https://github.com/k8s-operatorhub/community-operators
- in the forked repo copy/sync the files generated by the `release.sh` script from the `mondoo-operator/bundle` directory into `community-operators/operators/mondoo-operator/0.2.0/`.
- update the `community-operators/operators/mondoo-operator/mondoo-operator.package.yaml` to point to latest version of the operator
- commit and open PR

## FAQ

**I want to run tests locally before committing**

The following tests are the core of the CI that is being used on the `k8s-operatorhub` repo, and can be run locally:

```bash
bash <(curl -sL https://raw.githubusercontent.com/redhat-openshift-ecosystem/community-operators-pipeline/ci/latest/ci/scripts/opp.sh) kiwi, lemon, orange operators/mondoo-operator/0.0.10
```
