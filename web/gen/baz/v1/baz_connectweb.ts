// @generated by protoc-gen-connect-web v0.8.3 with parameter "target=ts"
// @generated from file baz/v1/baz.proto (package baz.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { DoResponse } from "./baz_pb.js";

/**
 * @generated from service baz.v1.BazService
 */
export const BazService = {
  typeName: "baz.v1.BazService",
  methods: {
    /**
     * @generated from rpc baz.v1.BazService.Do
     */
    do: {
      name: "Do",
      I: Empty,
      O: DoResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
