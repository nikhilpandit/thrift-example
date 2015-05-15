// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hello

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Hello interface {
	Ping() (r bool, err error)
	// Parameters:
	//  - Username
	Hello(username string) (r string, err error)
}

type HelloClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewHelloClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HelloClient {
	return &HelloClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewHelloClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HelloClient {
	return &HelloClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *HelloClient) Ping() (r bool, err error) {
	if err = p.sendPing(); err != nil {
		return
	}
	return p.recvPing()
}

func (p *HelloClient) sendPing() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ping", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := PingArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HelloClient) recvPing() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error1 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error2 error
		error2, err = error1.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error2
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result := PingResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.ErrorA1 != nil {
		err = result.ErrorA1
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Username
func (p *HelloClient) Hello(username string) (r string, err error) {
	if err = p.sendHello(username); err != nil {
		return
	}
	return p.recvHello()
}

func (p *HelloClient) sendHello(username string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("hello", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := HelloArgs{
		Username: username,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HelloClient) recvHello() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error3 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error4 error
		error4, err = error3.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error4
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "hello failed: out of sequence response")
		return
	}
	result := HelloResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.ErrorA1 != nil {
		err = result.ErrorA1
		return
	}
	value = result.GetSuccess()
	return
}

type HelloProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Hello
}

func (p *HelloProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *HelloProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *HelloProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewHelloProcessor(handler Hello) *HelloProcessor {

	self5 := &HelloProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self5.processorMap["ping"] = &helloProcessorPing{handler: handler}
	self5.processorMap["hello"] = &helloProcessorHello{handler: handler}
	return self5
}

func (p *HelloProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x6 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x6.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x6

}

type helloProcessorPing struct {
	handler Hello
}

func (p *helloProcessorPing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := PingResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.Ping(); err2 != nil {
		switch v := err2.(type) {
		case *HelloError:
			result.ErrorA1 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: "+err2.Error())
			oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("ping", thrift.REPLY, seqId); err2 != nil {
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
	if err != nil {
		return
	}
	return true, err
}

type helloProcessorHello struct {
	handler Hello
}

func (p *helloProcessorHello) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := HelloArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("hello", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := HelloResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Hello(args.Username); err2 != nil {
		switch v := err2.(type) {
		case *HelloError:
			result.ErrorA1 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing hello: "+err2.Error())
			oprot.WriteMessageBegin("hello", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("hello", thrift.REPLY, seqId); err2 != nil {
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
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type PingArgs struct {
}

func NewPingArgs() *PingArgs {
	return &PingArgs{}
}

func (p *PingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingArgs(%+v)", *p)
}

type PingResult struct {
	Success *bool       `thrift:"success,0" json:"success"`
	ErrorA1 *HelloError `thrift:"error,1" json:"error"`
}

func NewPingResult() *PingResult {
	return &PingResult{}
}

var PingResult_Success_DEFAULT bool

func (p *PingResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return *p.Success
}

var PingResult_ErrorA1_DEFAULT *HelloError

func (p *PingResult) GetErrorA1() *HelloError {
	if !p.IsSetErrorA1() {
		return PingResult_ErrorA1_DEFAULT
	}
	return p.ErrorA1
}
func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PingResult) IsSetErrorA1() bool {
	return p.ErrorA1 != nil
}

func (p *PingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 0: %s", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *PingResult) ReadField1(iprot thrift.TProtocol) error {
	p.ErrorA1 = &HelloError{}
	if err := p.ErrorA1.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ErrorA1, err)
	}
	return nil
}

func (p *PingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *PingResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorA1() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:error: %s", p, err)
		}
		if err := p.ErrorA1.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ErrorA1, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:error: %s", p, err)
		}
	}
	return err
}

func (p *PingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingResult(%+v)", *p)
}

type HelloArgs struct {
	Username string `thrift:"username,1" json:"username"`
}

func NewHelloArgs() *HelloArgs {
	return &HelloArgs{}
}

func (p *HelloArgs) GetUsername() string {
	return p.Username
}
func (p *HelloArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *HelloArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Username = v
	}
	return nil
}

func (p *HelloArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("hello_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *HelloArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:username: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Username)); err != nil {
		return fmt.Errorf("%T.username (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:username: %s", p, err)
	}
	return err
}

func (p *HelloArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloArgs(%+v)", *p)
}

type HelloResult struct {
	Success *string     `thrift:"success,0" json:"success"`
	ErrorA1 *HelloError `thrift:"error,1" json:"error"`
}

func NewHelloResult() *HelloResult {
	return &HelloResult{}
}

var HelloResult_Success_DEFAULT string

func (p *HelloResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return HelloResult_Success_DEFAULT
	}
	return *p.Success
}

var HelloResult_ErrorA1_DEFAULT *HelloError

func (p *HelloResult) GetErrorA1() *HelloError {
	if !p.IsSetErrorA1() {
		return HelloResult_ErrorA1_DEFAULT
	}
	return p.ErrorA1
}
func (p *HelloResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelloResult) IsSetErrorA1() bool {
	return p.ErrorA1 != nil
}

func (p *HelloResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *HelloResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 0: %s", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *HelloResult) ReadField1(iprot thrift.TProtocol) error {
	p.ErrorA1 = &HelloError{}
	if err := p.ErrorA1.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ErrorA1, err)
	}
	return nil
}

func (p *HelloResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("hello_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *HelloResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *HelloResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorA1() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:error: %s", p, err)
		}
		if err := p.ErrorA1.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ErrorA1, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:error: %s", p, err)
		}
	}
	return err
}

func (p *HelloResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloResult(%+v)", *p)
}
