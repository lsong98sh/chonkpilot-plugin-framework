// Package topics 收口插件侧需要引用的**宿主消息主题常量**（值必须与宿主
// `docs/spec/60-reference/61-消息一览.md` 逐字一致；插件不得自定义主题）。
//
// 2026-10-25（PLUGIN-SPLIT-1）：插件仓不得依赖主仓 `chonkpilot-lib/msgkeys`，
// 故把插件实际引用到的少量主题常量内化到本包（宿主侧仍是 msgkeys 单一生成源）。
package topics

const (
	// TopicDataPrjConfigLoad 读项目配置单键（请求-应答，应答回到同主题）。
	TopicDataPrjConfigLoad = "data-prj-config-load"
	// TopicDataPrjConfigSave 写项目配置单键（请求-应答，应答回到同主题）。
	TopicDataPrjConfigSave = "data-prj-config-save"
	// TopicSessionComplete 一轮会话结束（宿主广播，插件订阅）。
	TopicSessionComplete = "session-complete"
	// TopicToolNotify 工具/插件进度通知面（插件发布，前端与宿主订阅）。
	TopicToolNotify = "tool-notify"
)
