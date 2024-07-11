package com.topview.blc;

import lombok.Getter;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.codec.datatypes.Address;
import org.fisco.bcos.sdk.v3.codec.datatypes.Function;
import org.fisco.bcos.sdk.v3.codec.datatypes.Type;
import org.fisco.bcos.sdk.v3.codec.datatypes.TypeReference;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple1;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple2;
import org.fisco.bcos.sdk.v3.contract.Contract;
import org.fisco.bcos.sdk.v3.crypto.CryptoSuite;
import org.fisco.bcos.sdk.v3.crypto.keypair.CryptoKeyPair;
import org.fisco.bcos.sdk.v3.model.CryptoType;
import org.fisco.bcos.sdk.v3.model.TransactionReceipt;
import org.fisco.bcos.sdk.v3.model.callback.CallCallback;
import org.fisco.bcos.sdk.v3.model.callback.TransactionCallback;
import org.fisco.bcos.sdk.v3.transaction.model.exception.ContractException;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

@SuppressWarnings("unchecked")
public class UserData extends Contract {
    public static final String[] BINARY_ARRAY = {"608060405234801561001057600080fd5b5061052d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806347ff2fc814610067578063603bbcd01461008357806399eb9d6b1461009f5780639b96eece146100bb578063b46310f6146100eb578063ea0d5dcd14610107575b600080fd5b610081600480360381019061007c919061038d565b610137565b005b61009d6004803603810190610098919061038d565b610200565b005b6100b960048036038101906100b4919061038d565b610245565b005b6100d560048036038101906100d0919061038d565b61028b565b6040516100e29190610461565b60405180910390f35b610105600480360381019061010091906103b6565b6102d3565b005b610121600480360381019061011c919061038d565b61031a565b60405161012e9190610461565b60405180910390f35b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146101b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b090610441565b60405180910390fd5b42600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550565b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600081359050610372816104c9565b92915050565b600081359050610387816104e0565b92915050565b60006020828403121561039f57600080fd5b60006103ad84828501610363565b91505092915050565b600080604083850312156103c957600080fd5b60006103d785828601610363565b92505060206103e885828601610378565b9150509250929050565b60006103ff601e8361047c565b91507fe5908ce4b880e59cb0e59d80e58faae883bde6b3a8e5868ce4b880e6aca100006000830152602082019050919050565b61043b816104bf565b82525050565b6000602082019050818103600083015261045a816103f2565b9050919050565b60006020820190506104766000830184610432565b92915050565b600082825260208201905092915050565b60006104988261049f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6104d28161048d565b81146104dd57600080fd5b50565b6104e9816104bf565b81146104f457600080fd5b5056fea264697066735822122045a2aa86a7011317172760bf7fa47af98e3e2e3535604009896e234a2d95185d64736f6c634300060a0033"};

    public static final String BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", BINARY_ARRAY);

    public static final String[] SM_BINARY_ARRAY = {};

    public static final String SM_BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", SM_BINARY_ARRAY);

    public static final String[] ABI_ARRAY = {"[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetBalanceOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetUserStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"setBalanceOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"setUserStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"};

    @Getter
    public static final String ABI = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", ABI_ARRAY);

    public static final String FUNC_GETBALANCEOF = "getBalanceOf";

    public static final String FUNC_GETUSERSTATUS = "getUserStatus";

    public static final String FUNC_RESETBALANCEOF = "resetBalanceOf";

    public static final String FUNC_RESETUSERSTATUS = "resetUserStatus";

    public static final String FUNC_SETBALANCEOF = "setBalanceOf";

    public static final String FUNC_SETUSERSTATUS = "setUserStatus";

    protected UserData(String contractAddress, Client client, CryptoKeyPair credential) {
        super(getBinary(client.getCryptoSuite()), contractAddress, client, credential);
    }

    public static String getBinary(CryptoSuite cryptoSuite) {
        return (cryptoSuite.getCryptoTypeConfig() == CryptoType.ECDSA_TYPE ? BINARY : SM_BINARY);
    }

