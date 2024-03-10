package com.topview.blc;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import lombok.Getter;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.codec.abi.FunctionEncoder;
import org.fisco.bcos.sdk.v3.codec.datatypes.Address;
import org.fisco.bcos.sdk.v3.codec.datatypes.Bool;
import org.fisco.bcos.sdk.v3.codec.datatypes.Event;
import org.fisco.bcos.sdk.v3.codec.datatypes.Function;
import org.fisco.bcos.sdk.v3.codec.datatypes.Type;
import org.fisco.bcos.sdk.v3.codec.datatypes.TypeReference;
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

@SuppressWarnings("unchecked")
public class UserLogic extends Contract {
    public static final String NAME = "UserLogic:";
    public static final String[] BINARY_ARRAY = {
        "608060405260008060146101000a81548160ff0219169083151502179055503480156200002b57600080fd5b50604051620021b0380380620021b08339818101604052810190620000519190620000af565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000129565b600081519050620000a9816200010f565b92915050565b600060208284031215620000c257600080fd5b6000620000d28482850162000098565b91505092915050565b6000620000e882620000ef565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200011a81620000db565b81146200012657600080fd5b50565b61207780620001396000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80635b86f599116100665780635b86f5991461015957806383197ef014610175578063a9059cbb1461017f578063cb9429b51461019b578063ff056949146101b75761009e565b80631d1ac66a146100a357806322e8c87d146100bf57806323b872dd146100dd578063410e9385146100f95780634773489214610129575b600080fd5b6100bd60048036038101906100b89190611aa3565b6101d3565b005b6100c76104df565b6040516100d49190611e38565b60405180910390f35b6100f760048036038101906100f29190611af5565b610508565b005b610113600480360381019061010e9190611aa3565b61089d565b6040516101209190611e1d565b60405180910390f35b610143600480360381019061013e9190611aa3565b610953565b6040516101509190611ef3565b60405180910390f35b610173600480360381019061016e9190611b44565b610a06565b005b61017d610cf1565b005b61019960048036038101906101949190611b44565b611094565b005b6101b560048036038101906101b09190611aa3565b611428565b005b6101d160048036038101906101cc9190611b44565b61173a565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b15801561023a57600080fd5b505afa15801561024e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102729190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016102aa9190611db0565b60206040518083038186803b1580156102c257600080fd5b505afa1580156102d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fa9190611b80565b610339576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033090611e93565b60405180910390fd5b6103428161089d565b610381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037890611e73565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166399eb9d6b826040518263ffffffff1660e01b81526004016103db9190611d95565b600060405180830381600087803b1580156103f557600080fd5b505af1158015610409573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663603bbcd0826040518263ffffffff1660e01b81526004016104679190611d95565b600060405180830381600087803b15801561048157600080fd5b505af1158015610495573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167fdcaca4b2811e788974d2e06d9eea56573b1a21e19aad43bbf58ede7d1b0114dd60405160405180910390a250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b15801561056f57600080fd5b505afa158015610583573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a79190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016105df9190611db0565b60206040518083038186803b1580156105f757600080fd5b505afa15801561060b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062f9190611b80565b61066e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066590611e93565b60405180910390fd5b6001600060146101000a81548160ff0219169083151502179055506106928361089d565b6106d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c890611e53565b60405180910390fd5b6106da8261089d565b610719576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071090611e53565b60405180910390fd5b806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece856040518263ffffffff1660e01b81526004016107749190611d95565b60206040518083038186803b15801561078c57600080fd5b505afa1580156107a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c49190611bd2565b1015610805576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fc90611eb3565b60405180910390fd5b61080f838261173a565b6108198282610a06565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f0a85107a334eae0d22d21cdf13af0f8e8125039ec60baaa843d2c4c5b0680174836040516108769190611ef3565b60405180910390a360008060146101000a81548160ff021916908315150217905550505050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ea0d5dcd846040518263ffffffff1660e01b81526004016108fa9190611d95565b60206040518083038186803b15801561091257600080fd5b505afa158015610926573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094a9190611bd2565b14159050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece836040518263ffffffff1660e01b81526004016109af9190611d95565b60206040518083038186803b1580156109c757600080fd5b505afa1580156109db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ff9190611bd2565b9050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b158015610a6d57600080fd5b505afa158015610a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa59190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b8152600401610add9190611db0565b60206040518083038186803b158015610af557600080fd5b505afa158015610b09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2d9190611b80565b610b6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6390611e93565b60405180910390fd5b610b758261089d565b610bb4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bab90611e53565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b46310f683836000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece876040518263ffffffff1660e01b8152600401610c4d9190611d95565b60206040518083038186803b158015610c6557600080fd5b505afa158015610c79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9d9190611bd2565b016040518363ffffffff1660e01b8152600401610cbb929190611df4565b600060405180830381600087803b158015610cd557600080fd5b505af1158015610ce9573d6000803e3d6000fd5b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b158015610d5857600080fd5b505afa158015610d6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d909190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b8152600401610dc89190611db0565b60206040518083038186803b158015610de057600080fd5b505afa158015610df4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e189190611b80565b610e57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4e90611e93565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b158015610ebe57600080fd5b505afa158015610e", "d2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef69190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663fd945a6b306040518263ffffffff1660e01b8152600401610f2e9190611d95565b600060405180830381600087803b158015610f4857600080fd5b505af1158015610f5c573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b158015610fc757600080fd5b505afa158015610fdb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fff9190611ba9565b73ffffffffffffffffffffffffffffffffffffffff16636e9960c36040518163ffffffff1660e01b815260040160206040518083038186803b15801561104457600080fd5b505afa158015611058573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107c9190611acc565b73ffffffffffffffffffffffffffffffffffffffff16ff5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b1580156110fb57600080fd5b505afa15801561110f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111339190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b815260040161116b9190611db0565b60206040518083038186803b15801561118357600080fd5b505afa158015611197573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111bb9190611b80565b6111fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f190611e93565b60405180910390fd5b6001600060146101000a81548160ff02191690831515021790555061121e3361089d565b61125d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125490611e53565b60405180910390fd5b6112668261089d565b6112a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161129c90611e53565b60405180910390fd5b806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece336040518263ffffffff1660e01b81526004016113009190611db0565b60206040518083038186803b15801561131857600080fd5b505afa15801561132c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113509190611bd2565b1015611391576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161138890611eb3565b60405180910390fd5b61139b338261173a565b6113a58282610a06565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f0a85107a334eae0d22d21cdf13af0f8e8125039ec60baaa843d2c4c5b0680174836040516114029190611ef3565b60405180910390a360008060146101000a81548160ff0219169083151502179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b15801561148f57600080fd5b505afa1580156114a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c79190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016114ff9190611db0565b60206040518083038186803b15801561151757600080fd5b505afa15801561152b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154f9190611b80565b61158e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161158590611e93565b60405180910390fd5b6115978161089d565b156115d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ce90611ed3565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347ff2fc8826040518263ffffffff1660e01b81526004016116319190611d95565b600060405180830381600087803b15801561164b57600080fd5b505af115801561165f573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b46310f682620186a06040518363ffffffff1660e01b81526004016116c2929190611dcb565b600060405180830381600087803b1580156116dc57600080fd5b505af11580156116f0573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167f1e29f390a93bfed613d86b3664c2fab2e58b951c49fba222e4b7a33106c5172d60405160405180910390a250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166316d6b5f66040518163ffffffff1660e01b815260040160206040518083038186803b1580156117a157600080fd5b505afa1580156117b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d99190611ba9565b73ffffffffffffffffffffffffffffffffffffffff1663a1f0c941336040518263ffffffff1660e01b81526004016118119190611db0565b60206040518083038186803b15801561182957600080fd5b505afa15801561183d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118619190611b80565b6118a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161189790611e93565b60405180910390fd5b6118a98261089d565b6118e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118df90611e53565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b46310f683836000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b96eece876040518263ffffffff1660e01b81526004016119819190611d95565b60206040518083038186803b15801561199957600080fd5b505afa1580156119ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119d19190611bd2565b036040518363ffffffff1660e01b81526004016119ef929190611df4565b600060405180830381600087803b158015611a0957600080fd5b505af1158015611a1d573d6000803e3d6000fd5b505050505050565b600081359050611a3481611fe5565b92915050565b600081519050611a4981611fe5565b92915050565b600081519050611a5e81611ffc565b92915050565b600081519050611a7381612013565b92915050565b600081359050611a888161202a565b92915050565b600081519050611a9d8161202a565b92915050565b600060208284031215611ab557600080fd5b6000611ac384828501611a25565b91505092915050565b600060208284031215611ade57600080fd5b6000611aec84828501611a3a565b91505092915050565b600080600060608486031215611b0a57600080fd5b6000611b1886828701611a25565b9350506020611b2986828701611a25565b9250506040611b3a86828701611a79565b9150509250925092565b60008060408385031215611b5757600080fd5b6000611b6585828601611a25565b9250506020611b7685828601611a79565b9150509250929050565b600060208284031215611b9257600080fd5b6000611ba084828501611a4f565b91505092915050565b600060208284031215611bbb57600080fd5b6000611bc984828501611a64565b91505092915050565b600060208284031215611be457600080fd5b6000611bf284828501611a8e565b91505092915050565b611c0481611f79565b82525050565b611c1381611f1f565b82525050565b611c2281611f31565b82525050565b611c3181611f8b565b82525050565b611c4081611faf565b82525050565b6000611c53600f83611f0e565b91507fe794a8e688b7e69caae6b3a8e5868c00000000000000000000000000000000006000830152602082019050919050565b6000611c93600f83611f0e565b91507fe794a8e688b7e4b88de5ad98e59ca800000000000000000000000000000000006000830152602082019050919050565b6000611cd3601583611f0e565b91507fe4bda0e6b2a1e69c89e69d83e99990e8aebfe997ae00000000000000000000006000830152602082019050919050565b6000611d13600c83611f0e565b91507fe4bd99e9a29de4b88de8b6b300000000000000000000000000000000000000006000830152602082019050919050565b6000611d53600f83611f0e565b91507fe794a8e688b7e5b7b2e5ad98e59ca800000000000000000000000000000000006000830152602082019050919050565b611d8f81611f6f565b82525050565b6000602082019050611daa6000830184611c0a565b92915050565b6000602082019050611dc56000830184611bfb565b92915050565b6000604082019050611de06000830185611c0a565b611ded6020830184611c37565b9392505050565b6000604082019050611e096000830185611c0a565b611e166020830184611d86565b9392505050565b6000602082019050611e326000830184611c19565b92915050565b6000602082019050611e4d6000830184611c28565b92915050565b60006020820190508181036000830152611e6c81611c46565b9050919050565b60006020820190508181036000830152611e8c81611c86565b9050919050565b60006020820190508181036000830152611eac81611cc6565b9050919050565b60006020820190508181036000830152611ecc", "81611d06565b9050919050565b60006020820190508181036000830152611eec81611d46565b9050919050565b6000602082019050611f086000830184611d86565b92915050565b600082825260208201905092915050565b6000611f2a82611f4f565b9050919050565b60008115159050919050565b6000611f4882611f1f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611f8482611fc1565b9050919050565b6000611f9682611f9d565b9050919050565b6000611fa882611f4f565b9050919050565b6000611fba82611f6f565b9050919050565b6000611fcc82611fd3565b9050919050565b6000611fde82611f4f565b9050919050565b611fee81611f1f565b8114611ff957600080fd5b50565b61200581611f31565b811461201057600080fd5b50565b61201c81611f3d565b811461202757600080fd5b50565b61203381611f6f565b811461203e57600080fd5b5056fea2646970667358221220c5a88666f69f3d3564988904460778f50f53e965829cb0821238f170e952b77e64736f6c634300060a0033"};

