package com.topview.blc;

import lombok.Getter;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.codec.abi.FunctionEncoder;
import org.fisco.bcos.sdk.v3.codec.datatypes.*;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple1;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple2;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple3;
import org.fisco.bcos.sdk.v3.contract.Contract;
import org.fisco.bcos.sdk.v3.crypto.CryptoSuite;
import org.fisco.bcos.sdk.v3.crypto.keypair.CryptoKeyPair;
import org.fisco.bcos.sdk.v3.model.CryptoType;
import org.fisco.bcos.sdk.v3.model.TransactionReceipt;
import org.fisco.bcos.sdk.v3.model.callback.CallCallback;
import org.fisco.bcos.sdk.v3.model.callback.TransactionCallback;
import org.fisco.bcos.sdk.v3.transaction.model.exception.ContractException;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

@SuppressWarnings("unchecked")
public class UserLogic extends Contract {
    public static final String[] BINARY_ARRAY = {
        "608060405260008060146101000a81548160ff0219169083151502179055503480156200002b57600080fd5b506040516200143a3803806200143a8339818101604052810190620000519190620000af565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000129565b600081519050620000a9816200010f565b92915050565b600060208284031215620000c257600080fd5b6000620000d28482850162000098565b91505092915050565b6000620000e882620000ef565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200011a81620000db565b81146200012657600080fd5b50565b61130180620001396000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80634773489211610066578063477348921461011e5780635b86f5991461014e578063a9059cbb1461016a578063cb9429b514610186578063ff056949146101a257610093565b80631d1ac66a1461009857806322e8c87d146100b457806323b872dd146100d2578063410e9385146100ee575b600080fd5b6100b260048036038101906100ad9190610e48565b6101be565b005b6100bc610364565b6040516100c99190611122565b60405180910390f35b6100ec60048036038101906100e79190610e71565b61038d565b005b61010860048036038101906101039190610e48565b6105bc565b6040516101159190611107565b60405180910390f35b61013860048036038101906101339190610e48565b610672565b60405161014591906111bd565b60405180910390f35b61016860048036038101906101639190610ec0565b610725565b005b610184600480360381019061017f9190610ec0565b6108aa565b005b6101a0600480360381019061019b9190610e48565b610ad8565b005b6101bc60048036038101906101b79190610ec0565b610c84565b005b6101c7816105bc565b610206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fd9061115d565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166399eb9d6b826040518263ffffffff1660e01b8152600401610260919061107f565b600060405180830381600087803b15801561027a57600080fd5b505af115801561028e573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663603bbcd0826040518263ffffffff1660e01b81526004016102ec919061107f565b600060405180830381600087803b15801561030657600080fd5b505af115801561031a573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167fdcaca4b2811e788974d2e06d9eea56573b1a21e19aad43bbf58ede7d1b0114dd60405160405180910390a250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6001600060146101000a81548160ff0219169083151502179055506103b1836105bc565b6103f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e79061113d565b60405180910390fd5b6103f9826105bc565b610438576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042f9061113d565b60405180910390fd5b806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece856040518263ffffffff1660e01b8152600401610493919061107f565b60206040518083038186803b1580156104ab57600080fd5b505afa1580156104bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e39190610efc565b1015610524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051b9061117d565b60405180910390fd5b61052e8382610c84565b6105388282610725565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f0a85107a334eae0d22d21cdf13af0f8e8125039ec60baaa843d2c4c5b06801748360405161059591906111bd565b60405180910390a360008060146101000a81548160ff021916908315150217905550505050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ea0d5dcd846040518263ffffffff1660e01b8152600401610619919061107f565b60206040518083038186803b15801561063157600080fd5b505afa158015610645573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106699190610efc565b14159050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece836040518263ffffffff1660e01b81526004016106ce919061107f565b60206040518083038186803b1580156106e657600080fd5b505afa1580156106fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071e9190610efc565b9050919050565b61072e826105bc565b61076d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107649061113d565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b46310f683836000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece876040518263ffffffff1660e01b8152600401610806919061107f565b60206040518083038186803b15801561081e57600080fd5b505afa158015610832573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108569190610efc565b016040518363ffffffff1660e01b81526004016108749291906110de565b600060405180830381600087803b15801561088e57600080fd5b505af11580156108a2573d6000803e3d6000fd5b505050505050565b6001600060146101000a81548160ff0219169083151502179055506108ce336105bc565b61090d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109049061113d565b60405180910390fd5b610916826105bc565b610955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094c9061113d565b60405180910390fd5b806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece336040518263ffffffff1660e01b81526004016109b0919061109a565b60206040518083038186803b1580156109c857600080fd5b505afa1580156109dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a009190610efc565b1015610a41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a389061117d565b60405180910390fd5b610a4b3382610c84565b610a558282610725565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f0a85107a334eae0d22d21cdf13af0f8e8125039ec60baaa843d2c4c5b068017483604051610ab291906111bd565b60405180910390a360008060146101000a81548160ff0219169083151502179055505050565b610ae1816105bc565b15610b21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b189061119d565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347ff2fc8826040518263ffffffff1660e01b8152600401610b7b919061107f565b600060405180830381600087803b158015610b9557600080fd5b505af1158015610ba9573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b46310f682620186a06040518363ffffffff1660e01b8152600401610c0c9291906110b5565b600060405180830381600087803b158015610c2657600080fd5b505af1158015610c3a573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167f1e29f390a93bfed613d86b3664c2fab2e58b951c49fba222e4b7a33106c5172d60405160405180910390a250565b610c8d826105bc565b610ccc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc39061113d565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b46310f683836000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece876040518263ffffffff1660e01b8152600401610d65919061107f565b60206040518083038186803b158015610d7d57600080fd5b505afa158015610d91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db59190610efc565b036040518363ffffffff1660e01b8152600401610dd39291906110de565b600060405180830381600087803b158015610ded57600080fd5b505af1158015610e01573d6000803e3d6000fd5b505050505050565b600081359050610e188161129d565b92915050565b600081359050610e2d816112b4565b92915050565b600081519050610e42816112b4565b92915050565b600060208284031215610e5a57600080fd5b6000610e6884828501610e09565b91505092915050565b600080600060608486031215610e8657600080fd5b6000610e9486828701610e09565b9350506020610ea586828701610e09565b9250506040610eb686828701610e1e565b9150509250925092565b600080604083", "85031215610ed357600080fd5b6000610ee185828601610e09565b9250506020610ef285828601610e1e565b9150509250929050565b600060208284031215610f0e57600080fd5b6000610f1c84828501610e33565b91505092915050565b610f2e81611231565b82525050565b610f3d816111e9565b82525050565b610f4c816111fb565b82525050565b610f5b81611243565b82525050565b610f6a81611267565b82525050565b6000610f7d600f836111d8565b91507fe794a8e688b7e69caae6b3a8e5868c00000000000000000000000000000000006000830152602082019050919050565b6000610fbd600f836111d8565b91507fe794a8e688b7e4b88de5ad98e59ca800000000000000000000000000000000006000830152602082019050919050565b6000610ffd600c836111d8565b91507fe4bd99e9a29de4b88de8b6b300000000000000000000000000000000000000006000830152602082019050919050565b600061103d600f836111d8565b91507fe794a8e688b7e5b7b2e5ad98e59ca800000000000000000000000000000000006000830152602082019050919050565b61107981611227565b82525050565b60006020820190506110946000830184610f34565b92915050565b60006020820190506110af6000830184610f25565b92915050565b60006040820190506110ca6000830185610f34565b6110d76020830184610f61565b9392505050565b60006040820190506110f36000830185610f34565b6111006020830184611070565b9392505050565b600060208201905061111c6000830184610f43565b92915050565b60006020820190506111376000830184610f52565b92915050565b6000602082019050818103600083015261115681610f70565b9050919050565b6000602082019050818103600083015261117681610fb0565b9050919050565b6000602082019050818103600083015261119681610ff0565b9050919050565b600060208201905081810360008301526111b681611030565b9050919050565b60006020820190506111d26000830184611070565b92915050565b600082825260208201905092915050565b60006111f482611207565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061123c82611279565b9050919050565b600061124e82611255565b9050919050565b600061126082611207565b9050919050565b600061127282611227565b9050919050565b60006112848261128b565b9050919050565b600061129682611207565b9050919050565b6112a6816111e9565b81146112b157600080fd5b50565b6112bd81611227565b81146112c857600080fd5b5056fea26469706673582212205bf4b0e824ded655f0bb8a645801fd48e1b51c5289e84de3391761b6f5dc44f164736f6c634300060a0033"};

