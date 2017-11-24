package toy7

import (
	"bytes"
	"os/exec"
	"syscall"
)

// unix -> fork (vfork, clone)
// go -> linux clone

// fork execve wait exit

// os.Process
type theProcess struct {
	pid    int
	handle uintptr // handle is accessed atomically on Windows
	isdone uint32  // process has been successfully waited on, non zero if true
}

// 低级接口
// func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
// 高级接口  os/exec

// Process -> Kill(SIGKILL) Signal Wait Release
// func (p *Process) Wait() (*ProcessState, error)
// 父进程需要知道子进程何时改变了状态 -- 子进程终止或因受到信号而停止
// Wait 方法阻塞直到进程退出，返回一个 ProcessState 描述进程的状态和可能的错误。
// Wait 方法会释放绑定到 Process 的所有资源
// 大多数操作系统中，Process 必须是当前进程的子进程，否则会返回错误
type theProcessState struct {
	pid    int
	status syscall.WaitStatus // system-dependent status info 记录状态原因 -> Exited() Signaled() CoreDump() Stopped() Continued()
	rusage *syscall.Rusage
}

// Command -> 通过 exe.Command 产生 Cmd 实例

// func (c *Cmd) Start() error -> 执行c包含的命令，不等待命令完成就返回。Wait 方法会返回命令的退出状态码
// 并且在命令执行完毕后释放相关资源。内部调用 os.StartProcess，执行 forkExec

// func (c *Cmd) Wait() error -> 阻塞到命令完成，该命令先通过 Start 执行
// 命令返回后释放资源

// ...

// CMD demo
func DemoFillStd(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out = new(bytes.Buffer)

	cmd.Stdout = out
	cmd.Stderr = out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func DemoUseOutput(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}

func demoUsePipe(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	var out = make([]byte, 0, 1024)
	for {
		tmp := make([]byte, 128)
		n, err := stdout.Read(tmp)
		out = append(out, tmp[:n]...)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return nil, err
	}
	return out, nil
}

// os.Exit -> 系统调用的不是 _exit 而是  exit_group
// Exit 让当前进程以给出的状态码 code 退出。
// 进程会立刻终止，defer 的函数不会被执行。
