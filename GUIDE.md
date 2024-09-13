### 실행 가이드

# macos 에서 실행

-  homebrew 설치
   -  https://brew.sh/ko/
-  docker 설치
   -  터미널에서 `brew install --cask docker` 실행
   ```shell
   brew install --cask docker
   ```
-  docker 실행
   -  launchpad 에서 docker app 실행
-  docker 실행 화면에서 로그인 혹은 로그인없이 시작
-  터미널 실행 - 프로젝트 경로로 이동
-  image build 및 docker container 실행(`docker-compose up -d` 명령어 실행)
   ```shell
   docker-compose up -d
   ```
-  browser 실행 후 접속
   -  http://localhost:8080

# windows 에서 실행

-  chocolatey 설치
   -  https://chocolatey.org/install
-  docker 설치
   -  powershell 관리자권한 실행 후 `choco install docker-desktop` 명령어 실행
   ```shell
   choco install docker-desktop
   ```
-  docker 실행(아래 두 가지 중 한 가지 방법으로 실행)
   -  `시작` 메뉴에서 docker 검색 후 실행
   -  `C:\Program Files\Docker\Docker\Docker Desktop.exe` 실행
-  docker 실행 화면에서 로그인 혹은 로그인없이 시작
-  터미널 실행 - 프로젝트 경로로 이동
-  image build 및 docker container 실행(`docker-compose up -d` 명령어 실행)
   ```shell
   docker-compose up -d
   ```
-  browser 실행 후 접속
   -  http://localhost:8080
