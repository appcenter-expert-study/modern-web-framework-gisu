# Go 프로젝트 초기화 & Echo 설치 가이드

아래 순서대로만 따라 하면 1주차 과제를 진행할 수 있습니다.

---

## 1️⃣ Go 설치 확인

터미널(또는 명령 프롬프트)에서 아래 명령어를 실행하세요.

```bash
go version
```

정상 출력 예시:

```text
go version go1.22.0 darwin/arm64
```

- 위와 같이 나오면 Go가 이미 설치된 상태입니다.
- 명령어를 찾을 수 없다고 나오면 Go 설치가 필요합니다.

### Go 설치 방법
- 공식 사이트: https://go.dev/dl
- 운영체제에 맞는 설치 파일 다운로드 후 설치
- 설치 후 다시 `go version`으로 확인

---

## 2️⃣ 프로젝트 폴더로 이동

프로젝트 폴더로 이동합니다.

```bash
cd modern-web-framework
```

⚠️ 반드시 `modern-web-framework` 디렉토리 안에서 다음 단계를 진행해야 합니다.

---

## 3️⃣ Go 프로젝트 초기화 (go.mod 생성)

아래 명령어를 실행하세요.

```bash
go mod init modern-web-framework
```

실행 결과:

```text
go: creating new go.mod: module modern-web-framework
```

이 명령어는:
- 이 폴더를 하나의 Go 프로젝트로 선언하고
- 의존성 관리를 Go에게 맡기기 위한 설정입니다.

📌 `go.mod` 파일이 생성되면 성공입니다.

---

## 4️⃣ Echo 의존성 설치

이번 스터디에서는 Echo를 **HTTP 엔진**으로만 사용합니다.

아래 명령어를 실행하세요.

```bash
go get github.com/labstack/echo/v4
```

실행 후:
- `go.mod`와 `go.sum` 파일이 자동으로 업데이트됩니다.
- 별도의 설정은 필요하지 않습니다.

---

## 5️⃣ 프로젝트 실행

과제 구현이 끝난 후, 아래 명령어로 서버를 실행합니다.

```bash
go run .
```

정상 실행되면:
- 터미널이 에러 없이 대기 상태가 됩니다.
- 브라우저에서 아래 주소로 접속할 수 있습니다.

```text
http://localhost:8080
```


아무 경로로 요청 시 Dispatcher 로그가 출력되면 성공입니다.

---

## 6️⃣ 자주 발생하는 실수

### ❌ echo 설치 후 import 에러 발생

아래 명령어로 해결할 수 있습니다.

```bash
go mod tidy
```

Go가 의존성을 다시 정리해줍니다.

---

## 7️⃣ 지금 단계에서 알면 충분한 Go 개념

- Go는 **폴더 단위로 프로젝트를 관리**합니다.
- `go.mod`는 프로젝트 설정 파일입니다.
- `go run .`로 프로그램을 실행합니다.
- import 경로는 폴더 구조를 기준으로 결정됩니다.

---