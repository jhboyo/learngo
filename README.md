# Learn Go Programming

이 저장소는 Go 프로그래밍 언어를 학습하고 실습하기 위한 프로젝트입니다.

## 프로젝트 소개

이 프로젝트는 Go 언어의 기본 개념부터 실전 응용까지 단계별로 학습할 수 있는 예제와 실습 코드를 포함하고 있습니다. 각 예제는 설명 주석을 포함하여 학습 과정에 도움이 되도록 구성되어 있습니다.

## 목차

- [설치 및 환경 설정](#설치-및-환경-설정)
- [기본 문법](#기본-문법)
- [자료구조](#자료구조)
- [함수와 메서드](#함수와-메서드)
- [구조체와 인터페이스](#구조체와-인터페이스)
- [동시성 프로그래밍](#동시성-프로그래밍)
- [실전 예제](#실전-예제)
- [참고자료](#참고자료)

## 설치 및 환경 설정

### Go 설치하기

1. [Go 공식 웹사이트](https://golang.org/dl/)에서 운영체제에 맞는 Go를 다운로드하여 설치합니다.
2. 설치 확인:
   ```bash
   go version
   ```

### GOPATH 설정

```bash
# Linux/macOS
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Windows
set GOPATH=%USERPROFILE%\go
set PATH=%PATH%;%GOPATH%\bin
```

## 기본 문법

- 변수와 상수
- 기본 데이터 타입
- 조건문과 반복문
- 배열, 슬라이스, 맵

## 자료구조

- 배열과 슬라이스 심화
- 맵 활용
- 포인터
- 사용자 정의 타입

## 함수와 메서드

- 함수 정의와 호출
- 다중 반환값
- 가변 인자
- 익명 함수와 클로저
- 메서드 정의

## 구조체와 인터페이스

- 구조체 정의와 사용
- 임베딩
- 인터페이스 정의
- 타입 단언
- 인터페이스 활용 예제

## 동시성 프로그래밍

- 고루틴(Goroutine)
- 채널(Channel)
- 동기화 기법
- 컨텍스트(Context)
- 동시성 패턴

## 실전 예제

- CLI 애플리케이션
- REST API 서버
- 데이터베이스 연동
- 웹 크롤러
- 파일 처리

## 프로젝트 실행 방법

각 예제는 다음과 같이 실행할 수 있습니다:

```bash
# 특정 예제 실행
cd examples/hello-world
go run main.go

# 또는
go run examples/hello-world/main.go
```

## 테스트 실행

```bash
# 모든 테스트 실행
go test ./...

# 특정 패키지 테스트 실행
go test ./examples/calculator
```

## 참고자료

- [Go 공식 문서](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Go 언어 커뮤니티](https://forum.golangbridge.org/)
- [Effective Go](https://golang.org/doc/effective_go)

## 기여하기

1. 이 저장소를 포크(Fork)합니다.
2. 새로운 브랜치를 생성합니다.
3. 변경사항을 커밋합니다.
4. 브랜치에 푸시합니다.
5. Pull Request를 생성합니다.

## 라이센스

이 프로젝트는 MIT 라이센스를 따릅니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.
