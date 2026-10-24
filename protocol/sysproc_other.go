//go:build !windows

package protocol

import "syscall"

// SysProcAttr 非 Windows 平台无「控制台窗口」概念，返回 nil（不设置进程属性），
// 保证交叉编译与测试不受影响。
func SysProcAttr() *syscall.SysProcAttr {
	return nil
}
