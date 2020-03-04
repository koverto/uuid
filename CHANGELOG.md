# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][], and this project adheres to
[Semantic Versioning][].

## Unreleased

## v1.3.0 - 2020-03-04

### Added

- `func Parse(s string) (*UUID, error)`

### Changed

- Updated go.mongodb.org/mongo-driver from 1.3.0 to 1.3.1

## v1.2.1 - 2020-02-17

### Fixed

- Implement JSON and GraphQL un/marshaling on the correct struct

## v1.2.0 - 2020-02-17

### Added

- Added JSON support
- Added GraphQL support

## v1.1.0 - 2020-02-12

### Added

- Added BSON support

## v1.0.0 - 2020-02-12

### Added

- `UUID` protobuf message type

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
