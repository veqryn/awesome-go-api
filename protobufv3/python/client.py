from __future__ import print_function

import logging

import grpc
from gen import awesome_pb2
from gen import awesome_pb2_grpc
from grpc_status import rpc_status


def run():
    with grpc.insecure_channel("localhost:8000") as channel:
        stub = awesome_pb2_grpc.DefaultStub(channel)

        print("--- Getting a Greeting:")
        resp = stub.Greeting(awesome_pb2.GreetingReq(name="Bob"))
        print(resp)

        print("--- Sending a Review:")
        resp = stub.Review(awesome_pb2.ReviewReq(
            author="Bob",
            message="foobar",
            rating=4,
        ))
        print(resp)

        print("--- Getting an Error:")
        try:
            resp = stub.Error(awesome_pb2.ErrorReq())
            print(resp)
        except grpc.RpcError as e:
            print(e)
            status = rpc_status.from_call(e)
            for detail in status.details:
                print(detail)


if __name__ == "__main__":
    logging.basicConfig()
    run()
