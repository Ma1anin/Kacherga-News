import mongoose, { Document, Schema } from "mongoose";

interface IUser extends Document {
  login: string;
  password: string;
  name: string;
  role: string;
  picture: string;
}

const UserSchema: Schema = new Schema({
  login: { type: String, required: true },
  password: { type: String, required: true },
  name: { type: String, required: true },
  role: { type: String, required: true },
  picture: { type: String },
});

const UserModel = mongoose.model<IUser>("User", UserSchema);

export default UserModel;
