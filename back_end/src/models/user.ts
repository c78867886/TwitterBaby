import { Document } from "mongoose";
import { User } from "../interfaces/user";

export interface UserModel extends User, Document {
  //custom methods for your model would be defined here
}