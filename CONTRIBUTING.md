## Local Development

To run beaver locally for development or even testing, please follow the following:

```bash
# Use src/github.com/clivern/beaver
$ mkdir -p $GOPATH/src/github.com/clivern/beaver
$ git clone https://github.com/Clivern/Beaver.git $GOPATH/src/github.com/clivern/beaver
$ cd $GOPATH/src/github.com/clivern/beaver

# Create a feature branch
$ git branch feature/x
$ git checkout feature/x

$ export GO111MODULE=on
$ cp config.yml config.dist.yml
$ cp config.yml config.test.yml

# Add redis to config.test.yml and config.dist.yml

# to run beaver
$ go run beaver.go
$ go build beaver.go

# To run test cases
$ make ci
```

Then Create a PR with the master branch.

## Contributing

- With issues:
  - Use the search tool before opening a new issue.
  - Please provide source code and commit sha if you found a bug.
  - Review existing issues and provide feedback or react to them.

- With pull requests:
  - Open your pull request against `master`
  - Your pull request should have no more than two commits, if not you should squash them.
  - It should pass all tests in the available continuous integrations systems such as TravisCI.
  - You should add/modify tests to cover your proposed code changes.
  - If your pull request contains a new feature, please document it on the README.
