# Pause container explained
- Pause container serves as the "parent container" for all the containers in your pod, with 2 responsibilities:
    - the basics of linux namespaces sharing in the pod
    - with PID namespace sharing enabled, it serves as PID = 1 for each pod and release zoombie processes
```text
$ docker ps
CONTAINER ID        IMAGE                           COMMAND ...
...
3b45e983c859        gcr.io/google_containers/pause-amd64:3.0    "/pause" ...
```
## References
- How to use chroot
- Understanding cgroups and namespaces

## 1. Sharing namespace
* ### Create a shell with new namespace
```text
unshare --pid --uts --ipc --mount -f chroot rootfs /bin/sh
```
* ### How to a pod from scratch by using pause container and sharing namespacecs
    - First, start pause container 
    ```text
    docker run -d --name pause --ipc=shareable -p 8080:80 gcr.io/google_containers/pause-amd64:3.0
    ```
    - Then, run containers for our pod. This is set up nginx to proxy requests to its localhost on port 2368
    ```text
    // --- nginx.conf ---
    error_log stderr;
    events { worker_connections  1024; }
    http {
    access_log /dev/stdout combined;
    server {
        listen 80 default_server;
        server_name example.com www.example.com;
        location / {
            proxy_pass http://127.0.0.1:2368;
        }
    }
    }

    docker run -d --name nginx -v `pwd`/nginx.conf:/etc/nginx/nginx.conf --net=container:pause --ipc=container:pause --pid=container:pause nginx
    ```
    - Process can start other process using the fork and exec syscalls 
        - fork syscall is used to start another copy of the running process and exec is used to replace the current process with a new one
        - each process has an entry in the OS process table, this is about the process's state and exit code
        - when child process has finished, its process table entry remains util the parent process has retrieved its exit code using the wait syscall -> this is called "recaping" zombie processes