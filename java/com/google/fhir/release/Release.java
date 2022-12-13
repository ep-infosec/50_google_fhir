//    Copyright 2020 Google Inc.
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        https://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package com.google.fhir.release;

/**
 * A dummy main class for a Java binary.
 *
 * In order to create a single JAR that includes the transitive closure of dependencies required for
 * FHIR protos, we use the *_deploy.jar generated by the java_binary rule. See
 * https://docs.bazel.build/versions/master/be/java.html#java_binary. This is the main class the
 * java_binary rule requires.
 */
final class Release {
  private Release() {
  }

  public static void main(String[] args) {
  }
}
