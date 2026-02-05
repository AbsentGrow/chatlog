package help

import (
	"fmt"

	"github.com/sjzar/chatlog/internal/ui/style"

	"github.com/rivo/tview"
)

const (
	Title     = "help"
	ShowTitle = "帮助"
	Content   = `[yellow]Chatlog 使用指南[white]

[green]基本操作:[white]
• 使用 [yellow]←→[white] 键在主菜单和帮助页面之间切换
• 使用 [yellow]↑↓[white] 键在菜单项之间移动
• 按 [yellow]Enter[white] 选择菜单项
• 按 [yellow]Esc[white] 返回上一级菜单
• 按 [yellow]Ctrl+C[white] 退出程序

[green]主菜单功能:[white]
1. [yellow]检测微信版本[white] - 检测当前系统的微信版本号
   • 显示微信版本和平台信息
   • Windows 支持最高版本: 4.0.3.36
   • macOS 支持最高版本: 4.0.3.80
   • 版本过高时会提示降级建议

2. [yellow]获取密钥[white] - 从进程获取数据密钥 & 图片密钥
   • 自动从微信进程内存中读取加密密钥
   • macOS 需要约 20 秒，期间微信会短暂卡住
   • 获取成功后密钥会自动保存

3. [yellow]解密数据[white] - 解密数据文件
   • 使用获取的密钥解密微信数据库
   • 解密后的文件保存到工作目录
   • 支持聊天记录、联系人、群聊等数据

4. [yellow]启动 HTTP 服务[white] - 启动本地 HTTP & MCP 服务器
   • 默认地址: http://localhost:5030
   • 提供 RESTful API 查询聊天记录
   • 支持 MCP 协议与 AI 助手集成

5. [yellow]开启自动解密[white] - 自动解密新增的数据文件
   • 监控数据目录变化
   • 自动解密新增消息
   • 支持 Webhook 回调通知

6. [yellow]设置[white] - 设置应用程序选项
   • HTTP 服务地址 - 配置监听地址和端口
   • 工作目录 - 设置解密数据存储位置
   • 数据密钥 - 手动配置数据解密密钥
   • 图片密钥 - 手动配置图片解密密钥
   • 数据目录 - 指定微信数据文件位置

7. [yellow]切换账号[white] - 切换当前操作的账号
   • 支持多账号管理
   • 可选择微信进程或历史账号
   • 自动加载对应账号的配置和数据

8. [yellow]退出[white] - 退出程序

[green]快速开始:[white]
[yellow]步骤 1:[white] 确保微信客户端正在运行
   • Windows 最高支持版本: 4.0.3.36
   • macOS 最高支持版本: 4.0.3.80
[yellow]步骤 2:[white] 选择"检测微信版本"检查版本兼容性
[yellow]步骤 3:[white] 选择"获取密钥"从微信进程获取加密密钥
[yellow]步骤 4:[white] 选择"解密数据"解密微信数据库文件
[yellow]步骤 5:[white] 选择"启动 HTTP 服务"启动 API 服务
[yellow]步骤 6:[white] 通过浏览器访问 http://localhost:5030 查看数据

[green]HTTP API 使用:[white]
• 聊天记录: [yellow]GET /api/v1/chatlog?time=2023-01-01&talker=wxid_xxx[white]
• 联系人列表: [yellow]GET /api/v1/contact[white]
• 群聊列表: [yellow]GET /api/v1/chatroom[white]
• 会话列表: [yellow]GET /api/v1/session[white]

[green]MCP 集成:[white]
Chatlog 支持 Model Context Protocol，可与支持 MCP 的 AI 助手集成。
• Streamable HTTP: [yellow]http://localhost:5030/mcp[white]
• SSE: [yellow]http://localhost:5030/sse[white]

[green]常见问题:[white]
• 版本兼容性 - Windows 最高 4.0.3.36，macOS 最高 4.0.3.80
• 版本过高 - 解密失败时会自动检测并提示降级
• 获取密钥失败 - 确保微信程序正在运行且版本兼容
• 解密失败 - 检查密钥是否正确获取，或尝试重新获取
• HTTP 服务启动失败 - 检查端口 5030 是否被占用
• 数据目录和工作目录会自动保存，下次启动时自动加载

[green]数据安全:[white]
• 所有数据处理均在本地完成，不会上传到任何外部服务器
• 请妥善保管解密后的数据，避免隐私泄露
• 建议在安全的网络环境中使用 HTTP 服务
`
)

type Help struct {
	*tview.TextView
	title string
}

func New() *Help {
	help := &Help{
		TextView: tview.NewTextView(),
		title:    Title,
	}

	help.SetDynamicColors(true)
	help.SetRegions(true)
	help.SetWrap(true)
	help.SetTextAlign(tview.AlignLeft)
	help.SetBorder(true)
	help.SetBorderColor(style.BorderColor)
	help.SetTitle(ShowTitle)

	fmt.Fprint(help, Content)

	return help
}
