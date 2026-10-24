// Package datafmt 收口插件侧需要的**数据面形状与纯函数**（不需要数据库/门面）：
// 会话消息形状（与宿主 session store 的 JSON 逐字一致）+ token 估算 + llmref 折算。
//
// 2026-10-25（PLUGIN-SPLIT-1）：插件仓不得依赖主仓 `chonkpilot-data`，
// 故把插件实际用到的三件纯逻辑内化到本包（口径与宿主实现逐字等价）。
package datafmt

import (
	"strings"
	"unicode/utf8"
)

// TokenDivisor 是估算口径除数：token 数 ≈ 字符数 / TokenDivisor。
const TokenDivisor = 2

// ChatMsg 是 chat 消息（对齐 OpenAI 协议；快照格式）。
// JSON 字段与宿主 `session/store.ChatMsg` 完全一致（插件按消息面读写）。
type ChatMsg struct {
	Role       string         `json:"role"`
	Kind       string         `json:"kind,omitempty"`
	Content    string         `json:"content"`
	ToolCallID string         `json:"tool_call_id,omitempty"`
	ToolCalls  []ToolCall     `json:"tool_calls,omitempty"`
	Reasoning  string         `json:"reasoning,omitempty"`
	Meta       map[string]any `json:"_meta,omitempty"`
}

// ToolCall 是 LLM 请求的工具调用（快照格式）。
type ToolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

// EstimateTokens 粗略估算文本 token 数（字符数 / TokenDivisor，非精确计数）。
func EstimateTokens(text string) int {
	return utf8.RuneCountInString(text) / TokenDivisor
}

// EstimateTokensOfMessages 估算会话消息 token 总量（正文 + 工具调用名与参数）。
func EstimateTokensOfMessages(msgs []ChatMsg) int {
	n := 0
	for _, m := range msgs {
		n += utf8.RuneCountInString(m.Content)
		for _, tc := range m.ToolCalls {
			n += utf8.RuneCountInString(tc.Function.Name)
			n += utf8.RuneCountInString(tc.Function.Arguments)
		}
	}
	return n / TokenDivisor
}

// LLMRefName 把 llmref 取值折算为 provider name：
//   - 字符串 → 原样（trim）；空串保持空串（= 显式「系统默认（启动参数）」）；
//   - 数字（int / int64 / float64；旧 int 索引）→ `llms[idx]` 的 name；越界 / 负数 → 下标 0；
//   - 其余（键缺失 / nil / 无 llms）→ 同「下标 0」语义；llms 为空 → ""。
func LLMRefName(v any, llms any) string {
	if s, ok := v.(string); ok {
		return strings.TrimSpace(s)
	}
	idx := 0
	switch n := v.(type) {
	case int:
		idx = n
	case int64:
		idx = int(n)
	case float64:
		idx = int(n)
	}
	switch list := llms.(type) {
	case []any:
		if len(list) == 0 {
			return ""
		}
		if idx < 0 || idx >= len(list) {
			idx = 0
		}
		m, _ := list[idx].(map[string]any)
		name, _ := m["name"].(string)
		return strings.TrimSpace(name)
	case []map[string]any:
		if len(list) == 0 {
			return ""
		}
		if idx < 0 || idx >= len(list) {
			idx = 0
		}
		name, _ := list[idx]["name"].(string)
		return strings.TrimSpace(name)
	}
	return ""
}
