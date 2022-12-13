/*
 * Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef GOOGLE_FHIR_STU3_PROFILES_H_
#define GOOGLE_FHIR_STU3_PROFILES_H_

#include "google/protobuf/message.h"
#include "absl/status/status.h"
#include "absl/status/statusor.h"
#include "google/fhir/error_reporter.h"
#include "google/fhir/status/status.h"
#include "google/fhir/status/statusor.h"
#include "proto/google/fhir/proto/stu3/resources.pb.h"

namespace google {
namespace fhir {
namespace stu3 {

// Converts FHIR resources between profiled and unprofiled, reporting any
// conversion or validation errors to the given ScopedErrorReporter.  Conversion
// will continue processing as long as the ScopedErrorReporter returns an Ok
// status for all errors it is given.
//
// If <target> is a profiled type of <source>:
// Converts a resource to a profiled version of that resource.
// If the profile adds new inlined fields for Extensions or Codings within
// CodeableConcepts, and those extensions or codings are present on the base
// message, the data will be reorganized into the new fields in the profile.
// Finally, this runs the Fhir resource validation code on the resulting
// message, to ensure the result complies with the requirements of the profile
// (e.g., fields that are considered required by the profile).
//
// If <target> is a base type of <source>:
// Performs the inverse operation to the above.
// Any data that is in an inlined field in the profiled message,
// that does not exists in the base message will be converted back to the
// original format (e.g., extension).
absl::Status ConvertToProfile(
    const ::google::protobuf::Message& source, ::google::protobuf::Message* target,
    const ::google::fhir::ScopedErrorReporter& reporter);

// Converts FHIR resources between profiled and unprofiled, returning any
// conversion errors or warnings in an OperationOutcome. Users should check
// the returned outcome and properly handle or report issues.
//
// See the above ConvertToProfile method for details on the source and target
// parameters.
absl::StatusOr<::google::fhir::stu3::proto::OperationOutcome> ConvertToProfile(
    const ::google::protobuf::Message& source, ::google::protobuf::Message* target);
}  // namespace stu3

// Deprecated. Use ::google::fhir::stu3::ConvertToProfile instead.
//
// If <target> is a profiled type of <source>:
// Converts a resource to a profiled version of that resource.
// If the profile adds new inlined fields for Extensions or Codings within
// CodeableConcepts, and those extensions or codings are present on the base
// message, the data will be reorganized into the new fields in the profile.
// Finally, this runs the Fhir resource validation code on the resulting
// message, to ensure the result complies with the requirements of the profile
// (e.g., fields that are considered required by the profile).
//
// If <target> is a base type of <source>:
// Performs the inverse operation to the above.
// Any data that is in an inlined field in the profiled message,
// that does not exists in the base message will be converted back to the
// original format (e.g., extension).
absl::Status ConvertToProfileStu3(const ::google::protobuf::Message& source,
                                  ::google::protobuf::Message* target);

// Deprecated. Lenient conversions never had well-defined semantics,
// so users should move to other ConvertToProfile functions where they can
// chose whether to ignore specific validation errors.
//
// Identical to ConvertToProfile, except does not run the validation step.
absl::Status ConvertToProfileLenientStu3(const ::google::protobuf::Message& source,
                                         ::google::protobuf::Message* target);

// Given a Message, returns a copy with all data is stored in typed fields where
// possible.
// E.g., if the message contains an extension in the raw extension field that
// has a corresponding typed field, the return copy will have the data in the
// typed field.
template <typename T>
absl::StatusOr<T> NormalizeStu3(const T& message) {
  T normalized;
  FHIR_RETURN_IF_ERROR(ConvertToProfileLenientStu3(message, &normalized));
  return normalized;
}

// Normalizes a resource, and then validates the output.
// Returns any status error, either from normalizing or validating.
// This guarantees that if this function returns a resource, that resource
// is both in normalized form, and valid according to all restrictions of
// the profile.
template <typename T>
absl::StatusOr<T> NormalizeAndValidateStu3(const T& message) {
  T normalized;
  FHIR_RETURN_IF_ERROR(ConvertToProfileStu3(message, &normalized));
  return normalized;
}

}  // namespace fhir
}  // namespace google

#endif  // GOOGLE_FHIR_STU3_PROFILES_H_
