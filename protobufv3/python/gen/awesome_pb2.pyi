from google.protobuf import any_pb2 as _any_pb2
from google.api import annotations_pb2 as _annotations_pb2
from protoc_gen_openapiv2.options import annotations_pb2 as _annotations_pb2_1
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class GreetingReq(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class GreetingResp(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class ReviewReq(_message.Message):
    __slots__ = ("author", "message", "rating")
    AUTHOR_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    author: str
    message: str
    rating: int
    def __init__(self, author: _Optional[str] = ..., message: _Optional[str] = ..., rating: _Optional[int] = ...) -> None: ...

class ReviewResp(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ErrorReq(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ErrorResp(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
