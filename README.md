Cos
===

Visualize the coverage as Pull-Request checks.

[![Travis](https://img.shields.io/travis/linyows/cos.svg?style=flat-square)][travis]
[![GitHub release](http://img.shields.io/github/release/linyows/cos.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: https://travis-ci.org/linyows/cos
[release]: https://github.com/linyows/cos/releases
[license]: https://github.com/linyows/cos/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/linyows/cos

Description
-----------

The Cos creates commit status for coverage.

Usage
-----

### SimpleCov and Rspec

```sh
$ cos -p "cat coverage/.last_run.json | grep covered_percent | awk '{print $2}'" bin/rspec
```

### PHP CodeCoverage and PHPUnit

```sh
$ cos -p "cat coverage/report.txt | grep -i lines | awk '{print $2}' | sed 's/%//'" vendor/bin/phpunit
```

Config File
-----------

```sh
cat << EOF > .cos
test = "bin/rspec"
percent = "cat coverage/.last_run.json | grep covered_percent | awk '{print $2}'"
github_token = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
base_branch = "master"
EOF
```

Install
-------

To install, use `go get`:

```sh
$ go get -d github.com/linyows/cos
```

Contribution
------------

1. Fork ([https://github.com/linyows/cos/fork](https://github.com/linyows/cos/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

Author
------

[linyows](https://github.com/linyows)
