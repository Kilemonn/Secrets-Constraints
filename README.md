# Secrets-Constraints

An commandline application that is used to monitor and notify you when constraints onto your stored secrets are not met.

## Constraints

The constraints that are supported are:

- **Unique** - That the value of this property is unique across all matching credentials.
- **HasPrefix** - Check it has the supplied prefix.
- **HasSuffix** - Check it has the supplied suffix.
- **IsNumber** - Check value is numeric.
- **IsBoolean** - Check value is a boolean (true/false/t/f/0/1) - make this case insensitive.

Examples:
``` yaml
credential-providers:
    - GCP
    - AWS
    - Kubernetes:
        ...
constraints:
    - all-are-unique:
        pattern: ALL
        condition: UNIQUE
    - db-strings-have-prefix:
        pattern: db-connection-string-*+
        condition: HasPrefix(jdbc://a.domain.to.db/)
    - has-test-suffix:
        pattern: test-param-*+
        condition: HasSuffix(test)
    - ports-are-numeric:
        pattern: *+port
        condition: IsNumeric
    - debug-is-boolean:
        pattern: debug
        condition: IsBoolean
```
