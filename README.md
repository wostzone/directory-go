# WoST Directory Service 

Golang implementation of a Directory Service plugin for the WoST Gateway

## Overview

The directory service provides the means to query a list of registered Things.

The [WoT Directory Service API](https://www.w3.org/TR/2020/WD-wot-discovery-20201124/#exploration-directory-api) describes the requirements for a WoT compatible directory service. This implementation intents to be compliant with these requirements, specifically:

1. HTTP API exposed over HTTPS using TLS
2. Follow RFC7807 to carry error details
3. Provide Thing Descriptions compatible with the TD specifications
4. TBD: OpenAPI specification See also [Issue 82](https://github.com/w3c/wot-discovery/issues/82)

This is a WoST gateway plugin that records TD's from Things that register their TD with the gateway.

## Status

This plugin is in the design phase. Currently identifying the WoT requirements for this service.
The issue tracking is used to track design and implementation.


## Usage

Coming soon :)

## Authentication

Coming soon :)


