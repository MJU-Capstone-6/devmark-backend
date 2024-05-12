# 1. go 환경 설정
FROM golang:latest AS build

WORKDIR /app

# 2. 빌드하기 위한 사전 환경 설정
COPY go.mod .
COPY go.sum .
RUN go mod download

# 3. 빌드
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o devmark .

# 4. config와 config.yml을 빌드 파일과 동일한 디렉토리로 복사
COPY config config

# 5. 최종 이미지 생성
FROM build

WORKDIR /app

# 필요한 런타임 의존성 복사
COPY --from=build /app/devmark .
COPY --from=build /app/config config

RUN chmod +x ./devmark

EXPOSE 8080 
# 어플리케이션 실행
CMD ["./puddlee","serve"]
