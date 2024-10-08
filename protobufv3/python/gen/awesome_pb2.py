# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: awesome.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'awesome.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from protoc_gen_openapiv2.options import annotations_pb2 as protoc__gen__openapiv2_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rawesome.proto\x12\x19\x63om.github.veqryn.awesome\x1a\x19google/protobuf/any.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"@\n\x0bGreetingReq\x12\x31\n\x04name\x18\x01 \x01(\tB\x1d\x92\x41\x1a\x32\rName to greetJ\x07\"world\"x\x1eR\x04name\"a\n\x0cGreetingResp\x12@\n\x07message\x18\x01 \x01(\tB&\x92\x41#2\x10Greeting messageJ\x0f\"Hello, world!\"R\x07message:\x0f\x92\x41\x0c\n\n\xd2\x01\x07message\"\xe2\x01\n\tReviewReq\x12\x33\n\x06\x61uthor\x18\x01 \x01(\tB\x1b\x92\x41\x18\x32\x14\x41uthor of the reviewx\nR\x06\x61uthor\x12\x34\n\x07message\x18\x02 \x01(\tB\x15\x92\x41\x12\x32\x0eReview messagexdH\x00R\x07message\x88\x01\x01\x12\x45\n\x06rating\x18\x03 \x01(\x03\x42-\x92\x41*2\x12Rating from 1 to 5Y\x00\x00\x00\x00\x00\x00\x14@i\x00\x00\x00\x00\x00\x00\xf0?\x9a\x02\x01\x03R\x06rating:\x17\x92\x41\x14\n\x12\xd2\x01\x06\x61uthor\xd2\x01\x06ratingB\n\n\x08_message\"\x0c\n\nReviewResp\"\n\n\x08\x45rrorReq\"\x0b\n\tErrorResp2\xdc\x03\n\x07\x44\x65\x66\x61ult\x12\xa8\x01\n\x08Greeting\x12&.com.github.veqryn.awesome.GreetingReq\x1a\'.com.github.veqryn.awesome.GreetingResp\"K\x92\x41\x30\x12\x14Say hello to someone\x1a\x18Responds with a greeting\x82\xd3\xe4\x93\x02\x12\x12\x10/greeting/{name}\x12\x97\x01\n\x06Review\x12$.com.github.veqryn.awesome.ReviewReq\x1a%.com.github.veqryn.awesome.ReviewResp\"@\x92\x41*\x12\rSend a review\x1a\x19Post a review to be saved\x82\xd3\xe4\x93\x02\r\"\x08/reviews:\x01*\x12\x8b\x01\n\x05\x45rror\x12#.com.github.veqryn.awesome.ErrorReq\x1a$.com.github.veqryn.awesome.ErrorResp\"7\x92\x41&\x12\x0cGet an error\x1a\x16Responds with an error\x82\xd3\xe4\x93\x02\x08\x12\x06/errorB\xa3\x02Z2github.com/veqryn/awesome-go-api/protobufv3/go/gen\x92\x41\xeb\x01\x12\xb1\x01\n\x0e\x41wesome GO API\x12MActual example use cases for a curated list of golang api generator libraries*I\n\x0bMIT License\x12:https://github.com/veqryn/awesome-go-api/blob/main/LICENSE2\x05\x31.0.0\x1a\x0elocalhost:8080*\x01\x01\x32\x10\x61pplication/json:\x10\x61pplication/jsonb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'awesome_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z2github.com/veqryn/awesome-go-api/protobufv3/go/gen\222A\353\001\022\261\001\n\016Awesome GO API\022MActual example use cases for a curated list of golang api generator libraries*I\n\013MIT License\022:https://github.com/veqryn/awesome-go-api/blob/main/LICENSE2\0051.0.0\032\016localhost:8080*\001\0012\020application/json:\020application/json'
  _globals['_GREETINGREQ'].fields_by_name['name']._loaded_options = None
  _globals['_GREETINGREQ'].fields_by_name['name']._serialized_options = b'\222A\0322\rName to greetJ\007\"world\"x\036'
  _globals['_GREETINGRESP'].fields_by_name['message']._loaded_options = None
  _globals['_GREETINGRESP'].fields_by_name['message']._serialized_options = b'\222A#2\020Greeting messageJ\017\"Hello, world!\"'
  _globals['_GREETINGRESP']._loaded_options = None
  _globals['_GREETINGRESP']._serialized_options = b'\222A\014\n\n\322\001\007message'
  _globals['_REVIEWREQ'].fields_by_name['author']._loaded_options = None
  _globals['_REVIEWREQ'].fields_by_name['author']._serialized_options = b'\222A\0302\024Author of the reviewx\n'
  _globals['_REVIEWREQ'].fields_by_name['message']._loaded_options = None
  _globals['_REVIEWREQ'].fields_by_name['message']._serialized_options = b'\222A\0222\016Review messagexd'
  _globals['_REVIEWREQ'].fields_by_name['rating']._loaded_options = None
  _globals['_REVIEWREQ'].fields_by_name['rating']._serialized_options = b'\222A*2\022Rating from 1 to 5Y\000\000\000\000\000\000\024@i\000\000\000\000\000\000\360?\232\002\001\003'
  _globals['_REVIEWREQ']._loaded_options = None
  _globals['_REVIEWREQ']._serialized_options = b'\222A\024\n\022\322\001\006author\322\001\006rating'
  _globals['_DEFAULT'].methods_by_name['Greeting']._loaded_options = None
  _globals['_DEFAULT'].methods_by_name['Greeting']._serialized_options = b'\222A0\022\024Say hello to someone\032\030Responds with a greeting\202\323\344\223\002\022\022\020/greeting/{name}'
  _globals['_DEFAULT'].methods_by_name['Review']._loaded_options = None
  _globals['_DEFAULT'].methods_by_name['Review']._serialized_options = b'\222A*\022\rSend a review\032\031Post a review to be saved\202\323\344\223\002\r\"\010/reviews:\001*'
  _globals['_DEFAULT'].methods_by_name['Error']._loaded_options = None
  _globals['_DEFAULT'].methods_by_name['Error']._serialized_options = b'\222A&\022\014Get an error\032\026Responds with an error\202\323\344\223\002\010\022\006/error'
  _globals['_GREETINGREQ']._serialized_start=149
  _globals['_GREETINGREQ']._serialized_end=213
  _globals['_GREETINGRESP']._serialized_start=215
  _globals['_GREETINGRESP']._serialized_end=312
  _globals['_REVIEWREQ']._serialized_start=315
  _globals['_REVIEWREQ']._serialized_end=541
  _globals['_REVIEWRESP']._serialized_start=543
  _globals['_REVIEWRESP']._serialized_end=555
  _globals['_ERRORREQ']._serialized_start=557
  _globals['_ERRORREQ']._serialized_end=567
  _globals['_ERRORRESP']._serialized_start=569
  _globals['_ERRORRESP']._serialized_end=580
  _globals['_DEFAULT']._serialized_start=583
  _globals['_DEFAULT']._serialized_end=1059
# @@protoc_insertion_point(module_scope)
