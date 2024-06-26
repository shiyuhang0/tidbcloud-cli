## ticloud serverless export create

Export data from a TiDB Serverless cluster

```
ticloud serverless export create [flags]
```

### Examples

```
  Create an export in interactive mode:
  $ ticloud serverless export create

  Create an export with local type in non-interactive mode:
  $ ticloud serverless export create -c <cluster-id> --database <database> --table <table>

  Create an export with s3 type in non-interactive mode:
  $ ticloud serverless export create -c <cluster-id> --s3.uri <s3-uri> --s3.access-key-id <access-key-id> --s3.secret-access-key <secret-access-key>
```

### Options

```
  -c, --cluster-id string             The ID of the cluster, in which the export will be created.
      --compression string            The compression algorithm of the export file. One of ["GZIP" "SNAPPY" "ZSTD" "NONE"]. (default "GZIP")
      --database string               The database name you want to export. (default "*")
      --file-type string              The export file type. One of ["CSV" "SQL"]. (default "CSV")
  -h, --help                          help for create
      --s3.access-key-id string       The access key ID of the S3. Required when target type is S3.
      --s3.secret-access-key string   The secret access key of the S3. Required when target type is S3.
      --s3.uri string                 The s3 uri in s3://<bucket>/<path> format. Required when target type is S3.
      --table string                  The table name you want to export. (default "*")
      --target-type string            The export target. One of ["LOCAL" "S3"]. (default "LOCAL")
```

### Options inherited from parent commands

```
  -D, --debug            Enable debug mode
      --no-color         Disable color output
  -P, --profile string   Profile to use from your configuration file
```

### SEE ALSO

* [ticloud serverless export](ticloud_serverless_export.md)	 - Manage TiDB Serverless exports

