import mongoose, { Document, Schema } from "mongoose";

interface IEvent extends Document {
  title: string;
  content: string;
  createdAt: Date;
  authorID: string;
}

const EventSchema: Schema = new Schema({
  title: { type: String, required: true },
  content: { type: String, required: true },
  createdAt: { type: Date, required: true },
  authorID: { type: String, required: true },
});

const EventModel = mongoose.model<IEvent>("Event", EventSchema);

export default EventModel;
