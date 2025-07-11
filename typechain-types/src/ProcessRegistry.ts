/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "../common";

export declare namespace IProcessRegistry {
  export type EncryptionKeyStruct = { x: BigNumberish; y: BigNumberish };

  export type EncryptionKeyStructOutput = [x: bigint, y: bigint] & {
    x: bigint;
    y: bigint;
  };

  export type BallotModeStruct = {
    costFromWeight: boolean;
    forceUniqueness: boolean;
    maxCount: BigNumberish;
    costExponent: BigNumberish;
    maxValue: BigNumberish;
    minValue: BigNumberish;
    maxTotalCost: BigNumberish;
    minTotalCost: BigNumberish;
  };

  export type BallotModeStructOutput = [
    costFromWeight: boolean,
    forceUniqueness: boolean,
    maxCount: bigint,
    costExponent: bigint,
    maxValue: bigint,
    minValue: bigint,
    maxTotalCost: bigint,
    minTotalCost: bigint
  ] & {
    costFromWeight: boolean;
    forceUniqueness: boolean;
    maxCount: bigint;
    costExponent: bigint;
    maxValue: bigint;
    minValue: bigint;
    maxTotalCost: bigint;
    minTotalCost: bigint;
  };

  export type CensusStruct = {
    censusOrigin: BigNumberish;
    maxVotes: BigNumberish;
    censusRoot: BytesLike;
    censusURI: string;
  };

  export type CensusStructOutput = [
    censusOrigin: bigint,
    maxVotes: bigint,
    censusRoot: string,
    censusURI: string
  ] & {
    censusOrigin: bigint;
    maxVotes: bigint;
    censusRoot: string;
    censusURI: string;
  };

  export type ProcessStruct = {
    status: BigNumberish;
    organizationId: AddressLike;
    encryptionKey: IProcessRegistry.EncryptionKeyStruct;
    latestStateRoot: BigNumberish;
    result: BigNumberish[];
    startTime: BigNumberish;
    duration: BigNumberish;
    voteCount: BigNumberish;
    voteOverwriteCount: BigNumberish;
    metadataURI: string;
    ballotMode: IProcessRegistry.BallotModeStruct;
    census: IProcessRegistry.CensusStruct;
  };

  export type ProcessStructOutput = [
    status: bigint,
    organizationId: string,
    encryptionKey: IProcessRegistry.EncryptionKeyStructOutput,
    latestStateRoot: bigint,
    result: bigint[],
    startTime: bigint,
    duration: bigint,
    voteCount: bigint,
    voteOverwriteCount: bigint,
    metadataURI: string,
    ballotMode: IProcessRegistry.BallotModeStructOutput,
    census: IProcessRegistry.CensusStructOutput
  ] & {
    status: bigint;
    organizationId: string;
    encryptionKey: IProcessRegistry.EncryptionKeyStructOutput;
    latestStateRoot: bigint;
    result: bigint[];
    startTime: bigint;
    duration: bigint;
    voteCount: bigint;
    voteOverwriteCount: bigint;
    metadataURI: string;
    ballotMode: IProcessRegistry.BallotModeStructOutput;
    census: IProcessRegistry.CensusStructOutput;
  };
}

