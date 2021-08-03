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

> Go 프로그래밍을 위해 일반적으로 `Workspace` 폴더 (작업 폴더)를 생성하는데, 이 폴더 안에는 3개의 서브디렉토리 (bin, pkg, src)를 생성합니다. 예를 들어, `~/GoApp 디렉토리`를 Workspace 폴더로 정했다면, `~/GoApp` 안에 bin, pkg, src 서브 폴더를 만들어 줍니다.

# 다양한 명령어
```sh
$ go build 
```
* 소스 파일 자체의 정보만을 사용하여 Go 바이너리를 빌드합니다. 별도의 makefile은 없습니다.

```sh
go test
```
* 유닛 테스트 및 마이크로벤치마크

```sh
go fmt 
```
* 코드 서식 지정

```sh
go get
```
* 원격 패키지의 검색 및 설치

```sh
go vet 
```
* 코드 내의 잠재적인 오류를 찾아내는 정적 분석기

```sh
go doc 
```
* HTTP를 통해 문서 확인

```sh
go rename 
```
* 변수, 함수 등을 type-safe 방식으로 이름 변경

```sh
go generate 
```
* 코드 생성기를 호출하는 표준 방식