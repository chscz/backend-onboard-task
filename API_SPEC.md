# API 명세

### 사용자 인증

| method | endpoint  |      task       |
| :----: | :-------- | :-------------: |
|  GET   | /register | 회원가입 페이지 |
|  POST  | /register |    회원가입     |
|  GET   | /login    |  로그인 페이지  |
|  POST  | /login    |     로그인      |
|  POST  | /logout   |    로그아웃     |

### 게시글

| method | endpoint         |        task        |
| :----: | :--------------- | :----------------: |
|  GET   | /                |  게시글 목록 조회  |
|  GET   | /post/create     | 게시글 생성 페이지 |
|  POST  | /post/create     |    게시글 생성     |
|  GET   | /post/update/:id | 게시글 수정 페이지 |
|  POST  | /post/update/:id |    게시글 수정     |
|  GET   | /post/delete/:id |    게시글 삭제     |
|  GET   | /post/detail/:id |    게시글 보기     |

### 상세

## 사용자 인증

-  /register
   -  GET
   -  회원가입 내용을 작성할 수 있는 페이지 요청
-  /register
   -  POST
   -  작성된 내용을 통해 회원가입 요청
   -  요청 파라미터:
      -  Form: `email`, `password`, `password_confirm`, `name`
   -  비밀번호는 `bcrypt` 암호화하여 저장
-  /login
   -  GET
   -  로그인 페이지 요청
-  /login
   -  POST
   -  로그인 요청
   -  로그인 성공시 JWT를 발급받고 쿠키에 저장(토큰 만료: 로그인 시점부터 15분)
   -  요청 파라미터:
      -  Form: `email`, `password`
-  /logout
   -  POST
   -  로그아웃 요청
   -  서버에서 로그아웃하여 쿠키에 저장된 JWT토큰 삭제

## 게시글

-  /
   -  GET
   -  게시글 목록 조회
   -  전체 게시글 목록에 대한 작성일, 수정일, 제목, 작성자를 표시
   -  한 페이지당 10개의 게시글 목록 표시
   -  응답:
      -  성공:
         ```json
         {"title":"게시글 수정하기", "posts":{"title":"제목",...,"view_count":3},...,"totalPostCount":17}
         ```
      -  실패예시:
         ```json
         { "code": 500, "message": "목록 조회 실패" }
         ```
-  /post/create
   -  GET
   -  게시글 작성 페이지 요청
   -  응답: 게시물 작성 페이지로 이동
      -  성공:
         ```json
         "title": "게시글 작성하기"
         ```
-  /post/create
   -  POST
   -  작성한 게시글에 대한 내용 등록 요청
   -  요청 파라미터:
      -  Form: `title`, `content`
   -  응답
      -  성공: "/" 리디렉션
      -  실패예시:
         ```json
         { "code": 400, "message": "제목과 내용을 모두 입력하세요" }
         ```
-  /post/update/:id
   -  GET
   -  요청 파라미터:
      -  Path: `id`(수정할 게시글의 ID)
   -  특정 `id`의 게시글 수정화면 요청
   -  응답: 게시물 수정 페이지로 이동
      -  성공: 게시물 수정 페이지로 이동
         ```json
         { "title": "제목~", ... , "content": "내용~" }
         ```
      -  실패예시:
         ```json
         { "code": 401, "message": "작성자와 불일치" }
         ```
-  /post/update/:id
   -  POST
   -  요청 파라미터:
      -  Path: `id`(수정할 게시글의 ID)
      -  Form: `title`, `content`
   -  특정 `id`의 게시글에 대한 수정 요청
   -  응답:
      -  성공: "/" 리디렉션
      -  실패:
         ```json
         {
            "code": 500,
            "message": "게시글 수정 실패",
            "err_mgs": "storage error"
         }
         ```
-  /post/delete/:id
   -  GET
   -  요청 파라미터:
      -  Path: `id`(삭제할 게시글의 ID)
   -  특정 `id`의 게시글에 대한 삭제 요청
   -  응답:
      -  성공: "/" 리디렉션
      -  실패예시:
         ```json
         { "code": 500, "message": "삭제할 게시글 조회 실패" }
         ```
-  /post/detail/:id
   -  GET
   -  요청 파라미터:
      -  Path: `id`(열람할 게시글의 ID)
   -  특정 `id`의 게시글에 대한 내용 보기 요청
   -  응답:
      -  성공: 게시물 보기 페이지로 이동
         ```json
         { "title":"게시글 보기", "post":{"title":"제목", ... , "content":"내용"} }
         ```
      -  실패예시:
         ```json
         { "code": 500, "message": "게시글 보기 실패" }
         ```
