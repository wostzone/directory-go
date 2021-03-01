# WoST Directory Service 

Golang implementation of a Directory Service plugin for the WoST Gateway.

## Objective

Provide a WoST Gateway plugin that lets consumers query Thing Descriptions of registered Things.

## Status

This plugin is in the design phase. Currently identifying the WoT requirements for this service.
The issue tracking is used to track design and implementation.

## Audience

This project is aimed at software developers and system implementors with knowledge of operating systems and computing devices.

A Binary distribution of the plugin can be installed and used by users with basic Linux skills.

## Summary

This directory service provides the means to query a list of registered Things.

The [WoT Discovery Specification](https://w3c.github.io/wot-discovery/) describes the requirements and is used to guide this implementation. The intent is to be compliant where possible. Note that at the time of implementation this specification is still a draft and subject to change. Compatibility with future changes would be nice but requires time travel which is out of scope.

The specification covers both service discovery and a directory service. This plugin focuses exclusively on the directory service. For discovery see the '[discovery-go](https://github.com/wostzone/discovery-go)' plugin.

This service is considered to be a Thing and publishes its own TD to discover and configure.


## Installation

### System Requirements

### Install From Package Manager

### Manual Installation

## Usage

Like all WoST plugins, it can be run from the commandline or via the WoST gateway.
The plugin loads the gateway configuration and the plugin configuration to determine how to connect to the gateway and configure plugin specific settings.

### Authentication

The directory service supports authentication and access control. This is essential when needing to access TD's with private information over the public internet.

### Configuration

Coming soon :)


