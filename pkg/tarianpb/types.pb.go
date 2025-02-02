// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tarianpb/types.proto

package tarianpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MatchLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MatchLabel) Reset() {
	*x = MatchLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchLabel) ProtoMessage() {}

func (x *MatchLabel) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchLabel.ProtoReflect.Descriptor instead.
func (*MatchLabel) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{0}
}

func (x *MatchLabel) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MatchLabel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchLabels []*MatchLabel `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{1}
}

func (x *Selector) GetMatchLabels() []*MatchLabel {
	if x != nil {
		return x.MatchLabels
	}
	return nil
}

type AllowedProcessRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regex *string `protobuf:"bytes,1,opt,name=regex,proto3,oneof" json:"regex,omitempty"`
}

func (x *AllowedProcessRule) Reset() {
	*x = AllowedProcessRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedProcessRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedProcessRule) ProtoMessage() {}

func (x *AllowedProcessRule) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedProcessRule.ProtoReflect.Descriptor instead.
func (*AllowedProcessRule) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{2}
}

func (x *AllowedProcessRule) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

type AllowedFileRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sha256Sum *string `protobuf:"bytes,2,opt,name=sha256sum,proto3,oneof" json:"sha256sum,omitempty"`
}

func (x *AllowedFileRule) Reset() {
	*x = AllowedFileRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedFileRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedFileRule) ProtoMessage() {}

func (x *AllowedFileRule) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedFileRule.ProtoReflect.Descriptor instead.
func (*AllowedFileRule) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{3}
}

func (x *AllowedFileRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllowedFileRule) GetSha256Sum() string {
	if x != nil && x.Sha256Sum != nil {
		return *x.Sha256Sum
	}
	return ""
}

type Constraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind             string                `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Namespace        string                `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name             string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Selector         *Selector             `protobuf:"bytes,4,opt,name=selector,proto3" json:"selector,omitempty"`
	AllowedProcesses []*AllowedProcessRule `protobuf:"bytes,5,rep,name=allowed_processes,json=allowedProcesses,proto3" json:"allowed_processes,omitempty"`
	AllowedFiles     []*AllowedFileRule    `protobuf:"bytes,6,rep,name=allowed_files,json=allowedFiles,proto3" json:"allowed_files,omitempty"`
}

func (x *Constraint) Reset() {
	*x = Constraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraint) ProtoMessage() {}

func (x *Constraint) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraint.ProtoReflect.Descriptor instead.
func (*Constraint) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{4}
}

func (x *Constraint) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Constraint) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Constraint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Constraint) GetSelector() *Selector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *Constraint) GetAllowedProcesses() []*AllowedProcessRule {
	if x != nil {
		return x.AllowedProcesses
	}
	return nil
}

func (x *Constraint) GetAllowedFiles() []*AllowedFileRule {
	if x != nil {
		return x.AllowedFiles
	}
	return nil
}

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  int32  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{5}
}

func (x *Process) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Process) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ViolatedFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ActualSha256Sum   string `protobuf:"bytes,2,opt,name=actualSha256Sum,proto3" json:"actualSha256Sum,omitempty"`
	ExpectedSha256Sum string `protobuf:"bytes,3,opt,name=expectedSha256Sum,proto3" json:"expectedSha256Sum,omitempty"`
}

func (x *ViolatedFile) Reset() {
	*x = ViolatedFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViolatedFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViolatedFile) ProtoMessage() {}

func (x *ViolatedFile) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViolatedFile.ProtoReflect.Descriptor instead.
func (*ViolatedFile) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{6}
}

func (x *ViolatedFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViolatedFile) GetActualSha256Sum() string {
	if x != nil {
		return x.ActualSha256Sum
	}
	return ""
}

func (x *ViolatedFile) GetExpectedSha256Sum() string {
	if x != nil {
		return x.ExpectedSha256Sum
	}
	return ""
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{7}
}

func (x *Label) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Namespace string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Labels    []*Label `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{8}
}

func (x *Pod) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Pod) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pod) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type FalcoAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule         string            `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	Priority     string            `protobuf:"bytes,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Output       string            `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	OutputFields map[string]string `protobuf:"bytes,4,rep,name=outputFields,proto3" json:"outputFields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FalcoAlert) Reset() {
	*x = FalcoAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FalcoAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FalcoAlert) ProtoMessage() {}

func (x *FalcoAlert) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FalcoAlert.ProtoReflect.Descriptor instead.
func (*FalcoAlert) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{9}
}

func (x *FalcoAlert) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *FalcoAlert) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *FalcoAlert) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *FalcoAlert) GetOutputFields() map[string]string {
	if x != nil {
		return x.OutputFields
	}
	return nil
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod               *Pod            `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	ViolatedProcesses []*Process      `protobuf:"bytes,2,rep,name=violatedProcesses,proto3" json:"violatedProcesses,omitempty"`
	ViolatedFiles     []*ViolatedFile `protobuf:"bytes,3,rep,name=violatedFiles,proto3" json:"violatedFiles,omitempty"`
	FalcoAlert        *FalcoAlert     `protobuf:"bytes,4,opt,name=falcoAlert,proto3,oneof" json:"falcoAlert,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{10}
}

func (x *Target) GetPod() *Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

func (x *Target) GetViolatedProcesses() []*Process {
	if x != nil {
		return x.ViolatedProcesses
	}
	return nil
}

func (x *Target) GetViolatedFiles() []*ViolatedFile {
	if x != nil {
		return x.ViolatedFiles
	}
	return nil
}