    public static UserData load(String contractAddress, Client client, CryptoKeyPair credential) {
        return new UserData(contractAddress, client, credential);
    }

    public static UserData deploy(Client client, CryptoKeyPair credential) throws
        ContractException {
        return deploy(UserData.class, client, credential, getBinary(client.getCryptoSuite()), getABI(), null, null);
    }

    public BigInteger getBalanceOf(String _user) throws ContractException {
        final Function function = new Function(FUNC_GETBALANCEOF,
                List.of(new Address(_user)),
                List.of(new TypeReference<Uint256>() {
                }));
        return executeCallWithSingleValueReturn(function, BigInteger.class);
    }

    public void getBalanceOf(String _user, CallCallback callback) {
        final Function function = new Function(FUNC_GETBALANCEOF,
                List.of(new Address(_user)),
                List.of(new TypeReference<Uint256>() {
                }));
        asyncExecuteCall(function, callback);
    }

    public BigInteger getUserStatus(String _user) throws ContractException {
        final Function function = new Function(FUNC_GETUSERSTATUS,
                List.of(new Address(_user)),
                List.of(new TypeReference<Uint256>() {
                }));
        return executeCallWithSingleValueReturn(function, BigInteger.class);
    }

    public void getUserStatus(String _user, CallCallback callback) {
        final Function function = new Function(FUNC_GETUSERSTATUS,
                List.of(new Address(_user)),
                List.of(new TypeReference<Uint256>() {
                }));
        asyncExecuteCall(function, callback);
    }

    public TransactionReceipt resetBalanceOf(String _user) {
        final Function function = new Function(
            FUNC_RESETBALANCEOF,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForResetBalanceOf(String _user) {
        final Function function = new Function(
            FUNC_RESETBALANCEOF,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String resetBalanceOf(String _user, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_RESETBALANCEOF,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple1<String> getResetBalanceOfInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_RESETBALANCEOF,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple1<>(

                (String) results.get(0).getValue()
        );
    }

    public TransactionReceipt resetUserStatus(String _user) {
        final Function function = new Function(
            FUNC_RESETUSERSTATUS,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForResetUserStatus(String _user) {
        final Function function = new Function(
            FUNC_RESETUSERSTATUS,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String resetUserStatus(String _user, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_RESETUSERSTATUS,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple1<String> getResetUserStatusInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_RESETUSERSTATUS,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple1<>(

                (String) results.get(0).getValue()
        );
    }

    public TransactionReceipt setBalanceOf(String _user, BigInteger _balance) {
        final Function function = new Function(
            FUNC_SETBALANCEOF,
                Arrays.asList(new Address(_user),
                new Uint256(_balance)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForSetBalanceOf(String _user, BigInteger _balance) {
        final Function function = new Function(
            FUNC_SETBALANCEOF,
                Arrays.asList(new Address(_user),
                new Uint256(_balance)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String setBalanceOf(String _user, BigInteger _balance, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_SETBALANCEOF,
                Arrays.asList(new Address(_user),
                new Uint256(_balance)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple2<String, BigInteger> getSetBalanceOfInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_SETBALANCEOF,
                List.of(),
                Arrays.asList(new TypeReference<Address>() {
            }, new TypeReference<Uint256>() {
            }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple2<>(

                (String) results.get(0).getValue(),
                (BigInteger) results.get(1).getValue()
        );
    }

    public TransactionReceipt setUserStatus(String _user) {
        final Function function = new Function(
            FUNC_SETUSERSTATUS,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForSetUserStatus(String _user) {
        final Function function = new Function(
            FUNC_SETUSERSTATUS,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String setUserStatus(String _user, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_SETUSERSTATUS,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple1<String> getSetUserStatusInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_SETUSERSTATUS,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple1<>(

                (String) results.get(0).getValue()
        );
    }
}
