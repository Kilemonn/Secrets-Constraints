# Secrets-Constraints

## Overview

An commandline application that is used to verify the value/format of your stored secrets across multiple managers.

The current credential providers that are supported are:
- Google Cloud Secrets Manager
- AWS Secrets Manager
- Kubernetes Secrets
- The local machine environment variables

### Please refer to the [Wiki](https://github.com/Kilemonn/Secrets-Constraints/wiki) for more details!

## Quick Start

Installation of the commandline tool can be done with the following command:

> go install github.com/Kilemonn/Secrets-Constraints@latest

## Usage

The application requires a `.yaml` configuration file that defines the credential providers along with the constraints that you want to perform on each credential.

Using the environment as an example we can define the following `yaml` configuration file to check that the database properties are set correctly (this is an example to demonstrate what kind of validation is available).

``` yaml
credential-providers:
    - Env: # Registers the environment as a credential provider
constraints:
    - database-connection-string-prefix-is-development: # Create a new constraint with arbitrary name
        pattern: db-host-name # This is the regex that will be matched against the credential name (in this case, environment variable name) - if this matches successfully then the following condition will be evaluated against this secret's value (environment variable value)
        condition: HasPrefix(jdbc://path-to-development-db)
    - database-port-number:
        pattern: db-port-number
        condition: IsNumeric
    - all-properties-are-unqiue:
        pattern: ALL
        condition: Unique
```

The above configuration defines the "environment" as the only credential provider.
In this case, all environment variables are loaded and will be run against each of the defined constraints.
In the `constraints` definition, the `pattern` is a **Regular expression (Regex)** pattern that is run against the secret name (in this case the environment variable name). If the pattern matches then the application will attempt to perform the condition against the secret's value (in this case the environment varriable's value).

There is an "ALL" `pattern` keyword that will force match against all entries. In this case, the `all-properties-are-unique` will most likely fail, as generally environment variables do have duplicated values etc.

## Constraints

The constraints that are supported are:

- **Unique** - That the value of this property is unique across all matching credentials.
- **HasPrefix(<prefix-string>)** - Check it has the supplied prefix.
- **HasSuffix(<suffix-string>)** - Check it has the supplied suffix.
- **IsNumber** - Check value is numeric.
- **IsBoolean** - Check value is a boolean.

## Further Documentation in the Wiki

Please refer to the [Wiki](https://github.com/Kilemonn/Secrets-Constraints/wiki) for more documentation about how to configure different credential providers and their different usages.
