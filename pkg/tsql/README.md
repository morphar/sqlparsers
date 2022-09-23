# Generation

To regenerate the tsql parser, execute:

```sh
antlr2participle tsql.g4
```

The antlr2participle tool is experimental.

## Not Supported

- `ALTER DATABASE`
- `PIVOT` and `UNPIVOT`
- `sp_xml_preparedocument`
- Materialized columns
- Some table options
- Some hint options
- Updating data in a remote table by using a linked server
- OPENQUERY
- OPENROWSET