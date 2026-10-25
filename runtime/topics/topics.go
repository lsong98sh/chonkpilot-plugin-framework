// Package topics 收口插件侧需要引用的**宿主消息主题常量**（值必须与宿主
// `docs/spec/60-reference/61-消息一览.md` 逐字一致；插件不得自定义主题）。
//
// 2026-10-25（PLUGIN-SPLIT-1）：插件仓不得依赖主仓 `chonkpilot-lib/msgkeys`，
// 故把插件 / 宿主桥实际引用到的主题常量内化到本包（宿主侧仍是 msgkeys 单一生成源）。
package topics

const (
	// ── 会话 / 快照 ──
	TopicDataSessionList         = "data-session-list"
	TopicDataSessionGet          = "data-session-get"
	TopicDataSessionHistory      = "data-session-history"
	TopicDataSessionContent      = "data-session-content"
	TopicDataSessionContext      = "data-session-context"
	TopicDataSessionLoadMessages = "data-session-load-messages"
	TopicDataSessionSetSummary   = "data-session-set-summary"
	TopicDataSnapshotGet         = "data-snapshot-get"
	TopicDataSnapshotSet         = "data-snapshot-set"
	// TopicSessionComplete 一轮会话结束（宿主广播，插件订阅）。
	TopicSessionComplete = "session-complete"

	// ── 中途插话（inbox）──
	TopicLlmInboxList    = "llm-inbox-list"
	TopicLlmInboxAdd     = "llm-inbox-add"
	TopicLlmInboxDiscard = "llm-inbox-discard"
	TopicLlmInboxPromote = "llm-inbox-promote"

	// ── 配置 ──
	// TopicDataPrjConfigLoad 读项目配置单键（请求-应答，应答回到同主题）。
	TopicDataPrjConfigLoad   = "data-prj-config-load"
	TopicDataPrjConfigList   = "data-prj-config-list"
	TopicDataPrjConfigSave   = "data-prj-config-save"
	TopicDataPrjConfigDelete = "data-prj-config-delete"
	TopicDataUserConfigLoad   = "data-user-config-load"
	TopicDataUserConfigSave   = "data-user-config-save"
	TopicDataUserConfigDelete = "data-user-config-delete"
	TopicDataPromptLoad      = "data-prompt-load"
	TopicDataPromptSave      = "data-prompt-save"
	TopicDataFilelistList    = "data-filelist-list"
	TopicDataFilelistPut     = "data-filelist-put"
	TopicDataFilelistDel     = "data-filelist-del"

	// ── 记忆 / 知识 ──
	TopicDataMemoryList        = "data-memory-list"
	TopicDataMemoryRead        = "data-memory-read"
	TopicDataMemorySave        = "data-memory-save"
	TopicDataMemoryExtractLoad = "data-memory-extract-load"
	TopicDataMemoryExtractSave = "data-memory-extract-save"
	TopicDataKnowledgeList     = "data-knowledge-list"
	TopicDataKnowledgeRead     = "data-knowledge-read"

	// ── 任务 ──
	TopicDataTasktreeList = "data-tasktree-list"
	TopicDataTasktreeTasks = "data-tasktree-tasks"
	TopicTaskStop         = "task-stop"
	TopicTaskBackground   = "task-background"
	TopicTaskVerify       = "task-verify"

	// TopicToolNotify 工具/插件进度通知面（插件发布，前端与宿主订阅）。
	TopicToolNotify = "tool-notify"
)
