package com.topview.blc;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import lombok.Getter;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.codec.abi.FunctionEncoder;
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

@SuppressWarnings("unchecked")
public class UserData extends Contract {
    public static final String NAME = "UserData:";
    public static final String[] BINARY_ARRAY = {"608060405234801561001057600080fd5b50604051610bb3380380610bb38339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b610aa58061010e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806399eb9d6b1161005b57806399eb9d6b146100d85780639b96eece146100f4578063b46310f614610124578063ea0d5dcd146101405761007d565b806316d6b5f61461008257806347ff2fc8146100a0578063603bbcd0146100bc575b600080fd5b61008a610170565b6040516100979190610901565b60405180910390f35b6100ba60048036038101906100b591906107ab565b610199565b005b6100d660048036038101906100d191906107ab565b61034b565b005b6100f260048036038101906100ed91906107ab565b61047a565b005b61010e600480360381019061010991906107ab565b6105a9565b60405161011b919061095c565b60405180910390f35b61013e600480360381019061013991906107d4565b6105f2565b005b61015a600480360381019061015591906107ab565b610723565b604051610167919061095c565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016101f391906108e6565b60206040518083038186803b15801561020b57600080fd5b505afa15801561021f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102439190610810565b610282576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102799061091c565b60405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610304576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fb9061093c565b60405180910390fd5b42600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016103a591906108e6565b60206040518083038186803b1580156103bd57600080fd5b505afa1580156103d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f59190610810565b610434576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042b9061091c565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016104d491906108e6565b60206040518083038186803b1580156104ec57600080fd5b505afa158015610500573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105249190610810565b610563576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055a9061091c565b60405180910390fd5b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b815260040161064c91906108e6565b60206040518083038186803b15801561066457600080fd5b505afa158015610678573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069c9190610810565b6106db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d29061091c565b60405180910390fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008135905061077b81610a2a565b92915050565b60008151905061079081610a41565b92915050565b6000813590506107a581610a58565b92915050565b6000602082840312156107bd57600080fd5b60006107cb8482850161076c565b91505092915050565b600080604083850312156107e757600080fd5b60006107f58582860161076c565b925050602061080685828601610796565b9150509250929050565b60006020828403121561082257600080fd5b600061083084828501610781565b91505092915050565b610842816109d0565b82525050565b610851816109e2565b82525050565b6000610864601583610977565b91507fe4bda0e6b2a1e69c89e69d83e99990e8aebfe997ae00000000000000000000006000830152602082019050919050565b60006108a4601e83610977565b91507fe5908ce4b880e59cb0e59d80e58faae883bde6b3a8e5868ce4b880e6aca100006000830152602082019050919050565b6108e0816109c6565b82525050565b60006020820190506108fb6000830184610839565b92915050565b60006020820190506109166000830184610848565b92915050565b6000602082019050818103600083015261093581610857565b9050919050565b6000602082019050818103600083015261095581610897565b9050919050565b600060208201905061097160008301846108d7565b92915050565b600082825260208201905092915050565b6000610993826109a6565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006109db82610a06565b9050919050565b60006109ed826109f4565b9050919050565b60006109ff826109a6565b9050919050565b6000610a1182610a18565b9050919050565b6000610a23826109a6565b9050919050565b610a3381610988565b8114610a3e57600080fd5b50565b610a4a8161099a565b8114610a5557600080fd5b50565b610a61816109c6565b8114610a6c57600080fd5b5056fea2646970667358221220f2926d263dc85fe2e0c305090119b5852a262a8f549ebaf0f83fb6bf4f3416f964736f6c634300060a0033"};

    public static final String BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", BINARY_ARRAY);

    public static final String[] SM_BINARY_ARRAY = {};

    public static final String SM_BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", SM_BINARY_ARRAY);

    public static final String[] ABI_ARRAY = {"[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessControllerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getAccessController\",\"outputs\":[{\"internalType\":\"contract AccessController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetBalanceOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetUserStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"setBalanceOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"setUserStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"};

    @Getter public static final String ABI = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", ABI_ARRAY);

    public static final String FUNC_GETACCESSCONTROLLER = "getAccessController";

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

    public static UserData deploy(Client client, CryptoKeyPair credential,
        String _accessControllerAddress) throws ContractException {
        byte[] encodedConstructor = FunctionEncoder.encodeConstructor(List.of(new Address(_accessControllerAddress)));
        return deploy(UserData.class, client, credential, getBinary(client.getCryptoSuite()), getABI(), encodedConstructor, null);
    }

    public String getAccessController() throws ContractException {
        final Function function = new Function(FUNC_GETACCESSCONTROLLER,
            List.of(),
            List.of(new TypeReference<Address>() {
            }));
        return executeCallWithSingleValueReturn(function, String.class);
    }

    public void getAccessController(CallCallback callback) {
        final Function function = new Function(FUNC_GETACCESSCONTROLLER,
            List.of(),
            List.of(new TypeReference<Address>() {
            }));
        asyncExecuteCall(function, callback);
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
            Arrays.asList(new org.fisco.bcos.sdk.v3.codec.datatypes.Address(_user),
                new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(_balance)),
            Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForSetBalanceOf(String _user, BigInteger _balance) {
        final Function function = new Function(
            FUNC_SETBALANCEOF,
            Arrays.asList(new org.fisco.bcos.sdk.v3.codec.datatypes.Address(_user),
                new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(_balance)),
            Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String setBalanceOf(String _user, BigInteger _balance, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_SETBALANCEOF,
            Arrays.asList(new org.fisco.bcos.sdk.v3.codec.datatypes.Address(_user),
                new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(_balance)),
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
