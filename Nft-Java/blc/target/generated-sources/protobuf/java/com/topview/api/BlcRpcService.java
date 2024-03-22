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
import com.google.protobuf.Message;

import java.util.HashMap;
import java.util.Map;
import java.util.function.BiConsumer;
import java.util.concurrent.CompletableFuture;

public interface BlcRpcService extends org.apache.dubbo.rpc.model.DubboStub {

    String JAVA_SERVICE_NAME = "com.topview.api.BlcRpcService";
    String SERVICE_NAME = "com.topview.api.BlcRpcService";

    com.topview.api.SignUpResponse signUp(com.google.protobuf.Empty request);

    default CompletableFuture<com.topview.api.SignUpResponse> signUpAsync(com.google.protobuf.Empty request){
        return CompletableFuture.completedFuture(signUp(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void signUp(com.google.protobuf.Empty request, StreamObserver<com.topview.api.SignUpResponse> responseObserver){
        signUpAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.UserBalanceResponse getUserBalance(com.topview.api.UserBalanceRequest request);

    default CompletableFuture<com.topview.api.UserBalanceResponse> getUserBalanceAsync(com.topview.api.UserBalanceRequest request){
        return CompletableFuture.completedFuture(getUserBalance(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getUserBalance(com.topview.api.UserBalanceRequest request, StreamObserver<com.topview.api.UserBalanceResponse> responseObserver){
        getUserBalanceAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.ActivityAmountResponse getActivityAmount(com.google.protobuf.Empty request);

    default CompletableFuture<com.topview.api.ActivityAmountResponse> getActivityAmountAsync(com.google.protobuf.Empty request){
        return CompletableFuture.completedFuture(getActivityAmount(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getActivityAmount(com.google.protobuf.Empty request, StreamObserver<com.topview.api.ActivityAmountResponse> responseObserver){
        getActivityAmountAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.google.protobuf.Empty createActivity(com.topview.api.CreateActivityRequest request);

    default CompletableFuture<com.google.protobuf.Empty> createActivityAsync(com.topview.api.CreateActivityRequest request){
        return CompletableFuture.completedFuture(createActivity(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void createActivity(com.topview.api.CreateActivityRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
        createActivityAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.ActivityAndPool getIdToActivity(com.topview.api.GetIdToActivityRequest request);

    default CompletableFuture<com.topview.api.ActivityAndPool> getIdToActivityAsync(com.topview.api.GetIdToActivityRequest request){
        return CompletableFuture.completedFuture(getIdToActivity(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getIdToActivity(com.topview.api.GetIdToActivityRequest request, StreamObserver<com.topview.api.ActivityAndPool> responseObserver){
        getIdToActivityAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.BeforeMintDTO beforeMint(com.topview.api.BeforeMintRequest request);

    default CompletableFuture<com.topview.api.BeforeMintDTO> beforeMintAsync(com.topview.api.BeforeMintRequest request){
        return CompletableFuture.completedFuture(beforeMint(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void beforeMint(com.topview.api.BeforeMintRequest request, StreamObserver<com.topview.api.BeforeMintDTO> responseObserver){
        beforeMintAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.google.protobuf.Empty getDcFromActivity(com.topview.api.GetDcFromActivityRequest request);

    default CompletableFuture<com.google.protobuf.Empty> getDcFromActivityAsync(com.topview.api.GetDcFromActivityRequest request){
        return CompletableFuture.completedFuture(getDcFromActivity(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getDcFromActivity(com.topview.api.GetDcFromActivityRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
        getDcFromActivityAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.UserStatusResponse getUserStatus(com.topview.api.GetUserStatusRequest request);

    default CompletableFuture<com.topview.api.UserStatusResponse> getUserStatusAsync(com.topview.api.GetUserStatusRequest request){
        return CompletableFuture.completedFuture(getUserStatus(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getUserStatus(com.topview.api.GetUserStatusRequest request, StreamObserver<com.topview.api.UserStatusResponse> responseObserver){
        getUserStatusAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request);

    default CompletableFuture<com.topview.api.CheckDcAndReturnTimeOutputDTO> checkDcAndReturnTimeAsync(com.topview.api.CheckDcAndReturnTimeRequest request){
        return CompletableFuture.completedFuture(checkDcAndReturnTime(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void checkDcAndReturnTime(com.topview.api.CheckDcAndReturnTimeRequest request, StreamObserver<com.topview.api.CheckDcAndReturnTimeOutputDTO> responseObserver){
        checkDcAndReturnTimeAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.GetHashToDcIdResponse getHashToDcId(com.topview.api.GetHashToDcIdRequest request);

    default CompletableFuture<com.topview.api.GetHashToDcIdResponse> getHashToDcIdAsync(com.topview.api.GetHashToDcIdRequest request){
        return CompletableFuture.completedFuture(getHashToDcId(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getHashToDcId(com.topview.api.GetHashToDcIdRequest request, StreamObserver<com.topview.api.GetHashToDcIdResponse> responseObserver){
        getHashToDcIdAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.google.protobuf.Empty give(com.topview.api.GiveRequest request);

    default CompletableFuture<com.google.protobuf.Empty> giveAsync(com.topview.api.GiveRequest request){
        return CompletableFuture.completedFuture(give(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void give(com.topview.api.GiveRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
        giveAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request);

    default CompletableFuture<com.topview.api.DcHistoryAndMessageOutputDTO> getDcHistoryAndMessageAsync(com.topview.api.GetDcHistoryAndMessageRequest request){
        return CompletableFuture.completedFuture(getDcHistoryAndMessage(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getDcHistoryAndMessage(com.topview.api.GetDcHistoryAndMessageRequest request, StreamObserver<com.topview.api.DcHistoryAndMessageOutputDTO> responseObserver){
        getDcHistoryAndMessageAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.topview.api.PoolAmountResponse getPoolAmount(com.google.protobuf.Empty request);

    default CompletableFuture<com.topview.api.PoolAmountResponse> getPoolAmountAsync(com.google.protobuf.Empty request){
        return CompletableFuture.completedFuture(getPoolAmount(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void getPoolAmount(com.google.protobuf.Empty request, StreamObserver<com.topview.api.PoolAmountResponse> responseObserver){
        getPoolAmountAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.google.protobuf.Empty createPool(com.topview.api.CreatePoolRequest request);

    default CompletableFuture<com.google.protobuf.Empty> createPoolAsync(com.topview.api.CreatePoolRequest request){
        return CompletableFuture.completedFuture(createPool(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void createPool(com.topview.api.CreatePoolRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
        createPoolAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }

    com.google.protobuf.Empty mint(com.topview.api.MintRequest request);

    default CompletableFuture<com.google.protobuf.Empty> mintAsync(com.topview.api.MintRequest request){
        return CompletableFuture.completedFuture(mint(request));
    }

    /**
    * This server stream type unary method is <b>only</b> used for generated stub to support async unary method.
    * It will not be called if you are NOT using Dubbo3 generated triple stub and <b>DO NOT</b> implement this method.
    */
    default void mint(com.topview.api.MintRequest request, StreamObserver<com.google.protobuf.Empty> responseObserver){
        mintAsync(request).whenComplete((r, t) -> {
            if (t != null) {
                responseObserver.onError(t);
            } else {
                responseObserver.onNext(r);
                responseObserver.onCompleted();
            }
        });
    }






}
