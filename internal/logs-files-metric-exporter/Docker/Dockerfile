FROM openshift/origin-release:golang-1.15 AS build
RUN mkdir -p /logs-files-metric-exporter/cmd/logs-files-metric-exporter
RUN mkdir -p /logs-files-metric-exporter/pkg/symnotify
COPY ./symlink-logfiles-watcher.go  /logs-files-metric-exporter/cmd/logs-files-metric-exporter/.
COPY ./symnotify.go   /logs-files-metric-exporter/pkg/symnotify/.
COPY ./go.mod  /logs-files-metric-exporter/.
COPY ./go.sum  /logs-files-metric-exporter/.
WORKDIR /logs-files-metric-exporter
RUN go get -d -v ./...
RUN go mod vendor
EXPOSE 8080
RUN go build -o logs-files-metric-exporter ./cmd/logs-files-metric-exporter/symlink-logfiles-watcher.go
ENV pathoflogfiles=/var/log/containers/
ENV debugOn=true
ENV port=:2112
CMD echo "./logs-files-metric-exporter -logfilespathname=${pathoflogfiles} -debug=${debugOn} -listeningport=${port}"
CMD ./logs-files-metric-exporter -logfilespathname=${pathoflogfiles} -debug=${debugOn} -listeningport=${port}
