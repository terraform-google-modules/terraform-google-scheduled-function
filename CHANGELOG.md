# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][keepachangelog-site],
and this project adheres to [Semantic Versioning][semver-site].

## [1.5.0](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.2...v1.5.0) (2020-08-11)


### Features

* Update project cleaner function to clean up Cloud Endpoints ([#38](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/38)) ([4403dbe](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/4403dbe45e27c86f34928550ad44beff1e92a92b))

### [1.4.2](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.1...v1.4.2) (2020-05-06)


### Bug Fixes

* Add Terraform 0.12 support to bucket_force_destroy boolean variable ([#24](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/24)) ([2f11dfa](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/2f11dfab523d82fa925ddbdf69fe6df8299f98d6))

## [Unreleased]

## [1.4.1] - 2020-01-13

### Fixed

- Fixed issue with scheduler job output value inconsistency. [#29]

## [1.4.0] - 2019-12-19

### Added

- The `scheduler_job` variable, the `scheduler_job` output, and the `pubsub_topic_name` output to enable linking multiple
  instances of the module to the same Cloud Scheduler job. [#15]

## [1.3.0] - 2019-12-18

### Added

- The `function_source_dependent_files` variable is passed on to the `event-function` module's `source_dependent_files` variable. [#28]

### Changed

- The function implementation is provided by the Event Function module. [#6]

## [1.2.0] - 2019-11-20

### Added

- The `function_timeout_s` variable is exposed on the `project_cleanup` submodule.

## [1.1.1] - 2019-11-13

### Fixed

- The IAM module was replaced with IAM member resources to support dynamic members in additive mode. [#22]

## [1.1.0] - 2019-11-11

### Changed

- The `project_cleanup` submodule can be scheduled to remove labelled or unlabelled projects. [#20] [#21]

### Added

- The `logs-slack-alerts` example. [#13]

## [1.0.0] - 2019-07-30

### Changed

- Supported version of Terraform is 0.12. [#11]

## [0.4.1] - 2019-07-03

### Fixed

- Project and region are applied to the scheduler job. [#8]

## [0.4.0] - 2019-06-17

### Added

- A variable which configures the time zone of the scheduler job. [#5]

## [0.3.0] - 2019-06-11

### Added

- Submodule which cleans up old projects.

## [0.2.0] - 2019-04-02

### Added

- Ability to specify a service account for functions to run as

## [0.1.0] - 2019-03-14

### Added

- Initial release

[Unreleased]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.1...HEAD
[1.4.1]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.4.1...v1.0.0
[0.4.1]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/releases/tag/v0.1.0

[#29]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/29
[#28]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/28
[#22]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/22
[#21]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/21
[#20]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/20
[#15]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/15
[#13]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/13
[#11]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/11
[#8]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/8
[#6]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/6
[#5]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/5

[keepachangelog-site]: https://keepachangelog.com/en/1.0.0/
[semver-site]: https://semver.org/spec/v2.0.0.html
