package toy5

// const (
//     O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
//     O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
//     O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
//     O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
//     O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
//     O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
//     O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
//     O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
// )

// const (
//     // 单字符是被 String 方法用于格式化的属性缩写。
//     ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
//     ModeAppend                                     // a: 只能写入，且只能写入到末尾
//     ModeExclusive                                  // l: 用于执行
//     ModeTemporary                                  // T: 临时文件（非备份文件）
//     ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
//     ModeDevice                                     // D: 设备
//     ModeNamedPipe                                  // p: 命名管道（FIFO）
//     ModeSocket                                     // S: Unix域socket
//     ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
//     ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
//     ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
//     ModeSticky                                     // t: 只有root/创建者能删除/移动文件

//     // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
//     ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//     ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
// )

// Dir() / Base()
