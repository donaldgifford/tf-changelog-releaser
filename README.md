# tf-changelog-releaser

[![codecov](https://codecov.io/gh/donaldgifford/tf-changelog-releaser/graph/badge.svg?token=KQA7XTPRBT)](https://codecov.io/gh/donaldgifford/tf-changelog-releaser)
[![Go Report Card](https://goreportcard.com/badge/github.com/donaldgifford/tf-changelog-releaser)](https://goreportcard.com/report/github.com/donaldgifford/tf-changelog-releaser)
![OpenTofu Version](https://img.shields.io/badge/tofu-%3E%3D1.6.0-blue.svg)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D0.12.0-blue.svg)
![Terraform-docs Version](https://img.shields.io/badge/tf--docs-v0.20.0-blue)

tf-changelog-releaser is a golang based cli tool that creates a changelog and
releases for terraform modules in a mono-repo style configuration. The
inspiration for this tool comes from
[terraform module releaser](https://github.com/techpivot/terraform-module-releaser)
which is a similar tool that runs as a github action. This tool will allow you
to generate the documentation, changelog, and release in a similar way just
giving you the ability to run from the cli or github action.

This way you can use a `pre-commit-hook`, bash script, Makefile, Github Action,
or run it however you want.

## License

This code is released under the MIT License. See [LICENSE.txt](LICENSE.txt).
