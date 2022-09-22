# Generation

To regenerate the tsql parser, execute:

```sh
antlr2participle tsql.g4
```

The antlr2participle tool is experimental and lives in `internal`.

## Not Supported

- `ALTER DATABASE`
- `PIVOT` and `UNPIVOT`
- `sp_xml_preparedocument`
- Materialized columns
- Some table options
- Some hint options