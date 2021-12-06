// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"}],\"name\":\"OrganizationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"}],\"name\":\"OrganizationDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"}],\"name\":\"OrganizationModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadataURI\",\"type\":\"bytes\"}],\"name\":\"ServiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"ServiceDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadataURI\",\"type\":\"bytes\"}],\"name\":\"ServiceMetadataModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"ServiceTagsModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"orgMetadataURI\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"}],\"name\":\"createOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOrganizationOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"orgMetadataURI\",\"type\":\"bytes\"}],\"name\":\"changeOrganizationMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"newMembers\",\"type\":\"address[]\"}],\"name\":\"addOrganizationMembers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"existingMembers\",\"type\":\"address[]\"}],\"name\":\"removeOrganizationMembers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"}],\"name\":\"deleteOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"metadataURI\",\"type\":\"bytes\"}],\"name\":\"createServiceRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"metadataURI\",\"type\":\"bytes\"}],\"name\":\"updateServiceRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"deleteServiceRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listOrganizations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orgIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"}],\"name\":\"getOrganizationById\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"orgMetadataURI\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"serviceIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"}],\"name\":\"listServicesForOrganization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"bytes32[]\",\"name\":\"serviceIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orgId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"getServiceRegistrationById\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"metadataURI\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `"0x608060405234801561001057600080fd5b5061002a6301ffc9a760e01b6001600160e01b0361004816565b610043631f91217560e11b6001600160e01b0361004816565b6100cc565b6001600160e01b031980821614156100a7576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b611b9a806100db6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a4123f0f1161008c578063d2e544f711610066578063d2e544f71461059c578063d9219d2814610611578063e443852d14610686578063ef72a9af146106b2576100ea565b8063a4123f0f14610485578063bcb43444146104ff578063ca97758c14610579576100ea565b80636928848d116100c85780636928848d146101fe5780638895303f1461021d578063987b3f4d146102ca5780639a017e0a1461033f576100ea565b806301ffc9a7146100ef578063340d84e51461012a5780633a2cb86014610182575b600080fd5b6101166004803603602081101561010557600080fd5b50356001600160e01b031916610777565b604080519115158252519081900360200190f35b610132610796565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561016e578181015183820152602001610156565b505050509050019250505060405180910390f35b61019f6004803603602081101561019857600080fd5b50356107ef565b604051808315151515815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156101e95781810151838201526020016101d1565b50505050905001935050505060405180910390f35b61021b6004803603602081101561021457600080fd5b5035610876565b005b6102406004803603604081101561023357600080fd5b5080359060200135610a77565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561028d578181015183820152602001610275565b50505050905090810190601f1680156102ba5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61021b600480360360408110156102e057600080fd5b81359190810190604081016020820135600160201b81111561030157600080fd5b82018360208201111561031357600080fd5b803590602001918460018302840111600160201b8311171561033457600080fd5b509092509050610b78565b61035c6004803603602081101561035557600080fd5b5035610bdb565b604051808715151515815260200186815260200180602001856001600160a01b03166001600160a01b031681526020018060200180602001848103845288818151815260200191508051906020019080838360005b838110156103c95781810151838201526020016103b1565b50505050905090810190601f1680156103f65780820380516001836020036101000a031916815260200191505b508481038352865181528651602091820191808901910280838360005b8381101561042b578181015183820152602001610413565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561046a578181015183820152602001610452565b50505050905001995050505050505050505060405180910390f35b61021b6004803603606081101561049b57600080fd5b813591602081013591810190606081016040820135600160201b8111156104c157600080fd5b8201836020820111156104d357600080fd5b803590602001918460018302840111600160201b831117156104f457600080fd5b509092509050610d95565b61021b6004803603606081101561051557600080fd5b813591602081013591810190606081016040820135600160201b81111561053b57600080fd5b82018360208201111561054d57600080fd5b803590602001918460018302840111600160201b8311171561056e57600080fd5b509092509050610ed0565b61021b6004803603604081101561058f57600080fd5b5080359060200135610f85565b61021b600480360360408110156105b257600080fd5b81359190810190604081016020820135600160201b8111156105d357600080fd5b8201836020820111156105e557600080fd5b803590602001918460208302840111600160201b8311171561060657600080fd5b509092509050610fe2565b61021b6004803603604081101561062757600080fd5b81359190810190604081016020820135600160201b81111561064857600080fd5b82018360208201111561065a57600080fd5b803590602001918460208302840111600160201b8311171561067b57600080fd5b509092509050611030565b61021b6004803603604081101561069c57600080fd5b50803590602001356001600160a01b03166110b3565b61021b600480360360608110156106c857600080fd5b81359190810190604081016020820135600160201b8111156106e957600080fd5b8201836020820111156106fb57600080fd5b803590602001918460018302840111600160201b8311171561071c57600080fd5b919390929091602081019035600160201b81111561073957600080fd5b82018360208201111561074b57600080fd5b803590602001918460208302840111600160201b8311171561076c57600080fd5b509092509050611120565b6001600160e01b03191660009081526020819052604090205460ff1690565b606060018054806020026020016040519081016040528092919081815260200182805480156107e457602002820191906000526020600020905b8154815260200190600101908083116107d0575b505050505090505b90565b60008181526002602052604081205460609061080e5760009150610871565b600083815260026020908152604091829020600501805483518184028101840190945280845260019550909183018282801561086957602002820191906000526020600020905b815481526020019060010190808311610855575b505050505090505b915091565b6108818160016112b7565b61088c816000611371565b6000818152600260205260409020600501545b80156108e557600082815260026020526040902060050180546108dc91849160001985019081106108cc57fe5b906000526020600020015461140e565b6000190161089f565b506000818152600260205260409020600301545b80156109495760008281526002602052604090206003018054610940918491600019850190811061092657fe5b6000918252602090912001546001600160a01b0316611546565b600019016108f9565b5060008181526002602052604081206007015460018054919291600019810190811061097157fe5b90600052602060002001549050806001838154811061098c57fe5b9060005260206000200154146109ce5780600183815481106109aa57fe5b60009182526020808320909101929092558281526002909152604090206007018290555b60018054806109d957fe5b600082815260208082208301600019908101839055909201909255848252600290526040812081815590610a1060018301826118d4565b6002820180546001600160a01b0319169055610a3060038301600061191b565b610a3e60058301600061191b565b50600060079190910181905560405184917fb1dbc279d80967cb8113073651e5919753c82fad7b002581eb7a020eaf89e9a791a2505050565b6000828152600260205260408120548190606090610a985760009250610b71565b6000858152600260209081526040808320878452600601909152902054610ac25760009250610b71565b6000858152600260208181526040808420888552600601825292839020805460019182018054865181851615610100026000190190911695909504601f81018590048502860185019096528585529197509550919290830182828015610b695780601f10610b3e57610100808354040283529160200191610b69565b820191906000526020600020905b815481529060010190602001808311610b4c57829003601f168201915b505050505090505b9250925092565b610b838360016112b7565b610b8e836000611371565b6000838152600260205260409020610baa906001018383611939565b5060405183907f06ccb920be65231f5c9d04dd4883d3c7648ebe5f5317cc7177ee4f4a7cc2d03890600090a2505050565b6000818152600260205260408120548190606090829082908190610c025760009550610d8c565b600087815260026020818152604092839020805460019182018054865181851615610100026000190190911695909504601f8101859004850286018501909652858552919a509850919290830182828015610c9e5780601f10610c7357610100808354040283529160200191610c9e565b820191906000526020600020905b815481529060010190602001808311610c8157829003601f168201915b50505060008a8152600260208181526040928390209182015460039092018054845181840281018401909552808552969a506001600160a01b03909216985091949093509150830182828015610d1d57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610cff575b5050505050915060026000888152602001908152602001600020600501805480602002602001604051908101604052809291908181526020018280548015610d8457602002820191906000526020600020905b815481526020019060010190808311610d70575b505050505090505b91939550919395565b610da08460016112b7565b610dab846001611371565b610db7848460006116dc565b610dbf6119b7565b838152604080516020601f8501819004810282018101909252838152908490849081908401838280828437600092018290525060208087019586528a82526002815260408083206005810154828a01528b84526006018252909120865181559451805187969550610e3994506001860193509101906119db565b5060409182015160029182015560008781526020918252828120600501805460018101825590825290829020018690558151818152908101849052859187917f3229c7118d95880b1bb8abc6231f47f1a63d7b1e7e22fbd91a8ccffc9fa75df9918791879181908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a35050505050565b610edb8460016112b7565b610ee6846001611371565b610ef2848460016116dc565b60008481526002602090815260408083208684526006019091529020610f1c906001018383611939565b5082847fb7b13a2b2a9c0147b27815cbec2c7e5ed10588c9d5812211619614c379174c5a848460405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a350505050565b610f908260016112b7565b610f9b826001611371565b610fa7828260016116dc565b610fb1828261140e565b604051819083907f3caed2ddcd24bae20b0075a02d974ba3a229dace98271af93ce8c8f3ebd9d27290600090a35050565b610fed8360016112b7565b610ff8836001611371565b60005b81811015610baa576110288484848481811061101357fe5b905060200201356001600160a01b0316611546565b600101610ffb565b61103b8360016112b7565b611046836001611371565b611083838383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506117bb92505050565b60405183907f06ccb920be65231f5c9d04dd4883d3c7648ebe5f5317cc7177ee4f4a7cc2d03890600090a2505050565b6110be8260016112b7565b6110c9826000611371565b600082815260026020819052604080832090910180546001600160a01b0319166001600160a01b0385161790555183917f06ccb920be65231f5c9d04dd4883d3c7648ebe5f5317cc7177ee4f4a7cc2d03891a25050565b61112b8560006112b7565b611133611a49565b60008681526002602090815260409091208251815581830151805184936111619260018501929101906119db565b5060408201516002820180546001600160a01b0319166001600160a01b03909216919091179055606082015180516111a3916003840191602090910190611a8b565b50608082015180516111bf916005840191602090910190611aec565b5060a0919091015160079091015560008681526002602052604090208681556111ec906001018686611939565b50600086815260026020818152604080842092830180546001600160a01b0319163317905560018054600790940184905583810181559093527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690910188905581518482028181018301909352848152611284928992879187918291908501908490808284376000920191909152506117bb92505050565b60405186907f0e7857bfbd020070a2c8d2fe38c788de1e6adc88cb46f60cf6ec7cd385e81db190600090a2505050505050565b801561131757600082815260026020526040902054611312576040805162461bcd60e51b81526020600482015260126024820152711bdc99c8191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b61136d565b6000828152600260205260409020541561136d576040805162461bcd60e51b81526020600482015260126024820152716f726720616c72656164792065786973747360701b604482015290519081900360640190fd5b5050565b600082815260026020819052604090912001546001600160a01b03163314806113bd57508080156113bd5750600082815260026020908152604080832033845260040190915290205415155b61136d576040805162461bcd60e51b815260206004820152601760248201527f756e617574686f72697a656420696e766f636174696f6e000000000000000000604482015290519081900360640190fd5b6000828152600260208181526040808420858552600681018352908420830154868552929091526005018054919291600019810190811061144b57fe5b906000526020600020015490508060026000868152602001908152602001600020600501838154811061147a57fe5b9060005260206000200154146114d95760008481526002602052604090206005018054829190849081106114aa57fe5b600091825260208083209091019290925585815260028083526040808320858452600601909352919020018290555b60008481526002602052604090206005018054806114f357fe5b60008281526020808220830160001990810183905590920190925585825260028152604080832086845260060190915281208181559061153660018301826118d4565b6002820160009055505050505050565b60008281526002602090815260408083206001600160a01b03851684526004019091529020541561136d5760008281526002602081815260408084206001600160a01b03861685526004810183529084205486855292909152600301805491929160001981019081106115b557fe5b60009182526020808320909101548683526002909152604090912060030180546001600160a01b039092169250829160001985019081106115f257fe5b6000918252602090912001546001600160a01b0316146116755760008481526002602052604090206003018054829190600019850190811061163057fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b03948516179055868252600281526040808320938516835260049093019052208290555b600084815260026020526040902060030180548061168f57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092558582526002815260408083206001600160a01b038716845260040190915281205550505050565b801561174e576000838152600260209081526040808320858452600601909152902054611749576040805162461bcd60e51b81526020600482015260166024820152751cd95c9d9a58d948191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6117b6565b6000838152600260209081526040808320858452600601909152902054156117b6576040805162461bcd60e51b81526020600482015260166024820152757365727669636520616c72656164792065786973747360501b604482015290519081900360640190fd5b505050565b60005b81518110156117b6576002600084815260200190815260200160002060040160008383815181106117eb57fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054600014156118cc576002600084815260200190815260200160002060030182828151811061183e57fe5b602090810291909101810151825460018101845560009384528284200180546001600160a01b0319166001600160a01b0390921691909117905584825260029052604081206003810154845190926004909201919085908590811061189f57fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055505b6001016117be565b50805460018160011615610100020316600290046000825580601f106118fa5750611918565b601f0160209004906000526020600020908101906119189190611b26565b50565b50805460008255906000526020600020908101906119189190611b26565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061197a5782800160ff198235161785556119a7565b828001600101855582156119a7579182015b828111156119a757823582559160200191906001019061198c565b506119b3929150611b26565b5090565b60405180606001604052806000801916815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a1c57805160ff19168380011785556119a7565b828001600101855582156119a7579182015b828111156119a7578251825591602001919060010190611a2e565b6040518060c00160405280600080191681526020016060815260200160006001600160a01b031681526020016060815260200160608152602001600081525090565b828054828255906000526020600020908101928215611ae0579160200282015b82811115611ae057825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611aab565b506119b3929150611b40565b8280548282559060005260206000209081019282156119a757916020028201828111156119a7578251825591602001919060010190611a2e565b6107ec91905b808211156119b35760008155600101611b2c565b6107ec91905b808211156119b35780546001600160a01b0319168155600101611b4656fea264697066735822122024200563242f8da6291b3b980955449d1733202dc332f2c79bf517c8012f73bf64736f6c63430006020033"`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetOrganizationById is a free data retrieval call binding the contract method 0x9a017e0a.
//
// Solidity: function getOrganizationById(bytes32 orgId) constant returns(bool found, bytes32 id, bytes orgMetadataURI, address owner, address[] members, bytes32[] serviceIds)
func (_Registry *RegistryCaller) GetOrganizationById(opts *bind.CallOpts, orgId [32]byte) (struct {
	Found          bool
	Id             [32]byte
	OrgMetadataURI []byte
	Owner          common.Address
	Members        []common.Address
	ServiceIds     [][32]byte
}, error) {
	ret := new(struct {
		Found          bool
		Id             [32]byte
		OrgMetadataURI []byte
		Owner          common.Address
		Members        []common.Address
		ServiceIds     [][32]byte
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "getOrganizationById", orgId)
	return *ret, err
}

// GetOrganizationById is a free data retrieval call binding the contract method 0x9a017e0a.
//
// Solidity: function getOrganizationById(bytes32 orgId) constant returns(bool found, bytes32 id, bytes orgMetadataURI, address owner, address[] members, bytes32[] serviceIds)
func (_Registry *RegistrySession) GetOrganizationById(orgId [32]byte) (struct {
	Found          bool
	Id             [32]byte
	OrgMetadataURI []byte
	Owner          common.Address
	Members        []common.Address
	ServiceIds     [][32]byte
}, error) {
	return _Registry.Contract.GetOrganizationById(&_Registry.CallOpts, orgId)
}

// GetOrganizationById is a free data retrieval call binding the contract method 0x9a017e0a.
//
// Solidity: function getOrganizationById(bytes32 orgId) constant returns(bool found, bytes32 id, bytes orgMetadataURI, address owner, address[] members, bytes32[] serviceIds)
func (_Registry *RegistryCallerSession) GetOrganizationById(orgId [32]byte) (struct {
	Found          bool
	Id             [32]byte
	OrgMetadataURI []byte
	Owner          common.Address
	Members        []common.Address
	ServiceIds     [][32]byte
}, error) {
	return _Registry.Contract.GetOrganizationById(&_Registry.CallOpts, orgId)
}

// GetServiceRegistrationById is a free data retrieval call binding the contract method 0x8895303f.
//
// Solidity: function getServiceRegistrationById(bytes32 orgId, bytes32 serviceId) constant returns(bool found, bytes32 id, bytes metadataURI)
func (_Registry *RegistryCaller) GetServiceRegistrationById(opts *bind.CallOpts, orgId [32]byte, serviceId [32]byte) (struct {
	Found       bool
	Id          [32]byte
	MetadataURI []byte
}, error) {
	ret := new(struct {
		Found       bool
		Id          [32]byte
		MetadataURI []byte
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "getServiceRegistrationById", orgId, serviceId)
	return *ret, err
}

// GetServiceRegistrationById is a free data retrieval call binding the contract method 0x8895303f.
//
// Solidity: function getServiceRegistrationById(bytes32 orgId, bytes32 serviceId) constant returns(bool found, bytes32 id, bytes metadataURI)
func (_Registry *RegistrySession) GetServiceRegistrationById(orgId [32]byte, serviceId [32]byte) (struct {
	Found       bool
	Id          [32]byte
	MetadataURI []byte
}, error) {
	return _Registry.Contract.GetServiceRegistrationById(&_Registry.CallOpts, orgId, serviceId)
}

// GetServiceRegistrationById is a free data retrieval call binding the contract method 0x8895303f.
//
// Solidity: function getServiceRegistrationById(bytes32 orgId, bytes32 serviceId) constant returns(bool found, bytes32 id, bytes metadataURI)
func (_Registry *RegistryCallerSession) GetServiceRegistrationById(orgId [32]byte, serviceId [32]byte) (struct {
	Found       bool
	Id          [32]byte
	MetadataURI []byte
}, error) {
	return _Registry.Contract.GetServiceRegistrationById(&_Registry.CallOpts, orgId, serviceId)
}

// ListOrganizations is a free data retrieval call binding the contract method 0x340d84e5.
//
// Solidity: function listOrganizations() constant returns(bytes32[] orgIds)
func (_Registry *RegistryCaller) ListOrganizations(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "listOrganizations")
	return *ret0, err
}

// ListOrganizations is a free data retrieval call binding the contract method 0x340d84e5.
//
// Solidity: function listOrganizations() constant returns(bytes32[] orgIds)
func (_Registry *RegistrySession) ListOrganizations() ([][32]byte, error) {
	return _Registry.Contract.ListOrganizations(&_Registry.CallOpts)
}

// ListOrganizations is a free data retrieval call binding the contract method 0x340d84e5.
//
// Solidity: function listOrganizations() constant returns(bytes32[] orgIds)
func (_Registry *RegistryCallerSession) ListOrganizations() ([][32]byte, error) {
	return _Registry.Contract.ListOrganizations(&_Registry.CallOpts)
}

// ListServicesForOrganization is a free data retrieval call binding the contract method 0x3a2cb860.
//
// Solidity: function listServicesForOrganization(bytes32 orgId) constant returns(bool found, bytes32[] serviceIds)
func (_Registry *RegistryCaller) ListServicesForOrganization(opts *bind.CallOpts, orgId [32]byte) (struct {
	Found      bool
	ServiceIds [][32]byte
}, error) {
	ret := new(struct {
		Found      bool
		ServiceIds [][32]byte
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "listServicesForOrganization", orgId)
	return *ret, err
}

// ListServicesForOrganization is a free data retrieval call binding the contract method 0x3a2cb860.
//
// Solidity: function listServicesForOrganization(bytes32 orgId) constant returns(bool found, bytes32[] serviceIds)
func (_Registry *RegistrySession) ListServicesForOrganization(orgId [32]byte) (struct {
	Found      bool
	ServiceIds [][32]byte
}, error) {
	return _Registry.Contract.ListServicesForOrganization(&_Registry.CallOpts, orgId)
}

// ListServicesForOrganization is a free data retrieval call binding the contract method 0x3a2cb860.
//
// Solidity: function listServicesForOrganization(bytes32 orgId) constant returns(bool found, bytes32[] serviceIds)
func (_Registry *RegistryCallerSession) ListServicesForOrganization(orgId [32]byte) (struct {
	Found      bool
	ServiceIds [][32]byte
}, error) {
	return _Registry.Contract.ListServicesForOrganization(&_Registry.CallOpts, orgId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Registry *RegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Registry *RegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Registry *RegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// AddOrganizationMembers is a paid mutator transaction binding the contract method 0xd9219d28.
//
// Solidity: function addOrganizationMembers(bytes32 orgId, address[] newMembers) returns()
func (_Registry *RegistryTransactor) AddOrganizationMembers(opts *bind.TransactOpts, orgId [32]byte, newMembers []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addOrganizationMembers", orgId, newMembers)
}

// AddOrganizationMembers is a paid mutator transaction binding the contract method 0xd9219d28.
//
// Solidity: function addOrganizationMembers(bytes32 orgId, address[] newMembers) returns()
func (_Registry *RegistrySession) AddOrganizationMembers(orgId [32]byte, newMembers []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddOrganizationMembers(&_Registry.TransactOpts, orgId, newMembers)
}

// AddOrganizationMembers is a paid mutator transaction binding the contract method 0xd9219d28.
//
// Solidity: function addOrganizationMembers(bytes32 orgId, address[] newMembers) returns()
func (_Registry *RegistryTransactorSession) AddOrganizationMembers(orgId [32]byte, newMembers []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddOrganizationMembers(&_Registry.TransactOpts, orgId, newMembers)
}

// ChangeOrganizationMetadataURI is a paid mutator transaction binding the contract method 0x987b3f4d.
//
// Solidity: function changeOrganizationMetadataURI(bytes32 orgId, bytes orgMetadataURI) returns()
func (_Registry *RegistryTransactor) ChangeOrganizationMetadataURI(opts *bind.TransactOpts, orgId [32]byte, orgMetadataURI []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeOrganizationMetadataURI", orgId, orgMetadataURI)
}

// ChangeOrganizationMetadataURI is a paid mutator transaction binding the contract method 0x987b3f4d.
//
// Solidity: function changeOrganizationMetadataURI(bytes32 orgId, bytes orgMetadataURI) returns()
func (_Registry *RegistrySession) ChangeOrganizationMetadataURI(orgId [32]byte, orgMetadataURI []byte) (*types.Transaction, error) {
	return _Registry.Contract.ChangeOrganizationMetadataURI(&_Registry.TransactOpts, orgId, orgMetadataURI)
}

// ChangeOrganizationMetadataURI is a paid mutator transaction binding the contract method 0x987b3f4d.
//
// Solidity: function changeOrganizationMetadataURI(bytes32 orgId, bytes orgMetadataURI) returns()
func (_Registry *RegistryTransactorSession) ChangeOrganizationMetadataURI(orgId [32]byte, orgMetadataURI []byte) (*types.Transaction, error) {
	return _Registry.Contract.ChangeOrganizationMetadataURI(&_Registry.TransactOpts, orgId, orgMetadataURI)
}

// ChangeOrganizationOwner is a paid mutator transaction binding the contract method 0xe443852d.
//
// Solidity: function changeOrganizationOwner(bytes32 orgId, address newOwner) returns()
func (_Registry *RegistryTransactor) ChangeOrganizationOwner(opts *bind.TransactOpts, orgId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeOrganizationOwner", orgId, newOwner)
}

// ChangeOrganizationOwner is a paid mutator transaction binding the contract method 0xe443852d.
//
// Solidity: function changeOrganizationOwner(bytes32 orgId, address newOwner) returns()
func (_Registry *RegistrySession) ChangeOrganizationOwner(orgId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeOrganizationOwner(&_Registry.TransactOpts, orgId, newOwner)
}

// ChangeOrganizationOwner is a paid mutator transaction binding the contract method 0xe443852d.
//
// Solidity: function changeOrganizationOwner(bytes32 orgId, address newOwner) returns()
func (_Registry *RegistryTransactorSession) ChangeOrganizationOwner(orgId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeOrganizationOwner(&_Registry.TransactOpts, orgId, newOwner)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xef72a9af.
//
// Solidity: function createOrganization(bytes32 orgId, bytes orgMetadataURI, address[] members) returns()
func (_Registry *RegistryTransactor) CreateOrganization(opts *bind.TransactOpts, orgId [32]byte, orgMetadataURI []byte, members []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createOrganization", orgId, orgMetadataURI, members)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xef72a9af.
//
// Solidity: function createOrganization(bytes32 orgId, bytes orgMetadataURI, address[] members) returns()
func (_Registry *RegistrySession) CreateOrganization(orgId [32]byte, orgMetadataURI []byte, members []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.CreateOrganization(&_Registry.TransactOpts, orgId, orgMetadataURI, members)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xef72a9af.
//
// Solidity: function createOrganization(bytes32 orgId, bytes orgMetadataURI, address[] members) returns()
func (_Registry *RegistryTransactorSession) CreateOrganization(orgId [32]byte, orgMetadataURI []byte, members []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.CreateOrganization(&_Registry.TransactOpts, orgId, orgMetadataURI, members)
}

// CreateServiceRegistration is a paid mutator transaction binding the contract method 0xa4123f0f.
//
// Solidity: function createServiceRegistration(bytes32 orgId, bytes32 serviceId, bytes metadataURI) returns()
func (_Registry *RegistryTransactor) CreateServiceRegistration(opts *bind.TransactOpts, orgId [32]byte, serviceId [32]byte, metadataURI []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createServiceRegistration", orgId, serviceId, metadataURI)
}

// CreateServiceRegistration is a paid mutator transaction binding the contract method 0xa4123f0f.
//
// Solidity: function createServiceRegistration(bytes32 orgId, bytes32 serviceId, bytes metadataURI) returns()
func (_Registry *RegistrySession) CreateServiceRegistration(orgId [32]byte, serviceId [32]byte, metadataURI []byte) (*types.Transaction, error) {
	return _Registry.Contract.CreateServiceRegistration(&_Registry.TransactOpts, orgId, serviceId, metadataURI)
}

// CreateServiceRegistration is a paid mutator transaction binding the contract method 0xa4123f0f.
//
// Solidity: function createServiceRegistration(bytes32 orgId, bytes32 serviceId, bytes metadataURI) returns()
func (_Registry *RegistryTransactorSession) CreateServiceRegistration(orgId [32]byte, serviceId [32]byte, metadataURI []byte) (*types.Transaction, error) {
	return _Registry.Contract.CreateServiceRegistration(&_Registry.TransactOpts, orgId, serviceId, metadataURI)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x6928848d.
//
// Solidity: function deleteOrganization(bytes32 orgId) returns()
func (_Registry *RegistryTransactor) DeleteOrganization(opts *bind.TransactOpts, orgId [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deleteOrganization", orgId)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x6928848d.
//
// Solidity: function deleteOrganization(bytes32 orgId) returns()
func (_Registry *RegistrySession) DeleteOrganization(orgId [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.DeleteOrganization(&_Registry.TransactOpts, orgId)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x6928848d.
//
// Solidity: function deleteOrganization(bytes32 orgId) returns()
func (_Registry *RegistryTransactorSession) DeleteOrganization(orgId [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.DeleteOrganization(&_Registry.TransactOpts, orgId)
}

// DeleteServiceRegistration is a paid mutator transaction binding the contract method 0xca97758c.
//
// Solidity: function deleteServiceRegistration(bytes32 orgId, bytes32 serviceId) returns()
func (_Registry *RegistryTransactor) DeleteServiceRegistration(opts *bind.TransactOpts, orgId [32]byte, serviceId [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deleteServiceRegistration", orgId, serviceId)
}

// DeleteServiceRegistration is a paid mutator transaction binding the contract method 0xca97758c.
//
// Solidity: function deleteServiceRegistration(bytes32 orgId, bytes32 serviceId) returns()
func (_Registry *RegistrySession) DeleteServiceRegistration(orgId [32]byte, serviceId [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.DeleteServiceRegistration(&_Registry.TransactOpts, orgId, serviceId)
}

// DeleteServiceRegistration is a paid mutator transaction binding the contract method 0xca97758c.
//
// Solidity: function deleteServiceRegistration(bytes32 orgId, bytes32 serviceId) returns()
func (_Registry *RegistryTransactorSession) DeleteServiceRegistration(orgId [32]byte, serviceId [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.DeleteServiceRegistration(&_Registry.TransactOpts, orgId, serviceId)
}

// RemoveOrganizationMembers is a paid mutator transaction binding the contract method 0xd2e544f7.
//
// Solidity: function removeOrganizationMembers(bytes32 orgId, address[] existingMembers) returns()
func (_Registry *RegistryTransactor) RemoveOrganizationMembers(opts *bind.TransactOpts, orgId [32]byte, existingMembers []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "removeOrganizationMembers", orgId, existingMembers)
}

// RemoveOrganizationMembers is a paid mutator transaction binding the contract method 0xd2e544f7.
//
// Solidity: function removeOrganizationMembers(bytes32 orgId, address[] existingMembers) returns()
func (_Registry *RegistrySession) RemoveOrganizationMembers(orgId [32]byte, existingMembers []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveOrganizationMembers(&_Registry.TransactOpts, orgId, existingMembers)
}

// RemoveOrganizationMembers is a paid mutator transaction binding the contract method 0xd2e544f7.
//
// Solidity: function removeOrganizationMembers(bytes32 orgId, address[] existingMembers) returns()
func (_Registry *RegistryTransactorSession) RemoveOrganizationMembers(orgId [32]byte, existingMembers []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveOrganizationMembers(&_Registry.TransactOpts, orgId, existingMembers)
}

// UpdateServiceRegistration is a paid mutator transaction binding the contract method 0xbcb43444.
//
// Solidity: function updateServiceRegistration(bytes32 orgId, bytes32 serviceId, bytes metadataURI) returns()
func (_Registry *RegistryTransactor) UpdateServiceRegistration(opts *bind.TransactOpts, orgId [32]byte, serviceId [32]byte, metadataURI []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateServiceRegistration", orgId, serviceId, metadataURI)
}

// UpdateServiceRegistration is a paid mutator transaction binding the contract method 0xbcb43444.
//
// Solidity: function updateServiceRegistration(bytes32 orgId, bytes32 serviceId, bytes metadataURI) returns()
func (_Registry *RegistrySession) UpdateServiceRegistration(orgId [32]byte, serviceId [32]byte, metadataURI []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateServiceRegistration(&_Registry.TransactOpts, orgId, serviceId, metadataURI)
}

// UpdateServiceRegistration is a paid mutator transaction binding the contract method 0xbcb43444.
//
// Solidity: function updateServiceRegistration(bytes32 orgId, bytes32 serviceId, bytes metadataURI) returns()
func (_Registry *RegistryTransactorSession) UpdateServiceRegistration(orgId [32]byte, serviceId [32]byte, metadataURI []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateServiceRegistration(&_Registry.TransactOpts, orgId, serviceId, metadataURI)
}

// RegistryOrganizationCreatedIterator is returned from FilterOrganizationCreated and is used to iterate over the raw logs and unpacked data for OrganizationCreated events raised by the Registry contract.
type RegistryOrganizationCreatedIterator struct {
	Event *RegistryOrganizationCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOrganizationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOrganizationCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOrganizationCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOrganizationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOrganizationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOrganizationCreated represents a OrganizationCreated event raised by the Registry contract.
type RegistryOrganizationCreated struct {
	OrgId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrganizationCreated is a free log retrieval operation binding the contract event 0x0e7857bfbd020070a2c8d2fe38c788de1e6adc88cb46f60cf6ec7cd385e81db1.
//
// Solidity: event OrganizationCreated(bytes32 indexed orgId)
func (_Registry *RegistryFilterer) FilterOrganizationCreated(opts *bind.FilterOpts, orgId [][32]byte) (*RegistryOrganizationCreatedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OrganizationCreated", orgIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOrganizationCreatedIterator{contract: _Registry.contract, event: "OrganizationCreated", logs: logs, sub: sub}, nil
}

// WatchOrganizationCreated is a free log subscription operation binding the contract event 0x0e7857bfbd020070a2c8d2fe38c788de1e6adc88cb46f60cf6ec7cd385e81db1.
//
// Solidity: event OrganizationCreated(bytes32 indexed orgId)
func (_Registry *RegistryFilterer) WatchOrganizationCreated(opts *bind.WatchOpts, sink chan<- *RegistryOrganizationCreated, orgId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OrganizationCreated", orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOrganizationCreated)
				if err := _Registry.contract.UnpackLog(event, "OrganizationCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistryOrganizationDeletedIterator is returned from FilterOrganizationDeleted and is used to iterate over the raw logs and unpacked data for OrganizationDeleted events raised by the Registry contract.
type RegistryOrganizationDeletedIterator struct {
	Event *RegistryOrganizationDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOrganizationDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOrganizationDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOrganizationDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOrganizationDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOrganizationDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOrganizationDeleted represents a OrganizationDeleted event raised by the Registry contract.
type RegistryOrganizationDeleted struct {
	OrgId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrganizationDeleted is a free log retrieval operation binding the contract event 0xb1dbc279d80967cb8113073651e5919753c82fad7b002581eb7a020eaf89e9a7.
//
// Solidity: event OrganizationDeleted(bytes32 indexed orgId)
func (_Registry *RegistryFilterer) FilterOrganizationDeleted(opts *bind.FilterOpts, orgId [][32]byte) (*RegistryOrganizationDeletedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OrganizationDeleted", orgIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOrganizationDeletedIterator{contract: _Registry.contract, event: "OrganizationDeleted", logs: logs, sub: sub}, nil
}

// WatchOrganizationDeleted is a free log subscription operation binding the contract event 0xb1dbc279d80967cb8113073651e5919753c82fad7b002581eb7a020eaf89e9a7.
//
// Solidity: event OrganizationDeleted(bytes32 indexed orgId)
func (_Registry *RegistryFilterer) WatchOrganizationDeleted(opts *bind.WatchOpts, sink chan<- *RegistryOrganizationDeleted, orgId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OrganizationDeleted", orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOrganizationDeleted)
				if err := _Registry.contract.UnpackLog(event, "OrganizationDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistryOrganizationModifiedIterator is returned from FilterOrganizationModified and is used to iterate over the raw logs and unpacked data for OrganizationModified events raised by the Registry contract.
type RegistryOrganizationModifiedIterator struct {
	Event *RegistryOrganizationModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOrganizationModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOrganizationModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOrganizationModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOrganizationModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOrganizationModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOrganizationModified represents a OrganizationModified event raised by the Registry contract.
type RegistryOrganizationModified struct {
	OrgId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrganizationModified is a free log retrieval operation binding the contract event 0x06ccb920be65231f5c9d04dd4883d3c7648ebe5f5317cc7177ee4f4a7cc2d038.
//
// Solidity: event OrganizationModified(bytes32 indexed orgId)
func (_Registry *RegistryFilterer) FilterOrganizationModified(opts *bind.FilterOpts, orgId [][32]byte) (*RegistryOrganizationModifiedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OrganizationModified", orgIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOrganizationModifiedIterator{contract: _Registry.contract, event: "OrganizationModified", logs: logs, sub: sub}, nil
}

// WatchOrganizationModified is a free log subscription operation binding the contract event 0x06ccb920be65231f5c9d04dd4883d3c7648ebe5f5317cc7177ee4f4a7cc2d038.
//
// Solidity: event OrganizationModified(bytes32 indexed orgId)
func (_Registry *RegistryFilterer) WatchOrganizationModified(opts *bind.WatchOpts, sink chan<- *RegistryOrganizationModified, orgId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OrganizationModified", orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOrganizationModified)
				if err := _Registry.contract.UnpackLog(event, "OrganizationModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistryServiceCreatedIterator is returned from FilterServiceCreated and is used to iterate over the raw logs and unpacked data for ServiceCreated events raised by the Registry contract.
type RegistryServiceCreatedIterator struct {
	Event *RegistryServiceCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryServiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryServiceCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryServiceCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryServiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryServiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryServiceCreated represents a ServiceCreated event raised by the Registry contract.
type RegistryServiceCreated struct {
	OrgId       [32]byte
	ServiceId   [32]byte
	MetadataURI []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCreated is a free log retrieval operation binding the contract event 0x3229c7118d95880b1bb8abc6231f47f1a63d7b1e7e22fbd91a8ccffc9fa75df9.
//
// Solidity: event ServiceCreated(bytes32 indexed orgId, bytes32 indexed serviceId, bytes metadataURI)
func (_Registry *RegistryFilterer) FilterServiceCreated(opts *bind.FilterOpts, orgId [][32]byte, serviceId [][32]byte) (*RegistryServiceCreatedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ServiceCreated", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryServiceCreatedIterator{contract: _Registry.contract, event: "ServiceCreated", logs: logs, sub: sub}, nil
}

// WatchServiceCreated is a free log subscription operation binding the contract event 0x3229c7118d95880b1bb8abc6231f47f1a63d7b1e7e22fbd91a8ccffc9fa75df9.
//
// Solidity: event ServiceCreated(bytes32 indexed orgId, bytes32 indexed serviceId, bytes metadataURI)
func (_Registry *RegistryFilterer) WatchServiceCreated(opts *bind.WatchOpts, sink chan<- *RegistryServiceCreated, orgId [][32]byte, serviceId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ServiceCreated", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryServiceCreated)
				if err := _Registry.contract.UnpackLog(event, "ServiceCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistryServiceDeletedIterator is returned from FilterServiceDeleted and is used to iterate over the raw logs and unpacked data for ServiceDeleted events raised by the Registry contract.
type RegistryServiceDeletedIterator struct {
	Event *RegistryServiceDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryServiceDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryServiceDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryServiceDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryServiceDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryServiceDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryServiceDeleted represents a ServiceDeleted event raised by the Registry contract.
type RegistryServiceDeleted struct {
	OrgId     [32]byte
	ServiceId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceDeleted is a free log retrieval operation binding the contract event 0x3caed2ddcd24bae20b0075a02d974ba3a229dace98271af93ce8c8f3ebd9d272.
//
// Solidity: event ServiceDeleted(bytes32 indexed orgId, bytes32 indexed serviceId)
func (_Registry *RegistryFilterer) FilterServiceDeleted(opts *bind.FilterOpts, orgId [][32]byte, serviceId [][32]byte) (*RegistryServiceDeletedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ServiceDeleted", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryServiceDeletedIterator{contract: _Registry.contract, event: "ServiceDeleted", logs: logs, sub: sub}, nil
}

// WatchServiceDeleted is a free log subscription operation binding the contract event 0x3caed2ddcd24bae20b0075a02d974ba3a229dace98271af93ce8c8f3ebd9d272.
//
// Solidity: event ServiceDeleted(bytes32 indexed orgId, bytes32 indexed serviceId)
func (_Registry *RegistryFilterer) WatchServiceDeleted(opts *bind.WatchOpts, sink chan<- *RegistryServiceDeleted, orgId [][32]byte, serviceId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ServiceDeleted", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryServiceDeleted)
				if err := _Registry.contract.UnpackLog(event, "ServiceDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistryServiceMetadataModifiedIterator is returned from FilterServiceMetadataModified and is used to iterate over the raw logs and unpacked data for ServiceMetadataModified events raised by the Registry contract.
type RegistryServiceMetadataModifiedIterator struct {
	Event *RegistryServiceMetadataModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryServiceMetadataModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryServiceMetadataModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryServiceMetadataModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryServiceMetadataModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryServiceMetadataModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryServiceMetadataModified represents a ServiceMetadataModified event raised by the Registry contract.
type RegistryServiceMetadataModified struct {
	OrgId       [32]byte
	ServiceId   [32]byte
	MetadataURI []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceMetadataModified is a free log retrieval operation binding the contract event 0xb7b13a2b2a9c0147b27815cbec2c7e5ed10588c9d5812211619614c379174c5a.
//
// Solidity: event ServiceMetadataModified(bytes32 indexed orgId, bytes32 indexed serviceId, bytes metadataURI)
func (_Registry *RegistryFilterer) FilterServiceMetadataModified(opts *bind.FilterOpts, orgId [][32]byte, serviceId [][32]byte) (*RegistryServiceMetadataModifiedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ServiceMetadataModified", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryServiceMetadataModifiedIterator{contract: _Registry.contract, event: "ServiceMetadataModified", logs: logs, sub: sub}, nil
}

// WatchServiceMetadataModified is a free log subscription operation binding the contract event 0xb7b13a2b2a9c0147b27815cbec2c7e5ed10588c9d5812211619614c379174c5a.
//
// Solidity: event ServiceMetadataModified(bytes32 indexed orgId, bytes32 indexed serviceId, bytes metadataURI)
func (_Registry *RegistryFilterer) WatchServiceMetadataModified(opts *bind.WatchOpts, sink chan<- *RegistryServiceMetadataModified, orgId [][32]byte, serviceId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ServiceMetadataModified", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryServiceMetadataModified)
				if err := _Registry.contract.UnpackLog(event, "ServiceMetadataModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistryServiceTagsModifiedIterator is returned from FilterServiceTagsModified and is used to iterate over the raw logs and unpacked data for ServiceTagsModified events raised by the Registry contract.
type RegistryServiceTagsModifiedIterator struct {
	Event *RegistryServiceTagsModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryServiceTagsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryServiceTagsModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryServiceTagsModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryServiceTagsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryServiceTagsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryServiceTagsModified represents a ServiceTagsModified event raised by the Registry contract.
type RegistryServiceTagsModified struct {
	OrgId     [32]byte
	ServiceId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceTagsModified is a free log retrieval operation binding the contract event 0xd8b715d9bf49bca034a993b5b934475414e84c904dcbf90dbddd139808b97b05.
//
// Solidity: event ServiceTagsModified(bytes32 indexed orgId, bytes32 indexed serviceId)
func (_Registry *RegistryFilterer) FilterServiceTagsModified(opts *bind.FilterOpts, orgId [][32]byte, serviceId [][32]byte) (*RegistryServiceTagsModifiedIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ServiceTagsModified", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryServiceTagsModifiedIterator{contract: _Registry.contract, event: "ServiceTagsModified", logs: logs, sub: sub}, nil
}

// WatchServiceTagsModified is a free log subscription operation binding the contract event 0xd8b715d9bf49bca034a993b5b934475414e84c904dcbf90dbddd139808b97b05.
//
// Solidity: event ServiceTagsModified(bytes32 indexed orgId, bytes32 indexed serviceId)
func (_Registry *RegistryFilterer) WatchServiceTagsModified(opts *bind.WatchOpts, sink chan<- *RegistryServiceTagsModified, orgId [][32]byte, serviceId [][32]byte) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var serviceIdRule []interface{}
	for _, serviceIdItem := range serviceId {
		serviceIdRule = append(serviceIdRule, serviceIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ServiceTagsModified", orgIdRule, serviceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryServiceTagsModified)
				if err := _Registry.contract.UnpackLog(event, "ServiceTagsModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
