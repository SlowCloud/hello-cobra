# 참고 사이트

- [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md)
    - cobra cli 사용법을 알려줌.
    - 빠르게 cobra repo를 구성할 수 있는 도구
- [cobra user guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)
    - cobra 사용 방법에 대한 상세한 가이드
- [cobra.dev](https://cobra.dev)
    - 코브라 메인 페이지. 아래로 내리면 가이드가 있는데, user guide가 훨씬 상세하다.


# 사용 방법

cobra에서 `main.go`는 `cmd/root.go`의 `Execute` 함수를 실행하는 역할만 한다.
새로운 명령어는 `cmd` 디렉터리에서 `cobra.Command`로 새로 생성하며, `init`에서 `rootCmd`에 명령어를 추가하는 역할만 진행한다.

상세한 설명은 `root.go`에 존재한다.
