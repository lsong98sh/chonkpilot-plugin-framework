package mgr

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// BootEntry 是 plugin.ini 中的一行启动命令。
//
// 约定（D-71）：
//   - 文件固定为 <exeDir>/plugin.ini；
//   - 每行一个可执行命令行（exe [args...]），行序即启动顺序；
//   - 可写在 [bootstrap] 段内；文件里一旦出现任何段头，则只取 [bootstrap] 段；
//   - `#` / `;` 开头为注释，空行跳过；
//   - 相对路径由调用方按 exe 目录解析（本包不猜测工作目录）。
type BootEntry struct {
	// Command 为原始命令行（已去首尾空白）。
	Command string
	// Exe 为命令行首个 token。
	Exe string
	// Args 为其余 token。
	Args []string
	// Line 为在文件中的行号（从 1 开始）。
	Line int
}

// ErrEmptyCommand 表示命令行为空（不应发生，解析阶段已跳过）。
var ErrEmptyCommand = errors.New("plugin.ini: empty command")

// LoadBootstrap 读取 plugin.ini；文件不存在时返回 (nil, nil)。
func LoadBootstrap(path string) ([]BootEntry, error) {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()
	return ParseBootstrap(f)
}

// ParseBootstrap 解析 plugin.ini 内容。
func ParseBootstrap(r io.Reader) ([]BootEntry, error) {
	var (
		entries []BootEntry
		// hasSection 表示文件里出现过段头；section 为当前段名（小写）。
		hasSection bool
		section    string
	)
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	line := 0
	for sc.Scan() {
		line++
		raw := strings.TrimSpace(sc.Text())
		if raw == "" || strings.HasPrefix(raw, "#") || strings.HasPrefix(raw, ";") {
			continue
		}
		if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
			hasSection = true
			section = strings.ToLower(strings.TrimSpace(raw[1 : len(raw)-1]))
			continue
		}
		// 有段头时只认 [bootstrap]。
		if hasSection && section != "bootstrap" {
			continue
		}
		fields, err := splitCommand(raw)
		if err != nil {
			return nil, err
		}
		if len(fields) == 0 {
			continue
		}
		entries = append(entries, BootEntry{
			Command: raw,
			Exe:     fields[0],
			Args:    fields[1:],
			Line:    line,
		})
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

// splitCommand 按空白切分命令行，支持双引号包裹（不处理转义）。
func splitCommand(s string) ([]string, error) {
	var (
		fields  []string
		cur     strings.Builder
		inQuote bool
	)
	flush := func() {
		if cur.Len() > 0 {
			fields = append(fields, cur.String())
			cur.Reset()
		}
	}
	for _, r := range s {
		switch {
		case r == '"':
			inQuote = !inQuote
		case !inQuote && (r == ' ' || r == '\t'):
			flush()
		default:
			cur.WriteRune(r)
		}
	}
	if inQuote {
		return nil, errors.New("plugin.ini: unbalanced quote")
	}
	flush()
	return fields, nil
}
