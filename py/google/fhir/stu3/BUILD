load(
    "@rules_python//python:python.bzl",
    "py_library",
    "py_test",
)
load("@fhir_bazel_pip_dependencies//:requirements.bzl", "requirement")

package(
    default_visibility = [
        "//visibility:public",
    ],
)

licenses(["notice"])

py_library(
    name = "json_format",
    srcs = [
        "__init__.py",
        "json_format.py",
    ],
    srcs_version = "PY3",
    deps = [
        ":primitive_handler",
        "//py",
        "//py/google/fhir/core/internal:primitive_time_utils",
        "//py/google/fhir/core/internal:resource_validation",
        "//py/google/fhir/core/internal/json_format",
        "@com_google_protobuf//:protobuf_python",
    ],
)

py_library(
    name = "primitive_handler",
    srcs = [
        "__init__.py",
        "primitive_handler.py",
    ],
    srcs_version = "PY3",
    deps = [
        "//py",
        "//proto/google/fhir/proto/stu3:datatypes_py_pb2",
        "//proto/google/fhir/proto/stu3:fhirproto_extensions_py_pb2",
        "//proto/google/fhir/proto/stu3:resources_py_pb2",
        "//py/google/fhir/core/internal:primitive_handler",
        "//py/google/fhir/core/internal:primitive_time_utils",
        "//py/google/fhir/core/internal/json_format/wrappers:base64_binary",
        "//py/google/fhir/core/internal/json_format/wrappers:date",
        "//py/google/fhir/core/internal/json_format/wrappers:date_time",
        "//py/google/fhir/core/internal/json_format/wrappers:decimal",
        "//py/google/fhir/core/internal/json_format/wrappers:instant",
        "//py/google/fhir/core/internal/json_format/wrappers:primitive_wrappers",
        "//py/google/fhir/core/internal/json_format/wrappers:time",
        "//py/google/fhir/core/utils:annotation_utils",
        "//py/google/fhir/core/utils:fhir_types",
        "@com_google_protobuf//:protobuf_python",
    ],
)

py_test(
    name = "codes_test",
    size = "small",
    srcs = [
        "codes_test.py",
    ],
    data = [
        "//testdata/stu3/codes:testdata",
    ],
    legacy_create_init = 0,
    python_version = "PY3",
    deps = [
        "@com_google_protobuf//:protobuf_python",
        "//py",
        "//proto/google/fhir/proto/stu3:codes_py_pb2",
        "//proto/google/fhir/proto/stu3:datatypes_py_pb2",
        "//py/google/fhir/core:codes",
        "//py/google/fhir/core:fhir_errors",
        "//py/google/fhir/core/testing:testdata_utils",
        requirement("absl-py"),
        requirement("six"),
    ],
)

py_test(
    name = "extensions_test",
    srcs = [
        "extensions_test.py",
    ],
    data = [
        "//testdata/stu3:extensions",
    ],
    legacy_create_init = 0,
    python_version = "PY3",
    deps = [
        "@com_google_protobuf//:protobuf_python",
        "//py",
        "//proto/google/fhir/proto/stu3:datatypes_py_pb2",
        "//proto/google/fhir/proto/stu3:extensions_py_pb2",
        "//proto/google/fhir/proto/stu3:fhirproto_extensions_py_pb2",
        "//proto/google/fhir/proto/stu3:ml_extensions_py_pb2",
        "//proto/google/fhir/proto/stu3:resources_py_pb2",
        "//py/google/fhir/core:extensions",
        "//py/google/fhir/core:extensions_test",
        "//testdata/stu3/profiles:test_extensions_py_pb2",
        requirement("absl-py"),
        requirement("six"),
    ],
)

py_test(
    name = "json_format_test",
    size = "large",
    srcs = [
        "json_format_test.py",
    ],
    data = [
        "//spec:stu3",
        "//testdata/stu3:bigquery",
        "//testdata/stu3:examples",
        "//testdata/stu3:validation",
    ],
    legacy_create_init = 0,
    python_version = "PY3",
    shard_count = 16,
    deps = [
        "@com_google_protobuf//:protobuf_python",
        "//py",
        "//proto/google/fhir/proto/stu3:datatypes_py_pb2",
        "//proto/google/fhir/proto/stu3:resources_py_pb2",
        "//py/google/fhir/core/internal/json_format:json_format_test",
        ":json_format",
        "//py/google/fhir/core/testing:testdata_utils",
        "//py/google/fhir/core/utils:proto_utils",
        requirement("absl-py"),
        requirement("six"),
    ],
)

py_test(
    name = "primitive_handler_test",
    srcs = [
        "primitive_handler_test.py",
    ],
    data = [
        "//testdata/stu3:validation",
    ],
    legacy_create_init = 0,
    python_version = "PY3",
    deps = [
        "//py",
        "//py/google/fhir/core/internal:primitive_handler_test",
        ":primitive_handler",
        requirement("absl-py"),
        requirement("six"),
    ],
)

py_test(
    name = "references_test",
    srcs = [
        "references_test.py",
    ],
    data = [],
    legacy_create_init = 0,
    python_version = "PY3",
    deps = [
        "//py",
        "//proto/google/fhir/proto/stu3:datatypes_py_pb2",
        "//py/google/fhir/core:references",
        "//py/google/fhir/core/utils:proto_utils",
        requirement("absl-py"),
        requirement("six"),
    ],
)

py_library(
    name = "resource_validation",
    srcs = [
        "__init__.py",
        "resource_validation.py",
    ],
    srcs_version = "PY3",
    deps = [
        ":primitive_handler",
        "//py",
        "//py/google/fhir/core/internal:resource_validation",
        "@com_google_protobuf//:protobuf_python",
    ],
)

py_test(
    name = "resource_validation_test",
    srcs = [
        "resource_validation_test.py",
    ],
    data = [
        "//testdata/stu3:validation",
    ],
    legacy_create_init = 0,
    python_version = "PY3",
    deps = [
        ":resource_validation",
        "@com_google_protobuf//:protobuf_python",
        "//py",
        "//proto/google/fhir/proto/stu3:resources_py_pb2",
        "//py/google/fhir/core:fhir_errors",
        "//py/google/fhir/core/testing:testdata_utils",
        requirement("absl-py"),
        requirement("six"),
    ],
)
