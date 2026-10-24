//go:build windows

package protocol

import "syscall"

// createNoWindow 为 Windows CREATE_NO_WINDOW（0x08000000）：
// 控制台进程以「无控制台窗口」方式创建，从根源避免控制台被创建/闪现。
const createNoWindow = 0x08000000

// SysProcAttr 返回隐藏子进程控制台窗口的进程创建属性：
// HideWindow（STARTF_USESHOWWINDOW + SW_HIDE）抑制窗口显示，
// CREATE_NO_WINDOW 避免控制台本身被创建——两者并用，覆盖瞬时闪窗。
// 非控制台（GUI 子系统）子进程不受影响；stdin/stdout/stderr 管道通信保持不变。
func SysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: createNoWindow,
	}
}
