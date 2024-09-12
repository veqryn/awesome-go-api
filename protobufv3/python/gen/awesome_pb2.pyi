from google.protobuf import any_pb2 as _any_pb2
from google.api import annotations_pb2 as _annotations_pb2
from protoc_gen_openapiv2.options import annotations_pb2 as _annotations_pb2_1
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

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
    __slots__ = ("title", "details", "properties")
    class PropertiesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: _any_pb2.Any
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_any_pb2.Any, _Mapping]] = ...) -> None: ...
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DETAILS_FIELD_NUMBER: _ClassVar[int]
    PROPERTIES_FIELD_NUMBER: _ClassVar[int]
    title: str
    details: str
    properties: _containers.MessageMap[str, _any_pb2.Any]
    def __init__(self, title: _Optional[str] = ..., details: _Optional[str] = ..., properties: _Optional[_Mapping[str, _any_pb2.Any]] = ...) -> None: ...
