package com.topview.api;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.58.0)",
    comments = "Source: blc.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BlcRpcServiceGrpc {

  private BlcRpcServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "com.topview.api.BlcRpcService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.topview.api.SignUpResponse> getSignUpMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SignUp",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.topview.api.SignUpResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.topview.api.SignUpResponse> getSignUpMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.topview.api.SignUpResponse> getSignUpMethod;
    if ((getSignUpMethod = BlcRpcServiceGrpc.getSignUpMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getSignUpMethod = BlcRpcServiceGrpc.getSignUpMethod) == null) {
          BlcRpcServiceGrpc.getSignUpMethod = getSignUpMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.topview.api.SignUpResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SignUp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.SignUpResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("SignUp"))
              .build();
        }
      }
    }
    return getSignUpMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.UserBalanceRequest,
      com.topview.api.UserBalanceResponse> getGetUserBalanceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserBalance",
      requestType = com.topview.api.UserBalanceRequest.class,
      responseType = com.topview.api.UserBalanceResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.UserBalanceRequest,
      com.topview.api.UserBalanceResponse> getGetUserBalanceMethod() {
    io.grpc.MethodDescriptor<com.topview.api.UserBalanceRequest, com.topview.api.UserBalanceResponse> getGetUserBalanceMethod;
    if ((getGetUserBalanceMethod = BlcRpcServiceGrpc.getGetUserBalanceMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetUserBalanceMethod = BlcRpcServiceGrpc.getGetUserBalanceMethod) == null) {
          BlcRpcServiceGrpc.getGetUserBalanceMethod = getGetUserBalanceMethod =
              io.grpc.MethodDescriptor.<com.topview.api.UserBalanceRequest, com.topview.api.UserBalanceResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserBalance"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.UserBalanceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.UserBalanceResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetUserBalance"))
              .build();
        }
      }
    }
    return getGetUserBalanceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.topview.api.ActivityAmountResponse> getGetActivityAmountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetActivityAmount",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.topview.api.ActivityAmountResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.topview.api.ActivityAmountResponse> getGetActivityAmountMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.topview.api.ActivityAmountResponse> getGetActivityAmountMethod;
    if ((getGetActivityAmountMethod = BlcRpcServiceGrpc.getGetActivityAmountMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetActivityAmountMethod = BlcRpcServiceGrpc.getGetActivityAmountMethod) == null) {
          BlcRpcServiceGrpc.getGetActivityAmountMethod = getGetActivityAmountMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.topview.api.ActivityAmountResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActivityAmount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.ActivityAmountResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetActivityAmount"))
              .build();
        }
      }
    }
    return getGetActivityAmountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.CreateActivityRequest,
      com.google.protobuf.Empty> getCreateActivityMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateActivity",
      requestType = com.topview.api.CreateActivityRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.CreateActivityRequest,
      com.google.protobuf.Empty> getCreateActivityMethod() {
    io.grpc.MethodDescriptor<com.topview.api.CreateActivityRequest, com.google.protobuf.Empty> getCreateActivityMethod;
    if ((getCreateActivityMethod = BlcRpcServiceGrpc.getCreateActivityMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getCreateActivityMethod = BlcRpcServiceGrpc.getCreateActivityMethod) == null) {
          BlcRpcServiceGrpc.getCreateActivityMethod = getCreateActivityMethod =
              io.grpc.MethodDescriptor.<com.topview.api.CreateActivityRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateActivity"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.CreateActivityRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("CreateActivity"))
              .build();
        }
      }
    }
    return getCreateActivityMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.GetIdToActivityRequest,
      com.topview.api.ActivityAndPool> getGetIdToActivityMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetIdToActivity",
      requestType = com.topview.api.GetIdToActivityRequest.class,
      responseType = com.topview.api.ActivityAndPool.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.GetIdToActivityRequest,
      com.topview.api.ActivityAndPool> getGetIdToActivityMethod() {
    io.grpc.MethodDescriptor<com.topview.api.GetIdToActivityRequest, com.topview.api.ActivityAndPool> getGetIdToActivityMethod;
    if ((getGetIdToActivityMethod = BlcRpcServiceGrpc.getGetIdToActivityMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetIdToActivityMethod = BlcRpcServiceGrpc.getGetIdToActivityMethod) == null) {
          BlcRpcServiceGrpc.getGetIdToActivityMethod = getGetIdToActivityMethod =
              io.grpc.MethodDescriptor.<com.topview.api.GetIdToActivityRequest, com.topview.api.ActivityAndPool>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetIdToActivity"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GetIdToActivityRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.ActivityAndPool.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetIdToActivity"))
              .build();
        }
      }
    }
    return getGetIdToActivityMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.BeforeMintRequest,
      com.topview.api.BeforeMintDTO> getBeforeMintMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "BeforeMint",
      requestType = com.topview.api.BeforeMintRequest.class,
      responseType = com.topview.api.BeforeMintDTO.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.BeforeMintRequest,
      com.topview.api.BeforeMintDTO> getBeforeMintMethod() {
    io.grpc.MethodDescriptor<com.topview.api.BeforeMintRequest, com.topview.api.BeforeMintDTO> getBeforeMintMethod;
    if ((getBeforeMintMethod = BlcRpcServiceGrpc.getBeforeMintMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getBeforeMintMethod = BlcRpcServiceGrpc.getBeforeMintMethod) == null) {
          BlcRpcServiceGrpc.getBeforeMintMethod = getBeforeMintMethod =
              io.grpc.MethodDescriptor.<com.topview.api.BeforeMintRequest, com.topview.api.BeforeMintDTO>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "BeforeMint"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.BeforeMintRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.BeforeMintDTO.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("BeforeMint"))
              .build();
        }
      }
    }
    return getBeforeMintMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.GetDcFromActivityRequest,
      com.google.protobuf.Empty> getGetDcFromActivityMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDcFromActivity",
      requestType = com.topview.api.GetDcFromActivityRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.GetDcFromActivityRequest,
      com.google.protobuf.Empty> getGetDcFromActivityMethod() {
    io.grpc.MethodDescriptor<com.topview.api.GetDcFromActivityRequest, com.google.protobuf.Empty> getGetDcFromActivityMethod;
    if ((getGetDcFromActivityMethod = BlcRpcServiceGrpc.getGetDcFromActivityMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetDcFromActivityMethod = BlcRpcServiceGrpc.getGetDcFromActivityMethod) == null) {
          BlcRpcServiceGrpc.getGetDcFromActivityMethod = getGetDcFromActivityMethod =
              io.grpc.MethodDescriptor.<com.topview.api.GetDcFromActivityRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDcFromActivity"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GetDcFromActivityRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetDcFromActivity"))
              .build();
        }
      }
    }
    return getGetDcFromActivityMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.GetUserStatusRequest,
      com.topview.api.UserStatusResponse> getGetUserStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserStatus",
      requestType = com.topview.api.GetUserStatusRequest.class,
      responseType = com.topview.api.UserStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.GetUserStatusRequest,
      com.topview.api.UserStatusResponse> getGetUserStatusMethod() {
    io.grpc.MethodDescriptor<com.topview.api.GetUserStatusRequest, com.topview.api.UserStatusResponse> getGetUserStatusMethod;
    if ((getGetUserStatusMethod = BlcRpcServiceGrpc.getGetUserStatusMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetUserStatusMethod = BlcRpcServiceGrpc.getGetUserStatusMethod) == null) {
          BlcRpcServiceGrpc.getGetUserStatusMethod = getGetUserStatusMethod =
              io.grpc.MethodDescriptor.<com.topview.api.GetUserStatusRequest, com.topview.api.UserStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GetUserStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.UserStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetUserStatus"))
              .build();
        }
      }
    }
    return getGetUserStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.CheckDcAndReturnTimeRequest,
      com.topview.api.CheckDcAndReturnTimeOutputDTO> getCheckDcAndReturnTimeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CheckDcAndReturnTime",
      requestType = com.topview.api.CheckDcAndReturnTimeRequest.class,
      responseType = com.topview.api.CheckDcAndReturnTimeOutputDTO.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.CheckDcAndReturnTimeRequest,
      com.topview.api.CheckDcAndReturnTimeOutputDTO> getCheckDcAndReturnTimeMethod() {
    io.grpc.MethodDescriptor<com.topview.api.CheckDcAndReturnTimeRequest, com.topview.api.CheckDcAndReturnTimeOutputDTO> getCheckDcAndReturnTimeMethod;
    if ((getCheckDcAndReturnTimeMethod = BlcRpcServiceGrpc.getCheckDcAndReturnTimeMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getCheckDcAndReturnTimeMethod = BlcRpcServiceGrpc.getCheckDcAndReturnTimeMethod) == null) {
          BlcRpcServiceGrpc.getCheckDcAndReturnTimeMethod = getCheckDcAndReturnTimeMethod =
              io.grpc.MethodDescriptor.<com.topview.api.CheckDcAndReturnTimeRequest, com.topview.api.CheckDcAndReturnTimeOutputDTO>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CheckDcAndReturnTime"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.CheckDcAndReturnTimeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.CheckDcAndReturnTimeOutputDTO.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("CheckDcAndReturnTime"))
              .build();
        }
      }
    }
    return getCheckDcAndReturnTimeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.GetHashToDcIdRequest,
      com.topview.api.GetHashToDcIdResponse> getGetHashToDcIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetHashToDcId",
      requestType = com.topview.api.GetHashToDcIdRequest.class,
      responseType = com.topview.api.GetHashToDcIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.GetHashToDcIdRequest,
      com.topview.api.GetHashToDcIdResponse> getGetHashToDcIdMethod() {
    io.grpc.MethodDescriptor<com.topview.api.GetHashToDcIdRequest, com.topview.api.GetHashToDcIdResponse> getGetHashToDcIdMethod;
    if ((getGetHashToDcIdMethod = BlcRpcServiceGrpc.getGetHashToDcIdMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetHashToDcIdMethod = BlcRpcServiceGrpc.getGetHashToDcIdMethod) == null) {
          BlcRpcServiceGrpc.getGetHashToDcIdMethod = getGetHashToDcIdMethod =
              io.grpc.MethodDescriptor.<com.topview.api.GetHashToDcIdRequest, com.topview.api.GetHashToDcIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetHashToDcId"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GetHashToDcIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GetHashToDcIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetHashToDcId"))
              .build();
        }
      }
    }
    return getGetHashToDcIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.GiveRequest,
      com.google.protobuf.Empty> getGiveMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Give",
      requestType = com.topview.api.GiveRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.GiveRequest,
      com.google.protobuf.Empty> getGiveMethod() {
    io.grpc.MethodDescriptor<com.topview.api.GiveRequest, com.google.protobuf.Empty> getGiveMethod;
    if ((getGiveMethod = BlcRpcServiceGrpc.getGiveMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGiveMethod = BlcRpcServiceGrpc.getGiveMethod) == null) {
          BlcRpcServiceGrpc.getGiveMethod = getGiveMethod =
              io.grpc.MethodDescriptor.<com.topview.api.GiveRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Give"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GiveRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("Give"))
              .build();
        }
      }
    }
    return getGiveMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.GetDcHistoryAndMessageRequest,
      com.topview.api.DcHistoryAndMessageOutputDTO> getGetDcHistoryAndMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDcHistoryAndMessage",
      requestType = com.topview.api.GetDcHistoryAndMessageRequest.class,
      responseType = com.topview.api.DcHistoryAndMessageOutputDTO.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.GetDcHistoryAndMessageRequest,
      com.topview.api.DcHistoryAndMessageOutputDTO> getGetDcHistoryAndMessageMethod() {
    io.grpc.MethodDescriptor<com.topview.api.GetDcHistoryAndMessageRequest, com.topview.api.DcHistoryAndMessageOutputDTO> getGetDcHistoryAndMessageMethod;
    if ((getGetDcHistoryAndMessageMethod = BlcRpcServiceGrpc.getGetDcHistoryAndMessageMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetDcHistoryAndMessageMethod = BlcRpcServiceGrpc.getGetDcHistoryAndMessageMethod) == null) {
          BlcRpcServiceGrpc.getGetDcHistoryAndMessageMethod = getGetDcHistoryAndMessageMethod =
              io.grpc.MethodDescriptor.<com.topview.api.GetDcHistoryAndMessageRequest, com.topview.api.DcHistoryAndMessageOutputDTO>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDcHistoryAndMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.GetDcHistoryAndMessageRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.DcHistoryAndMessageOutputDTO.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetDcHistoryAndMessage"))
              .build();
        }
      }
    }
    return getGetDcHistoryAndMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.topview.api.PoolAmountResponse> getGetPoolAmountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPoolAmount",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.topview.api.PoolAmountResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.topview.api.PoolAmountResponse> getGetPoolAmountMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.topview.api.PoolAmountResponse> getGetPoolAmountMethod;
    if ((getGetPoolAmountMethod = BlcRpcServiceGrpc.getGetPoolAmountMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getGetPoolAmountMethod = BlcRpcServiceGrpc.getGetPoolAmountMethod) == null) {
          BlcRpcServiceGrpc.getGetPoolAmountMethod = getGetPoolAmountMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.topview.api.PoolAmountResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPoolAmount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.PoolAmountResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("GetPoolAmount"))
              .build();
        }
      }
    }
    return getGetPoolAmountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.CreatePoolRequest,
      com.google.protobuf.Empty> getCreatePoolMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreatePool",
      requestType = com.topview.api.CreatePoolRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.CreatePoolRequest,
      com.google.protobuf.Empty> getCreatePoolMethod() {
    io.grpc.MethodDescriptor<com.topview.api.CreatePoolRequest, com.google.protobuf.Empty> getCreatePoolMethod;
    if ((getCreatePoolMethod = BlcRpcServiceGrpc.getCreatePoolMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getCreatePoolMethod = BlcRpcServiceGrpc.getCreatePoolMethod) == null) {
          BlcRpcServiceGrpc.getCreatePoolMethod = getCreatePoolMethod =
              io.grpc.MethodDescriptor.<com.topview.api.CreatePoolRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreatePool"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.CreatePoolRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("CreatePool"))
              .build();
        }
      }
    }
    return getCreatePoolMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.topview.api.MintRequest,
      com.google.protobuf.Empty> getMintMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Mint",
      requestType = com.topview.api.MintRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.topview.api.MintRequest,
      com.google.protobuf.Empty> getMintMethod() {
    io.grpc.MethodDescriptor<com.topview.api.MintRequest, com.google.protobuf.Empty> getMintMethod;
    if ((getMintMethod = BlcRpcServiceGrpc.getMintMethod) == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        if ((getMintMethod = BlcRpcServiceGrpc.getMintMethod) == null) {
          BlcRpcServiceGrpc.getMintMethod = getMintMethod =
              io.grpc.MethodDescriptor.<com.topview.api.MintRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Mint"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.topview.api.MintRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BlcRpcServiceMethodDescriptorSupplier("Mint"))
              .build();
        }
      }
    }
    return getMintMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BlcRpcServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlcRpcServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlcRpcServiceStub>() {
        @java.lang.Override
        public BlcRpcServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlcRpcServiceStub(channel, callOptions);
        }
      };
    return BlcRpcServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BlcRpcServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlcRpcServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlcRpcServiceBlockingStub>() {
        @java.lang.Override
        public BlcRpcServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlcRpcServiceBlockingStub(channel, callOptions);
        }
      };
    return BlcRpcServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BlcRpcServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlcRpcServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlcRpcServiceFutureStub>() {
        @java.lang.Override
        public BlcRpcServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlcRpcServiceFutureStub(channel, callOptions);
        }
      };
    return BlcRpcServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void signUp(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.topview.api.SignUpResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSignUpMethod(), responseObserver);
    }

    /**
     */
    default void getUserBalance(com.topview.api.UserBalanceRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.UserBalanceResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserBalanceMethod(), responseObserver);
    }

    /**
     */
    default void getActivityAmount(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.topview.api.ActivityAmountResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActivityAmountMethod(), responseObserver);
    }

    /**
     */
    default void createActivity(com.topview.api.CreateActivityRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateActivityMethod(), responseObserver);
    }

    /**
     */
    default void getIdToActivity(com.topview.api.GetIdToActivityRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.ActivityAndPool> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetIdToActivityMethod(), responseObserver);
    }

    /**
     */
    default void beforeMint(com.topview.api.BeforeMintRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.BeforeMintDTO> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBeforeMintMethod(), responseObserver);
    }

    /**
     */
    default void getDcFromActivity(com.topview.api.GetDcFromActivityRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDcFromActivityMethod(), responseObserver);
    }

    /**
     */
    default void getUserStatus(com.topview.api.GetUserStatusRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.UserStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserStatusMethod(), responseObserver);
    }

    /**
     */
    default void checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCheckDcAndReturnTimeMethod(), responseObserver);
    }

    /**
     */
    default void getHashToDcId(com.topview.api.GetHashToDcIdRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.GetHashToDcIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetHashToDcIdMethod(), responseObserver);
    }

    /**
     */
    default void give(com.topview.api.GiveRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGiveMethod(), responseObserver);
    }

    /**
     */
    default void getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDcHistoryAndMessageMethod(), responseObserver);
    }

    /**
     */
    default void getPoolAmount(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.topview.api.PoolAmountResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPoolAmountMethod(), responseObserver);
    }

    /**
     */
    default void createPool(com.topview.api.CreatePoolRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreatePoolMethod(), responseObserver);
    }

    /**
     */
    default void mint(com.topview.api.MintRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getMintMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service BlcRpcService.
   */
  public static abstract class BlcRpcServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return BlcRpcServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service BlcRpcService.
   */
  public static final class BlcRpcServiceStub
      extends io.grpc.stub.AbstractAsyncStub<BlcRpcServiceStub> {
    private BlcRpcServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlcRpcServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlcRpcServiceStub(channel, callOptions);
    }

    /**
     */
    public void signUp(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.topview.api.SignUpResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSignUpMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserBalance(com.topview.api.UserBalanceRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.UserBalanceResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserBalanceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getActivityAmount(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.topview.api.ActivityAmountResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActivityAmountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createActivity(com.topview.api.CreateActivityRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateActivityMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getIdToActivity(com.topview.api.GetIdToActivityRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.ActivityAndPool> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetIdToActivityMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void beforeMint(com.topview.api.BeforeMintRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.BeforeMintDTO> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBeforeMintMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getDcFromActivity(com.topview.api.GetDcFromActivityRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDcFromActivityMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserStatus(com.topview.api.GetUserStatusRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.UserStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCheckDcAndReturnTimeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getHashToDcId(com.topview.api.GetHashToDcIdRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.GetHashToDcIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetHashToDcIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void give(com.topview.api.GiveRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGiveMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request,
        io.grpc.stub.StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDcHistoryAndMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPoolAmount(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.topview.api.PoolAmountResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPoolAmountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createPool(com.topview.api.CreatePoolRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreatePoolMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void mint(com.topview.api.MintRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getMintMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service BlcRpcService.
   */
  public static final class BlcRpcServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<BlcRpcServiceBlockingStub> {
    private BlcRpcServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlcRpcServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlcRpcServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.topview.api.SignUpResponse signUp(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSignUpMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.UserBalanceResponse getUserBalance(com.topview.api.UserBalanceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserBalanceMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.ActivityAmountResponse getActivityAmount(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActivityAmountMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty createActivity(com.topview.api.CreateActivityRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateActivityMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.ActivityAndPool getIdToActivity(com.topview.api.GetIdToActivityRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetIdToActivityMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.BeforeMintDTO beforeMint(com.topview.api.BeforeMintRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBeforeMintMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty getDcFromActivity(com.topview.api.GetDcFromActivityRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDcFromActivityMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.UserStatusResponse getUserStatus(com.topview.api.GetUserStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCheckDcAndReturnTimeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.GetHashToDcIdResponse getHashToDcId(com.topview.api.GetHashToDcIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetHashToDcIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty give(com.topview.api.GiveRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGiveMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDcHistoryAndMessageMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.topview.api.PoolAmountResponse getPoolAmount(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPoolAmountMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty createPool(com.topview.api.CreatePoolRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreatePoolMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty mint(com.topview.api.MintRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getMintMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service BlcRpcService.
   */
  public static final class BlcRpcServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<BlcRpcServiceFutureStub> {
    private BlcRpcServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlcRpcServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlcRpcServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.SignUpResponse> signUp(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSignUpMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.UserBalanceResponse> getUserBalance(
        com.topview.api.UserBalanceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserBalanceMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.ActivityAmountResponse> getActivityAmount(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActivityAmountMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> createActivity(
        com.topview.api.CreateActivityRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateActivityMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.ActivityAndPool> getIdToActivity(
        com.topview.api.GetIdToActivityRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetIdToActivityMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.BeforeMintDTO> beforeMint(
        com.topview.api.BeforeMintRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBeforeMintMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> getDcFromActivity(
        com.topview.api.GetDcFromActivityRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDcFromActivityMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.UserStatusResponse> getUserStatus(
        com.topview.api.GetUserStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserStatusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.CheckDcAndReturnTimeOutputDTO> checkDcAndReturnTime(
        com.topview.api.CheckDcAndReturnTimeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCheckDcAndReturnTimeMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.GetHashToDcIdResponse> getHashToDcId(
        com.topview.api.GetHashToDcIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetHashToDcIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> give(
        com.topview.api.GiveRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGiveMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.DcHistoryAndMessageOutputDTO> getDcHistoryAndMessage(
        com.topview.api.GetDcHistoryAndMessageRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDcHistoryAndMessageMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.topview.api.PoolAmountResponse> getPoolAmount(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPoolAmountMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> createPool(
        com.topview.api.CreatePoolRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreatePoolMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> mint(
        com.topview.api.MintRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getMintMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SIGN_UP = 0;
  private static final int METHODID_GET_USER_BALANCE = 1;
  private static final int METHODID_GET_ACTIVITY_AMOUNT = 2;
  private static final int METHODID_CREATE_ACTIVITY = 3;
  private static final int METHODID_GET_ID_TO_ACTIVITY = 4;
  private static final int METHODID_BEFORE_MINT = 5;
  private static final int METHODID_GET_DC_FROM_ACTIVITY = 6;
  private static final int METHODID_GET_USER_STATUS = 7;
  private static final int METHODID_CHECK_DC_AND_RETURN_TIME = 8;
  private static final int METHODID_GET_HASH_TO_DC_ID = 9;
  private static final int METHODID_GIVE = 10;
  private static final int METHODID_GET_DC_HISTORY_AND_MESSAGE = 11;
  private static final int METHODID_GET_POOL_AMOUNT = 12;
  private static final int METHODID_CREATE_POOL = 13;
  private static final int METHODID_MINT = 14;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SIGN_UP:
          serviceImpl.signUp((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.topview.api.SignUpResponse>) responseObserver);
          break;
        case METHODID_GET_USER_BALANCE:
          serviceImpl.getUserBalance((com.topview.api.UserBalanceRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.UserBalanceResponse>) responseObserver);
          break;
        case METHODID_GET_ACTIVITY_AMOUNT:
          serviceImpl.getActivityAmount((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.topview.api.ActivityAmountResponse>) responseObserver);
          break;
        case METHODID_CREATE_ACTIVITY:
          serviceImpl.createActivity((com.topview.api.CreateActivityRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GET_ID_TO_ACTIVITY:
          serviceImpl.getIdToActivity((com.topview.api.GetIdToActivityRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.ActivityAndPool>) responseObserver);
          break;
        case METHODID_BEFORE_MINT:
          serviceImpl.beforeMint((com.topview.api.BeforeMintRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.BeforeMintDTO>) responseObserver);
          break;
        case METHODID_GET_DC_FROM_ACTIVITY:
          serviceImpl.getDcFromActivity((com.topview.api.GetDcFromActivityRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GET_USER_STATUS:
          serviceImpl.getUserStatus((com.topview.api.GetUserStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.UserStatusResponse>) responseObserver);
          break;
        case METHODID_CHECK_DC_AND_RETURN_TIME:
          serviceImpl.checkDcAndReturnTime((com.topview.api.CheckDcAndReturnTimeRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO>) responseObserver);
          break;
        case METHODID_GET_HASH_TO_DC_ID:
          serviceImpl.getHashToDcId((com.topview.api.GetHashToDcIdRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.GetHashToDcIdResponse>) responseObserver);
          break;
        case METHODID_GIVE:
          serviceImpl.give((com.topview.api.GiveRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GET_DC_HISTORY_AND_MESSAGE:
          serviceImpl.getDcHistoryAndMessage((com.topview.api.GetDcHistoryAndMessageRequest) request,
              (io.grpc.stub.StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO>) responseObserver);
          break;
        case METHODID_GET_POOL_AMOUNT:
          serviceImpl.getPoolAmount((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.topview.api.PoolAmountResponse>) responseObserver);
          break;
        case METHODID_CREATE_POOL:
          serviceImpl.createPool((com.topview.api.CreatePoolRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_MINT:
          serviceImpl.mint((com.topview.api.MintRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getSignUpMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              com.topview.api.SignUpResponse>(
                service, METHODID_SIGN_UP)))
        .addMethod(
          getGetUserBalanceMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.UserBalanceRequest,
              com.topview.api.UserBalanceResponse>(
                service, METHODID_GET_USER_BALANCE)))
        .addMethod(
          getGetActivityAmountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              com.topview.api.ActivityAmountResponse>(
                service, METHODID_GET_ACTIVITY_AMOUNT)))
        .addMethod(
          getCreateActivityMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.CreateActivityRequest,
              com.google.protobuf.Empty>(
                service, METHODID_CREATE_ACTIVITY)))
        .addMethod(
          getGetIdToActivityMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.GetIdToActivityRequest,
              com.topview.api.ActivityAndPool>(
                service, METHODID_GET_ID_TO_ACTIVITY)))
        .addMethod(
          getBeforeMintMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.BeforeMintRequest,
              com.topview.api.BeforeMintDTO>(
                service, METHODID_BEFORE_MINT)))
        .addMethod(
          getGetDcFromActivityMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.GetDcFromActivityRequest,
              com.google.protobuf.Empty>(
                service, METHODID_GET_DC_FROM_ACTIVITY)))
        .addMethod(
          getGetUserStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.GetUserStatusRequest,
              com.topview.api.UserStatusResponse>(
                service, METHODID_GET_USER_STATUS)))
        .addMethod(
          getCheckDcAndReturnTimeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.CheckDcAndReturnTimeRequest,
              com.topview.api.CheckDcAndReturnTimeOutputDTO>(
                service, METHODID_CHECK_DC_AND_RETURN_TIME)))
        .addMethod(
          getGetHashToDcIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.GetHashToDcIdRequest,
              com.topview.api.GetHashToDcIdResponse>(
                service, METHODID_GET_HASH_TO_DC_ID)))
        .addMethod(
          getGiveMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.GiveRequest,
              com.google.protobuf.Empty>(
                service, METHODID_GIVE)))
        .addMethod(
          getGetDcHistoryAndMessageMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.GetDcHistoryAndMessageRequest,
              com.topview.api.DcHistoryAndMessageOutputDTO>(
                service, METHODID_GET_DC_HISTORY_AND_MESSAGE)))
        .addMethod(
          getGetPoolAmountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              com.topview.api.PoolAmountResponse>(
                service, METHODID_GET_POOL_AMOUNT)))
        .addMethod(
          getCreatePoolMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.CreatePoolRequest,
              com.google.protobuf.Empty>(
                service, METHODID_CREATE_POOL)))
        .addMethod(
          getMintMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.topview.api.MintRequest,
              com.google.protobuf.Empty>(
                service, METHODID_MINT)))
        .build();
  }

  private static abstract class BlcRpcServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BlcRpcServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.topview.api.BlcServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BlcRpcService");
    }
  }

  private static final class BlcRpcServiceFileDescriptorSupplier
      extends BlcRpcServiceBaseDescriptorSupplier {
    BlcRpcServiceFileDescriptorSupplier() {}
  }

  private static final class BlcRpcServiceMethodDescriptorSupplier
      extends BlcRpcServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    BlcRpcServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (BlcRpcServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BlcRpcServiceFileDescriptorSupplier())
              .addMethod(getSignUpMethod())
              .addMethod(getGetUserBalanceMethod())
              .addMethod(getGetActivityAmountMethod())
              .addMethod(getCreateActivityMethod())
              .addMethod(getGetIdToActivityMethod())
              .addMethod(getBeforeMintMethod())
              .addMethod(getGetDcFromActivityMethod())
              .addMethod(getGetUserStatusMethod())
              .addMethod(getCheckDcAndReturnTimeMethod())
              .addMethod(getGetHashToDcIdMethod())
              .addMethod(getGiveMethod())
              .addMethod(getGetDcHistoryAndMessageMethod())
              .addMethod(getGetPoolAmountMethod())
              .addMethod(getCreatePoolMethod())
              .addMethod(getMintMethod())
              .build();
        }
      }
    }
    return result;
  }
}
