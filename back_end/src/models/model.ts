import { Model } from "mongoose";
import { UserModel } from "./user";

export interface Model {
  user: Model<UserModel>;
}