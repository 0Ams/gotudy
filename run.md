# 실행 방법
```go
package main 
import "fmt" 

func main() { 
	fmt.Println("hello world") 
}
```

# 실행 명령어
```bash
$ go run test/hello.go
```

#
Go 프로그래밍을 위해 일반적으로 `Workspace` 폴더 (작업 폴더)를 생성하는데, 이 폴더 안에는 3개의 서브디렉토리 (bin, pkg, src)를 생성합니다.
예를 들어, `~/GoApp 디렉토리`를 Workspace 폴더로 정했다면, `~/GoApp` 안에 bin, pkg, src 서브 폴더를 만들어 줍니다.