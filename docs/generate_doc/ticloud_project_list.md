## ticloud project list

List all accessible projects

```
ticloud project list [flags]
```

### Examples

```
  List the projects:
  $ ticloud project list

  List the projects with json format:
  $ ticloud project list -o json
```

### Options

```
  -h, --help            help for list
  -o, --output string   Output format, one of ["human" "json"]. For the complete result, please use json format. (default "human")
```

### Options inherited from parent commands

```
  -D, --debug            Enable debug mode
      --no-color         Disable color output
  -P, --profile string   Profile to use from your configuration file
```

### SEE ALSO

* [ticloud project](ticloud_project.md)	 - Manage projects

