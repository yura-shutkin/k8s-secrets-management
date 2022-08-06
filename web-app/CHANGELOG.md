# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.4.0] - 2022-08-06
### Changed
- Use real logger instead of printing out to stdout
- Dockerfiles moved into single directory to not create several .dockerignores
- Makefile updated

### Added
- Added test data for local development
- Added dockerignore

## [1.3.0] - 2022-085
### Added
- Added /json location for printing data in json format (envs)
- Added /ping location, that will always return static text
- Added /net-check location to check hosts from HOSTS env. Hosts in env should follow template `http://<fqdn>:<port>;http://<fqdn>:<port>`

### Changed
- "Log" in JSON format
