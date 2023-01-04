package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "test.txt"
	// func Open(name string) (*File, error)
	fmt.Println("====== Open ==============")
	if file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		rbs := make([]byte, 10)
		total := 0
		for {
			n, err := file.Read(rbs) // 每次读取 len(rbs) 个字节
			if err == io.EOF {
				fmt.Println("\n文件读取完毕:", n)
				break
			} else if err != nil {
				log.Fatal("文件读取失败", err)
			} else {
				total += n
				fmt.Print(string(rbs[:n]))
			}
		}
	}
	//func Environ() []string
	fmt.Println("====== Environ ==============")
	for _, en := range os.Environ() {
		fmt.Println(en)
	}

	//fmt.Println("====== variable ==============")
	fmt.Println("PathSeparator", string(os.PathSeparator))
	fmt.Println("PathListSeparator", string(os.PathListSeparator))
	fmt.Println("ModeDir", os.ModeDir)
	fmt.Println("DevNull", os.DevNull)
	fmt.Println("PathSeparator", os.PathSeparator)
	fmt.Println("Getenv", os.Getenv("term")) // 获取某个系统变量
	fmt.Println("Getegid", os.Getegid())     // 获取调用者的有效组 id
	fmt.Println("Geteuid", os.Geteuid())     // 获取调用者的数字有效用户标识
	fmt.Println("Getgid", os.Getgid())       // 获取调用者的数字组 ID
	fmt.Println("Getpid", os.Getpid())       // 获取调用者父进程的 ID
	fmt.Println("Getppid", os.Getppid())     // 获取调用者父进程的 ID
	fmt.Println("Getuid", os.Getuid())       // 获取调用者的数字用户标识
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Getwd，命令所在目录", wd)
	}
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Hostname，主机名", host)
	}

	// func UserCacheDir() (string, error): 用户缓存目录
	dir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("UserCacheDir，cache目录", dir)
	}
	// func UserConfigDir() (string, error): 用户配置目录
	dir, err = os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("UserConfigDir，配置目录", dir)
	}
	// func UserHomeDir() (string, error): 用户根目录
	dir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("UserHomeDir，根目录", dir)
	}

	//func Getgroups() ([]int, error)
	gs, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Getgroups，用户组", gs)
	}
	//func Getpagesize() int
	fmt.Println("Getpagesize，返回底层系统的内存页面大小", os.Getpagesize())
}

//func Chdir(dir string) error
//func Chmod(name string, mode FileMode) error
//func Chown(name string, uid, gid int) error
//func Chtimes(name string, atime time.Time, mtime time.Time) error
//func DirFS(dir string) fs.FS
//func Executable() (string, error)
//func Expand(s string, mapping func(string) string) string
//func ExpandEnv(s string) string

//func IsExist(err error) bool
//func IsNotExist(err error) bool
//func IsPathSeparator(c uint8) bool
//func IsPermission(err error) bool
//func IsTimeout(err error) bool
//func Lchown(name string, uid, gid int) error
//func Link(oldname, newname string) error
//func LookupEnv(key string) (string, bool)
//func Mkdir(name string, perm FileMode) error
//func MkdirAll(path string, perm FileMode) error
//func MkdirTemp(dir, pattern string) (string, error)
//func NewSyscallError(syscall string, err error) error
//func Pipe() (r *File, w *File, err error)
//func ReadFile(name string) ([]byte, error)
//func Readlink(name string) (string, error)
//func Remove(name string) error
//func RemoveAll(path string) error
//func Rename(oldpath, newpath string) error
//func SameFile(fi1, fi2 FileInfo) bool
//func Setenv(key, value string) error
//func Symlink(oldname, newname string) error
//func TempDir() string
//func Truncate(name string, size int64) error
//func Unsetenv(key string) error
//func WriteFile(name string, data []byte, perm FileMode) error
//func ReadDir(name string) ([]DirEntry, error)
//func Create(name string) (*File, error)
//func CreateTemp(dir, pattern string) (*File, error)
//func NewFile(fd uintptr, name string) *File
//func Lstat(name string) (FileInfo, error)
//func Stat(name string) (FileInfo, error)
//func FindProcess(pid int) (*Process, error)
//func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

//var Interrupt Signal = syscall.SIGINT ...
//type DirEntry = fs.DirEntry
//type File struct{ ... }
//type FileInfo = fs.FileInfo
//type FileMode = fs.FileMode
//type LinkError struct{ ... }
//type PathError = fs.PathError
//type ProcAttr struct{ ... }
//type Process struct{ ... }
//type ProcessState struct{ ... }
//type Signal interface{ ... }
//type SyscallError struct{ ... }
