# go-skaffold-demo
Simple Demo of Launching Go Server 100% Locally With Skaffold

# Prerequisites

Before running this project, you must have
[Skaffold](https://github.com/GoogleContainerTools/skaffold#installation)
properly installed on your machine.

## Installation

```bash
git clone https://github.com/skaffold-microservices/go-skaffold-demo.git
```

## Running

```bash
cd go-skaffold-demo
skaffold dev
```

After a successful run, you should see an ouput that looks like:

```
[ms-go-hello-78dc7f6669-df8zn ms-go-hello] Starting server at: http://0.0.0.0:30303
[ms-go-hello-78dc7f6669-df8zn ms-go-hello] Meaningless log message #1
[ms-go-hello-78dc7f6669-df8zn ms-go-hello] Meaningless log message #2
[ms-go-hello-78dc7f6669-df8zn ms-go-hello] Meaningless log message #3
[ms-go-hello-78dc7f6669-df8zn ms-go-hello] Meaningless log message #4
[ms-go-hello-78dc7f6669-df8zn ms-go-hello] Meaningless log message #5
```

and the corresponding server should be accessible at `http://0.0.0.0:30303`, as
advertised.

Voila!
