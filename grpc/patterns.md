# gRPC Communication pattern
- communicate between client and server with simple request/ response style (unary RPC)
- inter-process communication: sever-side streaming, client-side streaming and bidirectional streaming 

## Simple RPC (Unary RPC)
![](../media/grpc_unary.png)
- Define order.proto
```text
syntax = "proto3";
import "google/protobuf/wrappers.proto";

package ecommerce;

service OrderManagement {
    rpc getOrder(google.protobuf.StringValue) returns (*pb.Order, error);
}

message Order {
    string id = 1;
    repeated string items = 2;
    string description = 3;
    float price = 4;
    string destination = 5;
}
```
- server/main.go
```text
func (s *server) GetOrder (ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
    ord, exists := orderMap[orderId.Value]
    if exists {
        return &ord, nil
    }
    return nil, errors.New("Not found")
}
```
- client/main.go
```text
orderMgtClient := pb.NewOrderManagementClient(conn)

order, err := orderMgtClient.GetOrder(ctx, &wrapper.StringValue{ Value: "123" })
```

## Server streaming RPC with Send(&order) method
- server sends back a sequence of responses after getting the client's request message
- this sequence of multiple responses is known as stream 
- end stream by sending the server's status details 
- For example: order search 
![](../media/rpc_server_streaming.png)
- server/order.proto
```text
service OrderManagement {
    rpc searchOrders(google.protobuf.StringValue) returns (stream Order);
}

message Order {
    string id = 1;
    repeated string items = 2;
    string description = 3;
    float price = 4;
    string destination = 5;
}
```  
- server/main.go
```text
func (s *server) SearchOrders(searchQuery *wrappers.StringValue), stream pb.OrderManagement_SearchOrdersServer) error {
    for key, order := range orderMap {
        log.Print(key, order)
        for _, itemStr := range order.Items {
            log.Print(itemStr)
            if strings.Contains(itemStr, searchQuery.Value) {
                err := stream.Send(&order)
                if err != nil {
                    return fmt.Errorf("Error sending message to stream")
                }
    
                break;
            }
        }
    }
    return nil
}
```
- client/main.go
```text
c := pb.NewOrderManagementClient(conn)

searchStream, _ := c.SearchOrders(ctx, &wrappers.StringValue{ Value: "Google" })

for {
    searchOrder, err := searchStream.Recv()
    // end stream
    if err == io.EOF {
        break
    }
    log.Print("Search result: %v", searchResult)
}
```

## Client streaming RPC with Recv() method
- the client send multiple messages to the server
- the server send back a single response to the client 
- For example: updateOrders in the OrderManagement service, send the order list as stream 
![](../media/rpc_client_streaming.png)
- server/order.proto
```text
syntax = "proto3";
import "google/protobuf/wrappers.proto"

service OrderManagement {
    rpc updateOrders(stream Order) returns (google.protobuf.StringValue);
}

message Order {
    string id = 1;
    repeated string items = 2
    string description = 3;
    float price = 4;
    string destination 5;
}
```
- client/main.go
```text
func (s *server) UpdateOrders (stream pb.OrderManagement_UpdateOrders) error {
    orderStr := "Update order IDs: "
    for {
        order, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(
                &wrapper.StringValue{ Value: "Orders processed" + orderStr}
            )
        }
        
        orderMap[order.Id] = *order
        log.Printf("Order ID", order.Id, ": Update")
        ordersStr += order.Id + ", "
    }
}
```
- main.go
```text
c := pb.NewOrderManagementClient(conn)

updateStream, err := c.UpdateOrders(ctx)
if err != nil {
    log.Fatalf("Error update stream: %v", err)
}

// update order 1
if err := updateStream.Send(&updateOrder1); err != nil {
    log.Fatalf("Error update order 1: %v", err)
}

// update order 2
if err := updateStream.Send(&updateOrder2); err != nil {
    log.Fatalf("Error update order 2: %v", err)
}

// end update 
updateRes, err := updateStream.CloseAndRecv()
if err != nil {
    log.Fatalf("Error close stream: %v", err)
}
```
## Bi-directional streaming 
- the client is sending a request to the server as a stream of messages 
- the server also responds with a stream of messages 
![](../media/grpc_bi_directional.png)
- server/order.proto
```text
service OrderManagement {
    rpc processOrders(stream google.protobuf.StringValue) returns (stream CombinedShipment)
}

message Order {
    string id = 1;
    repeated string items = 2;
    string description = 3;
    float price = 4;
    string destination = 5;
}

message CombinedShipment {
    string id = 1;
    string status = 2;
    repeated Order orderList = 3;
}
```
- server/main.go
```text
func (s *server) ProcessOrders(ctx context.Context, stream pb.OrderManagement_ProcessOrdersServer) error {
    ...
    for {
        orderId, err := stream.Recv()
        if err == io.EOF {
            for _, comb := range combinedShipment {
                stream.Send(&comb)
            }
            return nil
        }
        if err != nil {
            return err
        }
        // logic to organize orders into shipments based on the destination
        ...
        
        if batchMarker == orderBatchSize {
            for _, comb := range combinedShipmentMap {
                stream.Send(&comb)
            }   
            batchMarker = 0
            combinedShipmentMap = make(map[string]pb.CombinedShipment)
        } else {
            batchMarker++
        }
    }
}
```
- main.go
```text
// send message
streamOrder, _ := c.ProcessOrders(ctx)
if err := streamOrder.Send(
        &wrapper.StringValue{ Value: "123" }
    ); err != nil {
    log.Fatalf("Error: %v", err)    
}

if err := streamOrder.Send(
        &wrapper.StringValue{ Value: "124" }
    ); err != nil {
    log.Fatalf("Error: %v", err)    
}

if err := streamOrder.Send(
        &wrapper.StringValue{ Value: "125" }
    ); err != nil {
    log.Fatalf("Error: %v", err)    
}

//
channel := make(chan struct{})
go asncClientBidirectionlRPC(streamProcOrder, channel)
time.Sleep(time.Milisecond * 1000)
if err := streamProcOrder.Send(
        &wrapper.StringValue{ Value: "111" }
    ); err != nil {
    log.Fatalf("Error: %v", err)
}

if err := streamProcOrder.CloseSend(); err != nil {
    log.Fatalf("Error: %v", err)
}

<- channel

func asncClientBidirectionalRPC(
    streamProcOrder pb.OrderManagement_ProcessOrdersClient,
    c chan struct{}) {
    for {
        combinedShipment, errProcOrder := streamProcOrder.Recv()
        if errProcOrder == io.EOF {
            break        
        }
        log.Printf("Combined shipment: ", combinedShipment.OrdersList)
    }
}
```
## Use-case for gRPC in overall picture
![](../media/sa_sample.png)