    public static final String BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", BINARY_ARRAY);

    public static final String[] SM_BINARY_ARRAY = {};

    public static final String SM_BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", SM_BINARY_ARRAY);

    public static final String[] ABI_ARRAY = {"[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userDataAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"LogSignOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"LogSignUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"checkUserStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"decreaseBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserData\",\"outputs\":[{\"internalType\":\"contract UserData\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"increaseBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"signOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"signUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"};

    @Getter
    public static final String ABI = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", ABI_ARRAY);

    public static final String FUNC_CHECKUSERSTATUS = "checkUserStatus";

    public static final String FUNC_DECREASEBALANCE = "decreaseBalance";

    public static final String FUNC_GETUSERBALANCE = "getUserBalance";

    public static final String FUNC_GETUSERDATA = "getUserData";

    public static final String FUNC_INCREASEBALANCE = "increaseBalance";

    public static final String FUNC_SIGNOUT = "signOut";

    public static final String FUNC_SIGNUP = "signUp";

    public static final String FUNC_TRANSFER = "transfer";

    public static final String FUNC_TRANSFERFROM = "transferFrom";

    public static final Event LOGSIGNOUT_EVENT = new Event("LogSignOut",
            List.of(new TypeReference<Address>(true) {
            }));

    public static final Event LOGSIGNUP_EVENT = new Event("LogSignUp",
            List.of(new TypeReference<Address>(true) {
            }));

    public static final Event LOGTRANSFER_EVENT = new Event("LogTransfer",
            Arrays.asList(new TypeReference<Address>(true) {
        }, new TypeReference<Address>(true) {
        }, new TypeReference<Uint256>() {
        }));

    protected UserLogic(String contractAddress, Client client, CryptoKeyPair credential) {
        super(getBinary(client.getCryptoSuite()), contractAddress, client, credential);
    }

    public static String getBinary(CryptoSuite cryptoSuite) {
        return (cryptoSuite.getCryptoTypeConfig() == CryptoType.ECDSA_TYPE ? BINARY : SM_BINARY);
    }

    public static UserLogic load(String contractAddress, Client client, CryptoKeyPair credential) {
        return new UserLogic(contractAddress, client, credential);
    }

    public static UserLogic deploy(Client client, CryptoKeyPair credential, String _userDataAddress)
        throws ContractException {
        byte[] encodedConstructor = FunctionEncoder.encodeConstructor(List.of(new Address(_userDataAddress)));
        return deploy(UserLogic.class, client, credential, getBinary(client.getCryptoSuite()), getABI(), encodedConstructor, null);
    }

    public List<LogSignOutEventResponse> getLogSignOutEvents(
        TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(LOGSIGNOUT_EVENT, transactionReceipt);
        ArrayList<LogSignOutEventResponse> responses = new ArrayList<>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            LogSignOutEventResponse typedResponse = new LogSignOutEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.userAddress = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public List<LogSignUpEventResponse> getLogSignUpEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(LOGSIGNUP_EVENT, transactionReceipt);
        ArrayList<LogSignUpEventResponse> responses = new ArrayList<>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            LogSignUpEventResponse typedResponse = new LogSignUpEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.userAddress = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public List<LogTransferEventResponse> getLogTransferEvents(
        TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(LOGTRANSFER_EVENT, transactionReceipt);
        ArrayList<LogTransferEventResponse> responses = new ArrayList<>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            LogTransferEventResponse typedResponse = new LogTransferEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.from = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.to = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse.value = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Boolean checkUserStatus(String _user) throws ContractException {
        final Function function = new Function(FUNC_CHECKUSERSTATUS,
                List.of(new Address(_user)),
                List.of(new TypeReference<Bool>() {
                }));
        return executeCallWithSingleValueReturn(function, Boolean.class);
    }

    public void checkUserStatus(String _user, CallCallback callback) {
        final Function function = new Function(FUNC_CHECKUSERSTATUS,
                List.of(new Address(_user)),
                List.of(new TypeReference<Bool>() {
                }));
        asyncExecuteCall(function, callback);
    }

    public TransactionReceipt decreaseBalance(String _target, BigInteger _value) {
        final Function function = new Function(
            FUNC_DECREASEBALANCE,
                Arrays.asList(new Address(_target),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForDecreaseBalance(String _target, BigInteger _value) {
        final Function function = new Function(
            FUNC_DECREASEBALANCE,
                Arrays.asList(new Address(_target),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String decreaseBalance(String _target, BigInteger _value, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_DECREASEBALANCE,
                Arrays.asList(new Address(_target),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple2<String, BigInteger> getDecreaseBalanceInput(
        TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_DECREASEBALANCE,
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

    public BigInteger getUserBalance(String _user) throws ContractException {
        final Function function = new Function(FUNC_GETUSERBALANCE,
                List.of(new Address(_user)),
                List.of(new TypeReference<Uint256>() {
                }));
        return executeCallWithSingleValueReturn(function, BigInteger.class);
    }

    public void getUserBalance(String _user, CallCallback callback) {
        final Function function = new Function(FUNC_GETUSERBALANCE,
                List.of(new Address(_user)),
                List.of(new TypeReference<Uint256>() {
                }));
        asyncExecuteCall(function, callback);
    }

    public String getUserData() throws ContractException {
        final Function function = new Function(FUNC_GETUSERDATA,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        return executeCallWithSingleValueReturn(function, String.class);
    }

    public void getUserData(CallCallback callback) {
        final Function function = new Function(FUNC_GETUSERDATA,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        asyncExecuteCall(function, callback);
    }

    public TransactionReceipt increaseBalance(String _target, BigInteger _value) {
        final Function function = new Function(
            FUNC_INCREASEBALANCE,
                Arrays.asList(new Address(_target),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForIncreaseBalance(String _target, BigInteger _value) {
        final Function function = new Function(
            FUNC_INCREASEBALANCE,
                Arrays.asList(new Address(_target),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String increaseBalance(String _target, BigInteger _value, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_INCREASEBALANCE,
                Arrays.asList(new Address(_target),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple2<String, BigInteger> getIncreaseBalanceInput(
        TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_INCREASEBALANCE,
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

    public TransactionReceipt signOut(String _user) {
        final Function function = new Function(
            FUNC_SIGNOUT,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForSignOut(String _user) {
        final Function function = new Function(
            FUNC_SIGNOUT,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String signOut(String _user, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_SIGNOUT,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple1<String> getSignOutInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_SIGNOUT,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple1<>(

                (String) results.get(0).getValue()
        );
    }

    public TransactionReceipt signUp(String _user) {
        final Function function = new Function(
            FUNC_SIGNUP,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForSignUp(String _user) {
        final Function function = new Function(
            FUNC_SIGNUP,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String signUp(String _user, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_SIGNUP,
                List.of(new Address(_user)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple1<String> getSignUpInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_SIGNUP,
                List.of(),
                List.of(new TypeReference<Address>() {
                }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple1<>(

                (String) results.get(0).getValue()
        );
    }

    public TransactionReceipt transfer(String _to, BigInteger _value) {
        final Function function = new Function(
            FUNC_TRANSFER,
                Arrays.asList(new Address(_to),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForTransfer(String _to, BigInteger _value) {
        final Function function = new Function(
            FUNC_TRANSFER,
                Arrays.asList(new Address(_to),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String transfer(String _to, BigInteger _value, TransactionCallback callback) {
        final Function function = new Function(
            FUNC_TRANSFER,
                Arrays.asList(new Address(_to),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple2<String, BigInteger> getTransferInput(TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_TRANSFER,
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

    public TransactionReceipt transferFrom(String _from, String _to, BigInteger _value) {
        final Function function = new Function(
            FUNC_TRANSFERFROM,
                Arrays.asList(new Address(_from),
                new Address(_to),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForTransferFrom(String _from, String _to, BigInteger _value) {
        final Function function = new Function(
            FUNC_TRANSFERFROM,
                Arrays.asList(new Address(_from),
                new Address(_to),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String transferFrom(String _from, String _to, BigInteger _value,
        TransactionCallback callback) {
        final Function function = new Function(
            FUNC_TRANSFERFROM,
                Arrays.asList(new Address(_from),
                new Address(_to),
                new Uint256(_value)),
                Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
    }

    public Tuple3<String, String, BigInteger> getTransferFromInput(
        TransactionReceipt transactionReceipt) {
        String data = transactionReceipt.getInput().substring(10);
        final Function function = new Function(FUNC_TRANSFERFROM,
                List.of(),
                Arrays.asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }, new TypeReference<Uint256>() {
            }));
        List<Type> results = this.functionReturnDecoder.decode(data, function.getOutputParameters());
        return new Tuple3<>(

                (String) results.get(0).getValue(),
                (String) results.get(1).getValue(),
                (BigInteger) results.get(2).getValue()
        );
    }

    public static class LogSignOutEventResponse {
        public TransactionReceipt.Logs log;

        public String userAddress;
    }

    public static class LogSignUpEventResponse {
        public TransactionReceipt.Logs log;

        public String userAddress;
    }

    public static class LogTransferEventResponse {
        public TransactionReceipt.Logs log;

        public String from;

        public String to;

        public BigInteger value;
    }
}
