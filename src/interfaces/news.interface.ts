export default interface News {
  _id: string;
  title: string;
  content: string;
  picture: string;
  createdAt: Date;
  authorID: string;
}
