**NOTE**: This tool is no longer needed because Elastic Fleet packages are now
automatically published after changes are merged into elastic/integrations.

# elastic-package-promoter

This tool outputs a list of packages that are pending promotion in
[elastic/package-storage](https://github.com/elastic/package-storage).
It outputs a summary of the changes ([example](/docs/example-output.md)) for
each package based on the version difference in the origin and target branches.
The output is markdown that can be pasted into a pull request description.

The output is most meaningful when the target is the `production` branch due
to how packages are deleted from `staging` after being promoted.

This tool *does not* do the promotion. It will output the
[`elastic-package promote`](https://www.elastic.co/guide/en/integrations-developer/master/elastic-package.html#_elastic_package_promote)
command to use, but it does not run it for you.

## Installation

```
# Go 1.16+
go install github.com/andrewkroh/elastic-package-promoter@latest
```

## Usage

```
$(go env GOPATH)/bin/elastic-package-promoter -h
Usage of elastic-package-promoter:
  -d    include deprecated packages
  -md-to-html
        convert markdown to HTML
  -origin string
        origin branch (default "snapshot")
  -target string
        target branch (default "production")
  -team string
        select packages owned by this team (e.g. elastic/security-external-integrations)
  -tmpl string
        template file to render (default "embed:summary_of_changes.md.gohtml")
```
