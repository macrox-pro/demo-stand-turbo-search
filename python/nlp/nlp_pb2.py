# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nlp.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tnlp.proto\"\x13\n\x03\x44oc\x12\x0c\n\x04text\x18\x01 \x01(\t\"{\n\x08\x44ocToken\x12\r\n\x05start\x18\x01 \x01(\r\x12\x0c\n\x04stop\x18\x02 \x01(\r\x12\x0c\n\x04text\x18\x03 \x01(\t\x12\x0b\n\x03pos\x18\x04 \x01(\t\x12\n\n\x02id\x18\x05 \x01(\t\x12\x0f\n\x07head_id\x18\x06 \x01(\t\x12\x0b\n\x03rel\x18\x07 \x01(\t\x12\r\n\x05lemma\x18\x08 \x01(\t\"m\n\x07\x44ocSpan\x12\r\n\x05start\x18\x01 \x01(\r\x12\x0c\n\x04stop\x18\x02 \x01(\r\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x0c\n\x04text\x18\x04 \x01(\t\x12\x0e\n\x06normal\x18\x05 \x01(\t\x12\x19\n\x06tokens\x18\x06 \x03(\x0b\x32\t.DocToken\"h\n\x07\x44ocSent\x12\r\n\x05start\x18\x01 \x01(\r\x12\x0c\n\x04stop\x18\x02 \x01(\r\x12\x0c\n\x04text\x18\x03 \x01(\t\x12\x17\n\x05spans\x18\x04 \x03(\x0b\x32\x08.DocSpan\x12\x19\n\x06tokens\x18\x05 \x03(\x0b\x32\t.DocToken\"U\n\x06Result\x12\x19\n\x06tokens\x18\x01 \x03(\x0b\x32\t.DocToken\x12\x17\n\x05sents\x18\x02 \x03(\x0b\x32\x08.DocSent\x12\x17\n\x05spans\x18\x03 \x03(\x0b\x32\x08.DocSpan2\x1d\n\x03NLP\x12\x16\n\x05Parse\x12\x04.Doc\x1a\x07.ResultB\x0eZ\x0cinternal/nlpb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nlp_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\014internal/nlp'
  _DOC._serialized_start=13
  _DOC._serialized_end=32
  _DOCTOKEN._serialized_start=34
  _DOCTOKEN._serialized_end=157
  _DOCSPAN._serialized_start=159
  _DOCSPAN._serialized_end=268
  _DOCSENT._serialized_start=270
  _DOCSENT._serialized_end=374
  _RESULT._serialized_start=376
  _RESULT._serialized_end=461
  _NLP._serialized_start=463
  _NLP._serialized_end=492
# @@protoc_insertion_point(module_scope)