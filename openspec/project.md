# Project Context

## Purpose

- 聊天记录工具，面向个人本地数据使用场景，提供“发现微信数据 → 提取密钥 → 解密数据库/多媒体 → 查询/导出/回调”的一体化能力。
- 目标：
  - 本地优先与隐私优先，所有处理在用户设备完成；
  - 跨平台兼容 Windows/macOS，覆盖微信 3.x 与 4.x 版本；
  - 以 TUI、CLI、HTTP API、MCP 四种形态暴露能力，便于人机与 AI 助手集成；
  - 支持多账号与自动解密，提供新消息 Webhook 回调；
  - 尽量零外部服务依赖，部署简单（含 Docker 方案）。

## Tech Stack

- 语言与运行时：Go 1.24（模块化构建，跨平台交叉编译）
- CLI 与配置：`spf13/cobra`（命令行）+ `spf13/viper`（环境变量/文件配置，支持 `$HOME/.chatlog/chatlog.json` 合并与 `CHATLOG_*` 环境变量）
- HTTP 服务：`gin-gonic/gin`（REST API、SSE）
- MCP 协议：`mark3labs/mcp-go`（支持 Streamable HTTP/SSE/stdio 模式；端点包含 `/mcp` 与 `/sse`，消息路由 `/message`）
- TUI：`rivo/tview` + `gdamore/tcell`（终端界面组件与绘制）
- 数据库：`mattn/go-sqlite3`（只读解密后的 WeChat 数据库）
- 文件与系统：`fsnotify/fsnotify`（文件监控）、`gopsutil`（进程与系统信息）
- 多媒体处理：`sjzar/go-silk`（SILK→MP3 实时转码）、`sjzar/go-lame`、`Eyevinn/mp4ff`、自研 `dat2img`（含 `wxgf` 解析）
- 压缩/编码：`klauspost/compress`、`pierrec/lz4/v4`、`klauspost/zstd`
- 序列化：`google.golang.org/protobuf`、`howett.net/plist`
- 日志：`rs/zerolog`（结构化日志）
- 构建与发布：`Makefile`、`Dockerfile`、`docker-compose.yml`

## Project Conventions

### Code Style

- 统一使用 Go 官方风格（`gofmt`/`go vet`），尽量保持包内简洁导出面；
- 包结构：核心逻辑位于 `internal/*`，可复用工具位于 `pkg/*`；
- 平台与版本命名：按 `darwin|windows` + `v3|v4` 后缀划分（例如 `internal/wechat/key/darwin/v4.go`）；
- 日志规范：默认 Info，调试使用 Debug；错误通过 `internal/errors` 统一构造与中间件输出（HTTP/MCP 一致化）；
- 配置优先级：环境变量 > 工作目录/数据目录内配置文件 > 默认值，尽量做到"开箱即用"；
- **TUI 文档同步规则**：
  - 当 `internal/chatlog/app.go` 中的菜单逻辑（`initMenu` 函数）发生变化时，**必须同步更新** `internal/ui/help/help.go` 中的帮助文档；
  - 帮助文档需要准确反映：菜单项顺序、功能描述、操作步骤、版本限制等信息；
  - 特别关注版本兼容性说明（Windows 最高 4.0.3.36，macOS 最高 4.0.3.80）。

### Architecture Patterns

- 分层与按领域/平台组合：
  - `internal/wechat`：进程探测、密钥提取、数据库解密（跨平台 + 版本统一流程）；
  - `internal/wechatdb`：数据源发现与仓库（Repository），提供消息/联系人/群聊/会话/媒体的统一查询接口；
  - `internal/mcp`：最小 MCP 抽象，封装 JSON-RPC、SSE/stdio 与工具接口；
  - `internal/ui`：TUI 组件（菜单/信息栏/表单/样式等）；
  - `internal/errors`：错误类型、HTTP/MCP 错误转码与 Gin 中间件；
  - `pkg/*`：通用工具（文件监控、拷贝缓存、压缩/编码、SILK 转码、`dat2img` 等）；
- DataSource + Repository：
  - DataSource 负责按平台/版本定位解密后的数据库与文件系统事件（用于自动解密与回调）；
  - Repository 聚合并缓存联系人/群聊索引，提供富查询与补全能力；