func (x *Target) GetFalcoAlert() *FalcoAlert {
	if x != nil {
		return x.FalcoAlert
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind            string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Type            string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Uid             string                 `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	ServerTimestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=serverTimestamp,proto3" json:"serverTimestamp,omitempty"`
	ClientTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=clientTimestamp,proto3" json:"clientTimestamp,omitempty"`
	AlertSentAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=alertSentAt,proto3" json:"alertSentAt,omitempty"`
	Targets         []*Target              `protobuf:"bytes,7,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_types_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_types_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_tarianpb_types_proto_rawDescGZIP(), []int{11}
}

func (x *Event) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Event) GetServerTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ServerTimestamp
	}
	return nil
}

func (x *Event) GetClientTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ClientTimestamp
	}
	return nil
}

func (x *Event) GetAlertSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AlertSentAt
	}
	return nil
}

func (x *Event) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

var File_tarianpb_types_proto protoreflect.FileDescriptor

var file_tarianpb_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a,
	0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0b, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x39, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x22, 0x56, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x68,
	0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x22, 0x9f, 0x02, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x4f, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x2f, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7a,
	0x0a, 0x0c, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x32,
	0x35, 0x36, 0x53, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x53, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x11,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x53, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x53, 0x75, 0x6d, 0x22, 0x2f, 0x0a, 0x05, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x78, 0x0a, 0x03, 0x50,
	0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70,
	0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x0a, 0x46, 0x61, 0x6c, 0x63, 0x6f, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x50, 0x0a, 0x0c,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6c, 0x63, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x3f,
	0x0a, 0x11, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x8a, 0x02, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x70, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x03, 0x70, 0x6f,
	0x64, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x11, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0d, 0x76, 0x69, 0x6f, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0d, 0x76,
	0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0a,
	0x66, 0x61, 0x6c, 0x63, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x46, 0x61, 0x6c, 0x63, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x66, 0x61, 0x6c, 0x63, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x22, 0xbd, 0x02, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x44, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x44, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x0b,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6f, 0x70,
	0x73, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x31, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tarianpb_types_proto_rawDescOnce sync.Once
	file_tarianpb_types_proto_rawDescData = file_tarianpb_types_proto_rawDesc
)

func file_tarianpb_types_proto_rawDescGZIP() []byte {
	file_tarianpb_types_proto_rawDescOnce.Do(func() {
		file_tarianpb_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tarianpb_types_proto_rawDescData)
	})
	return file_tarianpb_types_proto_rawDescData
}

var file_tarianpb_types_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tarianpb_types_proto_goTypes = []interface{}{
	(*MatchLabel)(nil),            // 0: tarianpb.types.MatchLabel
	(*Selector)(nil),              // 1: tarianpb.types.Selector
	(*AllowedProcessRule)(nil),    // 2: tarianpb.types.AllowedProcessRule
	(*AllowedFileRule)(nil),       // 3: tarianpb.types.AllowedFileRule
	(*Constraint)(nil),            // 4: tarianpb.types.Constraint
	(*Process)(nil),               // 5: tarianpb.types.Process
	(*ViolatedFile)(nil),          // 6: tarianpb.types.ViolatedFile
	(*Label)(nil),                 // 7: tarianpb.types.Label
	(*Pod)(nil),                   // 8: tarianpb.types.Pod
	(*FalcoAlert)(nil),            // 9: tarianpb.types.FalcoAlert
	(*Target)(nil),                // 10: tarianpb.types.Target
	(*Event)(nil),                 // 11: tarianpb.types.Event
	nil,                           // 12: tarianpb.types.FalcoAlert.OutputFieldsEntry
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_tarianpb_types_proto_depIdxs = []int32{
	0,  // 0: tarianpb.types.Selector.match_labels:type_name -> tarianpb.types.MatchLabel
	1,  // 1: tarianpb.types.Constraint.selector:type_name -> tarianpb.types.Selector
	2,  // 2: tarianpb.types.Constraint.allowed_processes:type_name -> tarianpb.types.AllowedProcessRule
	3,  // 3: tarianpb.types.Constraint.allowed_files:type_name -> tarianpb.types.AllowedFileRule
	7,  // 4: tarianpb.types.Pod.labels:type_name -> tarianpb.types.Label
	12, // 5: tarianpb.types.FalcoAlert.outputFields:type_name -> tarianpb.types.FalcoAlert.OutputFieldsEntry
	8,  // 6: tarianpb.types.Target.pod:type_name -> tarianpb.types.Pod
	5,  // 7: tarianpb.types.Target.violatedProcesses:type_name -> tarianpb.types.Process
	6,  // 8: tarianpb.types.Target.violatedFiles:type_name -> tarianpb.types.ViolatedFile
	9,  // 9: tarianpb.types.Target.falcoAlert:type_name -> tarianpb.types.FalcoAlert
	13, // 10: tarianpb.types.Event.serverTimestamp:type_name -> google.protobuf.Timestamp
	13, // 11: tarianpb.types.Event.clientTimestamp:type_name -> google.protobuf.Timestamp
	13, // 12: tarianpb.types.Event.alertSentAt:type_name -> google.protobuf.Timestamp
	10, // 13: tarianpb.types.Event.targets:type_name -> tarianpb.types.Target
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_tarianpb_types_proto_init() }
func file_tarianpb_types_proto_init() {
	if File_tarianpb_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tarianpb_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchLabel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedProcessRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedFileRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViolatedFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FalcoAlert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tarianpb_types_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_tarianpb_types_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_tarianpb_types_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_tarianpb_types_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tarianpb_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tarianpb_types_proto_goTypes,
		DependencyIndexes: file_tarianpb_types_proto_depIdxs,
		MessageInfos:      file_tarianpb_types_proto_msgTypes,
	}.Build()
	File_tarianpb_types_proto = out.File
	file_tarianpb_types_proto_rawDesc = nil
	file_tarianpb_types_proto_goTypes = nil
	file_tarianpb_types_proto_depIdxs = nil
}
