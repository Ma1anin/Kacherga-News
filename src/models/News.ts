import mongoose, { Document, Schema } from "mongoose";

interface INews extends Document {
  title: string;
  content: string;
  picture: string;
  createdAt: Date;
  authorID: string;
}

const NewsSchema: Schema = new Schema({
  title: { type: String, required: true },
  content: { type: String, required: true },
  picture: { type: String },
  createdAt: { type: Date, required: true },
  authorID: { type: String, required: true },
});

const NewsModel = mongoose.model<INews>("News", NewsSchema);

export default NewsModel;
