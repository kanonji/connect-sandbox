// @generated by protoc-gen-connect-web v0.8.1 with parameter "target=ts"
// @generated from file protobuf/bar/v1/bar.proto (package protobuf.bar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { GetRequest, GetResponse, InvokeResponse } from "./bar_pb.js";

/**
 * @generated from service protobuf.bar.v1.BarService
 */
export const BarService = {
  typeName: "protobuf.bar.v1.BarService",
  methods: {
    /**
     * @generated from rpc protobuf.bar.v1.BarService.Invoke
     */
    invoke: {
      name: "Invoke",
      I: Empty,
      O: InvokeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc protobuf.bar.v1.BarService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