export interface ProcessRegistryInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "MAX_CENSUS_ORIGIN"
      | "MAX_STATUS"
      | "UPGRADE_INTERFACE_VERSION"
      | "chainID"
      | "getNextProcessId"
      | "getProcess"
      | "getProcessEndTime"
      | "getRVerifierVKeyHash"
      | "getSTVerifierVKeyHash"
      | "initialize"
      | "newProcess"
      | "owner"
      | "processCount"
      | "processNonce"
      | "processes"
      | "proxiableUUID"
      | "rVerifier"
      | "renounceOwnership"
      | "setProcessCensus"
      | "setProcessDuration"
      | "setProcessResults"
      | "setProcessStatus"
      | "stVerifier"
      | "submitStateTransition"
      | "transferOwnership"
      | "upgradeToAndCall"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "CensusUpdated"
      | "Initialized"
      | "OwnershipTransferred"
      | "ProcessCreated"
      | "ProcessDurationChanged"
      | "ProcessResultsSet"
      | "ProcessStateRootUpdated"
      | "ProcessStatusChanged"
      | "Upgraded"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "MAX_CENSUS_ORIGIN",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "MAX_STATUS",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "chainID", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "getNextProcessId",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getProcess",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getProcessEndTime",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getRVerifierVKeyHash",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getSTVerifierVKeyHash",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [BigNumberish, AddressLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "newProcess",
    values: [
      BigNumberish,
      BigNumberish,
      BigNumberish,
      IProcessRegistry.BallotModeStruct,
      IProcessRegistry.CensusStruct,
      string,
      IProcessRegistry.EncryptionKeyStruct,
      BigNumberish
    ]
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "processCount",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "processNonce",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "processes",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "proxiableUUID",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "rVerifier", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "setProcessCensus",
    values: [BytesLike, IProcessRegistry.CensusStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "setProcessDuration",
    values: [BytesLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "setProcessResults",
    values: [BytesLike, BytesLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setProcessStatus",
    values: [BytesLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "stVerifier",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "submitStateTransition",
    values: [BytesLike, BytesLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall",
    values: [AddressLike, BytesLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "MAX_CENSUS_ORIGIN",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "MAX_STATUS", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "chainID", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getNextProcessId",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "getProcess", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getProcessEndTime",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getRVerifierVKeyHash",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getSTVerifierVKeyHash",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "newProcess", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "processCount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "processNonce",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "processes", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "proxiableUUID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "rVerifier", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setProcessCensus",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setProcessDuration",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setProcessResults",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setProcessStatus",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "stVerifier", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "submitStateTransition",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall",
    data: BytesLike
  ): Result;
}

export namespace CensusUpdatedEvent {
  export type InputTuple = [
    processId: BytesLike,
    censusRoot: BytesLike,
    censusURI: string,
    maxVotes: BigNumberish
  ];
  export type OutputTuple = [
    processId: string,
    censusRoot: string,
    censusURI: string,
    maxVotes: bigint
  ];
  export interface OutputObject {
    processId: string;
    censusRoot: string;
    censusURI: string;
    maxVotes: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace InitializedEvent {
  export type InputTuple = [version: BigNumberish];
  export type OutputTuple = [version: bigint];
  export interface OutputObject {
    version: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace OwnershipTransferredEvent {
  export type InputTuple = [previousOwner: AddressLike, newOwner: AddressLike];
  export type OutputTuple = [previousOwner: string, newOwner: string];
  export interface OutputObject {
    previousOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ProcessCreatedEvent {
  export type InputTuple = [processId: BytesLike, creator: AddressLike];
  export type OutputTuple = [processId: string, creator: string];
  export interface OutputObject {
    processId: string;
    creator: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ProcessDurationChangedEvent {
  export type InputTuple = [processId: BytesLike, duration: BigNumberish];
  export type OutputTuple = [processId: string, duration: bigint];
  export interface OutputObject {
    processId: string;
    duration: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ProcessResultsSetEvent {
  export type InputTuple = [
    processId: BytesLike,
    sender: AddressLike,
    result: BigNumberish[]
  ];
  export type OutputTuple = [
    processId: string,
    sender: string,
    result: bigint[]
  ];
  export interface OutputObject {
    processId: string;
    sender: string;
    result: bigint[];
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ProcessStateRootUpdatedEvent {
  export type InputTuple = [
    processId: BytesLike,
    sender: AddressLike,
    newStateRoot: BigNumberish
  ];
  export type OutputTuple = [
    processId: string,
    sender: string,
    newStateRoot: bigint
  ];
  export interface OutputObject {
    processId: string;
    sender: string;
    newStateRoot: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ProcessStatusChangedEvent {
  export type InputTuple = [
    processId: BytesLike,
    oldStatus: BigNumberish,
    newStatus: BigNumberish
  ];
  export type OutputTuple = [
    processId: string,
    oldStatus: bigint,
    newStatus: bigint
  ];
  export interface OutputObject {
    processId: string;
    oldStatus: bigint;
    newStatus: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace UpgradedEvent {
  export type InputTuple = [implementation: AddressLike];
  export type OutputTuple = [implementation: string];
  export interface OutputObject {
    implementation: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface ProcessRegistry extends BaseContract {
  connect(runner?: ContractRunner | null): ProcessRegistry;
  waitForDeployment(): Promise<this>;

  interface: ProcessRegistryInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  MAX_CENSUS_ORIGIN: TypedContractMethod<[], [bigint], "view">;

  MAX_STATUS: TypedContractMethod<[], [bigint], "view">;

  UPGRADE_INTERFACE_VERSION: TypedContractMethod<[], [string], "view">;

  chainID: TypedContractMethod<[], [bigint], "view">;

  getNextProcessId: TypedContractMethod<
    [organizationId: AddressLike],
    [string],
    "view"
  >;

  getProcess: TypedContractMethod<
    [processId: BytesLike],
    [IProcessRegistry.ProcessStructOutput],
    "view"
  >;

  getProcessEndTime: TypedContractMethod<
    [processId: BytesLike],
    [bigint],
    "view"
  >;

  getRVerifierVKeyHash: TypedContractMethod<[], [string], "view">;

  getSTVerifierVKeyHash: TypedContractMethod<[], [string], "view">;

  initialize: TypedContractMethod<
    [_chainID: BigNumberish, _stVerifier: AddressLike, _rVerifier: AddressLike],
    [void],
    "nonpayable"
  >;

  newProcess: TypedContractMethod<
    [
      status: BigNumberish,
      startTime: BigNumberish,
      duration: BigNumberish,
      ballotMode: IProcessRegistry.BallotModeStruct,
      census: IProcessRegistry.CensusStruct,
      metadata: string,
      encryptionKey: IProcessRegistry.EncryptionKeyStruct,
      initStateRoot: BigNumberish
    ],
    [string],
    "nonpayable"
  >;

  owner: TypedContractMethod<[], [string], "view">;

  processCount: TypedContractMethod<[], [bigint], "view">;

  processNonce: TypedContractMethod<[arg0: AddressLike], [bigint], "view">;

  processes: TypedContractMethod<
    [arg0: BytesLike],
    [
      [
        bigint,
        string,
        IProcessRegistry.EncryptionKeyStructOutput,
        bigint,
        bigint,
        bigint,
        bigint,
        bigint,
        string,
        IProcessRegistry.BallotModeStructOutput,
        IProcessRegistry.CensusStructOutput
      ] & {
        status: bigint;
        organizationId: string;
        encryptionKey: IProcessRegistry.EncryptionKeyStructOutput;
        latestStateRoot: bigint;
        startTime: bigint;
        duration: bigint;
        voteCount: bigint;
        voteOverwriteCount: bigint;
        metadataURI: string;
        ballotMode: IProcessRegistry.BallotModeStructOutput;
        census: IProcessRegistry.CensusStructOutput;
      }
    ],
    "view"
  >;

  proxiableUUID: TypedContractMethod<[], [string], "view">;

  rVerifier: TypedContractMethod<[], [string], "view">;

  renounceOwnership: TypedContractMethod<[], [void], "nonpayable">;

  setProcessCensus: TypedContractMethod<
    [processId: BytesLike, census: IProcessRegistry.CensusStruct],
    [void],
    "nonpayable"
  >;

  setProcessDuration: TypedContractMethod<
    [processId: BytesLike, _duration: BigNumberish],
    [void],
    "nonpayable"
  >;

  setProcessResults: TypedContractMethod<
    [processId: BytesLike, proof: BytesLike, input: BytesLike],
    [void],
    "nonpayable"
  >;

  setProcessStatus: TypedContractMethod<
    [processId: BytesLike, newStatus: BigNumberish],
    [void],
    "nonpayable"
  >;

  stVerifier: TypedContractMethod<[], [string], "view">;

  submitStateTransition: TypedContractMethod<
    [processId: BytesLike, proof: BytesLike, input: BytesLike],
    [void],
    "nonpayable"
  >;

  transferOwnership: TypedContractMethod<
    [newOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  upgradeToAndCall: TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "MAX_CENSUS_ORIGIN"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "MAX_STATUS"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "UPGRADE_INTERFACE_VERSION"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "chainID"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "getNextProcessId"
  ): TypedContractMethod<[organizationId: AddressLike], [string], "view">;
  getFunction(
    nameOrSignature: "getProcess"
  ): TypedContractMethod<
    [processId: BytesLike],
    [IProcessRegistry.ProcessStructOutput],
    "view"
  >;
  getFunction(
    nameOrSignature: "getProcessEndTime"
  ): TypedContractMethod<[processId: BytesLike], [bigint], "view">;
  getFunction(
    nameOrSignature: "getRVerifierVKeyHash"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getSTVerifierVKeyHash"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "initialize"
  ): TypedContractMethod<
    [_chainID: BigNumberish, _stVerifier: AddressLike, _rVerifier: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "newProcess"
  ): TypedContractMethod<
    [
      status: BigNumberish,
      startTime: BigNumberish,
      duration: BigNumberish,
      ballotMode: IProcessRegistry.BallotModeStruct,
      census: IProcessRegistry.CensusStruct,
      metadata: string,
      encryptionKey: IProcessRegistry.EncryptionKeyStruct,
      initStateRoot: BigNumberish
    ],
    [string],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "owner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "processCount"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "processNonce"
  ): TypedContractMethod<[arg0: AddressLike], [bigint], "view">;
  getFunction(
    nameOrSignature: "processes"
  ): TypedContractMethod<
    [arg0: BytesLike],
    [
      [
        bigint,
        string,
        IProcessRegistry.EncryptionKeyStructOutput,
        bigint,
        bigint,
        bigint,
        bigint,
        bigint,
        string,
        IProcessRegistry.BallotModeStructOutput,
        IProcessRegistry.CensusStructOutput
      ] & {
        status: bigint;
        organizationId: string;
        encryptionKey: IProcessRegistry.EncryptionKeyStructOutput;
        latestStateRoot: bigint;
        startTime: bigint;
        duration: bigint;
        voteCount: bigint;
        voteOverwriteCount: bigint;
        metadataURI: string;
        ballotMode: IProcessRegistry.BallotModeStructOutput;
        census: IProcessRegistry.CensusStructOutput;
      }
    ],
    "view"
  >;
  getFunction(
    nameOrSignature: "proxiableUUID"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "rVerifier"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "renounceOwnership"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "setProcessCensus"
  ): TypedContractMethod<
    [processId: BytesLike, census: IProcessRegistry.CensusStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setProcessDuration"
  ): TypedContractMethod<
    [processId: BytesLike, _duration: BigNumberish],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setProcessResults"
  ): TypedContractMethod<
    [processId: BytesLike, proof: BytesLike, input: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setProcessStatus"
  ): TypedContractMethod<
    [processId: BytesLike, newStatus: BigNumberish],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "stVerifier"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "submitStateTransition"
  ): TypedContractMethod<
    [processId: BytesLike, proof: BytesLike, input: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "transferOwnership"
  ): TypedContractMethod<[newOwner: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "upgradeToAndCall"
  ): TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;

  getEvent(
    key: "CensusUpdated"
  ): TypedContractEvent<
    CensusUpdatedEvent.InputTuple,
    CensusUpdatedEvent.OutputTuple,
    CensusUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "Initialized"
  ): TypedContractEvent<
    InitializedEvent.InputTuple,
    InitializedEvent.OutputTuple,
    InitializedEvent.OutputObject
  >;
  getEvent(
    key: "OwnershipTransferred"
  ): TypedContractEvent<
    OwnershipTransferredEvent.InputTuple,
    OwnershipTransferredEvent.OutputTuple,
    OwnershipTransferredEvent.OutputObject
  >;
  getEvent(
    key: "ProcessCreated"
  ): TypedContractEvent<
    ProcessCreatedEvent.InputTuple,
    ProcessCreatedEvent.OutputTuple,
    ProcessCreatedEvent.OutputObject
  >;
  getEvent(
    key: "ProcessDurationChanged"
  ): TypedContractEvent<
    ProcessDurationChangedEvent.InputTuple,
    ProcessDurationChangedEvent.OutputTuple,
    ProcessDurationChangedEvent.OutputObject
  >;
  getEvent(
    key: "ProcessResultsSet"
  ): TypedContractEvent<
    ProcessResultsSetEvent.InputTuple,
    ProcessResultsSetEvent.OutputTuple,
    ProcessResultsSetEvent.OutputObject
  >;
  getEvent(
    key: "ProcessStateRootUpdated"
  ): TypedContractEvent<
    ProcessStateRootUpdatedEvent.InputTuple,
    ProcessStateRootUpdatedEvent.OutputTuple,
    ProcessStateRootUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "ProcessStatusChanged"
  ): TypedContractEvent<
    ProcessStatusChangedEvent.InputTuple,
    ProcessStatusChangedEvent.OutputTuple,
    ProcessStatusChangedEvent.OutputObject
  >;
  getEvent(
    key: "Upgraded"
  ): TypedContractEvent<
    UpgradedEvent.InputTuple,
    UpgradedEvent.OutputTuple,
    UpgradedEvent.OutputObject
  >;

  filters: {
    "CensusUpdated(bytes32,bytes32,string,uint256)": TypedContractEvent<
      CensusUpdatedEvent.InputTuple,
      CensusUpdatedEvent.OutputTuple,
      CensusUpdatedEvent.OutputObject
    >;
    CensusUpdated: TypedContractEvent<
      CensusUpdatedEvent.InputTuple,
      CensusUpdatedEvent.OutputTuple,
      CensusUpdatedEvent.OutputObject
    >;

    "Initialized(uint64)": TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;
    Initialized: TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;

    "OwnershipTransferred(address,address)": TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;
    OwnershipTransferred: TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;

    "ProcessCreated(bytes32,address)": TypedContractEvent<
      ProcessCreatedEvent.InputTuple,
      ProcessCreatedEvent.OutputTuple,
      ProcessCreatedEvent.OutputObject
    >;
    ProcessCreated: TypedContractEvent<
      ProcessCreatedEvent.InputTuple,
      ProcessCreatedEvent.OutputTuple,
      ProcessCreatedEvent.OutputObject
    >;

    "ProcessDurationChanged(bytes32,uint256)": TypedContractEvent<
      ProcessDurationChangedEvent.InputTuple,
      ProcessDurationChangedEvent.OutputTuple,
      ProcessDurationChangedEvent.OutputObject
    >;
    ProcessDurationChanged: TypedContractEvent<
      ProcessDurationChangedEvent.InputTuple,
      ProcessDurationChangedEvent.OutputTuple,
      ProcessDurationChangedEvent.OutputObject
    >;

    "ProcessResultsSet(bytes32,address,uint256[])": TypedContractEvent<
      ProcessResultsSetEvent.InputTuple,
      ProcessResultsSetEvent.OutputTuple,
      ProcessResultsSetEvent.OutputObject
    >;
    ProcessResultsSet: TypedContractEvent<
      ProcessResultsSetEvent.InputTuple,
      ProcessResultsSetEvent.OutputTuple,
      ProcessResultsSetEvent.OutputObject
    >;

    "ProcessStateRootUpdated(bytes32,address,uint256)": TypedContractEvent<
      ProcessStateRootUpdatedEvent.InputTuple,
      ProcessStateRootUpdatedEvent.OutputTuple,
      ProcessStateRootUpdatedEvent.OutputObject
    >;
    ProcessStateRootUpdated: TypedContractEvent<
      ProcessStateRootUpdatedEvent.InputTuple,
      ProcessStateRootUpdatedEvent.OutputTuple,
      ProcessStateRootUpdatedEvent.OutputObject
    >;

    "ProcessStatusChanged(bytes32,uint8,uint8)": TypedContractEvent<
      ProcessStatusChangedEvent.InputTuple,
      ProcessStatusChangedEvent.OutputTuple,
      ProcessStatusChangedEvent.OutputObject
    >;
    ProcessStatusChanged: TypedContractEvent<
      ProcessStatusChangedEvent.InputTuple,
      ProcessStatusChangedEvent.OutputTuple,
      ProcessStatusChangedEvent.OutputObject
    >;

    "Upgraded(address)": TypedContractEvent<
      UpgradedEvent.InputTuple,
      UpgradedEvent.OutputTuple,
      UpgradedEvent.OutputObject
    >;
    Upgraded: TypedContractEvent<
      UpgradedEvent.InputTuple,
      UpgradedEvent.OutputTuple,
      UpgradedEvent.OutputObject
    >;
  };
}