    public static final String BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", BINARY_ARRAY);

    public static final String[] SM_BINARY_ARRAY = {};

    public static final String SM_BINARY = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", SM_BINARY_ARRAY);

    public static final String[] ABI_ARRAY = {"[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userDataAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"LogSignOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"LogSignUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"checkUserStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"decreaseBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserData\",\"outputs\":[{\"internalType\":\"contract UserData\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"increaseBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"signOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"signUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"};

    @Getter public static final String ABI = org.fisco.bcos.sdk.v3.utils.StringUtils.joinAll("", ABI_ARRAY);

    public static final String FUNC_CHECKUSERSTATUS = "checkUserStatus";

    public static final String FUNC_DECREASEBALANCE = "decreaseBalance";

    public static final String FUNC_DESTROY = "destroy";

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

    public TransactionReceipt destroy() {
        final Function function = new Function(
            FUNC_DESTROY,
            List.of(),
            Collections.emptyList(), 0);
        return executeTransaction(function);
    }

    public String getSignedTransactionForDestroy() {
        final Function function = new Function(
            FUNC_DESTROY,
            List.of(),
            Collections.emptyList(), 0);
        return createSignedTransaction(function);
    }

    public String destroy(TransactionCallback callback) {
        final Function function = new Function(
            FUNC_DESTROY,
            List.of(),
            Collections.emptyList(), 0);
        return asyncExecuteTransaction(function, callback);
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
