//    Copyright 2021 Google Inc.
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

package com.google.fhir.common;

import static com.google.common.truth.Truth.assertThat;
import static com.google.fhir.common.ProtoUtils.findField;
import static org.junit.Assert.assertThrows;

import com.google.common.collect.ImmutableCollection;
import com.google.common.collect.ImmutableList;
import com.google.protobuf.Descriptors.EnumDescriptor;
import com.google.protobuf.Message;
import com.google.protobuf.MessageOrBuilder;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;

@RunWith(Parameterized.class)
@SuppressWarnings("UnnecessarilyFullyQualified")
public final class CodesTest {
  private final Message rawCodeType;
  private final Message administrativeGenderType;
  private final Message comparatorCodeType;
  private final Message stringTypeCode;

  public CodesTest(
      Message rawCodeType,
      Message administrativeGenderType,
      Message comparatorCodeType,
      Message stringTypeCode) {
    this.rawCodeType = rawCodeType;
    this.administrativeGenderType = administrativeGenderType;
    this.comparatorCodeType = comparatorCodeType;
    this.stringTypeCode = stringTypeCode;
  }

  private static EnumDescriptor getEnumDescriptor(MessageOrBuilder message) throws Exception {
    return findField(message, "value").getEnumType();
  }

  @Parameterized.Parameters
  public static ImmutableCollection<Object[]> params() {
    return ImmutableList.of(
        new Object[] {
          com.google.fhir.r4.core.Code.getDefaultInstance(),
          com.google.fhir.r4.core.Patient.GenderCode.getDefaultInstance(),
          com.google.fhir.r4.core.Age.ComparatorCode.getDefaultInstance(),
          com.google.fhir.r4.core.Binary.ContentTypeCode.getDefaultInstance()
        },
        new Object[] {
          com.google.fhir.stu3.proto.Code.getDefaultInstance(),
          com.google.fhir.stu3.proto.AdministrativeGenderCode.getDefaultInstance(),
          com.google.fhir.stu3.proto.QuantityComparatorCode.getDefaultInstance(),
          com.google.fhir.stu3.proto.MimeTypeCode.getDefaultInstance()
        });
  }

  @Test
  public void testEnumValueToStringCode_withDefaultValue() throws Exception {
    assertThat(
            Codes.enumValueToCodeString(
                getEnumDescriptor(administrativeGenderType).getValues().get(1)))
        .isEqualTo("male");
  }

  @Test
  public void testEnumValueToStringCode_withJsonNameOverride() throws Exception {
    assertThat(
            Codes.enumValueToCodeString(getEnumDescriptor(comparatorCodeType).getValues().get(1)))
        .isEqualTo("<");
  }

  @Test
  public void testGetCodeAsString_rawCode() throws Exception {
    Message.Builder builder = rawCodeType.newBuilderForType();
    builder.setField(findField(builder, "value"), "flower");
    assertThat(Codes.getCodeAsString(builder)).isEqualTo("flower");
  }

  @Test
  public void testGetCodeAsString_enumTypeCode() throws Exception {
    Message.Builder builder = comparatorCodeType.newBuilderForType();
    builder.setField(
        findField(builder, "value"), getEnumDescriptor(comparatorCodeType).findValueByNumber(3));

    assertThat(Codes.getCodeAsString(builder)).isEqualTo(">=");
  }

  @Test
  public void testGetCodeAsString_stringTypeCode() throws Exception {
    Message.Builder builder = stringTypeCode.newBuilderForType();
    builder.setField(findField(builder, "value"), "flower");
    assertThat(Codes.getCodeAsString(builder)).isEqualTo("flower");
  }

  @Test
  public void testCodeStringToEnumValue_success() throws Exception {
    assertThat(
            Codes.codeStringToEnumValue(
                comparatorCodeType.getDescriptorForType().findFieldByName("value").getEnumType(),
                ">="))
        .isEqualTo(getEnumDescriptor(comparatorCodeType).findValueByNumber(3));
  }

  @Test
  public void testCodeStringToEnumValue_notFound() {
    assertThrows(
        InvalidFhirException.class,
        () ->
            Codes.codeStringToEnumValue(
                comparatorCodeType.getDescriptorForType().findFieldByName("value").getEnumType(),
                "monkey"));
  }
}
