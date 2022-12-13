// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "google/fhir/annotations.h"

#include <string>

#include "google/protobuf/descriptor.pb.h"
#include "google/protobuf/message.h"
#include "absl/status/status.h"
#include "absl/strings/substitute.h"
#include "proto/google/fhir/proto/annotations.pb.h"

namespace google::fhir {

const std::string& GetStructureDefinitionUrl(
    const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(
      proto::fhir_structure_definition_url);
}

const bool IsProfileOf(const ::google::protobuf::Descriptor* descriptor,
                       const ::google::protobuf::Descriptor* potential_base) {
  const std::string& base_url = GetStructureDefinitionUrl(potential_base);
  for (int i = 0;
       i < descriptor->options().ExtensionSize(proto::fhir_profile_base); i++) {
    if (descriptor->options().GetExtension(proto::fhir_profile_base, i) ==
        base_url) {
      return true;
    }
  }
  return false;
}

const bool IsProfile(const ::google::protobuf::Descriptor* descriptor) {
  // Note ContainedResource is not a true FHIR type, an as such, profiles of
  // contained resources don't have a fhir_profile_base
  // TODO(b/244184211): Use an annotation for this.
  return descriptor->options().ExtensionSize(proto::fhir_profile_base) > 0 ||
         (descriptor->name() == "ContainedResource" &&
          (descriptor->full_name() !=
               "google.fhir.stu3.proto.ContainedResource" &&
           descriptor->full_name() != "google.fhir.r4.core.ContainedResource" &&
           descriptor->full_name() !=
               "google.fhir.r5.proto.ContainedResource" &&
           descriptor->full_name() !=
               "google.fhir.dstu2.proto.ContainedResource"));
}

const bool IsChoiceTypeContainer(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::is_choice_type);
}

const bool IsChoiceType(const ::google::protobuf::FieldDescriptor* field) {
  return field->type() == google::protobuf::FieldDescriptor::Type::TYPE_MESSAGE &&
         IsChoiceTypeContainer(field->message_type());
}

const bool IsPrimitive(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::structure_definition_kind) ==
         proto::StructureDefinitionKindValue::KIND_PRIMITIVE_TYPE;
}

const bool IsComplex(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::structure_definition_kind) ==
         proto::StructureDefinitionKindValue::KIND_COMPLEX_TYPE;
}

const bool IsResource(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::structure_definition_kind) ==
         proto::StructureDefinitionKindValue::KIND_RESOURCE;
}

const bool IsReference(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().ExtensionSize(proto::fhir_reference_type) > 0;
}

const std::string& GetValueset(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::fhir_valueset_url);
}

const bool HasValueset(const ::google::protobuf::Descriptor* descriptor) {
  return !GetValueset(descriptor).empty();
}

const std::string& GetFixedSystem(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::fhir_fixed_system);
}

const bool HasFixedSystem(const ::google::protobuf::Descriptor* descriptor) {
  return !GetFixedSystem(descriptor).empty();
}

const std::string& GetInlinedCodingSystem(
    const ::google::protobuf::FieldDescriptor* field) {
  return field->options().GetExtension(proto::fhir_inlined_coding_system);
}

const std::string& GetInlinedCodingCode(
    const ::google::protobuf::FieldDescriptor* field) {
  return field->options().GetExtension(proto::fhir_inlined_coding_code);
}

const std::string& GetFixedCodingSystem(
    const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::fhir_fixed_system);
}

const bool HasFixedCodingSystem(const ::google::protobuf::Descriptor* descriptor) {
  return !GetFixedCodingSystem(descriptor).empty();
}

const std::string& GetSourceCodeSystem(
    const ::google::protobuf::EnumValueDescriptor* descriptor) {
  return descriptor->options().GetExtension(proto::source_code_system);
}

const bool HasSourceCodeSystem(
    const ::google::protobuf::EnumValueDescriptor* descriptor) {
  return !GetSourceCodeSystem(descriptor).empty();
}

const std::string& GetValueRegex(const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->options().GetExtension(proto::value_regex);
}

const bool HasInlinedExtensionUrl(const ::google::protobuf::FieldDescriptor* field) {
  return field->options().HasExtension(proto::fhir_inlined_extension_url);
}

const proto::FhirVersion GetFhirVersion(
    const ::google::protobuf::Descriptor* descriptor) {
  return descriptor->file()->options().GetExtension(proto::fhir_version);
}

const proto::FhirVersion GetFhirVersion(const ::google::protobuf::Message& message) {
  return GetFhirVersion(message.GetDescriptor());
}

const bool IsContainedResource(const ::google::protobuf::Message& message) {
  return IsContainedResource(message.GetDescriptor());
}

const bool IsContainedResource(const ::google::protobuf::Descriptor* descriptor) {
  // TODO(b/244184211): Add a contained-resource specific annotation and read
  // that instead.
  return descriptor->name() == "ContainedResource";
}

absl::Status CheckVersion(const google::protobuf::Message& message,
                          const google::fhir::proto::FhirVersion version) {
  if (google::fhir::GetFhirVersion(message) != version) {
    return absl::InvalidArgumentError(
        absl::Substitute("Called $0 API with non-$0 proto: $1",
                         google::fhir::proto::FhirVersion_Name(version),
                         message.GetDescriptor()->full_name()));
  }
  return absl::OkStatus();
}

}  // namespace google::fhir
