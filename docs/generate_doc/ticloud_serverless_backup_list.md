## ticloud serverless backup list

List serverless cluster backups

```
ticloud serverless backup list [flags]
```

### Examples

```
  List the backups in interactive mode:
  $ ticloud serverless backup list

  List the backups in non-interactive mode:
  $ ticloud serverless backup list -c <cluster-id> 

  List the backups with json format in non-interactive mode:
  $ ticloud serverless backup list -c <cluster-id> -o json
```

### Options

```
  -c, --cluster-id string   The cluster ID of the backup to be listed.
  -h, --help                help for list
  -o, --output string       Output format, one of ["human" "json"]. For the complete result, please use json format. (default "human")
```

### Options inherited from parent commands

```
  -D, --debug            Enable debug mode
      --no-color         Disable color output
  -P, --profile string   Profile to use from your configuration file
```

### SEE ALSO

* [ticloud serverless backup](ticloud_serverless_backup.md)	 - Manage serverless cluster backups

###### Auto generated by spf13/cobra on 15-Apr-2024
