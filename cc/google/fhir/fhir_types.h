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

#ifndef GOOGLE_FHIR_FHIR_TYPES_H_
#define GOOGLE_FHIR_FHIR_TYPES_H_

#include "google/protobuf/message.h"

namespace google {
namespace fhir {

// Provides version-independent tests for resource type without pulling in
// resource dependencies.

bool IsBundle(const ::google::protobuf::Message& message);
bool IsProfileOfBundle(const ::google::protobuf::Message& message);
bool IsTypeOrProfileOfBundle(const ::google::protobuf::Message& message);
bool IsBundle(const ::google::protobuf::Descriptor* descriptor);
bool IsProfileOfBundle(const ::google::protobuf::Descriptor* descriptor);
bool IsTypeOrProfileOfBundle(const ::google::protobuf::Descriptor* descriptor);

bool IsCodeableConcept(const ::google::protobuf::Message& message);
bool IsProfileOfCodeableConcept(const ::google::protobuf::Message& message);
bool IsTypeOrProfileOfCodeableConcept(const ::google::protobuf::Message& message);
bool IsCodeableConcept(const ::google::protobuf::Descriptor* descriptor);
bool IsProfileOfCodeableConcept(const ::google::protobuf::Descriptor* descriptor);
bool IsTypeOrProfileOfCodeableConcept(const ::google::protobuf::Descriptor* descriptor);

bool IsCoding(const ::google::protobuf::Message& message);
bool IsProfileOfCoding(const ::google::protobuf::Message& message);
bool IsTypeOrProfileOfCoding(const ::google::protobuf::Message& message);
bool IsCoding(const ::google::protobuf::Descriptor* descriptor);
bool IsProfileOfCoding(const ::google::protobuf::Descriptor* descriptor);
bool IsTypeOrProfileOfCoding(const ::google::protobuf::Descriptor* descriptor);

bool IsCode(const ::google::protobuf::Message& message);
bool IsProfileOfCode(const ::google::protobuf::Message& message);
bool IsTypeOrProfileOfCode(const ::google::protobuf::Message& message);
bool IsCode(const ::google::protobuf::Descriptor* descriptor);
bool IsProfileOfCode(const ::google::protobuf::Descriptor* descriptor);
bool IsTypeOrProfileOfCode(const ::google::protobuf::Descriptor* descriptor);

bool IsExtension(const ::google::protobuf::Message& message);
bool IsProfileOfExtension(const ::google::protobuf::Message& message);
bool IsTypeOrProfileOfExtension(const ::google::protobuf::Message& message);
bool IsExtension(const ::google::protobuf::Descriptor* descriptor);
bool IsProfileOfExtension(const ::google::protobuf::Descriptor* descriptor);
bool IsTypeOrProfileOfExtension(const ::google::protobuf::Descriptor* descriptor);

bool IsBoolean(const ::google::protobuf::Message& message);
bool IsBoolean(const ::google::protobuf::Descriptor* descriptor);

bool IsBase64Binary(const ::google::protobuf::Message& message);
bool IsBase64Binary(const ::google::protobuf::Descriptor* descriptor);

bool IsString(const ::google::protobuf::Message& message);
bool IsString(const ::google::protobuf::Descriptor* descriptor);

bool IsInteger(const ::google::protobuf::Message& message);
bool IsInteger(const ::google::protobuf::Descriptor* descriptor);

bool IsPositiveInt(const ::google::protobuf::Message& message);
bool IsPositiveInt(const ::google::protobuf::Descriptor* descriptor);

bool IsUnsignedInt(const ::google::protobuf::Message& message);
bool IsUnsignedInt(const ::google::protobuf::Descriptor* descriptor);

bool IsDecimal(const ::google::protobuf::Message& message);
bool IsDecimal(const ::google::protobuf::Descriptor* descriptor);

bool IsDateTime(const ::google::protobuf::Message& message);
bool IsDateTime(const ::google::protobuf::Descriptor* descriptor);

bool IsDate(const ::google::protobuf::Message& message);
bool IsDate(const ::google::protobuf::Descriptor* descriptor);

bool IsTime(const ::google::protobuf::Message& message);
bool IsTime(const ::google::protobuf::Descriptor* descriptor);

bool IsId(const ::google::protobuf::Message& message);
bool IsId(const ::google::protobuf::Descriptor* descriptor);

bool IsUuid(const ::google::protobuf::Message& message);
bool IsUuid(const ::google::protobuf::Descriptor* descriptor);

bool IsIdentifier(const ::google::protobuf::Message& message);
bool IsIdentifier(const ::google::protobuf::Descriptor* descriptor);

bool IsInstant(const ::google::protobuf::Message& message);
bool IsInstant(const ::google::protobuf::Descriptor* descriptor);

bool IsUri(const ::google::protobuf::Message& message);
bool IsUri(const ::google::protobuf::Descriptor* descriptor);

bool IsPeriod(const ::google::protobuf::Message& message);
bool IsPeriod(const ::google::protobuf::Descriptor* descriptor);

bool IsUrl(const ::google::protobuf::Message& message);
bool IsUrl(const ::google::protobuf::Descriptor* descriptor);

bool IsCanonical(const ::google::protobuf::Message& message);
bool IsCanonical(const ::google::protobuf::Descriptor* descriptor);

bool IsOid(const ::google::protobuf::Message& message);
bool IsOid(const ::google::protobuf::Descriptor* descriptor);

bool IsQuantity(const ::google::protobuf::Message& message);
bool IsQuantity(const ::google::protobuf::Descriptor* descriptor);

bool IsSimpleQuantity(const ::google::protobuf::Message& message);
bool IsSimpleQuantity(const ::google::protobuf::Descriptor* descriptor);

bool IsXhtml(const ::google::protobuf::Message& message);
bool IsXhtml(const ::google::protobuf::Descriptor* descriptor);

bool IsMarkdown(const ::google::protobuf::Message& message);
bool IsMarkdown(const ::google::protobuf::Descriptor* descriptor);

}  // namespace fhir
}  // namespace google

#endif  // GOOGLE_FHIR_FHIR_TYPES_H_
