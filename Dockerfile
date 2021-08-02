FROM registry.cn-shenzhen.aliyuncs.com/mengine/golang:stretch as golang-builder
# 将工作目录指定为与项目代码位置一致
WORKDIR /go/src/com.github.com/jdc
# 将代码从代码库复制到打包环境的WORKDIR
COPY . .
# 将main文件，从cmd中复制到WORKDIR的根目录
COPY cmd/main.go ./main.go


# 设置module模式的go打包环境
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go env -w GONOPROXY="*.code.meikeland.com"
RUN go env -w GOSUMDB="off"
RUN go env -w GOPRIVATE="*.code.meikeland.com"
# 将vendor包优先在打包环节使用
# RUN go env -w GOFLAGS="-mod=vendor"

RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app . 

# 采用alpine作为部署镜像的基础环境
FROM registry.cn-shenzhen.aliyuncs.com/mengine/alpine:latest

# 使用国内镜像，加快打包速度
RUN echo "http://mirrors.ustc.edu.cn/alpine/v3.10/main" > /etc/apk/repositories
RUN echo "http://mirrors.ustc.edu.cn/alpine/v3.10/community" >> /etc/apk/repositories
RUN apk --no-cache add ca-certificates

# 将程序部署在镜像的/root/目录下
WORKDIR /root/

# 复制部署内容
COPY --from=golang-builder /go/src/code.github.com/jdc/app .


# 设置端口和执行程序
EXPOSE 8062
ENTRYPOINT ["./app"]