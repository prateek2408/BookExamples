import grpc
import os
import hello_world_pb2_grpc
import hello_world_pb2

def invoke():
    server_addr =  os.getenv("SERVER_URL")
    channel = grpc.insecure_channel(server_addr)
    stub = hello_world_pb2_grpc.MyServiceStub(channel)
    while True:
        response = stub.MyFunc(hello_world_pb2.Request())
        print("Greeter client received: " + response.message)

invoke()
