### Pub/Sub Websocket

This is a WebSocket pub/sub implementation for building real-time applications. Pretty simple and effective. Check the client "**pubsub-websocket-client**" folder for explore. This application very straight-forward.

Broadcast will be used for send data to the topic of "orders" and client will subscribe to "orders" to get data

```go
helper.Broadcast <- helper.Message{
    Topic: "orders",
    Data:  order,
}
```

Client side:
```js
useEffect(() => {
    if (readyState === 1) {
        sendMessage(JSON.stringify({ topic: 'orders' }));
    }
}, [readyState, sendMessage]);
```

#### Order request

```http
POST /order HTTP/1.1
Host: 0.0.0.0:8080
Content-Type: application/json
Content-Length: 68

{
  "item": "Aloo paratha",
  "quantity": 6,
  "status": "Pending"
}
```

#### Bill request

```http
POST /bill HTTP/1.1
Host: 0.0.0.0:8080
Content-Type: application/json
Content-Length: 40

{
  "customer_id": 10,
  "amount": 10
}
```

> [!NOTE]  
> Didn't worked will bill endpoint in the react client

#### Tools and version for client

```
npm -v
9.6.7

node --version
v18.17.1
```