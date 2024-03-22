/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

    package com.topview.api;

import org.apache.dubbo.common.stream.StreamObserver;
import org.apache.dubbo.common.URL;
import org.apache.dubbo.rpc.Invoker;
import org.apache.dubbo.rpc.PathResolver;
import org.apache.dubbo.rpc.RpcException;
import org.apache.dubbo.rpc.ServerService;
import org.apache.dubbo.rpc.TriRpcStatus;
import org.apache.dubbo.rpc.model.MethodDescriptor;
import org.apache.dubbo.rpc.model.ServiceDescriptor;
import org.apache.dubbo.rpc.model.StubMethodDescriptor;
import org.apache.dubbo.rpc.model.StubServiceDescriptor;
import org.apache.dubbo.rpc.stub.BiStreamMethodHandler;
import org.apache.dubbo.rpc.stub.ServerStreamMethodHandler;
import org.apache.dubbo.rpc.stub.StubInvocationUtil;
import org.apache.dubbo.rpc.stub.StubInvoker;
import org.apache.dubbo.rpc.stub.StubMethodHandler;
import org.apache.dubbo.rpc.stub.StubSuppliers;
import org.apache.dubbo.rpc.stub.UnaryStubMethodHandler;

import com.google.protobuf.Message;

import java.util.HashMap;
import java.util.Map;
import java.util.function.BiConsumer;
import java.util.concurrent.CompletableFuture;

public final class DubboBlcRpcServiceTriple {

    public static final String SERVICE_NAME = BlcRpcService.SERVICE_NAME;

    private static final StubServiceDescriptor serviceDescriptor = new StubServiceDescriptor(SERVICE_NAME,BlcRpcService.class);

    static {
        org.apache.dubbo.rpc.protocol.tri.service.SchemaDescriptorRegistry.addSchemaDescriptor(SERVICE_NAME,BlcServiceProto.getDescriptor());
        StubSuppliers.addSupplier(SERVICE_NAME, DubboBlcRpcServiceTriple::newStub);
        StubSuppliers.addSupplier(BlcRpcService.JAVA_SERVICE_NAME,  DubboBlcRpcServiceTriple::newStub);
        StubSuppliers.addDescriptor(SERVICE_NAME, serviceDescriptor);
        StubSuppliers.addDescriptor(BlcRpcService.JAVA_SERVICE_NAME, serviceDescriptor);
    }

    @SuppressWarnings("all")
    public static BlcRpcService newStub(Invoker<?> invoker) {
        return new BlcRpcServiceStub((Invoker<BlcRpcService>)invoker);
    }

    private static final StubMethodDescriptor signUpMethod = new StubMethodDescriptor("SignUp",
    com.google.protobuf.Empty.class, com.topview.api.SignUpResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.SignUpResponse::parseFrom);

    private static final StubMethodDescriptor signUpAsyncMethod = new StubMethodDescriptor("SignUp",
    com.google.protobuf.Empty.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.SignUpResponse::parseFrom);

    private static final StubMethodDescriptor signUpProxyAsyncMethod = new StubMethodDescriptor("SignUpAsync",
    com.google.protobuf.Empty.class, com.topview.api.SignUpResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.SignUpResponse::parseFrom);

    private static final StubMethodDescriptor getUserBalanceMethod = new StubMethodDescriptor("GetUserBalance",
    com.topview.api.UserBalanceRequest.class, com.topview.api.UserBalanceResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.UserBalanceRequest::parseFrom,
    com.topview.api.UserBalanceResponse::parseFrom);

    private static final StubMethodDescriptor getUserBalanceAsyncMethod = new StubMethodDescriptor("GetUserBalance",
    com.topview.api.UserBalanceRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.UserBalanceRequest::parseFrom,
    com.topview.api.UserBalanceResponse::parseFrom);

    private static final StubMethodDescriptor getUserBalanceProxyAsyncMethod = new StubMethodDescriptor("GetUserBalanceAsync",
    com.topview.api.UserBalanceRequest.class, com.topview.api.UserBalanceResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.UserBalanceRequest::parseFrom,
    com.topview.api.UserBalanceResponse::parseFrom);

    private static final StubMethodDescriptor getActivityAmountMethod = new StubMethodDescriptor("GetActivityAmount",
    com.google.protobuf.Empty.class, com.topview.api.ActivityAmountResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.ActivityAmountResponse::parseFrom);

    private static final StubMethodDescriptor getActivityAmountAsyncMethod = new StubMethodDescriptor("GetActivityAmount",
    com.google.protobuf.Empty.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.ActivityAmountResponse::parseFrom);

