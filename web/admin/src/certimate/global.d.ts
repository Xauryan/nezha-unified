import { type BaseModel as PbBaseModel } from "pocketbase";

declare global {
  type ISO8601String = string;

  interface BaseModel extends PbBaseModel {
    created: ISO8601String;
    updated: ISO8601String;
    deleted?: ISO8601String;
  }

  type MaybeModelRecord<T extends BaseModel = BaseModel> = T | Omit<T, "id" | "created" | "updated" | "deleted">;

  type MaybeModelRecordWithId<T extends BaseModel = BaseModel> = T | Pick<T, "id">;

  interface BaseResponse<T = any> {
    code: number;
    msg: string;
    data: T;
  }
}

export {};
