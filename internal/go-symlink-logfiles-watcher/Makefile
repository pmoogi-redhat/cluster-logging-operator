all: build

copy_files: 
	cp  ../cmd/go-symlink-logfiles-watcher/*.go Docker/.
	cp  ../pkg/symnotify/*.go Docker/.
	cp  ./go.*  Docker/.

build:  copy_files
	hack/build-component-image.sh Docker  docker.io/cognetive/go-symlink-logfiles-watcher-in-clo_v1
