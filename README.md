# WoST Directory Service 

Golang implementation of a Directory Service plugin for the WoST Hub.

## Objective

Provide a WoST Hub plugin that lets consumers query Thing Descriptions of registered Things.

## Status

The status of this plugin is Under Development.



## Audience

This project is aimed at software developers and system implementors with knowledge of operating systems and computing devices.

A Binary distribution of the plugin can be installed and used by users with basic Linux skills.

## Summary

This directory service provides the means to query a list of registered Things.

The [WoT Discovery Specification](https://w3c.github.io/wot-discovery/) describes the requirements and is used to guide this implementation. The intent is to be compliant where possible. Note that at the time of implementation this specification is still a draft and subject to change. Compatibility with future changes would be nice but requires time travel which is out of scope.

The specification covers both service discovery and a directory service. This plugin focuses exclusively on the directory service. For discovery see the '[discovery-go](https://github.com/wostzone/discovery-go)' plugin.

This service is considered to be a Thing and publishes its own TD to discover and configure.



## Build and Installation

### System Requirements

This plugin runs as part of the WoST hub. It has no additional requirements other than a working hub.

### Manual Installation

See the hub README on plugin installation.


### Build From Source

Build with:
```
make all
```

The plugin can be found in dist/bin for 64bit intel or amd processors, or dist/arm for 64 bit ARM processors. Copy this to the hub bin or arm directory.
An example configuration file is provided in config/directory.yaml. Copy this to the hub config directory.


## Usage

Like all WoST plugins, it can be run from the commandline or via the WoST hub.
The plugin loads the hub configuration and the plugin configuration to determine how to connect to the hub and configure plugin specific settings.

Once started consumers can use the service API to query discovered Things and subscribe to events.

### Authentication

The directory service supports authentication and access control. The required authentication method is configured in the directory.yaml configuration file.

TODO: further detailing of authentication.



