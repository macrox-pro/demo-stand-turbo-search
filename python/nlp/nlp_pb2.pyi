from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Doc(_message.Message):
    __slots__ = ["text"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class DocSent(_message.Message):
    __slots__ = ["spans", "start", "stop", "text", "tokens"]
    SPANS_FIELD_NUMBER: _ClassVar[int]
    START_FIELD_NUMBER: _ClassVar[int]
    STOP_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    TOKENS_FIELD_NUMBER: _ClassVar[int]
    spans: _containers.RepeatedCompositeFieldContainer[DocSpan]
    start: int
    stop: int
    text: str
    tokens: _containers.RepeatedCompositeFieldContainer[DocToken]
    def __init__(self, start: _Optional[int] = ..., stop: _Optional[int] = ..., text: _Optional[str] = ..., spans: _Optional[_Iterable[_Union[DocSpan, _Mapping]]] = ..., tokens: _Optional[_Iterable[_Union[DocToken, _Mapping]]] = ...) -> None: ...

class DocSpan(_message.Message):
    __slots__ = ["normal", "start", "stop", "text", "tokens", "type"]
    NORMAL_FIELD_NUMBER: _ClassVar[int]
    START_FIELD_NUMBER: _ClassVar[int]
    STOP_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    TOKENS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    normal: str
    start: int
    stop: int
    text: str
    tokens: _containers.RepeatedCompositeFieldContainer[DocToken]
    type: str
    def __init__(self, start: _Optional[int] = ..., stop: _Optional[int] = ..., type: _Optional[str] = ..., text: _Optional[str] = ..., normal: _Optional[str] = ..., tokens: _Optional[_Iterable[_Union[DocToken, _Mapping]]] = ...) -> None: ...

class DocToken(_message.Message):
    __slots__ = ["head_id", "id", "lemma", "pos", "rel", "start", "stop", "text"]
    HEAD_ID_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    LEMMA_FIELD_NUMBER: _ClassVar[int]
    POS_FIELD_NUMBER: _ClassVar[int]
    REL_FIELD_NUMBER: _ClassVar[int]
    START_FIELD_NUMBER: _ClassVar[int]
    STOP_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    head_id: str
    id: str
    lemma: str
    pos: str
    rel: str
    start: int
    stop: int
    text: str
    def __init__(self, start: _Optional[int] = ..., stop: _Optional[int] = ..., text: _Optional[str] = ..., pos: _Optional[str] = ..., id: _Optional[str] = ..., head_id: _Optional[str] = ..., rel: _Optional[str] = ..., lemma: _Optional[str] = ...) -> None: ...

class Result(_message.Message):
    __slots__ = ["sents", "spans", "tokens"]
    SENTS_FIELD_NUMBER: _ClassVar[int]
    SPANS_FIELD_NUMBER: _ClassVar[int]
    TOKENS_FIELD_NUMBER: _ClassVar[int]
    sents: _containers.RepeatedCompositeFieldContainer[DocSent]
    spans: _containers.RepeatedCompositeFieldContainer[DocSpan]
    tokens: _containers.RepeatedCompositeFieldContainer[DocToken]
    def __init__(self, tokens: _Optional[_Iterable[_Union[DocToken, _Mapping]]] = ..., sents: _Optional[_Iterable[_Union[DocSent, _Mapping]]] = ..., spans: _Optional[_Iterable[_Union[DocSpan, _Mapping]]] = ...) -> None: ...
