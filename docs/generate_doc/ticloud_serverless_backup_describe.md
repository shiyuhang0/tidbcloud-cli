## ticloud serverless backup describe

Describe a serverless cluster backup

```
ticloud serverless backup describe [flags]
```

### Examples

```
  Get the backup in interactive mode:
  $ ticloud serverless backup describe

  Get the backup in non-interactive mode:
  $ ticloud serverless backup describe --backup-id <backup-id>
```

### Options

```
      --backup-id string   The ID of the backup to be described.
  -h, --help               help for describe
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
