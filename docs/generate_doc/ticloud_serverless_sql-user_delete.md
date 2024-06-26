## ticloud serverless sql-user delete

Delete a SQL user

```
ticloud serverless sql-user delete [flags]
```

### Examples

```
  Delete a SQL user in interactive mode:
  $ ticloud serverless sql-user delete

  Delete a SQL user in non-interactive mode:
  $ ticloud serverless sql-user delete -c <cluster-id> --user <user-name>
```

### Options

```
  -c, --cluster-id string   The cluster ID of the SQL user to be deleted.
      --force               Delete a SQL user without confirmation.
  -h, --help                help for delete
  -u, --user string         The name of the SQL user to be deleted.
```

### Options inherited from parent commands

```
  -D, --debug            Enable debug mode
      --no-color         Disable color output
  -P, --profile string   Profile to use from your configuration file
```

### SEE ALSO

* [ticloud serverless sql-user](ticloud_serverless_sql-user.md)	 - Manage cluster SQL users

