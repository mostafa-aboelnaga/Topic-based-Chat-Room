# Topic-based Chat Room
 Topic-based implemented Chat Room using RabbitMQ in Go


<br>

## 😎 **Quick start**

<br>

* ###  First, [download](https://go.dev/dl/) and install **Go**.

* ###  Now, this needs RabbitMQ to be [installed](https://www.rabbitmq.com/download.html) and is running on `localhost` on the [standard port](https://www.rabbitmq.com/networking.html#ports). 

    #### (In case you use a different host, port or credentials, connections settings would require adjusting)

* ### To try out RabbitMQ you may use the [community Docker image](https://registry.hub.docker.com/_/rabbitmq/) as follows:
    ```bash
    # for RabbitMQ 3.9, the latest series
    docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management
    ```

<br>

## Next, make sure you are in the root directory, then just run the following:

<br>

```bash
go run emit_logs_topic.go
```
### and in other tab (instance of your terminal or something)

<br>

```bash
go run receive_logs_topic.go
```

<br>

# Note that

* ### You can push to **only one topic**, you may modify this code to emit to multiple topics, looping on publish, etc.

<br>

* ### **The `emit_logs_topic.go` file can't use** \* OR \# when trying to run it
    ### Here are some variations for the run command:
    `sends/pushes to anonymous.info`

        go run emit_logs_topic.go anonymous.info

    `sends/pushes to kernel.info`
    
        go run emit_logs_topic.go kernel.info

<br>

* ### **For the `receive_logs_topic.go` file**, we can use \* to simulate fetching every topic, etc.
    ### Here are some variations for the run command:
    `receives from any topic having info in the end`

        go run receive_logs_topic.go *.info

    `receives from any topic that starts with anonymous`

        go run receive_logs_topic.go anonymous.*

    `receives from kernel.error topic`

        go run receive_logs_topic.go kernel.error

<br>



## That's all you need to know to start, have fun! ✅

<br>


---

<br>

##  ⚒️ **Built with**

<br>


- [Go](https://go.dev/) - Go is an open source programming language supported by Google. Used to build fast, reliable, and efficient software at scale.

- [Docker](https://docs.docker.com/get-docker/) - Docker is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly. With Docker, you can manage your infrastructure in the same ways you manage your applications. By taking advantage of Docker’s methodologies for shipping, testing, and deploying code quickly, you can significantly reduce the delay between writing code and running it in production.

- [RabbitMQ](https://www.rabbitmq.com/download.html) - RabbitMQ is the most widely deployed open source message broker.

<br> 


## 🚩 [License](https://github.com/mostafa-aboelnaga/Topic-based-Chat-Room/blob/main/LICENSE)

MIT © [Mostafa Aboelnaga](https://github.com/mostafa-aboelnaga/)