    private static final StubMethodDescriptor getActivityAmountProxyAsyncMethod = new StubMethodDescriptor("GetActivityAmountAsync",
    com.google.protobuf.Empty.class, com.topview.api.ActivityAmountResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.ActivityAmountResponse::parseFrom);

    private static final StubMethodDescriptor createActivityMethod = new StubMethodDescriptor("CreateActivity",
    com.topview.api.CreateActivityRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CreateActivityRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor createActivityAsyncMethod = new StubMethodDescriptor("CreateActivity",
    com.topview.api.CreateActivityRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CreateActivityRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor createActivityProxyAsyncMethod = new StubMethodDescriptor("CreateActivityAsync",
    com.topview.api.CreateActivityRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CreateActivityRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor getIdToActivityMethod = new StubMethodDescriptor("GetIdToActivity",
    com.topview.api.GetIdToActivityRequest.class, com.topview.api.ActivityAndPool.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetIdToActivityRequest::parseFrom,
    com.topview.api.ActivityAndPool::parseFrom);

    private static final StubMethodDescriptor getIdToActivityAsyncMethod = new StubMethodDescriptor("GetIdToActivity",
    com.topview.api.GetIdToActivityRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetIdToActivityRequest::parseFrom,
    com.topview.api.ActivityAndPool::parseFrom);

    private static final StubMethodDescriptor getIdToActivityProxyAsyncMethod = new StubMethodDescriptor("GetIdToActivityAsync",
    com.topview.api.GetIdToActivityRequest.class, com.topview.api.ActivityAndPool.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetIdToActivityRequest::parseFrom,
    com.topview.api.ActivityAndPool::parseFrom);

    private static final StubMethodDescriptor beforeMintMethod = new StubMethodDescriptor("BeforeMint",
    com.topview.api.BeforeMintRequest.class, com.topview.api.BeforeMintDTO.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.BeforeMintRequest::parseFrom,
    com.topview.api.BeforeMintDTO::parseFrom);

    private static final StubMethodDescriptor beforeMintAsyncMethod = new StubMethodDescriptor("BeforeMint",
    com.topview.api.BeforeMintRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.BeforeMintRequest::parseFrom,
    com.topview.api.BeforeMintDTO::parseFrom);

    private static final StubMethodDescriptor beforeMintProxyAsyncMethod = new StubMethodDescriptor("BeforeMintAsync",
    com.topview.api.BeforeMintRequest.class, com.topview.api.BeforeMintDTO.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.BeforeMintRequest::parseFrom,
    com.topview.api.BeforeMintDTO::parseFrom);

    private static final StubMethodDescriptor getDcFromActivityMethod = new StubMethodDescriptor("GetDcFromActivity",
    com.topview.api.GetDcFromActivityRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetDcFromActivityRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor getDcFromActivityAsyncMethod = new StubMethodDescriptor("GetDcFromActivity",
    com.topview.api.GetDcFromActivityRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetDcFromActivityRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor getDcFromActivityProxyAsyncMethod = new StubMethodDescriptor("GetDcFromActivityAsync",
    com.topview.api.GetDcFromActivityRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetDcFromActivityRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor getUserStatusMethod = new StubMethodDescriptor("GetUserStatus",
    com.topview.api.GetUserStatusRequest.class, com.topview.api.UserStatusResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetUserStatusRequest::parseFrom,
    com.topview.api.UserStatusResponse::parseFrom);

    private static final StubMethodDescriptor getUserStatusAsyncMethod = new StubMethodDescriptor("GetUserStatus",
    com.topview.api.GetUserStatusRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetUserStatusRequest::parseFrom,
    com.topview.api.UserStatusResponse::parseFrom);

    private static final StubMethodDescriptor getUserStatusProxyAsyncMethod = new StubMethodDescriptor("GetUserStatusAsync",
    com.topview.api.GetUserStatusRequest.class, com.topview.api.UserStatusResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetUserStatusRequest::parseFrom,
    com.topview.api.UserStatusResponse::parseFrom);

    private static final StubMethodDescriptor checkDcAndReturnTimeMethod = new StubMethodDescriptor("CheckDcAndReturnTime",
    com.topview.api.CheckDcAndReturnTimeRequest.class, com.topview.api.CheckDcAndReturnTimeOutputDTO.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CheckDcAndReturnTimeRequest::parseFrom,
    com.topview.api.CheckDcAndReturnTimeOutputDTO::parseFrom);

    private static final StubMethodDescriptor checkDcAndReturnTimeAsyncMethod = new StubMethodDescriptor("CheckDcAndReturnTime",
    com.topview.api.CheckDcAndReturnTimeRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CheckDcAndReturnTimeRequest::parseFrom,
    com.topview.api.CheckDcAndReturnTimeOutputDTO::parseFrom);

    private static final StubMethodDescriptor checkDcAndReturnTimeProxyAsyncMethod = new StubMethodDescriptor("CheckDcAndReturnTimeAsync",
    com.topview.api.CheckDcAndReturnTimeRequest.class, com.topview.api.CheckDcAndReturnTimeOutputDTO.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CheckDcAndReturnTimeRequest::parseFrom,
    com.topview.api.CheckDcAndReturnTimeOutputDTO::parseFrom);

    private static final StubMethodDescriptor getHashToDcIdMethod = new StubMethodDescriptor("GetHashToDcId",
    com.topview.api.GetHashToDcIdRequest.class, com.topview.api.GetHashToDcIdResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetHashToDcIdRequest::parseFrom,
    com.topview.api.GetHashToDcIdResponse::parseFrom);

    private static final StubMethodDescriptor getHashToDcIdAsyncMethod = new StubMethodDescriptor("GetHashToDcId",
    com.topview.api.GetHashToDcIdRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetHashToDcIdRequest::parseFrom,
    com.topview.api.GetHashToDcIdResponse::parseFrom);

    private static final StubMethodDescriptor getHashToDcIdProxyAsyncMethod = new StubMethodDescriptor("GetHashToDcIdAsync",
    com.topview.api.GetHashToDcIdRequest.class, com.topview.api.GetHashToDcIdResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetHashToDcIdRequest::parseFrom,
    com.topview.api.GetHashToDcIdResponse::parseFrom);

    private static final StubMethodDescriptor giveMethod = new StubMethodDescriptor("Give",
    com.topview.api.GiveRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GiveRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor giveAsyncMethod = new StubMethodDescriptor("Give",
    com.topview.api.GiveRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GiveRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor giveProxyAsyncMethod = new StubMethodDescriptor("GiveAsync",
    com.topview.api.GiveRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GiveRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor getDcHistoryAndMessageMethod = new StubMethodDescriptor("GetDcHistoryAndMessage",
    com.topview.api.GetDcHistoryAndMessageRequest.class, com.topview.api.DcHistoryAndMessageOutputDTO.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetDcHistoryAndMessageRequest::parseFrom,
    com.topview.api.DcHistoryAndMessageOutputDTO::parseFrom);

    private static final StubMethodDescriptor getDcHistoryAndMessageAsyncMethod = new StubMethodDescriptor("GetDcHistoryAndMessage",
    com.topview.api.GetDcHistoryAndMessageRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetDcHistoryAndMessageRequest::parseFrom,
    com.topview.api.DcHistoryAndMessageOutputDTO::parseFrom);

    private static final StubMethodDescriptor getDcHistoryAndMessageProxyAsyncMethod = new StubMethodDescriptor("GetDcHistoryAndMessageAsync",
    com.topview.api.GetDcHistoryAndMessageRequest.class, com.topview.api.DcHistoryAndMessageOutputDTO.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.GetDcHistoryAndMessageRequest::parseFrom,
    com.topview.api.DcHistoryAndMessageOutputDTO::parseFrom);

    private static final StubMethodDescriptor getPoolAmountMethod = new StubMethodDescriptor("GetPoolAmount",
    com.google.protobuf.Empty.class, com.topview.api.PoolAmountResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.PoolAmountResponse::parseFrom);

    private static final StubMethodDescriptor getPoolAmountAsyncMethod = new StubMethodDescriptor("GetPoolAmount",
    com.google.protobuf.Empty.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.PoolAmountResponse::parseFrom);

    private static final StubMethodDescriptor getPoolAmountProxyAsyncMethod = new StubMethodDescriptor("GetPoolAmountAsync",
    com.google.protobuf.Empty.class, com.topview.api.PoolAmountResponse.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.google.protobuf.Empty::parseFrom,
    com.topview.api.PoolAmountResponse::parseFrom);

    private static final StubMethodDescriptor createPoolMethod = new StubMethodDescriptor("CreatePool",
    com.topview.api.CreatePoolRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CreatePoolRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor createPoolAsyncMethod = new StubMethodDescriptor("CreatePool",
    com.topview.api.CreatePoolRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CreatePoolRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor createPoolProxyAsyncMethod = new StubMethodDescriptor("CreatePoolAsync",
    com.topview.api.CreatePoolRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.CreatePoolRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor mintMethod = new StubMethodDescriptor("Mint",
    com.topview.api.MintRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.MintRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor mintAsyncMethod = new StubMethodDescriptor("Mint",
    com.topview.api.MintRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.MintRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);

    private static final StubMethodDescriptor mintProxyAsyncMethod = new StubMethodDescriptor("MintAsync",
    com.topview.api.MintRequest.class, com.google.protobuf.Empty.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), com.topview.api.MintRequest::parseFrom,
    com.google.protobuf.Empty::parseFrom);





    public static class BlcRpcServiceStub implements BlcRpcService{
        private final Invoker<BlcRpcService> invoker;

        public BlcRpcServiceStub(Invoker<BlcRpcService> invoker) {
            this.invoker = invoker;
        }

        @Override
        public com.topview.api.SignUpResponse signUp(com.google.protobuf.Empty request){
            return StubInvocationUtil.unaryCall(invoker, signUpMethod, request);
        }

        public CompletableFuture<com.topview.api.SignUpResponse> signUpAsync(com.google.protobuf.Empty request){
            return StubInvocationUtil.unaryCall(invoker, signUpAsyncMethod, request);
        }

        @Override
        public void signUp(com.google.protobuf.Empty request, StreamObserver<com.topview.api.SignUpResponse> responseObserver){
            StubInvocationUtil.unaryCall(invoker, signUpMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.UserBalanceResponse getUserBalance(com.topview.api.UserBalanceRequest request){
            return StubInvocationUtil.unaryCall(invoker, getUserBalanceMethod, request);
        }

        public CompletableFuture<com.topview.api.UserBalanceResponse> getUserBalanceAsync(com.topview.api.UserBalanceRequest request){
            return StubInvocationUtil.unaryCall(invoker, getUserBalanceAsyncMethod, request);
        }

        @Override
        public void getUserBalance(com.topview.api.UserBalanceRequest request, StreamObserver<com.topview.api.UserBalanceResponse> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getUserBalanceMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.ActivityAmountResponse getActivityAmount(com.google.protobuf.Empty request){
            return StubInvocationUtil.unaryCall(invoker, getActivityAmountMethod, request);
        }

        public CompletableFuture<com.topview.api.ActivityAmountResponse> getActivityAmountAsync(com.google.protobuf.Empty request){
            return StubInvocationUtil.unaryCall(invoker, getActivityAmountAsyncMethod, request);
        }

        @Override
        public void getActivityAmount(com.google.protobuf.Empty request, StreamObserver<com.topview.api.ActivityAmountResponse> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getActivityAmountMethod , request, responseObserver);
        }
        @Override
        public com.google.protobuf.Empty createActivity(com.topview.api.CreateActivityRequest request){
            return StubInvocationUtil.unaryCall(invoker, createActivityMethod, request);
        }

        public CompletableFuture<com.google.protobuf.Empty> createActivityAsync(com.topview.api.CreateActivityRequest request){
            return StubInvocationUtil.unaryCall(invoker, createActivityAsyncMethod, request);
        }

        @Override
        public void createActivity(com.topview.api.CreateActivityRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
            StubInvocationUtil.unaryCall(invoker, createActivityMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.ActivityAndPool getIdToActivity(com.topview.api.GetIdToActivityRequest request){
            return StubInvocationUtil.unaryCall(invoker, getIdToActivityMethod, request);
        }

        public CompletableFuture<com.topview.api.ActivityAndPool> getIdToActivityAsync(com.topview.api.GetIdToActivityRequest request){
            return StubInvocationUtil.unaryCall(invoker, getIdToActivityAsyncMethod, request);
        }

        @Override
        public void getIdToActivity(com.topview.api.GetIdToActivityRequest request, StreamObserver<com.topview.api.ActivityAndPool> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getIdToActivityMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.BeforeMintDTO beforeMint(com.topview.api.BeforeMintRequest request){
            return StubInvocationUtil.unaryCall(invoker, beforeMintMethod, request);
        }

        public CompletableFuture<com.topview.api.BeforeMintDTO> beforeMintAsync(com.topview.api.BeforeMintRequest request){
            return StubInvocationUtil.unaryCall(invoker, beforeMintAsyncMethod, request);
        }

        @Override
        public void beforeMint(com.topview.api.BeforeMintRequest request, StreamObserver<com.topview.api.BeforeMintDTO> responseObserver){
            StubInvocationUtil.unaryCall(invoker, beforeMintMethod , request, responseObserver);
        }
        @Override
        public com.google.protobuf.Empty getDcFromActivity(com.topview.api.GetDcFromActivityRequest request){
            return StubInvocationUtil.unaryCall(invoker, getDcFromActivityMethod, request);
        }

        public CompletableFuture<com.google.protobuf.Empty> getDcFromActivityAsync(com.topview.api.GetDcFromActivityRequest request){
            return StubInvocationUtil.unaryCall(invoker, getDcFromActivityAsyncMethod, request);
        }

        @Override
        public void getDcFromActivity(com.topview.api.GetDcFromActivityRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getDcFromActivityMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.UserStatusResponse getUserStatus(com.topview.api.GetUserStatusRequest request){
            return StubInvocationUtil.unaryCall(invoker, getUserStatusMethod, request);
        }

        public CompletableFuture<com.topview.api.UserStatusResponse> getUserStatusAsync(com.topview.api.GetUserStatusRequest request){
            return StubInvocationUtil.unaryCall(invoker, getUserStatusAsyncMethod, request);
        }

        @Override
        public void getUserStatus(com.topview.api.GetUserStatusRequest request, StreamObserver<com.topview.api.UserStatusResponse> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getUserStatusMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request){
            return StubInvocationUtil.unaryCall(invoker, checkDcAndReturnTimeMethod, request);
        }

        public CompletableFuture<com.topview.api.CheckDcAndReturnTimeOutputDTO> checkDcAndReturnTimeAsync(com.topview.api.CheckDcAndReturnTimeRequest request){
            return StubInvocationUtil.unaryCall(invoker, checkDcAndReturnTimeAsyncMethod, request);
        }

        @Override
        public void checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request, StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO> responseObserver){
            StubInvocationUtil.unaryCall(invoker, checkDcAndReturnTimeMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.GetHashToDcIdResponse getHashToDcId(com.topview.api.GetHashToDcIdRequest request){
            return StubInvocationUtil.unaryCall(invoker, getHashToDcIdMethod, request);
        }

        public CompletableFuture<com.topview.api.GetHashToDcIdResponse> getHashToDcIdAsync(com.topview.api.GetHashToDcIdRequest request){
            return StubInvocationUtil.unaryCall(invoker, getHashToDcIdAsyncMethod, request);
        }

        @Override
        public void getHashToDcId(com.topview.api.GetHashToDcIdRequest request, StreamObserver<com.topview.api.GetHashToDcIdResponse> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getHashToDcIdMethod , request, responseObserver);
        }
        @Override
        public com.google.protobuf.Empty give(com.topview.api.GiveRequest request){
            return StubInvocationUtil.unaryCall(invoker, giveMethod, request);
        }

        public CompletableFuture<com.google.protobuf.Empty> giveAsync(com.topview.api.GiveRequest request){
            return StubInvocationUtil.unaryCall(invoker, giveAsyncMethod, request);
        }

        @Override
        public void give(com.topview.api.GiveRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
            StubInvocationUtil.unaryCall(invoker, giveMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request){
            return StubInvocationUtil.unaryCall(invoker, getDcHistoryAndMessageMethod, request);
        }

        public CompletableFuture<com.topview.api.DcHistoryAndMessageOutputDTO> getDcHistoryAndMessageAsync(com.topview.api.GetDcHistoryAndMessageRequest request){
            return StubInvocationUtil.unaryCall(invoker, getDcHistoryAndMessageAsyncMethod, request);
        }

        @Override
        public void getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request, StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getDcHistoryAndMessageMethod , request, responseObserver);
        }
        @Override
        public com.topview.api.PoolAmountResponse getPoolAmount(com.google.protobuf.Empty request){
            return StubInvocationUtil.unaryCall(invoker, getPoolAmountMethod, request);
        }

        public CompletableFuture<com.topview.api.PoolAmountResponse> getPoolAmountAsync(com.google.protobuf.Empty request){
            return StubInvocationUtil.unaryCall(invoker, getPoolAmountAsyncMethod, request);
        }

        @Override
        public void getPoolAmount(com.google.protobuf.Empty request, StreamObserver<com.topview.api.PoolAmountResponse> responseObserver){
            StubInvocationUtil.unaryCall(invoker, getPoolAmountMethod , request, responseObserver);
        }
        @Override
        public com.google.protobuf.Empty createPool(com.topview.api.CreatePoolRequest request){
            return StubInvocationUtil.unaryCall(invoker, createPoolMethod, request);
        }

        public CompletableFuture<com.google.protobuf.Empty> createPoolAsync(com.topview.api.CreatePoolRequest request){
            return StubInvocationUtil.unaryCall(invoker, createPoolAsyncMethod, request);
        }

        @Override
        public void createPool(com.topview.api.CreatePoolRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
            StubInvocationUtil.unaryCall(invoker, createPoolMethod , request, responseObserver);
        }
        @Override
        public com.google.protobuf.Empty mint(com.topview.api.MintRequest request){
            return StubInvocationUtil.unaryCall(invoker, mintMethod, request);
        }

        public CompletableFuture<com.google.protobuf.Empty> mintAsync(com.topview.api.MintRequest request){
            return StubInvocationUtil.unaryCall(invoker, mintAsyncMethod, request);
        }

        @Override
        public void mint(com.topview.api.MintRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
            StubInvocationUtil.unaryCall(invoker, mintMethod , request, responseObserver);
        }



    }

    public static abstract class BlcRpcServiceImplBase implements BlcRpcService, ServerService<BlcRpcService> {

        private <T, R> BiConsumer<T, StreamObserver<R>> syncToAsync(java.util.function.Function<T, R> syncFun) {
            return new BiConsumer<T, StreamObserver<R>>() {
                @Override
                public void accept(T t, StreamObserver<R> observer) {
                    try {
                        R ret = syncFun.apply(t);
                        observer.onNext(ret);
                        observer.onCompleted();
                    } catch (Throwable e) {
                        observer.onError(e);
                    }
                }
            };
        }

        @Override
        public final Invoker<BlcRpcService> getInvoker(URL url) {
            PathResolver pathResolver = url.getOrDefaultFrameworkModel()
            .getExtensionLoader(PathResolver.class)
            .getDefaultExtension();
            Map<String,StubMethodHandler<?, ?>> handlers = new HashMap<>();

            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/SignUp" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/SignUpAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetUserBalance" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetUserBalanceAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetActivityAmount" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetActivityAmountAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/CreateActivity" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/CreateActivityAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetIdToActivity" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetIdToActivityAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/BeforeMint" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/BeforeMintAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetDcFromActivity" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetDcFromActivityAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetUserStatus" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetUserStatusAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/CheckDcAndReturnTime" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/CheckDcAndReturnTimeAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetHashToDcId" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetHashToDcIdAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/Give" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GiveAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetDcHistoryAndMessage" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetDcHistoryAndMessageAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetPoolAmount" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/GetPoolAmountAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/CreatePool" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/CreatePoolAsync" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/Mint" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/MintAsync" );

            BiConsumer<com.google.protobuf.Empty, StreamObserver<com.topview.api.SignUpResponse>> signUpFunc = this::signUp;
            handlers.put(signUpMethod.getMethodName(), new UnaryStubMethodHandler<>(signUpFunc));
            BiConsumer<com.google.protobuf.Empty, StreamObserver<com.topview.api.SignUpResponse>> signUpAsyncFunc = syncToAsync(this::signUp);
            handlers.put(signUpProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(signUpAsyncFunc));
            BiConsumer<com.topview.api.UserBalanceRequest, StreamObserver<com.topview.api.UserBalanceResponse>> getUserBalanceFunc = this::getUserBalance;
            handlers.put(getUserBalanceMethod.getMethodName(), new UnaryStubMethodHandler<>(getUserBalanceFunc));
            BiConsumer<com.topview.api.UserBalanceRequest, StreamObserver<com.topview.api.UserBalanceResponse>> getUserBalanceAsyncFunc = syncToAsync(this::getUserBalance);
            handlers.put(getUserBalanceProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getUserBalanceAsyncFunc));
            BiConsumer<com.google.protobuf.Empty, StreamObserver<com.topview.api.ActivityAmountResponse>> getActivityAmountFunc = this::getActivityAmount;
            handlers.put(getActivityAmountMethod.getMethodName(), new UnaryStubMethodHandler<>(getActivityAmountFunc));
            BiConsumer<com.google.protobuf.Empty, StreamObserver<com.topview.api.ActivityAmountResponse>> getActivityAmountAsyncFunc = syncToAsync(this::getActivityAmount);
            handlers.put(getActivityAmountProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getActivityAmountAsyncFunc));
            BiConsumer<com.topview.api.CreateActivityRequest, StreamObserver<com.google.protobuf.Empty>> createActivityFunc = this::createActivity;
            handlers.put(createActivityMethod.getMethodName(), new UnaryStubMethodHandler<>(createActivityFunc));
            BiConsumer<com.topview.api.CreateActivityRequest, StreamObserver<com.google.protobuf.Empty>> createActivityAsyncFunc = syncToAsync(this::createActivity);
            handlers.put(createActivityProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(createActivityAsyncFunc));
            BiConsumer<com.topview.api.GetIdToActivityRequest, StreamObserver<com.topview.api.ActivityAndPool>> getIdToActivityFunc = this::getIdToActivity;
            handlers.put(getIdToActivityMethod.getMethodName(), new UnaryStubMethodHandler<>(getIdToActivityFunc));
            BiConsumer<com.topview.api.GetIdToActivityRequest, StreamObserver<com.topview.api.ActivityAndPool>> getIdToActivityAsyncFunc = syncToAsync(this::getIdToActivity);
            handlers.put(getIdToActivityProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getIdToActivityAsyncFunc));
            BiConsumer<com.topview.api.BeforeMintRequest, StreamObserver<com.topview.api.BeforeMintDTO>> beforeMintFunc = this::beforeMint;
            handlers.put(beforeMintMethod.getMethodName(), new UnaryStubMethodHandler<>(beforeMintFunc));
            BiConsumer<com.topview.api.BeforeMintRequest, StreamObserver<com.topview.api.BeforeMintDTO>> beforeMintAsyncFunc = syncToAsync(this::beforeMint);
            handlers.put(beforeMintProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(beforeMintAsyncFunc));
            BiConsumer<com.topview.api.GetDcFromActivityRequest, StreamObserver<com.google.protobuf.Empty>> getDcFromActivityFunc = this::getDcFromActivity;
            handlers.put(getDcFromActivityMethod.getMethodName(), new UnaryStubMethodHandler<>(getDcFromActivityFunc));
            BiConsumer<com.topview.api.GetDcFromActivityRequest, StreamObserver<com.google.protobuf.Empty>> getDcFromActivityAsyncFunc = syncToAsync(this::getDcFromActivity);
            handlers.put(getDcFromActivityProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getDcFromActivityAsyncFunc));
            BiConsumer<com.topview.api.GetUserStatusRequest, StreamObserver<com.topview.api.UserStatusResponse>> getUserStatusFunc = this::getUserStatus;
            handlers.put(getUserStatusMethod.getMethodName(), new UnaryStubMethodHandler<>(getUserStatusFunc));
            BiConsumer<com.topview.api.GetUserStatusRequest, StreamObserver<com.topview.api.UserStatusResponse>> getUserStatusAsyncFunc = syncToAsync(this::getUserStatus);
            handlers.put(getUserStatusProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getUserStatusAsyncFunc));
            BiConsumer<com.topview.api.CheckDcAndReturnTimeRequest, StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO>> checkDcAndReturnTimeFunc = this::checkDcAndReturnTime;
            handlers.put(checkDcAndReturnTimeMethod.getMethodName(), new UnaryStubMethodHandler<>(checkDcAndReturnTimeFunc));
            BiConsumer<com.topview.api.CheckDcAndReturnTimeRequest, StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO>> checkDcAndReturnTimeAsyncFunc = syncToAsync(this::checkDcAndReturnTime);
            handlers.put(checkDcAndReturnTimeProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(checkDcAndReturnTimeAsyncFunc));
            BiConsumer<com.topview.api.GetHashToDcIdRequest, StreamObserver<com.topview.api.GetHashToDcIdResponse>> getHashToDcIdFunc = this::getHashToDcId;
            handlers.put(getHashToDcIdMethod.getMethodName(), new UnaryStubMethodHandler<>(getHashToDcIdFunc));
            BiConsumer<com.topview.api.GetHashToDcIdRequest, StreamObserver<com.topview.api.GetHashToDcIdResponse>> getHashToDcIdAsyncFunc = syncToAsync(this::getHashToDcId);
            handlers.put(getHashToDcIdProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getHashToDcIdAsyncFunc));
            BiConsumer<com.topview.api.GiveRequest, StreamObserver<com.google.protobuf.Empty>> giveFunc = this::give;
            handlers.put(giveMethod.getMethodName(), new UnaryStubMethodHandler<>(giveFunc));
            BiConsumer<com.topview.api.GiveRequest, StreamObserver<com.google.protobuf.Empty>> giveAsyncFunc = syncToAsync(this::give);
            handlers.put(giveProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(giveAsyncFunc));
            BiConsumer<com.topview.api.GetDcHistoryAndMessageRequest, StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO>> getDcHistoryAndMessageFunc = this::getDcHistoryAndMessage;
            handlers.put(getDcHistoryAndMessageMethod.getMethodName(), new UnaryStubMethodHandler<>(getDcHistoryAndMessageFunc));
            BiConsumer<com.topview.api.GetDcHistoryAndMessageRequest, StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO>> getDcHistoryAndMessageAsyncFunc = syncToAsync(this::getDcHistoryAndMessage);
            handlers.put(getDcHistoryAndMessageProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getDcHistoryAndMessageAsyncFunc));
            BiConsumer<com.google.protobuf.Empty, StreamObserver<com.topview.api.PoolAmountResponse>> getPoolAmountFunc = this::getPoolAmount;
            handlers.put(getPoolAmountMethod.getMethodName(), new UnaryStubMethodHandler<>(getPoolAmountFunc));
            BiConsumer<com.google.protobuf.Empty, StreamObserver<com.topview.api.PoolAmountResponse>> getPoolAmountAsyncFunc = syncToAsync(this::getPoolAmount);
            handlers.put(getPoolAmountProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(getPoolAmountAsyncFunc));
            BiConsumer<com.topview.api.CreatePoolRequest, StreamObserver<com.google.protobuf.Empty>> createPoolFunc = this::createPool;
            handlers.put(createPoolMethod.getMethodName(), new UnaryStubMethodHandler<>(createPoolFunc));
            BiConsumer<com.topview.api.CreatePoolRequest, StreamObserver<com.google.protobuf.Empty>> createPoolAsyncFunc = syncToAsync(this::createPool);
            handlers.put(createPoolProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(createPoolAsyncFunc));
            BiConsumer<com.topview.api.MintRequest, StreamObserver<com.google.protobuf.Empty>> mintFunc = this::mint;
            handlers.put(mintMethod.getMethodName(), new UnaryStubMethodHandler<>(mintFunc));
            BiConsumer<com.topview.api.MintRequest, StreamObserver<com.google.protobuf.Empty>> mintAsyncFunc = syncToAsync(this::mint);
            handlers.put(mintProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(mintAsyncFunc));




            return new StubInvoker<>(this, url, BlcRpcService.class, handlers);
        }


        @Override
        public com.topview.api.SignUpResponse signUp(com.google.protobuf.Empty request){
            throw unimplementedMethodException(signUpMethod);
        }

        @Override
        public com.topview.api.UserBalanceResponse getUserBalance(com.topview.api.UserBalanceRequest request){
            throw unimplementedMethodException(getUserBalanceMethod);
        }

        @Override
        public com.topview.api.ActivityAmountResponse getActivityAmount(com.google.protobuf.Empty request){
            throw unimplementedMethodException(getActivityAmountMethod);
        }

        @Override
        public com.google.protobuf.Empty createActivity(com.topview.api.CreateActivityRequest request){
            throw unimplementedMethodException(createActivityMethod);
        }

        @Override
        public com.topview.api.ActivityAndPool getIdToActivity(com.topview.api.GetIdToActivityRequest request){
            throw unimplementedMethodException(getIdToActivityMethod);
        }

        @Override
        public com.topview.api.BeforeMintDTO beforeMint(com.topview.api.BeforeMintRequest request){
            throw unimplementedMethodException(beforeMintMethod);
        }

        @Override
        public com.google.protobuf.Empty getDcFromActivity(com.topview.api.GetDcFromActivityRequest request){
            throw unimplementedMethodException(getDcFromActivityMethod);
        }

        @Override
        public com.topview.api.UserStatusResponse getUserStatus(com.topview.api.GetUserStatusRequest request){
            throw unimplementedMethodException(getUserStatusMethod);
        }

        @Override
        public com.topview.api.CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request){
            throw unimplementedMethodException(checkDcAndReturnTimeMethod);
        }

        @Override
        public com.topview.api.GetHashToDcIdResponse getHashToDcId(com.topview.api.GetHashToDcIdRequest request){
            throw unimplementedMethodException(getHashToDcIdMethod);
        }

        @Override
        public com.google.protobuf.Empty give(com.topview.api.GiveRequest request){
            throw unimplementedMethodException(giveMethod);
        }

        @Override
        public com.topview.api.DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request){
            throw unimplementedMethodException(getDcHistoryAndMessageMethod);
        }

        @Override
        public com.topview.api.PoolAmountResponse getPoolAmount(com.google.protobuf.Empty request){
            throw unimplementedMethodException(getPoolAmountMethod);
        }

        @Override
        public com.google.protobuf.Empty createPool(com.topview.api.CreatePoolRequest request){
            throw unimplementedMethodException(createPoolMethod);
        }

        @Override
        public com.google.protobuf.Empty mint(com.topview.api.MintRequest request){
            throw unimplementedMethodException(mintMethod);
        }





        @Override
        public final ServiceDescriptor getServiceDescriptor() {
            return serviceDescriptor;
        }
        private RpcException unimplementedMethodException(StubMethodDescriptor methodDescriptor) {
            return TriRpcStatus.UNIMPLEMENTED.withDescription(String.format("Method %s is unimplemented",
                "/" + serviceDescriptor.getInterfaceName() + "/" + methodDescriptor.getMethodName())).asException();
        }
    }

}
