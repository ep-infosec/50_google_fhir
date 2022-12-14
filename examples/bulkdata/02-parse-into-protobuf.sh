# !/bin/bash
# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

if [[ $# -eq 0 ]] ; then
    echo 'Missing argument: directory with ndjson input files'
    exit 1
fi

echo "converting FHIR files in $1"

bazel run //java/com/google/fhir/examples:ConvertNdJsonForBigQuery -- --output_directory $1 $1/*.ndjson
gzip $1/*.prototxt
