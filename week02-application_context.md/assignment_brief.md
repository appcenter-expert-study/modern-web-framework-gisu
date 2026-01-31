# 2주차 과제 설명 — ApplicationContext 만들기 (IoC의 본질)

## 이번 주차의 목표

이번 주차의 목표는 DI 문법을 흉내 내는 것이 아닙니다.
객체 생성은 누가 책임지는가?라는 질문에 프레임워크 차원의 답을 만드는 것이 이번 주차의 핵심입니다.

1주차에서 우리는 모든 HTTP 요청이 Dispatcher를 거치도록 만들었습니다
하지만 Dispatcher를 누가 만들고, 언제 만들며, 어떻게 조립하는지는 여전히 사람이 직접 책임지고 있습니다.

이번 주차에서는 이 책임을 프레임워크로 끌어올립니다.

---

## 과제

/week02-application_context/context에 있는 application_context.go에서 아직 구현되지 않은 RegisterBean, GetBean 메서드를 구현하세요