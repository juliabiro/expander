# expander
A commandline go tool that expands abbreviations based on a predefined list. The abbreviations can be automatically constructed with teh same tool, based on a mapping. 

## What is the problem that I am trying to solve?

In my work I often need to type long strings that only carry a few characters of information. As I can't type too well, I want to be able to type only a few characters and a tool that expands them into the valid input I need. 

## Installation

Check out hit repo:
```
$ git clone git@github.com:juliabiro/expander.git
```

Build binary:
```
$ cd expander
$ make build
```

This generates a binary called `expander`, that you may want to put in your PATH. 

## Use

### Expand abbreviations based on an already existing list

Expand all recognized words based on a previously defined mapping. 
```
$ expander ex "a23z" --config ./example_conf.json

apple-23-z
```

### Generate abbreviations

Generate consistent abbreviations of long expressions, based on a predefined list of abbreviations to apply.

Example:
```
$ expander map "apple-23-z" --config ./example_conf.json

Generated Abbreviations:
a23z: apple-23-z
```

The generated abbreviations can be saved to the same configfile for later use. 

### Configuration

`Expander` uses a single configuration file in json format, that contains both the mapping between short and log strings (for expanding) and the abbreviation rules to apply when mapping new long strings. The configfile contains 3 sections:
- `AbbreviationRules` for the rules based on which mappings can be generated
- `GeneratedConfig` into which generated mappings are saved
- `CustomConfig` where the user can define extra mappings, that are never overwritten by generated mappings. 

The last 2 sections need to be consistent with each other (they cannot contain the same key with different values.)

See `./example_conf.json` as an example. 

The configfile needs to be specified both for mapping and expanding, either with the --configfile flag, or by the EXPANDER_CONF environmental variable. 

### Generating and saving mappings

Generated mappings can be saved back to the same configfile that contained the abbreviation rules based on which they were created. 
By default, the generated mapping is not saved. When the `--dry-run` flag is set to true, the newly generated mapping is added to the existing one (so keys that are already there are overwitten, but other key value-pairs remain unchanged). 

Generating a mapping and saving it back to the config-file:
```
$ expander map "apple-23-z" --config ./example_conf.json --dry-run=false

Generated Abbreviations:
a23z: apple-23-z

Mapping saved to ./example_conf.json
```

The `--clear-existing-conf` flag can be used to disregard any generated configuration previously found in the file (regardless of whether the result is saved or not).

#### How are the abbreviations mapping used?

All occurances of the longer string are replaced with the shorter version. 
You can map any long string anything to an empty string, this way simply removing it. 
There is no regex support, there is only a simple string match. 
The abbreviations are executed in the order you define them, so if you have abbreviations to expressions that are prefixes to each other, make sure to specify the abbreviation for the more specific first. 

## Example workflow: expanding kubernetes contexts

1. generate a space-separated list of kubernetes contexts
```
kubectl config get-contexts --no-headers=true|tr -s " "|cut -d " " -f2
```
``` 
production-001-domain1.com
production-001-domain2.com 
staging-001-domain1.com 
staging-002-domain1.com
```

2. Create a configfile

See `example_conf.json` as an example. 

3. Generate the abbreviations and save them to a file
```
$ expander map `kubectl config get-contexts --no-headers=true|tr -s " "|cut -d " " -f2` --config ./example_conf.json
```

```
Generated Abbreviations:
a23z: apple-23-z
p01d1: production-001-domain1.com
p01d2: production-001-domain2.com
s01d1: staging-001-domain1.com
d02d1: staging-002-domain1.com

Mapping not saved. To save, use the --dry-run=false flag.
```

As you see, the mapping that was already in the file is still listed in the generated abbreviations. You can remove it with the `--clear-existing-conf` flag.
Check the output, and if it's right, rerun the command, this time `--dry-run=false` flag. 

```
$ expander map `kubectl config get-contexts --no-headers=true|tr -s " "|cut -d " " -f2` --config ./example_conf.json --dry-run=false
```

4. Use the expander to expand the context name
```
$ export  EXPANDER_CONF=./example_conf.json # this actually needs to be an absolute path
$ expander ex "d02dl"

staging-002-domain1.com
```

5. Create a shell function that uses the expander when `kubectl` is called, and make an alias for it (eg in your .aliases file). 
(This happens to be fish shell, not bash, but you get the gist)

```
function kubectl_context
    kubectl $argv[1..-2] --context ( expander ex "$argv[-1]" )
end

set -x EXPANED_CONF "/path/to/my/configfile"
alias k='kubectl_context'
```

If you source this, you can make commands like 
```
k get pods p01d1
```

And it will execute 
``` 
kubectl get pods --context production-001-domain1.com
```

Just don't forget to rerun the map generation when the set of contexts you have access to changes. 


## future work

- support for shell autocompletion
