# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## v0.2.2

### Fixed

- 修复解析邮件内容报错的问题

## v0.2.1

### Changed

- 优化`Parse`、`ParseHeader`、`ParseBody`函数的传参, 进一步封装, 使更易用

## v0.2.0

### Added

- 增加解析附件功能
- 增加 Header: `Reply-To`

## v0.1.1

### Fixed

- 修复针对`multipart/*`类型邮件内容解析不完全的问题

### Changed

- `Header.Date`由`string`类型改为`time.Time`类型
- 优化`Header.MessageID`获取的值: 去掉`<>`

## v0.1.0

### Added

- 支持解析邮件内容类型为`text/*`和`multipart/*`的邮件
- 支持解析中文内容, 比如邮件地址别名、邮件主题、邮件内容中的中文字符
- 支持解析邮件内容经过 `base64` 编码的邮件
- 支持分别解析邮件头和邮件内容, 或一次全部解析出来
