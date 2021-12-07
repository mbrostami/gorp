# Gorp
Repl for given cli command  

## Examples

#### 
```shell
> gorp git

git > --help
... 
git > status
...
git > commit
...
```

```shell
> gorp ansible -i inventory/a.yaml

ansible -i inventory/a.yaml > --help
ansible -i inventory/a.yaml > -l something
```

## TODO
[ ] Load autocompletes from config files 
[ ] Store history somewhere
[ ] Specify session name to have separate history
[ ] Search session history
[ ] Search all histories
