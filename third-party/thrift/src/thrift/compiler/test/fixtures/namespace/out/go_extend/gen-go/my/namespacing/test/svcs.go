// Autogenerated by Thrift for hsmodule.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package test


import (
    "context"
    "fmt"
    "reflect"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = context.Background
var _ = fmt.Printf
var _ = reflect.Ptr
var _ = thrift.ZERO
var _ = metadata.GoUnusedProtection__

type HsTestService interface {
    Init(ctx context.Context, int1 int64) (int64, error)
}

type HsTestServiceChannelClientInterface interface {
    thrift.ClientInterface
    HsTestService
}

type HsTestServiceClientInterface interface {
    thrift.ClientInterface
    Init(int1 int64) (int64, error)
}

type HsTestServiceContextClientInterface interface {
    HsTestServiceClientInterface
    InitContext(ctx context.Context, int1 int64) (int64, error)
}

type HsTestServiceChannelClient struct {
    ch thrift.RequestChannel
}
// Compile time interface enforcer
var _ HsTestServiceChannelClientInterface = (*HsTestServiceChannelClient)(nil)

func NewHsTestServiceChannelClient(channel thrift.RequestChannel) *HsTestServiceChannelClient {
    return &HsTestServiceChannelClient{
        ch: channel,
    }
}

func (c *HsTestServiceChannelClient) Close() error {
    return c.ch.Close()
}

type HsTestServiceClient struct {
    chClient *HsTestServiceChannelClient
}
// Compile time interface enforcer
var _ HsTestServiceClientInterface = (*HsTestServiceClient)(nil)
var _ HsTestServiceContextClientInterface = (*HsTestServiceClient)(nil)

func NewHsTestServiceClient(prot thrift.Protocol) *HsTestServiceClient {
    return &HsTestServiceClient{
        chClient: NewHsTestServiceChannelClient(
            thrift.NewSerialChannel(prot),
        ),
    }
}

func (c *HsTestServiceClient) Close() error {
    return c.chClient.Close()
}

func (c *HsTestServiceChannelClient) Init(ctx context.Context, int1 int64) (int64, error) {
    in := &reqHsTestServiceInit{
        Int1: int1,
    }
    out := newRespHsTestServiceInit()
    err := c.ch.Call(ctx, "init", in, out)
    if err != nil {
        return 0, err
    }
    return out.GetSuccess(), nil
}

func (c *HsTestServiceClient) Init(int1 int64) (int64, error) {
    return c.chClient.Init(context.Background(), int1)
}

func (c *HsTestServiceClient) InitContext(ctx context.Context, int1 int64) (int64, error) {
    return c.chClient.Init(ctx, int1)
}


type HsTestServiceProcessor struct {
    processorFunctionMap map[string]thrift.ProcessorFunction
    functionServiceMap   map[string]string
    handler            HsTestService
}

func NewHsTestServiceProcessor(handler HsTestService) *HsTestServiceProcessor {
    p := &HsTestServiceProcessor{
        handler:              handler,
        processorFunctionMap: make(map[string]thrift.ProcessorFunction),
        functionServiceMap:   make(map[string]string),
    }
    p.AddToProcessorFunctionMap("init", &procFuncHsTestServiceInit{handler: handler})
    p.AddToFunctionServiceMap("init", "HsTestService")

    return p
}

func (p *HsTestServiceProcessor) AddToProcessorFunctionMap(key string, processorFunction thrift.ProcessorFunction) {
    p.processorFunctionMap[key] = processorFunction
}

func (p *HsTestServiceProcessor) AddToFunctionServiceMap(key, service string) {
    p.functionServiceMap[key] = service
}

func (p *HsTestServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction) {
    return p.processorFunctionMap[key]
}

func (p *HsTestServiceProcessor) ProcessorFunctionMap() map[string]thrift.ProcessorFunction {
    return p.processorFunctionMap
}

func (p *HsTestServiceProcessor) FunctionServiceMap() map[string]string {
    return p.functionServiceMap
}

func (p *HsTestServiceProcessor) PackageName() string {
    return "test"
}

func (p *HsTestServiceProcessor) GetThriftMetadata() *metadata.ThriftMetadata {
    return GetThriftMetadataForService("hsmodule.HsTestService")
}


type procFuncHsTestServiceInit struct {
    handler HsTestService
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncHsTestServiceInit)(nil)

func (p *procFuncHsTestServiceInit) Read(iprot thrift.Decoder) (thrift.Struct, thrift.Exception) {
    args := newReqHsTestServiceInit()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncHsTestServiceInit) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("init", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncHsTestServiceInit) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqHsTestServiceInit)
    result := newRespHsTestServiceInit()
    retval, err := p.handler.Init(ctx, args.Int1)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing Init: " + err.Error(), err)
        return x, x
    }

    result.Success = &retval
    return result, nil
}


