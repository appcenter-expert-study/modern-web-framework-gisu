# 1주차 과제 설명 — DispatcherServlet 만들기

## 이번 주차의 목표

이번 주차의 목표는 기능 구현이 아닙니다.
웹 프레임워크의 시작점이 되는 단일 요청 진입점(Dispatcher) 구조를 직접 만들어 보는 것입니다.
아직은 컨트롤러도 없고, 라우팅도 없고, 파라미터 처리도 없지만
그럼에도 불구하고 모든 HTTP 요청이 반드시 하나의 지점(Dispatcher)을 거치도록 만드는 것이
이번 주의 핵심입니다.

---

## 이번 주차에서 반드시 지켜야 할 원칙

### 1️⃣ Echo는 “HTTP 엔진”일 뿐이다
- Echo는 서버를 띄우는 역할만 합니다.
- Echo에 비즈니스 로직이 들어가지 않습니다.
- Echo 라우팅은 `/*` 하나만 존재해야 합니다.

### 2️⃣ 모든 요청은 Dispatcher를 거친다
- 어떤 URL을 요청하든, 어떤 HTTP Method를 쓰든, 반드시 `Dispatcher`를 통과해야 합니다.

### 3️⃣ 프레임워크는 자기만의 Context를 가진다
- Echo의 `echo.Context`를 그대로 쓰지 않습니다. `echo.Context`를 우리 프레임워크의 `RequestContext`로 변환합니다. 이 시점부터 Echo는 프레임워크 바깥으로 생각합니다.

---

## 이번 주차에 구현하는 구조 (개념)

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

이 흐름이 만들어지면 이번 주차는 성공입니다.

---

## 구현이 끝났을 때 기대하는 상태

- 어떤 URL로 요청해도 서버가 죽지 않아야 합니다.
- 요청이 올 때마다 Dispatcher 로그가 찍혀야 합니다.
- Echo 코드에는 라우팅 로직이 없어야 합니다.
- Dispatcher는 Echo를 전혀 몰라야 합니다.

---
