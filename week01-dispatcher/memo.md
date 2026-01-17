## what is Go?
- 컴파일 언어이고 정적 타입을 가진다.
- 간결한 문법과 빠른 빌드가 강점이다.
- goroutine과 channel로 동시성을 쉽게 다룬다.

## what is mod & sum & main?
- go.mod: 모듈 이름과 의존성 버전을 정의한다.
- go.sum: 의존성 체크섬을 기록해 재현 가능한 빌드를 보장한다.
- main: 실행 파일의 진입점 패키지이며 `main()`에서 시작한다.

## what is bootstrap?
- Go에는 프레임워크 수준의 bootstrap이 없다.
- 이 프로젝트에서는 `bootstrap.Run()`이 조립(assemble)만 담당한다.
    - Engine 생성
    - Dispatcher 생성
    - Engine에 Dispatcher 등록
    - 서버 시작

### How to Use bootstrap.Run()?
- 디렉터리 구조
    week01-dispatcher/
        main.go
        bootstrap/
            bootstrap.go
- main.go에서 `bootstrap.Run()` 호출

## what is dispatcher?
- RequestContext를 받아 처리 책임을 위임/결정하는 계층이다.
- HTTP 구현(Echo 등)에 직접 의존하지 않는다.
- 현재는 요청 로그만 출력한다.

## what is http?
### engine
- Echo를 감싼 래퍼(EchoEngine)다.
- 모든 요청을 Dispatcher로 전달한다.

### request
- Echo Context를 `RequestContext`로 변환한다.
- 현재는 Method, Path만 전달한다.

## what is Echo?
- Echo는 Go용 HTTP 서버/라우팅 프레임워크다.
- 이 프로젝트에서는 Echo를 직접 쓰지 않고 EchoEngine으로 감싼다.
- 기본 기능: 라우팅, 미들웨어, 요청/응답 핸들링, 바인딩/검증.
- net/http 위에서 동작하며 `echo.Context`로 핸들러에 요청 정보를 전달한다.
- 지금 구조에서는 Echo의 라우팅을 쓰지 않고 모든 요청을 단일 엔드포인트로 받아 Dispatcher에 넘긴다.

## EchoEngine / EchoContext / RequestContext
- EchoEngine: Echo 인스턴스를 감싸고 Dispatcher에 연결하는 어댑터다.
- EchoContext: Echo가 제공하는 요청/응답 컨텍스트 타입이다.
- RequestContext: 프레임워크 내부 요청 모델이며 Method, Path만 가진다.

## Current System Architecture
```text
[HTTP 요청]
   ↓
Echo (HTTP 서버)
   ↓
EchoEngine (HTTP 엔진 래퍼)
   ↓
RequestContext 생성 (프레임워크 내부 요청 모델로 변환)
   ↓
Dispatcher.Dispatch() (RequestContext 처리)
```