- 服务形态：
  - TUI/CLI 直接操作；
  - HTTP API（`/api/v1/*`）提供查询与多媒体访问（图片/语音流式转码/文件直出）；
  - MCP 通过 `/mcp`（Streamable HTTP）与 `/sse`（SSE）对接 AI 客户端，结合 `/message` 路由实现 RPC；
  - Webhook 基于文件/消息事件回调（需开启自动解密）。

### Testing Strategy

- 单元测试优先：覆盖工具类方法（时间处理、编解码、路径与缓存逻辑等）与关键解析逻辑（见 `pkg/util/time_test.go` 等）；
- 组件/集成验证：`wechat`/`wechatdb` 涉及真实数据目录与平台差异，采用本地集成测试与手动验证（含不同微信版本/平台样本）；
- 运行命令：`go test ./...`；对性能敏感路径（大文件复制/索引）以基准数据进行对比测试；
- 测试数据：不纳入仓库，需开发者在本地准备合法的自有数据副本；
- 未来补充：增加仓库级 Mock（抽象文件层与进程接口）以覆盖更多跨平台分支。

### Git Workflow

- 分支策略：`main` 保持可发布状态，功能性工作走短生命周期分支（`feature/*`、`fix/*`）；
- 提交信息：推荐使用 Conventional Commits（`feat: `、`fix: `、`refactor: `、`chore: ` 等），范围可使用子模块名（如 `wechatdb:`）；
- 合并策略：建议通过 PR 合并并进行最少两人审核，PR 描述中引用相关 issue 或变更说明；
- 发布：打标签并产出多平台二进制与 Docker 镜像（参见 `Makefile` 与 `Dockerfile`）。

## Domain Context

- 微信版本与平台差异：
  - Windows 3.x/4.x 与 macOS 3.x/4.x 数据目录结构不同；
  - 数据库/多媒体路径、命名规则与加密方式在版本之间存在差异；
  - **版本支持限制**：Windows 最高支持 4.0.3.36，macOS 最高支持 4.0.3.80；
  - 版本检测功能会在解密失败时自动触发，提示用户降级；
- 密钥：数据密钥与图片密钥（支持获取数据与图片密钥：Windows < 4.0.3.36 / macOS < 4.0.3.80）；
- 多媒体：图片（含 `wxgf`）、语音（SILK 实时转 MP3 返回）、视频与文件直出，HTTP 层统一按相对路径映射；
- 业务对象：`talker`（对话方，含私聊/群聊）、`sender`（消息发送者）、消息类型与子类型、会话与时间范围查询；
- 自动解密与回调：文件监控触发新消息解析，按配置的 `webhook.items`（目标 URL、talker、sender、keyword）推送；
- 多账号：支持切换账号上下文，配置文件与数据目录独立隔离；
- 安全建议：HTTP 服务默认无鉴权，仅建议在受信网络运行或自行加前置代理/鉴权。

## Important Constraints

- 合法合规：仅处理用户自己合法拥有或已获授权的数据；严禁用于未授权访问/分析他人数据（详见 `DISCLAIMER.md`）；
- 平台限制：
  - macOS 获取密钥前需临时关闭 SIP（完成后可恢复）；
  - Docker 部署无法获取密钥，需预先在宿主机获取并以 `CHATLOG_DATA_KEY/CHATLOG_IMG_KEY` 配置；
- 权限与兼容：进程内存读取/文件监控在不同系统可能需要额外权限；微信版本升级可能导致路径/格式变化，需要及时适配；
- 资源消耗：大文件复制/解密与实时转码为 IO/CPU 密集型，需合理配置工作目录与缓存策略；
- 网络暴露：服务无内置鉴权与速率限制，不建议直接暴露公网。

## External Dependencies

- 操作系统与应用：Windows/macOS 桌面版微信（3.x/4.x）；
- AI 客户端（可选）：ChatWise、Cherry Studio（SSE 直连 `http://127.0.0.1:5030/sse`）、Claude Desktop、Monica Code（通过 `mcp-proxy`）；
- 容器/编排（可选）：Docker / Docker Compose；
- 第三方 Go 依赖（关键类目）：
  - Web/序列化：`gin`、`protobuf`、`plist`；
  - 存储/FS：`sqlite3`、`fsnotify`；
  - 系统：`gopsutil`；
  - 多媒体/压缩：`go-silk`、`go-lame`、`mp4ff`、`lz4`、`zstd`；
  - 终端与日志：`tview`/`tcell`、`zerolog`。
