# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from aggregatorpb import aggregator_pb2 as aggregatorpb_dot_aggregator__pb2


class AggregatorStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetWork = channel.unary_stream(
        '/aggregatorpb.Aggregator/GetWork',
        request_serializer=aggregatorpb_dot_aggregator__pb2.GetWorkRequest.SerializeToString,
        response_deserializer=aggregatorpb_dot_aggregator__pb2.Work.FromString,
        )
    self.ReportWork = channel.unary_unary(
        '/aggregatorpb.Aggregator/ReportWork',
        request_serializer=aggregatorpb_dot_aggregator__pb2.ReportWorkRequest.SerializeToString,
        response_deserializer=aggregatorpb_dot_aggregator__pb2.ReportWorkReply.FromString,
        )


class AggregatorServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetWork(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ReportWork(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AggregatorServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetWork': grpc.unary_stream_rpc_method_handler(
          servicer.GetWork,
          request_deserializer=aggregatorpb_dot_aggregator__pb2.GetWorkRequest.FromString,
          response_serializer=aggregatorpb_dot_aggregator__pb2.Work.SerializeToString,
      ),
      'ReportWork': grpc.unary_unary_rpc_method_handler(
          servicer.ReportWork,
          request_deserializer=aggregatorpb_dot_aggregator__pb2.ReportWorkRequest.FromString,
          response_serializer=aggregatorpb_dot_aggregator__pb2.ReportWorkReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'aggregatorpb.Aggregator', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
