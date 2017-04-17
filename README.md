# kce - kubectl Expanded

This is not an official Google product.

A simple wrapper to provide a kubectl experience for interacting with a
Kubernetes cluster that can be expanded and customized by the user.

TL;DR

1. Put `kce` on your path.
2. Create `~/.kce.config`
    ```toml
    [alias]
    ls = "get pods"
    ```
4. `kce get pods` will list pods just as `kubectl get pods` would.
3. `kce ls` will also list pods, just as `kubectl get pods` would.

## How To Create Aliases

There are two ways to expand `kce` commands.

1. Add a command line expansion alias.
2. Add a command alias.

### Command Line Expansion Alias

This is using a short `kce` command to produce the behavior of a long `kubectl`
command.

1. Add a line to the `alias` section of `~/.kce.config`.
    ```toml
    [alias]
    ls = "get pods"
    ```
2. Run the command.
    ```sh
    kce ls
    ```

### Command Alias

Responds to a `kce` command by running some arbitrary code. Only bash is
supported at this time.

1. Put a new alias with some associated bash code in the `~/.kce.config`.
    ```toml
    [alias]
    dt = """bash:
    date
    "
    ```
2. Run the command:
   ```sh
   kce dt
   ```

## Development

### Dependencies

```sh
go get github.com/BurntSushi/toml
go get -v github.com/spf13/cobra/cobra
```

### Build

```sh
go build
```

