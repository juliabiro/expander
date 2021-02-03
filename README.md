# expander
A commandline go tool that expands abbreviations bases on a predefined list. The abbreviations can be automatically constructed with teh same tool, based on a mapping. 

## What is the problem that I am tryong to solve?

In my work I often need to type long strings that really only carry a few characters of information. I want to be able to type only a few characters and a tool that expands them into the valid input I need. 

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

You can put the binary in your PATH. 

## Use

### expand abbreviations based on an already existing list

```
$ expander ex "a" --custom-config ./example.conf
```

You can specify 2 config files, with the flags`--generated-config` and `--custom-config` (the latter overwrites the former). Instead of the flags, you can use the `EXPANDER_GENERATED_CONF` and `EXPANDER_CUSTOM_CONF` environment variables. 

The config files need to take the following format:
```
short: long
a: apples
...
```
See example.conf as an example.

The expander only returns the long versions of the strings found in the mapping. For unrecognized strings, nothing is returned. 

### generate abbreviations

You can also use the tool generate consistent abbreviations of long expressions with the `map` command. This needs two inputs:
- the list of expressions to be abbreviated in a space-separated list, provided with the `--expressions` flag
- the list of abbreviations to apply, in a mapping file, provided with the `--abbreviations` flag

Example:
```
$ expander map "apple-23-z" --abbrevations example_mapping
```

The program will print the generated abbreviations list. If you want to, you can save the generated list to a file and use it later for expansion, by specifying the `--generated-congfig` flag or setting the `EXPANDER_GENERATED_CONF` environment variable. 


#### How is the abbreviations mapping used?

You can abbreviate anything to an empty string, this way simply removing it. 
There is no regex support, there is only a simple string match. 
The abbreviations are executed in the order you define them, so if you have abbreviations to expressions that are prefixes to each other, make sure to specify the abbreviation for the more specific first. 

## Example workflow: expanding kubernetes contexts

1. generate a space-separated list of kubernetes contexts
```
kubectl config get-contexts --no-headers=true|tr -s " "|cut -d " " -f2|tr  "\n" " "
```
``` 
production-001-domain1.com
production-001-domain2.com
staging-001-domain1.com
staging-002-domain1.com
```

2. Create a map of abbreviations and save it to a file

For example:
```
d1: domain1.com
d2: domain2.com
p: production
s: staging
-0:
-:
```

3. Generate the abbreviations and save them to a file
```
$ expander map `kubectl config get-contexts --no-headers=true|tr -s " "|cut -d " " -f2|tr  "\n" " "` 
```

```
Generated Abbreviations:
p01d1: production-001-domain1.com
p01d2: production-001-domain2.com
s01d1: staging-001-domain1.com
d02d1: staging-002-domain1.com

Mapping not saved. To save, use the --generated-config flag or set the EXPANDER_GENERATED_CONF env var.
```

Check the output, and if it's right, rerun the command, this time specifying the path to where the map should be saved with the `--generated-config` flag. 

```
$ expander map `kubectl config get-contexts --no-headers=true|tr -s " "|cut -d " " -f2|tr  "\n" " "` --generated-config /path/to/my/configiles 
```

4. Use the expander to expand the context name
```
$ export  EXPANDER_GENERATED_CONF=/path/to/my/configiles 
$ expander ex "d02dl"

staging-002-domain1.com
```

5. Create a function that uses the expander when `kubectl` is called

TBD
